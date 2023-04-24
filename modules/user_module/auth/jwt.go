package auth

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/util"
)

//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

const (
	IDKey    = "user_id"
	RoleKey  = "role"
	EmailKey = "email"
)

type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Store                 Store
	Clocker               util.Clocker
}

func NewJWTer(s Store, c util.Clocker) (*JWTer, error) {
	j := &JWTer{Store: s}
	privKey, err := parse(rawPrivKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}

	pubKey, err := parse(rawPubKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}
	j.PrivateKey = privKey
	j.PublicKey = pubKey
	j.Clocker = c
	return j, nil
}

func parse(rawKey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, err
}

func (j *JWTer) GenerateToken(ctx context.Context, u user_model.User) ([]byte, error) {
	tok, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`github.com/tatuya-web/go-modular-monolith`).
		Subject("access_token").
		IssuedAt(j.Clocker.Now()).
		Expiration(j.Clocker.Now().Add(30*time.Minute)).
		Claim(RoleKey, u.Role).
		Claim(EmailKey, u.Email).
		Build()
	if err != nil {
		return nil, fmt.Errorf("GetToken: failed to build token: %w", err)
	}
	if err := j.Store.Save(ctx, tok.JwtID(), u.ID); err != nil {
		return nil, err
	}

	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil {
		return nil, err
	}
	return signed, nil
}

func (j *JWTer) FillContext(ctx *gin.Context) error {
	token, err := j.GetToken(ctx.Request.Context(), ctx.Request)
	if err != nil {
		return err
	}
	uid, err := j.Store.Load(ctx.Request.Context(), token.JwtID())
	if err != nil {
		return err
	}
	SetUserID(ctx, uid)
	SetRole(ctx, token)
	return nil
}

func (j *JWTer) DeleteToken(ctx context.Context, r *http.Request, id user_model.UserID) error {
	token, err := j.GetToken(ctx, r)
	if err != nil {
		return err
	}

	if err := j.Store.Delete(ctx, token.JwtID()); err != nil {
		return err
	}
	return nil
}

func (j *JWTer) GetToken(ctx context.Context, r *http.Request) (jwt.Token, error) {
	token, err := jwt.ParseRequest(
		r,
		jwt.WithKey(jwa.RS256, j.PublicKey),
		jwt.WithValidate(false),
	)
	if err != nil {
		return nil, err
	}
	if err := jwt.Validate(token, jwt.WithClock(j.Clocker)); err != nil {
		return nil, fmt.Errorf("GetToken: failed to validate token: %w", err)
	}

	if _, err := j.Store.Load(ctx, token.JwtID()); err != nil {
		return nil, fmt.Errorf("GetToken: %q expired: %w", token.JwtID(), err)
	}
	return token, nil
}

func IsAdmin(ctx context.Context) bool {
	role, ok := GetRole(ctx)
	if !ok {
		return false
	}
	return role == "admin"
}

type userIDKey struct{}
type roleKey struct{}

func SetUserID(ctx *gin.Context, uid user_model.UserID) {
	ctx.Set(IDKey, uid)
}

func GetUserID(ctx context.Context) (user_model.UserID, bool) {
	id, ok := ctx.Value(IDKey).(user_model.UserID)
	return id, ok
}

func CheckOwn(ctx context.Context, userID user_model.UserID) bool {
	id, ok := ctx.Value(IDKey).(user_model.UserID)
	if !ok {
		return false
	}

	return id == userID
}

func SetRole(ctx *gin.Context, tok jwt.Token) {
	get, ok := tok.Get(RoleKey)
	if !ok {
		ctx.Set(RoleKey, "")
	}
	ctx.Set(RoleKey, get)
}

func GetRole(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(RoleKey).(string)
	return role, ok
}

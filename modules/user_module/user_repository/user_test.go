package user_repository

import (
	"context"
	"testing"

	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/repository"
	"github.com/tatuya-web/go-modular-monolith/testutil"
	"github.com/tatuya-web/go-modular-monolith/util"
	clock "github.com/tatuya-web/go-modular-monolith/util"
)

func TestRepository_RegisterUser(t *testing.T) {
	type fields struct {
		Clocker clock.Clocker
	}
	type args struct {
		ctx context.Context
		db  repository.Execer
		u   *user_model.User
	}
	ctx := context.Background()
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// このテストケースが完了したらもとに戻す
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ok user",
			fields: fields{
				Clocker: util.FixedClocker{},
			},
			args: args{
				ctx: ctx,
				db:  tx,
				u: &user_model.User{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserRepo{
				Clocker: tt.fields.Clocker,
			}
			if err := r.RegisterUser(tt.args.ctx, tt.args.db, tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("Repository.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package user_repository

import (
	"context"
	"testing"

	"github.com/tatuya-web/go-modular-monolith/modules/user_module/user_model"
	"github.com/tatuya-web/go-modular-monolith/testutil"
)

func TestKVS_Save(t *testing.T) {
	t.Parallel()

	cli := testutil.OpenRedisForTest(t)

	sut := &KVS{Cli: cli}
	key := "TestKVS_Save"
	uid := user_model.UserID(1234)
	ctx := context.Background()
	t.Cleanup(func() {
		cli.Del(ctx, key)
	})
	if err := sut.Save(ctx, key, uid); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}

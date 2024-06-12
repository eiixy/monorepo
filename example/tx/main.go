package main

import (
	"context"
	"github.com/eiixy/monorepo/internal/data/example/ent"
	"github.com/eiixy/monorepo/pkg/database/ent/enthelper"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/example?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	uc := userUseCase{
		ent:  client,
		user: userRepo{ent: client},
		role: roleRepo{ent: client},
	}
	ctx := context.Background()
	err = uc.UpdateUserName(ctx, 1, "user1")
	if err != nil {
		return
	}
	err = uc.UpdateUserAndRoles(ctx, 1, "test", []int{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}
}

type userUseCase struct {
	ent  *ent.Client
	user userRepo
	role roleRepo
}

// UpdateUserName 修改用户名 （不使用事务）
func (r userUseCase) UpdateUserName(ctx context.Context, id int, name string) error {
	return r.user.UpdateUserName(ctx, id, name)
}

// UpdateUserAndRoles 修改用户信息+关联角色 （使用事务）
func (r userUseCase) UpdateUserAndRoles(ctx context.Context, id int, name string, roleIDs []int) error {
	return enthelper.WithTx(ctx, r.ent, func(tx *ent.Tx) error {
		ctx = WithTx(ctx, tx)
		err := r.user.UpdateUserName(ctx, id, name)
		if err != nil {
			return err
		}
		err = r.role.UpdateUserRoles(ctx, id, roleIDs)
		if err != nil {
			return err
		}
		return nil
	})
}

type txKey struct{}

func WithTx(ctx context.Context, tx *ent.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func GetTxFromContext(ctx context.Context) *ent.Tx {
	v := ctx.Value(txKey{})
	if v == nil {
		return nil
	}
	return v.(*ent.Tx)
}

type userRepo struct {
	ent *ent.Client
}

// GetClient 获取 *ent.Client 优先从 context 中获取
func (r userRepo) GetClient(ctx context.Context) *ent.Client {
	if tx := GetTxFromContext(ctx); tx != nil {
		return tx.Client()
	}
	return r.ent
}

func (r userRepo) UpdateUserName(ctx context.Context, id int, name string) error {
	client := r.GetClient(ctx)
	return client.User.UpdateOneID(id).SetNickname(name).Exec(ctx)
}

type roleRepo struct {
	ent *ent.Client
}

func (r roleRepo) GetClient(ctx context.Context) *ent.Client {
	if tx := GetTxFromContext(ctx); tx != nil {
		return tx.Client()
	}
	return r.ent
}

func (r roleRepo) UpdateUserRoles(ctx context.Context, userID int, roleIDs []int) error {
	client := r.GetClient(ctx)
	return client.User.UpdateOneID(userID).AddRoleIDs(roleIDs...).Exec(ctx)
}

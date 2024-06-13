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
		user: userRepo{data{ent: client}},
		role: roleRepo{data{ent: client}},
	}
	ctx := context.Background()
	err = uc.UpdateUserName(ctx, 1, "user1")
	if err != nil {
		log.Fatal(err)
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
	return enthelper.WithTx(ctx, r.ent, func(ctx context.Context, tx *ent.Tx) error {
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

type data struct {
	ent *ent.Client
}

// Client 获取 *ent.Client 优先从 context 中获取
func (r data) Client(ctx context.Context) *ent.Client {
	if tx := enthelper.GetTxFromContext[*ent.Tx](ctx); tx != nil {
		return tx.Client()
	}
	return r.ent
}

type userRepo struct {
	data
}

func (r userRepo) UpdateUserName(ctx context.Context, id int, name string) error {
	return r.Client(ctx).User.UpdateOneID(id).SetNickname(name).Exec(ctx)
}

type roleRepo struct {
	data
}

func (r roleRepo) UpdateUserRoles(ctx context.Context, userID int, roleIDs []int) error {
	return r.Client(ctx).User.UpdateOneID(userID).AddRoleIDs(roleIDs...).Exec(ctx)
}

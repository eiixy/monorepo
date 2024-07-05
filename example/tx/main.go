package main

import (
	"context"
	"github.com/eiixy/monorepo/internal/app/admin/conf"
	"github.com/eiixy/monorepo/internal/app/admin/data"
	"github.com/eiixy/monorepo/internal/data/example/ent"
	"github.com/eiixy/monorepo/internal/pkg/config"
	"github.com/eiixy/monorepo/pkg/database/ent/enthelper"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var cfg conf.Config
	cfg.Data.Database = config.Database{
		Driver: "mysql",
		Dsn:    "root:12345678@tcp(127.0.0.1:3306)/example?parseTime=true",
	}
	client, err := data.NewEntClient(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	db, err := data.NewEntDatabase(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := repo{client, db}
	uc := userUseCase{
		ent:  client,
		db:   db,
		user: userRepo{r},
		role: roleRepo{r},
	}
	ctx := context.Background()
	err = uc.UpdateUserName(ctx, 1, "user1")
	if err != nil {
		log.Fatal(err)
	}
	// 使用 *ent.Database
	err = uc.UpdateUserAndRoles(ctx, 1, "test", []int{1, 2, 34})
	if err != nil {
		log.Fatal(err)
	}

	// 使用 *ent.Client
	//err = uc.UpdateUserAndRoles2(ctx, 1, "test", []int{1, 2, 34})
	//if err != nil {
	//	log.Fatal(err)
	//}
}

type userUseCase struct {
	ent  *ent.Client
	db   *ent.Database
	user userRepo
	role roleRepo
}

// UpdateUserName 修改用户名 （不使用事务）
func (r userUseCase) UpdateUserName(ctx context.Context, id int, name string) error {
	return r.user.UpdateUserName(ctx, id, name)
}

// UpdateUserAndRoles 修改用户信息+关联角色 （使用事务）
func (r userUseCase) UpdateUserAndRoles(ctx context.Context, id int, name string, roleIDs []int) error {
	return r.db.InTx(ctx, func(ctx context.Context) error {
		err := r.user.UpdateUserName(ctx, id, name)
		if err != nil {
			return err
		}
		return r.role.UpdateUserRoles(ctx, id, roleIDs)
	})
}

// UpdateUserAndRoles2 修改用户信息+关联角色 （使用事务，无需增加 ent template）
func (r userUseCase) UpdateUserAndRoles2(ctx context.Context, id int, name string, roleIDs []int) error {
	return enthelper.WithTx(ctx, r.ent, func(ctx context.Context, tx *ent.Tx) error {
		ctx = ent.NewTxContext(ctx, tx)
		err := r.user.UpdateUserName2(ctx, id, name)
		if err != nil {
			return err
		}
		return r.role.UpdateUserRoles2(ctx, id, roleIDs)
	})
}

type repo struct {
	ent *ent.Client
	db  *ent.Database
}

func (r repo) client(ctx context.Context) *ent.Client {
	if tx := ent.TxFromContext(ctx); tx != nil {
		return tx.Client()
	}
	return r.ent
}

type userRepo struct {
	repo
}

type roleRepo struct {
	repo
}

func (r userRepo) UpdateUserName(ctx context.Context, id int, name string) error {
	return r.db.User(ctx).UpdateOneID(id).SetNickname(name).Exec(ctx)
}

func (r roleRepo) UpdateUserRoles(ctx context.Context, userID int, roleIDs []int) error {
	return r.db.User(ctx).UpdateOneID(userID).AddRoleIDs(roleIDs...).Exec(ctx)
}

func (r userRepo) UpdateUserName2(ctx context.Context, id int, name string) error {
	return r.client(ctx).User.UpdateOneID(id).SetNickname(name).Exec(ctx)
}

func (r roleRepo) UpdateUserRoles2(ctx context.Context, userID int, roleIDs []int) error {
	return r.client(ctx).User.UpdateOneID(userID).AddRoleIDs(roleIDs...).Exec(ctx)
}

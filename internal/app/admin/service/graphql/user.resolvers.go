package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"errors"

	"github.com/eiixy/monorepo/internal/app/admin/server/auth"
	"github.com/eiixy/monorepo/internal/app/admin/service/graphql/dataloader"
	"github.com/eiixy/monorepo/internal/app/admin/service/graphql/model"
	"github.com/eiixy/monorepo/internal/data/example/ent"
	"github.com/eiixy/monorepo/internal/data/example/ent/permission"
	"github.com/eiixy/monorepo/internal/data/example/ent/role"
	"github.com/eiixy/monorepo/internal/data/example/ent/user"
	"golang.org/x/crypto/bcrypt"
)

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, oldPassword string, password string) (bool, error) {
	u, err := r.db.User(ctx).Get(ctx, auth.GetUserId(ctx))
	if err != nil {
		return false, err
	}
	if !r.accountUseCase.VerifyPassword(u, oldPassword) {
		return false, ErrAccountOrPasswordInvalid
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return false, errors.New("bcrypt: GenerateFromPassword Error")
	}
	err = u.Update().SetPassword(string(hashed)).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ForgetPassword is the resolver for the forgetPassword field.
func (r *mutationResolver) ForgetPassword(ctx context.Context, email string, code string, password string) (bool, error) {
	ok := r.accountUseCase.CheckEmailVerifyCode(email, model.VerifyCodeTypeForgetPassword, code)
	if !ok {
		return false, ErrVerifyCodeInvalid
	}
	first, err := r.db.User(ctx).Query().Where(user.Email(email)).First(ctx)
	if ent.IsNotFound(err) {
		return false, ErrAccountOrPasswordInvalid
	} else if err != nil {
		return false, err
	}
	err = first.Update().SetPassword(r.accountUseCase.GeneratePassword(password)).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateProfile is the resolver for the updateProfile field.
func (r *mutationResolver) UpdateProfile(ctx context.Context, input model.UpdateProfileInput) (*ent.User, error) {
	return r.db.User(ctx).UpdateOneID(auth.GetUserId(ctx)).
		SetNillableAvatar(input.Avatar).
		SetNillableNickname(input.Nickname).Save(ctx)
}

// Login is the resolver for the login field.
func (r *queryResolver) Login(ctx context.Context, email string, password string, captchaID *string, captchaValue *string) (*model.LoginReply, error) {
	if captchaID != nil && captchaValue != nil {
		matched := r.captcha.Verify(*captchaID, *captchaValue, true)
		if *captchaID == "" || *captchaValue == "" || !matched {
			return nil, errors.New("captcha verify failed")
		}
	}
	first, err := r.db.User(ctx).Query().Where(user.Email(email)).First(ctx)
	if ent.IsNotFound(err) {
		return nil, ErrAccountOrPasswordInvalid
	} else if err != nil {
		return nil, err
	}
	if !r.accountUseCase.VerifyPassword(first, password) {
		return nil, ErrAccountOrPasswordInvalid
	}
	token, exp, err := r.accountUseCase.GenerateToken(ctx, first.ID)
	if err != nil {
		return nil, err
	}
	return &model.LoginReply{
		Token: token,
		Exp:   int(exp),
		User:  first,
	}, nil
}

// Profile is the resolver for the profile field.
func (r *queryResolver) Profile(ctx context.Context) (*ent.User, error) {
	return r.db.User(ctx).Get(ctx, auth.GetUserId(ctx))
}

// Refresh is the resolver for the refresh field.
func (r *queryResolver) Refresh(ctx context.Context) (*model.LoginReply, error) {
	first, err := r.db.User(ctx).Get(ctx, auth.GetUserId(ctx))
	if err != nil {
		return nil, err
	}
	token, exp, err := r.accountUseCase.GenerateToken(ctx, first.ID)
	if err != nil {
		return nil, err
	}
	return &model.LoginReply{
		Token: token,
		Exp:   int(exp),
		User:  first,
	}, nil
}

// SendVerifyCode is the resolver for the sendVerifyCode field.
func (r *queryResolver) SendVerifyCode(ctx context.Context, email string, verifyType model.VerifyCodeType) (bool, error) {
	err := r.accountUseCase.SendEmail(email, verifyType)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Captcha is the resolver for the captcha field.
func (r *queryResolver) Captcha(ctx context.Context) (*model.CaptchaReply, error) {
	id, data, _, err := r.captcha.Generate()
	if err != nil {
		return nil, err
	}
	return &model.CaptchaReply{
		ID:      id,
		Captcha: data,
	}, nil
}

// RoleCount is the resolver for the roleCount field.
func (r *userResolver) RoleCount(ctx context.Context, obj *ent.User) (int, error) {
	return dataloader.For(ctx).GetUserRoleCount(ctx, obj.ID)
}

// Permissions is the resolver for the permissions field.
func (r *userResolver) Permissions(ctx context.Context, obj *ent.User) ([]*ent.Permission, error) {
	if obj.IsAdmin {
		return r.db.Permission(ctx).Query().Where().All(ctx)
	}
	return r.db.Permission(ctx).Query().Where(permission.HasRolesWith(role.HasUsersWith(user.ID(obj.ID)))).All(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

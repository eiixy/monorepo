// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/eiixy/monorepo/internal/data/admin/ent"
)

type LoginReply struct {
	Token string `json:"token"`
	// 过期时间（秒）
	Exp  int       `json:"exp"`
	User *ent.User `json:"user"`
}

type UpdateProfileInput struct {
	Nickname *string `json:"nickname,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
	Mobile   *string `json:"mobile,omitempty"`
}

type VerifyCodeType string

const (
	// 忘记密码
	VerifyCodeTypeForgetPassword VerifyCodeType = "ForgetPassword"
	// 注册验证
	VerifyCodeTypeRegister VerifyCodeType = "Register"
)

var AllVerifyCodeType = []VerifyCodeType{
	VerifyCodeTypeForgetPassword,
	VerifyCodeTypeRegister,
}

func (e VerifyCodeType) IsValid() bool {
	switch e {
	case VerifyCodeTypeForgetPassword, VerifyCodeTypeRegister:
		return true
	}
	return false
}

func (e VerifyCodeType) String() string {
	return string(e)
}

func (e *VerifyCodeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VerifyCodeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid verifyCodeType", str)
	}
	return nil
}

func (e VerifyCodeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
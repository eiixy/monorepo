package graphql

import "github.com/vektah/gqlparser/v2/gqlerror"

const CustomErrorKey = "custom"

var ErrDisabled = CustomError("Disabled", "disabled")

// ErrDeprecated 方法已弃用
var ErrDeprecated = CustomError("Deprecated", "deprecated")

// ErrNoPermission 无权限
var ErrNoPermission = CustomError("NoPermission", "no permission")

// ErrAccessDenied 拒绝访问，无角色或身份
var ErrAccessDenied = CustomError("AccessDenied", "access denied")

// ErrUnauthorized 未授权
var ErrUnauthorized = CustomError("Unauthorized", "Unauthorized")

// ErrAccountOrPasswordInvalid 账号或密码无效
var ErrAccountOrPasswordInvalid = CustomError("AccountOrPasswordInvalid", "account or password invalid")

// ErrVerifyCodeInvalid 验证码无效
var ErrVerifyCodeInvalid = CustomError("VerifyCodeInvalid", "verify code invalid")

func CustomError(code, message string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: message,
		Extensions: map[string]any{
			"code": code,
		},
		Rule: CustomErrorKey,
	}
}

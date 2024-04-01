package graphql

import "github.com/vektah/gqlparser/v2/gqlerror"

var ErrDisabled = &gqlerror.Error{
	Message: "disabled",
	Extensions: map[string]interface{}{
		"code": "Disabled",
	},
}

// ErrDeprecated 方法已弃用
var ErrDeprecated = &gqlerror.Error{
	Message: "deprecated",
	Extensions: map[string]interface{}{
		"code": "Deprecated",
	},
}

// ErrNoPermission 无权限
var ErrNoPermission = &gqlerror.Error{
	Message: "no permission",
	Extensions: map[string]interface{}{
		"code": "NoPermission",
	},
}

// ErrAccessDenied 拒绝访问，无角色或身份
var ErrAccessDenied = &gqlerror.Error{
	Message: "access denied",
	Extensions: map[string]interface{}{
		"code": "AccessDenied",
	},
}

// ErrUnauthorized 未授权
var ErrUnauthorized = &gqlerror.Error{
	Message: "Unauthorized",
	Extensions: map[string]interface{}{
		"code": "Unauthorized",
	},
}

// ErrAccountOrPasswordInvalid 账号或密码无效
var ErrAccountOrPasswordInvalid = &gqlerror.Error{
	Message: "account or password invalid",
	Extensions: map[string]interface{}{
		"code": "AccountOrPasswordInvalid",
	},
}

// ErrVerifyCodeInvalid 验证码无效
var ErrVerifyCodeInvalid = &gqlerror.Error{
	Message: "verify code invalid",
	Extensions: map[string]interface{}{
		"code": "VerifyCodeInvalid",
	},
}

// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/eiixy/monorepo/internal/data/admin/ent/operationlog"
	"github.com/eiixy/monorepo/internal/data/admin/ent/permission"
	"github.com/eiixy/monorepo/internal/data/admin/ent/role"
	"github.com/eiixy/monorepo/internal/data/admin/ent/schema"
	"github.com/eiixy/monorepo/internal/data/admin/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	operationlogMixin := schema.OperationLog{}.Mixin()
	operationlogMixinFields0 := operationlogMixin[0].Fields()
	_ = operationlogMixinFields0
	operationlogFields := schema.OperationLog{}.Fields()
	_ = operationlogFields
	// operationlogDescCreatedAt is the schema descriptor for created_at field.
	operationlogDescCreatedAt := operationlogMixinFields0[0].Descriptor()
	// operationlog.DefaultCreatedAt holds the default value on creation for the created_at field.
	operationlog.DefaultCreatedAt = operationlogDescCreatedAt.Default.(func() time.Time)
	// operationlogDescUpdatedAt is the schema descriptor for updated_at field.
	operationlogDescUpdatedAt := operationlogMixinFields0[1].Descriptor()
	// operationlog.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	operationlog.DefaultUpdatedAt = operationlogDescUpdatedAt.Default.(func() time.Time)
	// operationlog.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	operationlog.UpdateDefaultUpdatedAt = operationlogDescUpdatedAt.UpdateDefault.(func() time.Time)
	permissionMixin := schema.Permission{}.Mixin()
	permissionMixinFields0 := permissionMixin[0].Fields()
	_ = permissionMixinFields0
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescCreatedAt is the schema descriptor for created_at field.
	permissionDescCreatedAt := permissionMixinFields0[0].Descriptor()
	// permission.DefaultCreatedAt holds the default value on creation for the created_at field.
	permission.DefaultCreatedAt = permissionDescCreatedAt.Default.(func() time.Time)
	// permissionDescUpdatedAt is the schema descriptor for updated_at field.
	permissionDescUpdatedAt := permissionMixinFields0[1].Descriptor()
	// permission.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	permission.DefaultUpdatedAt = permissionDescUpdatedAt.Default.(func() time.Time)
	// permission.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	permission.UpdateDefaultUpdatedAt = permissionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// permissionDescSort is the schema descriptor for sort field.
	permissionDescSort := permissionFields[6].Descriptor()
	// permission.DefaultSort holds the default value on creation for the sort field.
	permission.DefaultSort = permissionDescSort.Default.(int)
	roleMixin := schema.Role{}.Mixin()
	roleMixinFields0 := roleMixin[0].Fields()
	_ = roleMixinFields0
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleMixinFields0[0].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleMixinFields0[1].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// role.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	role.UpdateDefaultUpdatedAt = roleDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roleDescSort is the schema descriptor for sort field.
	roleDescSort := roleFields[1].Descriptor()
	// role.DefaultSort holds the default value on creation for the sort field.
	role.DefaultSort = roleDescSort.Default.(int)
	userMixin := schema.User{}.Mixin()
	userMixinHooks1 := userMixin[1].Hooks()
	user.Hooks[0] = userMixinHooks1[0]
	userMixinInters1 := userMixin[1].Interceptors()
	user.Interceptors[0] = userMixinInters1[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescIsAdmin is the schema descriptor for is_admin field.
	userDescIsAdmin := userFields[5].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the is_admin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)

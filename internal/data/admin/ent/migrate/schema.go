// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MenusColumns holds the columns for the "menus" table.
	MenusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "icon", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "path", Type: field.TypeString},
		{Name: "sort", Type: field.TypeInt, Default: 1000},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true},
	}
	// MenusTable holds the schema information for the "menus" table.
	MenusTable = &schema.Table{
		Name:       "menus",
		Comment:    "菜单",
		Columns:    MenusColumns,
		PrimaryKey: []*schema.Column{MenusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "menus_menus_children",
				Columns:    []*schema.Column{MenusColumns[7]},
				RefColumns: []*schema.Column{MenusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "menu_created_at",
				Unique:  false,
				Columns: []*schema.Column{MenusColumns[1]},
			},
		},
	}
	// OperationLogsColumns holds the columns for the "operation_logs" table.
	OperationLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "type", Type: field.TypeString, Comment: "操作类型"},
		{Name: "context", Type: field.TypeJSON},
		{Name: "user_id", Type: field.TypeInt},
	}
	// OperationLogsTable holds the schema information for the "operation_logs" table.
	OperationLogsTable = &schema.Table{
		Name:       "operation_logs",
		Comment:    "操作日志",
		Columns:    OperationLogsColumns,
		PrimaryKey: []*schema.Column{OperationLogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "operation_logs_users_operation_logs",
				Columns:    []*schema.Column{OperationLogsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "operationlog_created_at",
				Unique:  false,
				Columns: []*schema.Column{OperationLogsColumns[1]},
			},
		},
	}
	// PermissionsColumns holds the columns for the "permissions" table.
	PermissionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString, Nullable: true},
		{Name: "sort", Type: field.TypeInt, Default: 1000},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true},
	}
	// PermissionsTable holds the schema information for the "permissions" table.
	PermissionsTable = &schema.Table{
		Name:       "permissions",
		Comment:    "操作日志",
		Columns:    PermissionsColumns,
		PrimaryKey: []*schema.Column{PermissionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "permissions_permissions_children",
				Columns:    []*schema.Column{PermissionsColumns[7]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "permission_created_at",
				Unique:  false,
				Columns: []*schema.Column{PermissionsColumns[1]},
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "sort", Type: field.TypeInt, Default: 1000},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Comment:    "角色",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "role_created_at",
				Unique:  false,
				Columns: []*schema.Column{RolesColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "nickname", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString},
		{Name: "status", Type: field.TypeEnum, Comment: "状态", Enums: []string{"normal", "freeze"}},
		{Name: "is_admin", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Comment:    "用户",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_created_at",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[1]},
			},
			{
				Name:    "user_deleted_at",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[3]},
			},
		},
	}
	// MenuPermissionsColumns holds the columns for the "menu_permissions" table.
	MenuPermissionsColumns = []*schema.Column{
		{Name: "menu_id", Type: field.TypeInt},
		{Name: "permission_id", Type: field.TypeInt},
	}
	// MenuPermissionsTable holds the schema information for the "menu_permissions" table.
	MenuPermissionsTable = &schema.Table{
		Name:       "menu_permissions",
		Columns:    MenuPermissionsColumns,
		PrimaryKey: []*schema.Column{MenuPermissionsColumns[0], MenuPermissionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "menu_permissions_menu_id",
				Columns:    []*schema.Column{MenuPermissionsColumns[0]},
				RefColumns: []*schema.Column{MenusColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "menu_permissions_permission_id",
				Columns:    []*schema.Column{MenuPermissionsColumns[1]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RoleMenusColumns holds the columns for the "role_menus" table.
	RoleMenusColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt},
		{Name: "menu_id", Type: field.TypeInt},
	}
	// RoleMenusTable holds the schema information for the "role_menus" table.
	RoleMenusTable = &schema.Table{
		Name:       "role_menus",
		Columns:    RoleMenusColumns,
		PrimaryKey: []*schema.Column{RoleMenusColumns[0], RoleMenusColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_menus_role_id",
				Columns:    []*schema.Column{RoleMenusColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_menus_menu_id",
				Columns:    []*schema.Column{RoleMenusColumns[1]},
				RefColumns: []*schema.Column{MenusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RolePermissionsColumns holds the columns for the "role_permissions" table.
	RolePermissionsColumns = []*schema.Column{
		{Name: "role_id", Type: field.TypeInt},
		{Name: "permission_id", Type: field.TypeInt},
	}
	// RolePermissionsTable holds the schema information for the "role_permissions" table.
	RolePermissionsTable = &schema.Table{
		Name:       "role_permissions",
		Columns:    RolePermissionsColumns,
		PrimaryKey: []*schema.Column{RolePermissionsColumns[0], RolePermissionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_permissions_role_id",
				Columns:    []*schema.Column{RolePermissionsColumns[0]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "role_permissions_permission_id",
				Columns:    []*schema.Column{RolePermissionsColumns[1]},
				RefColumns: []*schema.Column{PermissionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0], UserRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_user_id",
				Columns:    []*schema.Column{UserRolesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_roles_role_id",
				Columns:    []*schema.Column{UserRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MenusTable,
		OperationLogsTable,
		PermissionsTable,
		RolesTable,
		UsersTable,
		MenuPermissionsTable,
		RoleMenusTable,
		RolePermissionsTable,
		UserRolesTable,
	}
)

func init() {
	MenusTable.ForeignKeys[0].RefTable = MenusTable
	MenusTable.Annotation = &entsql.Annotation{}
	OperationLogsTable.ForeignKeys[0].RefTable = UsersTable
	OperationLogsTable.Annotation = &entsql.Annotation{}
	PermissionsTable.ForeignKeys[0].RefTable = PermissionsTable
	PermissionsTable.Annotation = &entsql.Annotation{}
	RolesTable.Annotation = &entsql.Annotation{}
	UsersTable.Annotation = &entsql.Annotation{}
	MenuPermissionsTable.ForeignKeys[0].RefTable = MenusTable
	MenuPermissionsTable.ForeignKeys[1].RefTable = PermissionsTable
	RoleMenusTable.ForeignKeys[0].RefTable = RolesTable
	RoleMenusTable.ForeignKeys[1].RefTable = MenusTable
	RolePermissionsTable.ForeignKeys[0].RefTable = RolesTable
	RolePermissionsTable.ForeignKeys[1].RefTable = PermissionsTable
	UserRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.ForeignKeys[1].RefTable = RolesTable
}

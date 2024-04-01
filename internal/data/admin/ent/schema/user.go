package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/eiixy/monorepo/internal/data/admin/mixin"
	"github.com/eiixy/monorepo/internal/data/annotations"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户"),
		entsql.WithComments(true),
		entgql.QueryField().Directives().Description("用户"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().Annotations(entgql.OrderField("EMAIL")),
		field.String("nickname"),
		field.String("avatar").Optional(),
		field.String("password").Sensitive(),
		field.Enum("status").Values("normal", "freeze").Comment("状态").Annotations(annotations.Enums{
			"normal": "正常",
			"freeze": "冻结",
		}),
		field.Bool("is_admin").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
		edge.To("operation_logs", OperationLog.Type),
	}
}

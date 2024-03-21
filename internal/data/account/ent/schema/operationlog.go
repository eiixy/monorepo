package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/eiixy/monorepo/internal/data/account/mixin"
)

// OperationLog holds the schema definition for the OperationLog entity.
type OperationLog struct {
	ent.Schema
}

func (OperationLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the OperationLog.
func (OperationLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("type").Comment("操作类型"),
		field.JSON("context", map[string]any{}),
	}
}

// Edges of the OperationLog.
func (OperationLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("operation_logs").Unique().Required().Field("user_id"),
	}
}

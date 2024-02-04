package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Chat holds the schema definition for the Chat entity.
type Chat struct {
	ent.Schema
}

// Annotations of the Chat.
func (Chat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "chat"},
	}
}

// Fields of the Chat.
func (Chat) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").Positive(),
		field.Int32("model_id"),
		field.String("session_id").Unique(),
		field.String("name"),
	}
}

// Edges of the Chat.
func (Chat) Edges() []ent.Edge {
	return nil
}

func (Chat) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

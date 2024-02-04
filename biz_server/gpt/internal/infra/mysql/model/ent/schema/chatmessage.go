package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ChatMessage holds the schema definition for the ChatMessage entity.
type ChatMessage struct {
	ent.Schema
}

// Annotations of the ChatMessage.
func (ChatMessage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "chat_message"},
	}
}

// Fields of the ChatMessage.
func (ChatMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").Positive(),
		field.Int32("chat_id"),
		field.String("request_id"),
		field.String("text"),
		field.Int8("version"),
	}
}

// Edges of the ChatMessage.
func (ChatMessage) Edges() []ent.Edge {
	return nil
}

func (ChatMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

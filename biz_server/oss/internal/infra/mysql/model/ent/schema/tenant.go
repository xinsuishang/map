package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Annotations of the User.
func (Tenant) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "tenants"},
	}
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("region"),
		field.String("type"),
		field.String("access_key"),
		field.String("secret_key"),
		field.String("desc"),
	}
}

func (Tenant) Indexes() []ent.Index {
	return []ent.Index{}
}

// Edges of the Tenant.
func (Tenant) Edges() []ent.Edge {
	return nil
}

func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

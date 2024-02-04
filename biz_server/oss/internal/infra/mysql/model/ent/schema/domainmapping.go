package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// DomainMapping holds the schema definition for the DomainMapping entity.
type DomainMapping struct {
	ent.Schema
}

// Annotations of the DomainMapping.
func (DomainMapping) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "domain_mapping"},
	}
}

// Fields of the DomainMapping.
func (DomainMapping) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id").Positive(),
		field.Int32("tenant_id").Unique(),
		field.String("region_id"),
		field.String("domain"),
		field.String("bucket_name"),
		field.String("desc"),
	}
}

func (DomainMapping) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

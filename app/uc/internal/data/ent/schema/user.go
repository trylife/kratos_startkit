package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect/entsql"
    "entgo.io/ent/schema"
    "entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
    return []schema.Annotation{
        entsql.Annotation{Table: "user"},
    }
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id"),
        field.String("uuid"),
        field.String("nickname"),
        field.String("phone_number"),
        field.Int("create_time"),
        field.Int("update_time"),
    }
}

// Edges of the User.
func (User) Edges() []ent.Edge {
    return nil
}

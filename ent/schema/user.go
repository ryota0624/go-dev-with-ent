package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New).
			StorageKey("id").Annotations(entproto.Field(1)),
		field.Int("age").
			Positive().Annotations(entproto.Field(2)).Annotations(entgql.OrderField("AGE")),
		field.String("name").
			Default("unknown").Annotations(entproto.Field(3)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type).
			Annotations(entproto.Skip()).
			Annotations(entgql.Bind(), entgql.RelayConnection()),
		edge.From("groups", Group.Type).
			Ref("users").Annotations(entproto.Skip()).
			Annotations(entgql.Bind()),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExpose)),
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExpose)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExpose)),
		entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExpose)),
		entproto.Message(),
		entproto.Service(entproto.Methods(entproto.MethodGet)),
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(),
	}
}

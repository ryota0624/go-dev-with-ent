package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New).
			StorageKey("id").Annotations(entproto.Field(1)),
		field.String("model").Annotations(entproto.Field(2)),
		field.Time("registered_at").Annotations(entproto.Field(3), entgql.OrderField("REGISTERED_AT")),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("owner", User.Type).
			Ref("cars").
			Required().
			Annotations(entgql.Bind()).
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique().Annotations(entproto.Field(4)),
	}
}

func (Car) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entgql.RelayConnection(),
	}
}

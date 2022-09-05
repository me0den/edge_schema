package schema

import (
	"entgo.io/ent/schema/edge"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("user_id"),
		field.String("name"),
		field.Int64("created_at").Default(time.Now().Unix()),
		field.Int64("updated_at").
			Default(time.Now().Unix()).
			UpdateDefault(time.Now().Unix),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", Item.Type).
			Through("user_items", UserItem.Type),
	}
}
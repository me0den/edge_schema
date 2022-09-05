package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Item struct {
	ent.Schema
}

func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.String("name"),
		field.String("description"),
		field.Int64("created_at").Default(time.Now().Unix()),
		field.Int64("updated_at").
			Default(time.Now().Unix()).
			UpdateDefault(time.Now().Unix),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("items").
			Through("item_user", UserItem.Type),
	}
}
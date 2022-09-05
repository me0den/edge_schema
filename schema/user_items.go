package schema

import (
	"entgo.io/ent/schema"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type UserItem struct {
	ent.Schema
}

func (UserItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id"),
		field.String("item_id"),
		field.Int("quantity"),
		field.Int64("created_at").Default(time.Now().Unix()),
		field.Int64("updated_at").
			Default(time.Now().Unix()).
			UpdateDefault(time.Now().Unix),
	}
}

func (UserItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("item", Item.Type).
			Unique().
			Required().
			Field("item_id"),
	}
}

func (UserItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "item_id").
			Unique(),
	}
}

func (UserItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "item_id"),
	}
}

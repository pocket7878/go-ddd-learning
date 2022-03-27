package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Todo.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().NotEmpty(),
		field.String("name").NotEmpty(),
		field.Enum("task_status").Values("undone", "done").Default("undone"),
		field.Int("postpone_count").Default(0),
		field.Time("due_date").Default(time.Now()),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}

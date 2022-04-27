package schema

import (
	"net/url"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
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
		field.Int("age").
			Positive(),
		field.String("name").NotEmpty(),
		field.Bool("active").Default(true),
		field.JSON("url", &url.URL{}).
			Optional(),
		field.JSON("tags", []string{}).
			Optional(),
		field.Enum("state").
			Values("on", "off").
			Optional(),
		field.UUID("uuid", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).Default(time.Now),
		field.Time("updated_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

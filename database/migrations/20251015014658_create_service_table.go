package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251015014658CreateServiceTable struct{}

// Signature The unique signature for the migration.
func (r *M20251015014658CreateServiceTable) Signature() string {
	return "20251015014658_create_service_table"
}

// Up Run the migrations.
func (r *M20251015014658CreateServiceTable) Up() error {
	if !facades.Schema().HasTable("service") {
		return facades.Schema().Create("service", func(table schema.Blueprint) {
			table.ID()
			table.String("name", 120)
			table.String("description", 255).Nullable()
			table.Float("price", 10, 2).Default(0)
			table.String("image", 255).Nullable()
			table.Enum("status", []any{"active", "inactive"}).Default("active")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20251015014658CreateServiceTable) Down() error {
	return facades.Schema().DropIfExists("service")
}

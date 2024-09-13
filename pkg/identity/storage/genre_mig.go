package storage

import "backend/pkg/infra/storage/migrator"

func addGenreMigration(mg *migrator.Migrator) {
	genreTable := migrator.Table{
		Name: "genre",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "name", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "description", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "created_by", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "created_at", Type: migrator.DB_DateTime, Nullable: false},
			{Name: "updated_by", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "updated_at", Type: migrator.DB_DateTime, Nullable: true},
		},
		Indices: []*migrator.Index{
			{Cols: []string{"name"}, Type: migrator.UniqueIndex},
		},
	}

	mg.AddMigration("create genre table", migrator.NewAddTableMigration(genreTable))
	mg.AddMigration("add index genre.name", migrator.NewAddIndexMigration(genreTable, genreTable.Indices[0]))
}

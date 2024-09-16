package storage

import "backend/pkg/infra/storage/migrator"

func addMovieMigration(mg *migrator.Migrator) {
	movieTable := migrator.Table{
		Name: "movie",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "title", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "description", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "poster_image", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "genre", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "duration", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "director", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "release_date", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "poster_url", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "created_by", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "created_at", Type: migrator.DB_DateTime, Nullable: false},
			{Name: "updated_by", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "updated_at", Type: migrator.DB_DateTime, Nullable: true},
		},
		Indices: []*migrator.Index{
			{Cols: []string{"title"}, Type: migrator.UniqueIndex},
		},
	}

	mg.AddMigration("create movie table", migrator.NewAddTableMigration(movieTable))
	mg.AddMigration("add index movie.title", migrator.NewAddIndexMigration(movieTable, movieTable.Indices[0]))
}

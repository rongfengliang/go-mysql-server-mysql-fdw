package main

import (
	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/auth"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
)

func main() {
	driver := sqle.NewDefault()
	driver.AddDatabase(createTestDatabase())

	config := server.Config{
		Protocol: "tcp",
		Address:  "0.0.0.0:3306",
		Auth:     auth.NewNativeSingle("root", "dalong", auth.AllPermissions),
	}

	s, err := server.NewDefaultServer(config, driver)
	if err != nil {
		panic(err)
	}

	s.Start()
}

func createTestDatabase() *memory.Database {
	const (
		dbName    = "test"
		tableName = "mytable"
	)

	db := memory.NewDatabase(dbName)
	table := memory.NewTable(tableName, sql.Schema{
		{Name: "name", Type: sql.Text, Nullable: false, Source: tableName},
		{Name: "email", Type: sql.Text, Nullable: false, Source: tableName},
	})

	db.AddTable(tableName, table)
	ctx := sql.NewEmptyContext()

	rows := []sql.Row{
		sql.NewRow("John Doe", "john@doe.com"),
		sql.NewRow("John Doe", "johnalt@doe.com"),
		sql.NewRow("Jane Doe", "jane@doe.com"),
		sql.NewRow("Evil Bob", "evilbob@gmail.com"),
	}

	for _, row := range rows {
		table.Insert(ctx, row)
	}

	return db
}

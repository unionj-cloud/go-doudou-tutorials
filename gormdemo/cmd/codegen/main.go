package main

import (
	"github.com/unionj-cloud/go-doudou/v2/framework/database"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:        "query",
		FieldNullable:  false,
		FieldCoverable: true,
		FieldSignable:  true,
		Mode:           gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(database.Db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

package main

import (
	"elsenova/db"
	"elsenova/models"
	"gorm.io/gen"
)

type Querier interface {
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db.Connection) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(models.AllModels...)

	g.ApplyInterface(func(Querier) {}, models.Vore{})

	// Generate the code
	g.Execute()
}

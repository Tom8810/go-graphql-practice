package main

import (
	"fmt"
	database "gorm_practice1/internal/db"
	getenv "gorm_practice1/internal/env"
	"log"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/gen",
		Mode: gen.WithDefaultQuery | gen.WithoutContext,
		FieldCoverable: true,
	})
	if err := getenv.Get("./internal/env/.env"); err != nil {
		fmt.Printf("cannnot read environmental variables")
	}
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	g.UseDB(database.DB)

	author := g.GenerateModel("authors")
	book := g.GenerateModel("books")

	g.ApplyBasic(
		g.GenerateModel("authors", gen.FieldRelate(field.HasMany, "Books", book, &field.RelateConfig{
			RelateSlicePointer: true,
		})),
		g.GenerateModel("books", gen.FieldRelate(field.BelongsTo, "Author", author, &field.RelateConfig{
			RelatePointer: true,
		})),
	)

	g.Execute()
}
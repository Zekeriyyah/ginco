package main

import (
	"github.com/zekeriyyah/ginco/internal/database"
	"github.com/zekeriyyah/ginco/migrations"
)

func main() {
	database.InitDB()
	database.TestSetup()

	migrations.Run()
	migrations.TestRun()
}

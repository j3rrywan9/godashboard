package controllers

import (
	"github.com/j3rrywan9/godashboard/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
)

const connectionString = "user=postgres dbname=postgres password=postgres host=172.16.100.9 port=5432 sslmode=disable"

var db *gorm.DB

func initDB() *gorm.DB {
	dbConn, err := gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	return dbConn
}

func init() {
	db = initDB()
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Database() revel.Result {
	var myDatabase []models.Mad_database
	myDatabase = models.Get_all_databases(*db)
	return c.Render(myDatabase)
}

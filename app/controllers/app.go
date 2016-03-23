package controllers

import (
	"github.com/j3rrywan9/godashboard/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
)

const connectionString = "user=postgres dbname=postgres password=postgres host=10.43.15.120 port=5432 sslmode=disable"

var db *gorm.DB

func initDB() *gorm.DB {
	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	return db
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
	myDatabase = models.Get_all_databases(db)
	return c.Render(myDatabase)
}

func (c App) BbRun() revel.Result {
	var myBbRun []models.Mad_bbrun
	myBbRun = models.Get_all_bbruns(db)
	return c.Render(myBbRun)
}

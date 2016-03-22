package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Mad_database struct {
	Id int
	Vendor string `sql:"size:50"`
	Version string `sql:"size:50"`
}

func Get_all_databases(db gorm.DB) []Mad_database {
	var records []Mad_database
	db.Table("mad_database").Select("*").Order("id").Scan(&records)
	return records
}

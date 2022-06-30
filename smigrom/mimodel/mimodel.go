package mimodel

import (
	"fmt"

	"gorm.io/gorm"
)

type MiTable struct {
	Struct interface{}
	DDL    string `manguDB:"CheckPool:'DROP_TABLE|DROP_COLUMN|ALTER_TABLE|ALTER_COLUMN_SIZE|ADD_CONSTRINT|DROP_CONSTRINT|ADD_INDEX_|DROP_INDEX_|_RENAME_INDEX_'"`
	DML    string
}

type Database struct {
	DB *gorm.DB
}

type IDatabase interface {
	MiGenerate(meta map[int64]MiTable) error
}

func (DB Database) MiGenerate(meta map[int64]MiTable) {
	fmt.Println(meta)
	for version, data := range meta {
		fmt.Println("🚀 ~ file: mimodel.go ~ line 26 ~ func ~ version : ", version)
		DB.DB.AutoMigrate(&data.Struct)
	}
}

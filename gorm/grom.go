package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Age       int
	Languages []*Language `gorm:"-"`
}

type Language struct {
	gorm.Model
	Name string
}

func main() {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	db.AutoMigrate(&User{}, &Language{})

	u := &User{
		Name: "zhangsan",
		Age:  18,
		Languages: []*Language{
			&Language{
				Name: "English",
			},
			&Language{
				Name: "Chinese",
			},
		},
	}

	if err := db.Save(u).Error; err != nil {
		panic(err)
	}

	u1 := &User{
		Name: "zhangsan",
	}

	if err := db.First(&u1).Error; err != nil {
		panic(err)
	}

	out, _ := json.MarshalIndent(u1, "", "  ")
	fmt.Println(string(out))
}

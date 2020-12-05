package main

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/rs/xid"
)

type TODOs struct {
	ID   string `xorm:"id"`
	Name string `xorm:"name"`
	TODO string `xorm:"todo"`
}

var (
	db *xorm.Engine
)

func main() {
	db = dbConfig()
	e := echo.New()
	e.Static("/", ".")
	e.GET("/todos/get", getTODO)
	e.POST("todos/post", createTODO)

	e.Start(":80")
}

func getTODO(c echo.Context) error {
	var todo []TODOs
	if err := db.Table("todos").Find(&todo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, todo)
}

func createTODO(c echo.Context) error {
	var todo TODOs
	c.Bind(&todo)
	todo.ID = xid.New().String()
	_, err := db.Table("todos").Insert(&todo)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, "ok")
}

func dbConfig() *xorm.Engine {
	db, err := xorm.NewEngine("mysql", "test:test@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

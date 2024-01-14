package main

import (
	"goweb/router"
	"goweb/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()

	r := router.Router()
	r.LoadHTMLFiles("templates/index.html", "templates/table.html", "templates/analyze.html", "templates/coffee.html", "templates/rule.html")
	r.Static("assets/bootstrap/css", "./templates/assets/bootstrap/css")
	r.Static("assets/fonts", "./templates/assets/fonts")
	r.Static("assets/css", "./templates/assets/css")
	r.Static("assets/js", "./templates/assets/js")
	r.Static("assets/bootstrap/js", "./templates/assets/bootstrap/js")
	r.Static("assets/img", "./templates/assets/img")
	r.Run(":8080")
}

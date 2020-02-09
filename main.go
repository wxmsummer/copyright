package main

import (
	_ "copyright/routers"
	"copyright/utils"
	"github.com/astaxie/beego"
)

func main() {

	utils.InitMysql()
	beego.Run()

}

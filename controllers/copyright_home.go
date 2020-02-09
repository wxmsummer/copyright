package controllers

import (
	"copyright/models"
	"fmt"
)

type CopyrightShowController struct {
	BaseController
}

func (this *CopyrightShowController) Get() {
	fmt.Println("IsLogin", this.Islogin, this.Loginuser)

	//tag := this.GetString("tag")
	//fmt.Println("tag", tag)
	page, _ := this.GetInt("page")
	var copyrightList []models.Copyright

	if page <= 0 {
		page = 1
	}
	copyrightList, _ = models.FindCopyrightWithPage(page)
	this.Data["pageCode"] = models.CopyrightConfigHomeFooterPageCode(page)
	this.Data["hasFooter"] = true

	fmt.Println("IsLogin:", this.Islogin, this.Loginuser)
	this.Data["Content"] = models.CopyrightMakeHomeBlocks(copyrightList, this.Islogin)

	this.TplName = "copyright_home.html"
}

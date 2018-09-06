package controllers

import (
	"github.com/astaxie/beego"
	"go_examples/RentHouse/models"
)

type HouseIndexController struct {
	beego.Controller
}

func (this *HouseIndexController) RetData(resp map[string]interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *HouseIndexController) GetHouseIndex() {
	beego.Info("============= /api/v1.0/house/index get success ==========")

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)
}

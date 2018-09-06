package controllers

import (
	"github.com/astaxie/beego"
	"go_examples/RentHouse/models"
)

type SessionController struct {
	beego.Controller
}

func (this *SessionController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *SessionController) GetSessionData() {
	beego.Info("======== /api/v1.0/session get success ========")

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_NODATA
	resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)

	defer this.RetData(resp)

	name_map := make(map[string]interface{})
	name := this.GetSession("name")

	if name != nil {
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		name_map["name"] = name.(string)
		resp["data"] = name_map
	}
	return
}

func (this *SessionController) DelSessionData() {
	beego.Info("======== /api/v1.0/session del success ========")

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	this.DelSession("name")
	this.DelSession("user_id")
	this.DelSession("mobile")

	return
}

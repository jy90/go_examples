package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_examples/RentHouse/models"
	"strconv"
)

type HouseController struct {
	beego.Controller
}

func (this *HouseController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *HouseController) GetHouseInfo() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)

	user_id := this.GetSession("user_id")
	houses := []models.House{}
	o := orm.NewOrm()
	qs := o.QueryTable("house")

	num, err := qs.Filter("user__id", user_id.(int)).All(&houses)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	if num == 0 {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}

	respData := make(map[string]interface{})
	respData["houses"] = houses
	resp["data"] = respData
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	return
}

func (this *HouseController) PublishHouseInfo() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	beego.Info("=========== post /api/v1.0/houses ==========")

	reqData := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &reqData)

	house := models.House{}
	house.Title = reqData["title"].(string)
	house.Price, _ = strconv.Atoi(reqData["price"].(string))
	house.Addess = reqData["address"].(string)
	house.Room_count, _ = strconv.Atoi(reqData["room_count"].(string))
	house.Acreage, _ = strconv.Atoi(reqData["acreage"].(string))
	house.Unit = reqData["unit"].(string)
	house.Capacity, _ = strconv.Atoi(reqData["capacity"].(string))
	house.Beds = reqData["beds"].(string)
	house.Deposit, _ = strconv.Atoi(reqData["deposit"].(string))
	house.Min_days, _ = strconv.Atoi(reqData["min_days"].(string))
	house.Max_days, _ = strconv.Atoi(reqData["max_days"].(string))

	facilities := []models.Facility{}

	for _, i := range reqData["facility"].([]interface{}) {
		f_id, _ := strconv.Atoi(i.(string))
		fac := models.Facility{Id: f_id}
		facilities = append(facilities, fac)
	}

	area_id, _ := strconv.Atoi(reqData["area_id"].(string))
	area := models.Area{Id: area_id}
	house.Area = &area
	user_id := this.GetSession("user_id")
	user := models.User{Id: user_id.(int)}
	house.User = &user
	beego.Info("house info --------- ", house)

	o := orm.NewOrm()
	_, err := o.Insert(&house)
	if err != nil {
		beego.Info("Insert house info error, ", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	m2m := o.QueryM2M(&house, "Facilities")
	num, err := m2m.Add(facilities)
	if err != nil || num == 0 {
		beego.Info("Insert facilities info error, ", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	resData := make(map[string]interface{})
	resData["hosue_id"] = house.Id
	resp["data"] = resData
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	return
}

func (this *HouseController) GetHouseDetailInfo() {
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	beego.Info("=========== get /api/v1.0/houses/?:id ==========")

	resData := make(map[string]interface{})

	user_id := this.GetSession("user_id")
	h_id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	beego.Info("*********************** ", h_id)
	o := orm.NewOrm()
	house := models.House{Id: h_id}
	user := models.User{Id: user_id.(int)}
	house.User = &user

	if err := o.Read(&house); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	o.LoadRelated(&house, "Area")
	o.LoadRelated(&house, "User")
	o.LoadRelated(&house, "Images")
	o.LoadRelated(&house, "Facilities")

	beego.Info("######## ", house)
	// for _, fac := range house.Facilities {
	// 	append(resD)
	// }

	resData["house"] = house
	resp["data"] = resData
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	return
}

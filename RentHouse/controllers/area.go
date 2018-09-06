package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"go_examples/RentHouse/models"
	_ "strings"
	"time"
)

type AreaController struct {
	beego.Controller
}

func (this *AreaController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *AreaController) GetAreaInfo() {
	beego.Info("============ /api/v1.0/areas get success! ===========")

	//返回给前端的map结构
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	cache_cli, err := cache.NewCache("redis", `{"key":"RentHouse","conn":"www.dog128.cn:8379","dbNum":"0"}`)
	if err != nil {
		beego.Info("cache redis conn err... err = ", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	area_info := cache_cli.Get("area_info")
	if area_info != nil {
		beego.Info("====== get area info from cache =======")
		fmt.Printf("area info = %s", area_info)

		var area_data interface{}
		json.Unmarshal(area_info.([]byte), &area_data)
		resp["data"] = area_data
		return
	}

	o := orm.NewOrm()
	var areas []models.Area

	qs := o.QueryTable("area")
	num, err := qs.All(&areas)
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

	resp["data"] = areas

	area_info_str, _ := json.Marshal(areas)
	if err := cache_cli.Put("area_info", area_info_str, time.Second*3600); err != nil {
		beego.Info("======= set area info into redis failed... err = ", err)
	}

	return
}

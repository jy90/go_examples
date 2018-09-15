package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_examples/RentHouse/models"
	"path"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UserController) Register() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	register_data := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &register_data)

	beego.Info(`register_data["mobile"] = `, register_data["mobile"])
	beego.Info(`register_data["password"] = `, register_data["password"])
	beego.Info(`register_data["sms_code"] = `, register_data["sms_code"])

	if register_data["mobile"] == "" || register_data["password"] == "" || register_data["sms_code"] == "" {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	user := models.User{}
	user.Mobile = register_data["mobile"].(string)
	user.Password_hash = register_data["password"].(string)
	user.Name = register_data["mobile"].(string)

	o := orm.NewOrm()

	id, err := o.Insert(&user)
	if err != nil {
		beego.Info("insert user data error = ", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	beego.Info("register success !!! user id = ", id)

	this.SetSession("name", user.Mobile)
	this.SetSession("user_id", id)
	this.SetSession("mobile", user.Mobile)
	return
}

func (this *UserController) Login() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	login_data := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &login_data)

	beego.Info("mobile = ", login_data["mobile"])
	beego.Info("password = ", login_data["password"])

	if login_data["mobile"] == "" || login_data["password"] == "" {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	var user models.User

	o := orm.NewOrm()
	qs := o.QueryTable("user")
	if err := qs.Filter("mobile", login_data["mobile"]).One(&user); err != nil {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}

	if user.Password_hash != login_data["password"].(string) {
		resp["errno"] = models.RECODE_PWDERR
		resp["errmsg"] = models.RecodeText(models.RECODE_PWDERR)
		return
	}

	beego.Info("======== login success =========", user.Name)

	this.SetSession("name", user.Mobile)
	this.SetSession("user_id", user.Id)
	this.SetSession("mobile", user.Mobile)
	return
}

func (this *UserController) UploadAvatar() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	file, header, err := this.GetFile("avatar")
	if err != nil {
		resp["errno"] = models.RECODE_SERVERERR
		resp["errmsg"] = models.RecodeText(models.RECODE_SERVERERR)
		return
	}

	fileBuffer := make([]byte, header.Size)
	if _, err := file.Read(fileBuffer); err != nil {
		resp["errno"] = models.RECODE_IOERR
		resp["errmsg"] = models.RecodeText(models.RECODE_IOERR)
		return
	}

	suffix := path.Ext(header.Filename)

	groupName, fileId, err := models.FDFSUploadByBuffer(fileBuffer, suffix[1:])
	if err != nil {
		resp["errno"] = models.RECODE_IOERR
		resp["errmsg"] = models.RecodeText(models.RECODE_IOERR)
		beego.Info("upload to fastdfs failed...", err)
		return
	}

	beego.Info("upload to fastdfs success groupName =", groupName, " fileId =", fileId)
	user_id := this.GetSession("user_id")

	o := orm.NewOrm()
	user := models.User{Id: user_id.(int)}
	if err := o.Read(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	user.Avatar_url = fileId
	if _, err := o.Update(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	avatar_url := "http://www.dog128.cn/" + fileId
	url_map := make(map[string]interface{})
	url_map["avatar_url"] = avatar_url
	resp["data"] = url_map
	return
}

func (this *UserController) GetUserInfo() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetData(resp)

	user_id := this.GetSession("user_id")
	user := models.User{Id: user_id.(int)}

	o := orm.NewOrm()
	// qs := o.QueryTable("user")
	if err := o.Read(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	user.Avatar_url = "http://www.dog128.cn/" + user.Avatar_url
	resp["data"] = &user
	return
}

func (this *UserController) UpdateUserName() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetData(resp)

	user_id := this.GetSession("user_id")
	UserName := make(map[string]string)
	json.Unmarshal(this.Ctx.Input.RequestBody, &UserName)
	beego.Info("get userName from RequestBody, UserName = ", UserName["name"])

	o := orm.NewOrm()
	user := models.User{Id: user_id.(int)}

	if err := o.Read(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	user.Name = UserName["name"]
	if _, err := o.Update(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	this.SetSession("name", user.Name)
	resp["data"] = &user
	return
}

func (this *UserController) UserAuth() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetData(resp)

	user_id := this.GetSession("user_id")
	realInfo := make(map[string]string)
	json.Unmarshal(this.Ctx.Input.RequestBody, &realInfo)
	beego.Info("get Real_name from RequestBody, Real_name = ", realInfo["real_name"])

	o := orm.NewOrm()
	user := models.User{Id: user_id.(int)}

	if err := o.Read(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	user.Real_name = realInfo["real_name"]
	user.Id_card = realInfo["id_card"]
	if _, err := o.Update(&user); err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	resp["data"] = &user
	return
}

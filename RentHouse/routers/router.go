package routers

import (
	"github.com/astaxie/beego"
	"go_examples/RentHouse/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetAreaInfo")

	beego.Router("/api/v1.0/index", &controllers.HouseIndexController{}, "get:GetHouseIndex")

	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSessionData;delete:DelSessionData")

	beego.Router("/api/v1.0/register", &controllers.UserController{}, "post:Register")

	beego.Router("/api/v1.0/login", &controllers.UserController{}, "post:Login")

	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{}, "post:UploadAvatar")

	beego.Router("/api/v1.0/userinfo", &controllers.UserController{}, "get:GetUserInfo")

	beego.Router("/api/v1.0/user/name", &controllers.UserController{}, "put:UpdateUserName")

	beego.Router("/api/v1.0/user/auth", &controllers.UserController{}, "get:GetUserInfo;post:UserAuth")

	beego.Router("/api/v1.0/user/houses", &controllers.HouseController{}, "get:GetHouseInfo")

	beego.Router("/api/v1.0/houses", &controllers.HouseController{}, "post:PublishHouseInfo")

	beego.Router("/api/v1.0/houses/?:id", &controllers.HouseController{}, "get:GetHouseDetailInfo")

}

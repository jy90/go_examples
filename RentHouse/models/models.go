package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id            int           `json:"user_id"`
	Name          string        `orm:"size(32);unique" json:"name"`
	Password_hash string        `orm:"size(128)" json:"password"`
	Mobile        string        `orm:"size(11);unique" json:"mobile"`
	Real_name     string        `orm:"size(32)" json:"real_name"`
	Id_card       string        `orm:"size(20)" json:"id_card"`
	Avatar_url    string        `orm:"size(256)" json:"avatar_url"`
	Houses        []*House      `orm:"reverse(many)" json:"houses"`
	Orders        []*OrderHouse `orm:"reverse(many)" json:"orders"`
}

type House struct {
	Id              int           `json:"house_id"`
	User            *User         `orm:"rel(fk)" json:"user_id"`
	Area            *Area         `orm:"rel(fk)" json:"area_id"`
	Title           string        `orm:"size(64)" json:"title"`
	Price           int           `orm:"default(0)" json:"price"`
	Addess          string        `orm:"size(512)" orm:"default("")" json:"address"`
	Room_count      int           `orm:"default(1)" json:"room_count"`
	Acreage         int           `orm:"default(0)" json:"acreage"` //房屋总面积
	Unit            string        `orm:"size(32)" orm:"default("")" json:"unit"`
	Capacity        int           `orm:"default(1)" json:"capacity"`
	Beds            string        `orm:"size(64)" orm:"default("") json:"beds"`
	Deposit         int           `orm:"default(0)" json:"deposit"` //押金
	Min_days        int           `orm:"dafault(1)" json:"min_days"`
	Max_days        int           `orm:"default(0)" json:"max_days"`
	Order_count     int           `orm:"default(0)" json:"order_couont"`
	Index_image_url string        `orm:"size(256)" orm:"default("")" json:"index_image_url"`
	Facilities      []*Facility   `orm:"reverse(many)" json:"facilities"` //房屋设施
	Images          []*HouseImage `orm:"reverse(many)" json:"img_urls"`
	Orders          []*OrderHouse `orm:"reverse(many)" json:"orders"`
	Ctime           time.Time     `orm:"auto_now_add;type(datetime)" json:"ctime"`
}

//首页最多展示房屋数量
var HOME_PAGE_MAX_HOUSES int = 5

//房屋列表页面每页显示条目数
var HOUSE_LIST_PAGE_CAPACITY int = 2

type Area struct {
	Id     int      `json:"aid"`
	Name   string   `orm:"size(32)" json:"aname"`
	Houses []*House `orm:"reverse(many)" json:"houses"`
}

type Facility struct {
	Id     int      `json:"fid"`
	Name   string   `orm:"size(32)"`
	Houses []*House `orm:"rel(m2m)"`
}

type HouseImage struct {
	Id    int    `json:"house_image_id"`
	Url   string `orm:"size(256)" json:"url"`
	House *House `orm:"rel(fk)" json:"house_id"`
}

const (
	ORDER_STATUS_WAIT_ACCEPT  = "WAIT_ACCEPT"  //待接单
	ORDER_STATUS_WAIT_PAYMENT = "WAIT_PAYMENT" //待支付
	ORDER_STATUS_PAID         = "PAID"         //已支付
	ORDER_STATUS_WAIT_COMMENT = "COMMENT"      //待评价
	ORDER_STATUS_COMPLETE     = "COMPLETE"     //已完成
	ORDER_STATUS_CANCELED     = "CANCELED"     //已取消
	ORDER_STATUS_REJECTED     = "REJECTED"     //已拒单
)

type OrderHouse struct {
	Id          int       `json:"order_id"`
	User        *User     `orm:"rel(fk)" json:"user_id"`
	House       *House    `orm:"rel(fk)" json:"house_id"`
	Begin_date  time.Time `orm:"type(datetime)"`
	End_time    time.Time `orm:"type(datetime)"`
	Days        int
	House_price int
	Amount      int
	Status      string    `orm:"default(WAIT_ACCEPT)"`
	Comment     string    `orm:"size(512)`
	Ctime       time.Time `orm:"auto_now_add;type(datetime)" json:"ctime"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(www.dog128.cn:8306)/renthouse?charset=utf8", 30)

	orm.RegisterModel(new(User), new(House), new(OrderHouse), new(Area), new(HouseImage), new(Facility))

	orm.RunSyncdb("default", false, true)
}

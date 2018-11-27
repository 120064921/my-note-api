package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}


//func init(this *UserController) {
//	this.user = new models.UserModel
//}

func (this *UserController) Get() {

	this.Data["json"] = map[string]interface{}{
		"code" : 200,
	}
	this.ServeJSON()

	//o := orm.NewOrm()
	//
	//user := models.UserModel{username: "slene"}
	//
	//// insert
	//id, err := o.Insert(&user)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
}

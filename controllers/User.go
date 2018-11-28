package controllers

import (
	"github.com/astaxie/beego"
	"my-note-api/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type UserController struct {
	beego.Controller
}

//func init(this *UserController) {
//	this.user = new models.UserModel
//}

func (this *UserController) Get() {

	o := orm.NewOrm()
	o.Using("default")
	//用户列表
	id, err := strconv.ParseInt(this.GetString("id"),10,64);
	if err != nil{
		fmt.Println("缺少参数id");
	}
	user, err2 := models.QueryById(id)
	if err2 == true{
		fmt.Println("获取模型失败");
		fmt.Println(err2);
	}else{
		fmt.Println("获取模型成功");
	}

	this.Data["json"] = map[string]interface{}{
		"code" : 200,
		"msg" : "success",
		"data" : user,
	}
	this.ServeJSON()
}

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
	idval,errId:=strconv.ParseInt(this.GetString("Id"),10,64);
	if errId!=nil{
		fmt.Println("缺少参数id");
	}
	user,err:=models.QueryById(idval)
	if err==true{
		fmt.Println("获取模型失败");
		fmt.Println(err);
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

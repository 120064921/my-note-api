package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"fmt"
)

// Model Struct
type UserModel struct {
	id   int
	username string `orm:"size(150)"`
	password string `orm:"size(255)"`
	status int
	create_time int
	last_time int

}

func init() {
	// set default database
	orm.RegisterDataBase(
		"default",
		"mysql",
		"username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8",
		30,
	)

	// register model
	orm.RegisterModel(new(UserModel))

	// create table
	//orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	user := UserModel{username: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.username = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := UserModel{id: user.id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// delete
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}

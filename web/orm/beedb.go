package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "test:123456@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	orm := beedb.New(db)
	beedb.OnDebug = true

	type Userinfo struct {
		Uid        int `PK` // 如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
		Username   string
		Departname string
		Created    time.Time
	}

	// 插入数据
	var saveone Userinfo
	saveone.Username = "Test Add User"
	saveone.Departname = "Test Add Departname"
	saveone.Created = time.Now()
	orm.Save(&saveone)
	//
	// add := make(map[string]interface{})
	// add["username"] = "astaxie"
	// add["departname"] = "cloud develop"
	// add["created"] = "2012-12-02"
	// orm.SetTable("userinfo").Insert(add)
	//
	// add := make(map[string]interface{})
	// add2 := make(map[string]interface{})
	// add["username"] = "astaxie"
	// add["departname"] = "cloud develop"
	// add["created"] = "2012-12-02"
	// add2["username"] = "astaxie2"
	// add2["departname"] = "cloud develop2"
	// add2["created"] = "2012-12-02"
	// addslice := []map[string]interface{}{add, add2}
	// orm.SetTable("userinfo").InsertBatch(addslice)

	// 更新数据
	saveone.Username = "Update Username"
	saveone.Departname = "Update Departname"
	saveone.Created = time.Now()
	orm.Save(&saveone)

	t := make(map[string]interface{})
	t["username"] = "astaxie"
	orm.SetTable("userinfo").SetPK("uid").Where(3).Update(t)

	// 查询数据
	var user Userinfo
	orm.Where("uid=?", 27).Find(&user)

	var user2 Userinfo
	orm.Where(3).Find(&user2)

	var user3 Userinfo
	orm.Where("name = ?", "john").Find(&user3)

	var user4 Userinfo
	// Where支持三个参数
	orm.Where("name = ? and age < ?", "john", 88).Find(&user4)

	// 例子1，根据条件id>3，获取20位置开始的10条数据的数据
	var allusers []Userinfo
	_ = orm.Where("uid > ?", "3").Limit(10,20).FindAll(&allusers)

	// 例子2，省略limit第二个参数，默认从0开始，获取10条数据
	var tenusers []Userinfo
	_ = orm.Where("uid > ?", "3").Limit(10).FindAll(&tenusers)

	// 例子3，获取全部数据
	var everyone []Userinfo
	_ = orm.OrderBy("uid desc,username asc").FindAll(&everyone)

	// 获取一些数据到map
	a, _ := orm.SetTable("userinfo").SetPK("uid").Where(3).Select("uid,username").FindMap()
	newA := make(map[string]interface{})
	for _, m := range a {
		for k, v := range m {
			newA[k] = string(v)
		}
	}
	fmt.Println(newA)

	// 删除数据
	// 例子1，删除单条数据
	orm.Delete(&saveone)

	// 例子2，删除多条数据
	// alluser就是上面定义的获取多条数据的slice
	orm.DeleteAll(&allusers)

	// 例子3，根据sql删除数据
	orm.SetTable("userinfo").Where("uid>?", 3).DeleteRow()

	// beedb还不支持struct的关联关系，但是有些应用却需要用到连接查询，所以现在beedb提供了一个简陋的实现方 案:
	// a, _ := orm.SetTable("userinfo").Join("LEFT", "userdeatail", "userinfo.uid=userdeatail.uid").Where()

	// 针对有些应用需要用到group by和having的功能，beedb也提供了一个简陋的实现
	// a, _  := orm.SetTable("userinfo").GroupBy("username").Having("username='astaxie'").FindMap()
}


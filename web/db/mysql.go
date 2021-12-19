package main

import (
	"database/sql"
	"fmt"

	. "study/util"

	_ "github.com/go-sql-driver/mysql"
)

// db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。
// db.Query()函数用来直接执行Sql返回Rows结果。
// stmt.Exec()函数用来执行stmt准备好的SQL语句
// 传入的参数都是=?对应的数据，这样做的方式可以一定程度上防止SQL注入
func main() {
	db, err := sql.Open("mysql", "test:123456@/test?charset=utf8")
	CheckErr(err)

	// 插入数据
	stmt, err := db.Prepare("insert userinfo set username=?, departname=?, created=?")
	CheckErr(err)

	res, err := stmt.Exec("junxi", "研发部门", "2019-01-14")
	CheckErr(err)

	id, err := res.LastInsertId()
	CheckErr(err)

	fmt.Println(id)
	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	CheckErr(err)

	res, err = stmt.Exec("junxiouba", id)
	CheckErr(err)

	affect, err := res.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("select * from userinfo")
	CheckErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		CheckErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// 删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	CheckErr(err)

	res, err = stmt.Exec(id)
	CheckErr(err)

	affect, err = res.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)
	db.Close()
}

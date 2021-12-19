package main

import (
	"database/sql"
	"fmt"

	. "study/util"

	_ "github.com/bmizerany/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=postgres sslmode=disable")
	CheckErr(err)

	// 插入数据
	insert := `INSERT INTO userinfo(username,departname,created) VALUES($1, $2, $3) returning uid`
	var id int
	err = db.QueryRow(insert,"astaxie", "研发部门", "2012-12-09").Scan(&id)
	CheckErr(err)
	fmt.Println(id)

	// 更新数据
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	CheckErr(err)

	res, err := stmt.Exec("astaxieupdate", id)
	CheckErr(err)

	affect, err := res.RowsAffected()
	CheckErr(err)
	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
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
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	CheckErr(err)

	res, err = stmt.Exec(id)
	CheckErr(err)

	affect, err = res.RowsAffected()
	CheckErr(err)

	fmt.Println(affect)

	db.Close()
}
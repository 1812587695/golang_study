package main

/*
数据库
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `profile_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `profile_id` (`profile_id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;



*/
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

//func init() {
//	db, err := sql.Open("mysql", "root:root@/beego?charset=utf8")
//	checkErr(err)
//}

func main() {
	//	insert()
	//	query()
	//	query2()
	update()
	//	remove()
}
func insert() {
	db, err := sql.Open("mysql", "root:root@/beego?charset=utf8")
	checkErr(err)
	stmt, err := db.Prepare(`INSERT user (name,profile_id) values (?,?)`)
	checkErr(err)
	res, err := stmt.Exec("to1222ny", 12)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

//func insert2() {
//	//Begin函数内部会去获取连接
//	tx, _ := db.Begin()
//	//每次循环用的都是tx内部的连接，没有新建连接，效率高

//	stmt, err := db.Prepare(`INSERT user (name,profile_id) values (?,?)`)
//	checkErr(err)
//	res, err := stmt.Exec("tony", 20)
//	checkErr(err)
//	id, err := res.LastInsertId()
//	checkErr(err)
//	fmt.Println(id)

//	//最后释放tx内部的连接
//	tx.Commit()
//}

func query() {
	db, err := sql.Open("mysql", "root:root@/beego?charset=utf8")
	checkErr(err)
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	for rows.Next() {
		var id int
		var name string
		var profile_id int
		rows.Columns()
		//		fmt.Println(rows.Columns())
		err = rows.Scan(&id, &name, &profile_id)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(profile_id)
	}

}

//这里查询的方式使用声明4个独立变量userId、userName、userAge、userSex来保存查询出来的每一行的值。
//在实际开发中通常会封装数据库的操作，对这样的查询通常会考虑返回字典类型。
func query2() {
	db, err := sql.Open("mysql", "root:root@/beego?charset=utf8")
	checkErr(err)
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}

}

// 修改
func update() {
	db, err := sql.Open("mysql", "root:root@/beego?charset=utf8")
	checkErr(err)
	stmt, err := db.Prepare(`UPDATE user SET name=? WHERE id=?`)
	checkErr(err)
	res, err := stmt.Exec("sdf12", 3)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	// num 返回1成功，0失败
	fmt.Println(num)
}

// 删除
func remove() {
	db, err := sql.Open("mysql", "root:root@/beego?charset=utf8")
	checkErr(err)
	stmt, err := db.Prepare(`DELETE FROM user WHERE id=?`)
	checkErr(err)
	res, err := stmt.Exec(1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	// num 返回1成功，0失败
	fmt.Println(num)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

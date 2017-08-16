package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var db =&sql.DB{}
func init(){
	db, _ = sql.Open("mysql", "root:123456@tcp(localhost:3306)/mytest?charset=utf8")
}
func main(){
	queryData()
}


func deleteData(){
	stmt, _ := db.Prepare(`DELETE FROM mytable WHERE id=?`)
	res, _ := stmt.Exec(2)
	num, _ := res.RowsAffected()
	fmt.Println(num)
}


func insertData(){
	stmt, _ := db.Prepare(`insert into mytable (name,place) values (?,?)`)
	myName:="skjfslkdjf"
	myPlace:="sdlkfjsdlkfj"
	res, _ := stmt.Exec(myName,myPlace)
	id, _ := res.LastInsertId()
	fmt.Println("suc",id)
}


func queryData(){
	rows, _ := db.Query("SELECT * FROM mytable where id=?",2)
	for rows.Next() {
		var name string
		var place string
		var id int
		rows.Scan(&name,&place,&id);
		fmt.Println(name,place,id)
	}
}
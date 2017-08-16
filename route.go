package main

import (
   "fmt"
   "net/http"
   "net/url"
   "github.com/drone/routes"
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
   "encoding/json"
)



var db =&sql.DB{}
func init(){
	db, _ = sql.Open("mysql", "root:123456@tcp(localhost:3306)/mytest?charset=utf8")
}

func getuser(w http.ResponseWriter, r *http.Request) {
   var params url.Values = r.URL.Query()
   var uid string = params.Get(":uid")
   fmt.Fprintln(w, "get a user ", uid, " success!")
}
func getuserAndAge(w http.ResponseWriter, r *http.Request) {
   var params url.Values = r.URL.Query()
   var uid string = params.Get(":uid")
   var age string = params.Get(":age")
   fmt.Fprintln(w, "get a user ", uid, " success! age is ", age)
}
func edituser(w http.ResponseWriter, r *http.Request) {
   var params url.Values = r.URL.Query()
   var uid string = params.Get(":uid")
   fmt.Fprintln(w, "edit a user ", uid, " success!")
}
func testRoute(w http.ResponseWriter,r *http.Request){
    cross(w)
	var params url.Values = r.URL.Query()
	var sid string =params.Get(":sid")
	fmt.Println(sid)
	fmt.Fprintln(w, sid)
}
func myRoute(w http.ResponseWriter,r *http.Request){
    cross(w)
    var params url.Values = r.URL.Query()
    var sid string =params.Get(":sid")
    var dataMap map[int]string
    /* 创建集合 */
    dataMap = make(map[int]string)
    rows, _ := db.Query("SELECT * FROM mytable where id=?",sid)
	for rows.Next() {
		var name string
		var place string
		var id int
        rows.Scan(&name,&place,&id)
        dataMap[id]=name
    }
    b, _ := json.Marshal(dataMap)
	fmt.Fprintln(w, string(b))
}
func cross(w http.ResponseWriter){
    w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
}
func main() {
   fmt.Println("正在启动WEB服务...")
   var mux *routes.RouteMux = routes.New()
   mux.Get("/:sid",testRoute)
   mux.Get("/mytest/:sid",myRoute)
   mux.Get("/user/:uid", getuser)
   mux.Get("/user/:uid/:age", getuserAndAge)
   mux.Post("/user/:uid", edituser)

   //http.Handle("/", mux)
   http.ListenAndServe(":8088", mux)
   fmt.Println("服务已停止")
}
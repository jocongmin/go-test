package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	router := httprouter.New()
	router.POST("/upload", Upload)
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	renderHTML(w, "upload.html", "no data")
}

func Upload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cross(w)
	formFile, header, err := r.FormFile("uploadfile")
	if err != nil {
		log.Printf("Get form file failed: %s\n", err)
		return
	}
	defer formFile.Close()
	f, err := CopyFile("./file/"+header.Filename, "./getfile/"+header.Filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(f)
	http.Redirect(w, r, "/", http.StatusFound)
}
func renderHTML(w http.ResponseWriter, file string, data interface{}) {
	// 获取页面内容
	t, _ := template.New(file).ParseFiles(file)
	// 将页面渲染后反馈给客户端
	t.Execute(w, data)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
func cross(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
}

func CopyFile(src, des string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()

	desFile, err := os.Create(des)
	if err != nil {
		fmt.Println(err)
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}

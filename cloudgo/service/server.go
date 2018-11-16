package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	// Classic 自动生成应用了基本的中间件 Logger
	n := negroni.Classic()
	mx := mux.NewRouter()

	// 初始化路由器
	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root

			fmt.Println(root)
		}
	}

	// 练习：处理表单提交
	mx.HandleFunc("/login", getLoginHandler(formatter)).Methods("POST")

	// example 教程中提供的使用使用模板来输出index 信息
	mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")

	// 练习： 使用模板生成手机销售表。
	mx.HandleFunc("/testmoblie", phoneSaleHandler(formatter)).Methods("GET")

	// mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
	mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")

	// 练习：模仿 404 NotFound ERROR 设置一个 501 NotImplemented的 handler
	mx.HandleFunc("/api/unknown", NotImplementedHandler()).Methods("GET")

	// 例子：使用 HTTP 包提供的静态服务器
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))

}

// NotImplementedHandler test
func NotImplementedHandler() http.HandlerFunc {
	// type conversion
	return NotImplemented
}

// NotImplemented handle the situation when server internal error occurs
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	// 如何路由到对应的函数
	http.Error(w, "ERROR : 501 not Implemented", 501)
}

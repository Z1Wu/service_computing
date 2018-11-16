package service

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/unrolled/render"
)

// func (formatter *render.Render) http.HandlerFunc {

// 	// 使用 formatter 来利用模板输出对应的 html 文件。
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		formatter.HTML(w, http.StatusOK, "index", struct {
// 			ID      string `json:"id"`
// 			Content string `json:"content"`
// 		}{ID: "111", Content: "Hello from Go!"})
// 	}
// }

// UserInfo struct constructs from the form passed when user login.
type UserInfo struct {
	Username string
	Password string
}

func getLoginHandler(formatter *render.Render) http.HandlerFunc {
	// fmt.Println("method:", r.Method) //获取请求的方法
	// if r.Method == "GET" {

	// } else {
	// 	//请求的是登录数据，那么执行登录的逻辑判断
	// 	fmt.Println("username:", r.Form["username"])
	// 	fmt.Println("password:", r.Form["password"])
	// }

	// values := map[string][]string{
	// 	"Name":  {"John"},
	// 	"Phone": {"999-999-999"},
	// }
	// 首先需要解析 url 中的参数，把post的表格的数据放到对应的 request 中。

	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// 拿到指针
		user := new(UserInfo)
		decoder := schema.NewDecoder()
		log.Print(r.Form)
		decoder.Decode(user, r.Form)
		// 测试拿到的信息。
		formatter.HTML(w, http.StatusOK, "index", struct {
			ID       string `json:"id"`
			Content  string `json:"content"`
			UserName string `json:"username"`
		}{ID: "111", Content: "Hello from Go!", UserName: user.Username})
	}

}

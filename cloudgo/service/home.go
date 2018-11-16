package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {

	// 使用 formatter 来利用模板输出对应的 html 文件。
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: "111", Content: "Hello from Go!"})
	}
}

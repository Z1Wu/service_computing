package service

import (
	"net/http"

	"github.com/unrolled/render"
)

// saleData repersent the date
type saleData struct {
	Name     string
	SaleNum  int
	Q1       float32
	Q2       float32
	Q3       float32
	Q3Before float32
	Q3Last   float32
}

type saleFormat struct {
	Items []saleData
}

// func toPercentage(num float32) string {
// 	return fmt.Sprintf("%.2f", num*100)
// }

func phoneSaleHandler(formatter *render.Render) http.HandlerFunc {
	// 使用 formatter 来利用模板输出对应的 html 文件。

	// myData := saleFormat{items: {[]saleData{
	// 	Name:     "mi",
	// 	Q1:       11.0,
	// 	Q2:       11.0,
	// 	Q3:       11.0,
	// 	Q3Before: 11.0,
	// 	Q3Last:   11.0}}}

	var myData saleFormat

	phone := saleData{
		Name:     "mi",
		SaleNum:  100,
		Q1:       11.0,
		Q2:       11.0,
		Q3:       11.0,
		Q3Before: 11.0,
		Q3Last:   11.0}

	myData.Items = append(myData.Items, phone)

	phone.Name = "HUAWEI"
	myData.Items = append(myData.Items, phone)

	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "moblie_sales", myData)
	}
}

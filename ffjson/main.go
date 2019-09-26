package main

import (
	"fmt"
	"github.com/xiazemin/json/ffjson/struct"
)
func main()  {
	news := model.NewsModel{110,"hello"}
	fmt.Println(news.ToJson())	// 打印：{"Id":110,"Title":"hello"}
}

package model

import "github.com/pquerna/ffjson/ffjson"
// 定义一个结构体
type NewsModel struct {
	Id int
	Title string
}

// 定义一个方法
func (news NewsModel) ToJson() string  {
	res,err := ffjson.Marshal(news)
	if err != nil {
		return  err.Error()
	}

	// 得到是字节数组，所以还有转为string
	return string(res)
}

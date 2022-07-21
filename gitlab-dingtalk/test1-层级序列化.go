package demo

import (
	//"strings"
    //"time"
	"encoding/json"
	"fmt"
)
type Test struct{
	Id int `json:id`
	Class *Class `json:"class"`
}
type Class struct {
    Name  string
    Grade int
}
func main(){
	//model := GitlabCIModel{  ObjectAttributes{Id: 1}, User{Name: "xiaoli"},Project{Id: 1}  }
	cla := new(Class)
	cla.Name= "1班"
	cla.Grade= 3
	model := Test{
		Id: 1,
	}
	model.Class=cla
	jsonStu, err := json.Marshal(model)
	
    if err != nil {
        fmt.Println("生成json字符串错误")
    }

    fmt.Println(string(jsonStu))
}

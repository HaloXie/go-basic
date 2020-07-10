package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := fmt.Sprint("成绩: ", s.Score, ", 好好学习，天天向上。")
	fmt.Println(msg)
	return msg
}

func (s student) Study2(name string) string {
	msg := fmt.Sprint("名字: ", name, ", 好好学习，天天向上。")
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("type:%s\n", t)
	fmt.Printf("Method Count:%v\n", t.NumMethod())

	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s type:%s\n", t.Method(i).Name, methodType)

		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		if t.Method(i).Name == "Study2" {
			scoreField := v.FieldByName("Name")
			// fmt.Println(scoreField)

			// var args = make([]reflect.Value, 1)
			// args[0] = reflect.ValueOf(scoreField.String())
			var args = []reflect.Value{reflect.ValueOf(scoreField.String())}
			v.Method(i).Call(args)
		} else {
			// 	v.Method(i).Call([]reflect.Value{})
			v.Method(i).Call(nil)
		}
	}
}

func main() {
	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}
	t := reflect.TypeOf(stu1)
	v := reflect.ValueOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct

	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldv := v.Field(i)
		// 结构体的类型判断
		if fieldv.Kind() == reflect.Int {
			fmt.Printf("name:%s value:%d index:%d type:%v json tag:%v\n", field.Name, fieldv.Int(),
				field.Index, field.Type, field.Tag.Get("json"))
		} else {
			fmt.Printf("name:%s value:%s index:%d type:%v json tag:%v\n", field.Name, fieldv.String(),
				field.Index, field.Type, field.Tag.Get("json"))
		}
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name,
			scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

	// 反射调用对应类型的方法
	printMethod(stu1)
}

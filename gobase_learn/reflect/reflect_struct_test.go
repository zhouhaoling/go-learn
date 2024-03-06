package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Id   int
	Name string
	Age  int
}

// 绑方法
func (u User) Hello() {
	fmt.Println("Hello")
}

func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型:", t)

	//获取值
	v := reflect.ValueOf(o)
	fmt.Println(v)
	//获取所有属性
	//获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		//取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v", f.Name, f.Type)
		//interface(), 获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
	}
	fmt.Println("=================方法====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}
}

func TestStruct(t *testing.T) {
	//u := User{1, "zs", 20}
	//Poni(u)
	//noStruct()
	GetTag()
}

type Boy struct {
	User
	Addr string
}

func noStruct() {
	m := Boy{User{1, "zs", 20}, "bj"}
	t := reflect.TypeOf(m)
	fmt.Println(t)
	// Anonymous：匿名
	//t.Field(i) : 返回结构体的第 i 个位置的类型
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", reflect.ValueOf(m).Field(0))
}

// 修改结构体值
func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	//获取指针指向的元素
	v = v.Elem()
	//取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("kuteng")
	}
}

type Student struct {
	Name string `json:"name1" db:"name2"`
}

func GetTag() {
	var s Student
	v := reflect.ValueOf(&s)
	//类型
	t := v.Type()
	//获取字段
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("db"))
}

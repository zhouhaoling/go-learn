package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// 反射获取interface类型信息
func reflectType(a interface{}) {
	//typ:返回接口的类型
	typ := reflect.TypeOf(a)
	fmt.Println("类型是：", typ)

	// kind() 可以获取具体类型
	k := typ.Kind()
	fmt.Println(k)
	//只能通过 reflect.数据类型 来进行类型判断，不能直接用float64判断
	switch k {
	case reflect.Float64:
		fmt.Println(" is float64")
	case reflect.String:
		fmt.Println("string")
	}
}

func TestName(t *testing.T) {
	var x float64 = 3.4
	//reflectType(x)
	//reflectValue(x)
	reflectSetValue(&x)
	fmt.Println("main:", x)
}

// 反射获取interface类型的值
func reflectValue(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	switch k {
	case reflect.Float64:
		fmt.Println("a是：", v.Float())
	}
}

// 通过反射设置值
func reflectSetValue(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		//通过反射修改值
		v.SetFloat(6.9)
		fmt.Println(v.Float())
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}
}

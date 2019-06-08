//// rtype is the common implementation of most values.
// It is embedded in other struct types.
//
// rtype must be kept in sync with ../runtime/type.go:/^type._type.
// type rtype struct {
//     size       uintptr
//     ptrdata    uintptr  // number of bytes in the type that can contain pointers
//     hash       uint32   // hash of type; avoids computation in hash tables
//     tflag      tflag    // extra type information flags
//     align      uint8    // alignment of variable with this type
//     fieldAlign uint8    // alignment of struct field with this type
//     kind       uint8    // enumeration for C
//     alg        *typeAlg // algorithm table
//     gcdata     *byte    // garbage collection data
//     str        nameOff  // string form
//     ptrToThis  typeOff  // type for pointer to this type, may be zero
// }

// type Value struct{
// 	typ *rtype
// 	ptr Unsafe.Pointer
// 	flag
// }

// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.

// func TypeOf(i interface{}) Type {
// 	eface := *(*emptyInterface)(unsafe.Pointer(&i))
// 	return toType(eface.typ)     // toType返回值是一个rtype的指针
// }

package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {

	// 调用kind后可以直接与reflect中定义的const值来比较
	if reflect.ValueOf(q).Kind() == reflect.Struct {

		// TypeOf返回的是接口，要想获取名字的字符串形式，使用Name()函数
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)

		// NumField和Field是通过reflect.Value类型调用
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {

			//Field()返回的是reflec.Value
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)

}

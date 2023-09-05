package main

import (
	"fmt"
	"reflect"
)

func main() {
	strs := []string{"xxx", "aaaa"}
	ints := []int{1, 2}
	printArray(strs)
	printArray(ints)
	println("====")
	printArrayAny(ints)
	printArrayAny(strs)
	folats := []float32{1.0, 2.0}
	printArrayAny(folats)

	println("==== 测试自定义范型类型")
	var s Slice[int] = []int{1, 2}
	fmt.Println(s.customSliceSum())
	printCustomSlice(s)
}
func printArray[T string | int](arr []T) {
	for _, a := range arr {
		fmt.Println(reflect.TypeOf(a).Kind(), " ", a)
	}
}

// 范型类型可以是具体的类型，string 、int 等
// 也可以是any ,表示go里面的所有基本类型
// 也可以是comparable ，标识go里面所有的可比较类型，int uint float bool struct 指针等
func printArrayAny[T any](arr []T) {
	arrOf := reflect.TypeOf(arr)
	fmt.Println(arrOf)
	// 通过反射获取范型的具体类型
	fmt.Println(arrOf.Elem())
	for _, a := range arr {
		fmt.Println(reflect.TypeOf(a).Kind(), " ", a)
	}
}

// Slice 自定义范型类型
type Slice[T int | float32] []T

// Map 自定义map类型
type Map[K int | string, V float32 | int] map[K]V

// Struct 自定义结构体
type Struct[T int | string] struct {
	Id   T
	Name string
}

// Interface 定义范型接口
type Interface[T int | float32 | string] interface {
	Print(data T)
}

// Map2 特殊的范型类型,标识该map只能接受int类型的value
type Map2[K int | string] map[K]int

// customSliceSum 给范型添加函数
func (s Slice[T]) customSliceSum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// Add 定义范型函数
func Add[T int | float32](a T, b T) T {
	return a + b
}

func printCustomSlice[T int | float32](s Slice[T]) {
	for _, v := range s {
		fmt.Print(v)
	}
	fmt.Println()
}

// MyInt 也可以自定义范型类型,~号的作用是，支持该类型的衍生类型~int 等效 int|int8|......
type MyInt interface {
	~int | int8 | int16
}

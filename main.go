package main

import (
	"fmt"
)

type Stringer interface{
	String() string
}

//レシーバにできる型には制限があるので、こちらで定義して用意する
type MyInt int
type MyString string
type Person struct{
	name string
	age int
}

func(i MyInt)String()string{
	return fmt.Sprintln("整数")
}

func(p Person)String()string{
	return fmt.Sprintln("Per")
}

func(s MyString)String()string{
	return fmt.Sprintf("文字列")
}

func switchType(s Stringer){
	switch v:=s.(type){
	case MyInt:
		fmt.Printf("%T:%v\n",v,v)
	case MyString:
		fmt.Printf("%T:%v\n",v,v)
	case Person:
		fmt.Printf("%v\n",v)
	}
}

func main(){
	var s Stringer=Person{name:"Taro",age:20}
	switchType(s)
}
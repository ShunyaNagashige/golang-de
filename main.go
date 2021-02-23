package main

import (
	"fmt"
	"os"
)

//任意の値を、Stringer型に変換したい！！

type Stringer interface{
	String() string
}

//戻り値errorインタフェースについて
//MyErrorはerrorインタフェースを実装しているため、MyErrorを返してもok
func ToStringer(v interface{})(Stringer,error){
	if s,ok:=v.(Stringer);ok{
		return s,nil
	}
	return nil,MyError("CastError")
}

type MyError string

//MyErrorはerrorインタフェースを実装する
func(e MyError)Error()string{
	return string(e)
}

type S string

//Sをstringにキャスト
func(s S)String()string{
	return string(s)
}

func main(){
	v:=S("hoge")
	if s,err:=ToStringer(v);err!=nil{
		fmt.Fprintln(os.Stderr,"ERROR:",err)
	}else{
		fmt.Println(s.String())
	}
}
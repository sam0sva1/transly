package tools

import "fmt"

func Chk(err error) {
	if err != nil {
		panic(err)
	}
}

func Log(data ...interface{}) {
	fmt.Println(data)
}

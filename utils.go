package main

import "fmt"


func chkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
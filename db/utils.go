package db

import "fmt"

func checkQueryError(err error, msg string) {
	if err != nil {
		fmt.Println(err.Error())
		panic(msg)
	}
}
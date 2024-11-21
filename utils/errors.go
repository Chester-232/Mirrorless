package utils

import "fmt"

func Checker(mess string, err error) {
	if err != nil {
		fmt.Println(mess, err);
	}
}
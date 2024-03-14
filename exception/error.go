package exception

import "fmt"

func PanicLogging(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

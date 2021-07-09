package main

import (
	"fmt"
	"restbase/model"
	"restbase/rest"
)

func main() {
	fmt.Println("\n------------------Main Function Starting--------")
	model.Migration()
	rest.Handler()

}

package def

import (
	"fmt"
	"log"
)

func DefFoo() {
	//fmt.Println("check error bos")

	if r := recover(); r != nil {
		fmt.Println("eh ada error bos ", r)
		log.Println("eh ada error bos ", r)
	}

	//fmt.Println("check error kelar bos")

}

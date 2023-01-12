package helper

import (
	"fmt"
	"os"
)

func DD(args ...interface{}) {
	fmt.Printf("%#v\n", args)
	os.Exit(0)
}

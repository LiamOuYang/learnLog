package main

import (
	"fmt"

	utime "github.com/apodemakeles/ugo/time"
)

func main() {
	fmt.Println("this is go dep demo utime:", utime.NowUnixTS())
}

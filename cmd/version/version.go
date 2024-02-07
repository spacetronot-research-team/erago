package version

import "fmt"

func Version() {
	fmt.Println(Current)
}

var Current = "v0.0.24"

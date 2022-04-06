package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	arr := []string{"Samuel", "Marc", "Samuel"}
	names := lo.Uniq[string](arr)
	// []string{"Samuel", "Marc"}
	fmt.Println(names)

}

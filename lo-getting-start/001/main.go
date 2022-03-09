package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	names := lo.Uniq[string]([]string{"Samuel", "Marc", "Samuel"})
	// []string{"Samuel", "Marc"}
	fmt.Println(names)

}

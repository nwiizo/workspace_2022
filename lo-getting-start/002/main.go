package main

import "fmt"

func main() {
	arr := []string{"Samuel", "Marc", "Samuel"}
	m := make(map[string]struct{}) // 空のstructを使う
	for _, ele := range arr {
		m[ele] = struct{}{} // m["a"] = struct{}{} が二度目は同じものとみなされて重複が消える。
	}

	fmt.Printf("%v", m) // ["Samuel", "Marc"]
}

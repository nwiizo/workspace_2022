package main

import (
	"fmt"
	"strings"

	"github.com/bitfield/script"
)

func main() {
	user, _ := script.Exec("whoami").String()

	user_file_path := "/Users/" + remove_line_breaks(user) + "/"
	fmt.Println(user_file_path)
	_, err := script.FindFiles(user_file_path).Stdout()
	if err != nil {
		fmt.Println(err)
	}
}

// 末尾の改行を削除する
func remove_line_breaks(s string) string {
	s = strings.TrimRight(s, "\n")
	if strings.HasSuffix(s, "\r") {
		s = strings.TrimRight(s, "\r")
	}

	return s
}

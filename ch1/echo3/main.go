// コマンドラインの引数を表示する
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// stringsパッケージのJoin関数を使用
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}
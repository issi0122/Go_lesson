// コマンドラインの引数を表示する
package main

import (
	"fmt"
	"os"
)

func main() {
	// 2つの変数がstring型であることを示す
	var s, sep string
	// os.Argsはコマンドラインの引数が入る
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		// separate文字
		sep = " "
	}
	fmt.Println(s)
}
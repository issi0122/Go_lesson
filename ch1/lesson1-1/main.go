// コマンドラインの引数を表示する
package main

import (
	"fmt"
	"os"
)

func main() {
  // 2つの変数がstring型であることを示す
	var s, sep string
	// os.Argsはコマンドラインの引数が入る(0には実行ファイル名が入る)
	// 実行時はbuildしておくこと
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		// separate文字
		sep = " "
	}
	fmt.Println(s)
}
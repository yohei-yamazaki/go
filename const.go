// コメント
/*
コメント
*/

// １文字目が大文字だったらグローバル変数、小文字だったらローカル変数
package main

import "fmt"

func main(){
  const (
    a = iota
    b
    c
	)
	fmt.Println(a, b,c)
}

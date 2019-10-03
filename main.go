// コメント
/*
コメント
*/

// １文字目が大文字だったらグローバル変数、小文字だったらローカル変数
package main

import "fmt"

func main(){
  msg := "hello world"
  fmt.Println(msg)

  var (
    a int
    s string
  )
  a = 1
  s = "s"
  fmt.Println(a,s)

  b:=10
  c:=13.5
  d:="hoge"
  var e bool
  fmt.Printf("b: %d, c:%f, d:%s, e:%t \n", b,c,d,e)
}

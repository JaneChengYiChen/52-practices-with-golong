
// 宣告程式屬於哪個 package
package main

// 引入套件
import (
    "fmt"
    "math/rand"
    "io"
    "net/http"
)

// 程式執行入口
func main() {
    start()
    habit()

    // router 讓網址可以對應處理函式
    http.HandleFunc("/", hello)
    // 監聽 8000 port
    http.ListenAndServe(":8000", nil)
}

// ref https://blog.kdchang.cc/2017/07/01/golang101-tutorial-introduction/
func start(){ 
    greetings := "Hello, Jane!"
    introduction := "My name is ___, is there anything I can help you?"
    fmt.Println(greetings, introduction)
}
// ref https://blog.kdchang.cc/2017/07/01/golang101-tutorial-introduction/
func habit() {
	fmt.Println("My favorite number is", rand.Intn(10)) 
}

// ref https://blog.kdchang.cc/2017/07/01/golang101-tutorial-introduction/
func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello Jane! It's your web page :))") 
}

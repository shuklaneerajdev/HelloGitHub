package main

import(
  "fmt"
  "net/http"
)

func main() {
  resp , err := http.Get("http://www.google.com")
  fmt.Println("The response from the server is ")
  fmt.Println(resp)
  fmt.Println("The error message obtained is ")
  fmt.Println(err)
}

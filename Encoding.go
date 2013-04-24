package main

import(
  "fmt"
  "encoding/base64"
  "crypto/sha512"
  "crypto/hmac"
)

// This has been heavily inspired fromt he code at http://pastebin.com/pK1L4x8V

func main(){
  fmt.Println("Here we would try to encode a string using the Go library")
  var key_b64 string
  var data string
  data = "1234567890"
  fmt.Println(data)
  key_b64 = "5HpXwHwrRxhQhXusxRu4jfuAZombb3GC3FeEQwhuH01exK8mdsrTC3hniMZ69lR5mDvCjN3kM9k8oBbTEM57tQ=="
  var key []byte
  var err error
  key , err = base64.StdEncoding.DecodeString(key_b64)
  fmt.Println(err)
  shash := hmac.New(sha512.New, key)
  fmt.Println(shash)
}



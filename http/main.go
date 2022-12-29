package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
type logWriter struct{
	
}
func main() {
// 
resp,err:= http.Get("https://www.invisionapp.com/")
// error handler 

if err!=nil{
	fmt.Println("Error:",err)
	os.Exit(1)
}
// bytesilce:= make([]byte,99999)

// resp.Body.Read(bytesilce)	
//  fmt.Println(string((bytesilce)))

// lw acts like a writer function in io.copy

lw:=logWriter{}

io.Copy(lw,resp.Body)

}

func (logWriter) Write(bs []byte) (int,error) {

	fmt.Println(string(bs))
	   fmt.Println("just wrote this may bytes:", len(bs)) 
	   return len(bs),nil
}

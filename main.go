package main 

 import (
         "fmt"
         "net"
         "os"
 )

 func main() {
 	ln, err := net.Listen("tcp", "0.0.0.0:"+os.Args[1])
 	if err != nil {
 		fmt.Printf("Port already being used! %s\n", err)
 	} else {
 		fmt.Printf("Port open\n")
 		ln.Close()
 	}
 }
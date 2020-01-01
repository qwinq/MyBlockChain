package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main()  {
	str1:=[]string{"hello","world","!"}
	res:=strings.Join(str1," ")
	fmt.Printf("res: %s\n",res)

	//Join(s [][]byte, sep []byte) []byte
	res1:=bytes.Join([][]byte{[]byte("hello"),[]byte("world")},[]byte(" "))
	fmt.Printf("res: %s\n",res1)
}

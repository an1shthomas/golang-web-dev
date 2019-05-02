package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main(){

	fmt.Println(os.Args[0])

	name := "Anish Thomas"
	str := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Hello World!</title>
</head>
<body>
<h1>` + name + `</h1>
</body>
</html>
`)
	nf,err := os.Create("index.html")
	if err != nil{
		log.Fatal("error creating file", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}

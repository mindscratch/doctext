package main

import (
	"fmt"
	"os"

	"github.com/sajari/docconv"
)

func main() {
	path := os.Args[1]
	fmt.Println("path:", path)
	res, err := docconv.ConvertPath(path)
	if err != nil {
		panic(fmt.Sprintf("unable to convert %s to text: %v", path, err))
	}
	if res.Body != "" {
		fmt.Println(res.Body)
	} else {
		fmt.Println("fail, no content")
	}
}

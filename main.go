package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/briantjt/uwuify/uwu"
)

func main() {
	reader, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uwu.UwUify(reader))
}

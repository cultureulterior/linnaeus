package main

import (
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"flag"
)

var out = flag.String("out", "-", "Outfile")

func main() {
	flag.Parse()
	data:=make(map[string]string)
	jsonString, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	if out != nil && *out != "-" {
		ioutil.WriteFile(*out, jsonString, 0666)
	} else {
		fmt.Print(string(jsonString))
	}
}

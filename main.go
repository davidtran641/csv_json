package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/davidtran641/csv_json/config"
	"github.com/davidtran641/csv_json/convert"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error, useage: main <file_path>")
		return
	}
	fileName := os.Args[1]

	reader, err := os.Open(fileName)
	if err != nil {
		log.Fatal("err", err)
		return
	}

	config := config.LoadConfig()
	res := convert.Parse(reader, convert.NewConverter(config.HeaderMapping))

	buf, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		log.Fatal("err encode json", err)
		return
	}

	fmt.Println(string(buf))

}

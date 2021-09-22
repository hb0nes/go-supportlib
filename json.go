package supportlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func PrintJSON(i interface{}) {
	pkgJSON, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		log.Println("PrintJSON: Could not convert interface to JSON.")
		return
	}
	fmt.Printf("%s\n", pkgJSON)
}

func WriteJSONToFile(i interface{}, fileName string) {
	json, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		log.Println("writeJSONToFile: Could not convert interface to JSON.")
		return
	}
	if err := ioutil.WriteFile(fileName, json, 0644); err != nil {
		log.Println("writeJSONToFile: Could not write file.")
		return
	}
}

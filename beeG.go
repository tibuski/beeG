package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	jsonFile, err := os.Open("beex.ini")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened bee.ini")
	defer jsonFile.Close()

	byteResult, _ := ioutil.ReadAll(jsonFile)

	var res map[string]interface{}
	json.Unmarshal([]byte(byteResult), &res)

	fmt.Println(res["options"])
}

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/qri-io/jsonschema"
)

func main() {
	var schemaData = []byte(`{
		"type": "object",
		"properties":{
			"name":	{	
				"type": "string",
				"pattern": "^[A-Z]"
			},
			"age":	{
				"type": "number",
				"minimum": 0
			},
			"numbers":{
				"type": "array",
				"minItems" : 2,
				"items":{
					"type": "number",
					"minimum": 0
				}
			}
		},
		"required":["name","numbers"]
	}`)
	rs := &jsonschema.RootSchema{}
	if err := json.Unmarshal(schemaData, rs); err != nil {
		panic("unmarshal schema: " + err.Error())
	}

	buf := bufio.NewReader(os.Stdin)

	valid, _, err := buf.ReadLine()
	if err != nil {
		fmt.Println(err.Error())
	}
	if validateErrors, jsonError := rs.ValidateBytes(valid); len(validateErrors) > 0 {
		fmt.Println(validateErrors[0].Error())
	} else if jsonError != nil {
		fmt.Println(jsonError.Error())
	} else {
		fmt.Println("Is valid ")

	}
}

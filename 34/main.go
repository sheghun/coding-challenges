package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var dat map[string]interface{}

	byt := scanner.Bytes()

	log.Println("I am here")

	if err := json.Unmarshal(byt, &dat); err != nil {
		log.Fatal("here", err)
	}

	log.Println("I am here")

	formatInput(dat, "")

}

func formatInput(arg any, s string) {
	argType := reflect.TypeOf(arg).Kind()
	if argType == reflect.Map {
		convertedArg := convertArg[map[string]any](arg)
		fmt.Println(s + "{")
		for key, val := range convertedArg {
			valType := reflect.TypeOf(val).Kind()
			if valType == reflect.Array || valType == reflect.Slice || valType == reflect.Map {
				fmt.Printf(s+"  %s: ", key)
				formatInput(val, "    ")
				continue
			}
			fmt.Printf(s+"  %s: %v\n", key, val)
		}
		fmt.Println(s + "}")
	}

	if argType == reflect.Array || argType == reflect.Slice {
		arrayArg, ok := arg.([]interface{})
		if !ok {
			log.Fatal("Invalid array supplied")
		}

		newSpace := strings.Replace(s, " ", "", 2)
		fmt.Println(newSpace + "[")
		for _, val := range arrayArg {
			valType := reflect.TypeOf(val).Kind()
			if valType == reflect.Map || valType == reflect.Array || valType == reflect.Map {
				formatInput(val, "    ")
				continue
			}
			fmt.Printf(s+"  %v", val)
		}
		fmt.Println(newSpace + "]")
	}
}

func convertArg[T any](arg any) T {
	return arg.(T)
}

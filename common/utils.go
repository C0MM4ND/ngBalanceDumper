package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadMap(filepath string)map[string]string {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	theMap := make(map[string]string)
	err = json.Unmarshal(content, &theMap)
	if err != nil {
		panic(err)
	}

	return theMap
}

func SaveMap(theMap map[string]string, fname string) {
	content, _ := json.MarshalIndent(theMap, "", "  ")
	err := ioutil.WriteFile(fname, content, 0644)
	if err != nil {
		panic(err)
	}
}
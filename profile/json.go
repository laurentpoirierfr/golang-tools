package profile

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/tidwall/gjson"
)

var bootstrap string
var application string
var values map[string]string

func init() {

	values = make(map[string]string)

	loadBootstrapValues()
	loadApplicationValues()

}

// GetValueString :
func GetValueString(key string) string {
	return values[key]
}

// GetValueInteger :
func GetValueInteger(key string) int {
	value := GetValueString(key)
	if val, err := strconv.Atoi(value); err == nil {
		return val
	} else {
		log.Println(value, "is not an integer.")
		return 0
	}
}

func getFieldsName(source string, parent string) []string {
	res := make([]string, 0)
	var arr map[string]gjson.Result
	if gjson.Parse(source).IsObject() {
		arr = gjson.Parse(source).Map()
	} else {
		return []string{parent}
	}
	for key, val := range arr {
		var temp []string
		if parent == "" {
			temp = getFieldsName(val.Raw, key)
		} else {
			temp = getFieldsName(val.Raw, parent+"."+key)
		}
		res = append(res, temp...)
	}
	return res
}

func loadBootstrapValues() {
	loadValues("bootstrap.json")
}

func loadApplicationValues() {
	filename := "application.json"
	if envProfile := os.Getenv("PROFILE"); len(envProfile) > 0 {
		filename = "application-" + envProfile + ".json"
		log.Println("Started with profile :", envProfile)
	} else {
		log.Println("Started with profile : default")
	}
	loadValues(filename)
}

func loadValues(filename string) {
	json, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(filename, "NOT loaded ...")
	} else {
		log.Println(filename, "loaded ...")
		j := string(json)
		fields := getFieldsName(string(j), "")
		for _, field := range fields {
			values[field] = gjson.Get(j, field).String()
		}
	}
}

package profile

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/laurentpoirierfr/golang-tools/tools"
	"github.com/tidwall/gjson"
)

var bootstrap string
var application string
var values map[string]string

func init() {

	values = make(map[string]string)

	json, err := ioutil.ReadFile("bootstrap.json")
	if err != nil {
		log.Println("bootstrap.json NOT loaded ...")
	} else {
		log.Println("bootstrap.json loaded ...")
		bootstrap = string(json)
		loadBootstrapValues()
	}

	profile := os.Getenv("PROFILE")
	filename := "application.json"
	if profile != "" {
		filename = "application-" + profile + ".json"
	}
	json, err = ioutil.ReadFile(filename)
	tools.FailOnError(err, "Failed to read file : "+filename)
	application = string(json)
	loadApplicationValues()
	log.Println(filename + " loaded ...")

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
	loadValues(bootstrap)
}

func loadApplicationValues() {
	loadValues(application)
}

func loadValues(json string) {
	fields := getFieldsName(json, "")
	for _, field := range fields {
		values[field] = gjson.Get(json, field).String()
	}
}

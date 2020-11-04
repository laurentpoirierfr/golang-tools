package config

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

var values map[string]string

func init() {

	values = make(map[string]string)

	loadBootstrapValues()
	loadApplicationValues()

}

// GetStringValue :
func GetStringValue(key string) string {
	result := values[key]
	var vars []string
	vars = referenceVar(result)
	if len(vars) > 0 {
		for i := 0; i < len(vars); i++ {
			value := extractVar(vars[i])
			svalue := strings.Split(value, ":")
			if len(svalue) > 0 {
				value = os.Getenv(svalue[0])
				if len(value) == 0 {
					value = svalue[1]
				}
			} else {
				value = GetStringValue(extractVar(vars[i]))
			}
			result = strings.Replace(result, vars[i], value, -1)
		}
	}
	return result
}

// GetIntegerValue :
func GetIntegerValue(key string) int {
	value := GetStringValue(key)
	if val, err := strconv.Atoi(value); err == nil {
		return val
	}
	log.Println(value + " is not an integer.")
	return 0
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
		log.Println("Started with profile :" + envProfile)
	} else {
		log.Println("Started with profile : default")
	}
	loadValues(filename)
}

func loadValues(filename string) {
	json, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(filename + " NOT loaded ...")
	} else {
		log.Println(filename + " loaded ...")
		j := string(json)
		fields := getFieldsName(string(j), "")
		for _, field := range fields {
			values[field] = gjson.Get(j, field).String()
		}
	}
}

func referenceVar(str string) []string {
	regex := regexp.MustCompile(`\$\{([^}]*)\}`)
	submatchall := regex.FindAllString(str, -1)
	var refer []string
	refer = make([]string, len(submatchall))
	index := 0
	for _, el := range submatchall {
		refer[index] = el
		index++
	}
	return refer
}

func extractVar(expres string) string {
	result := strings.Trim(expres, "${")
	result = strings.Trim(result, "}")
	return strings.TrimSpace(result)
}

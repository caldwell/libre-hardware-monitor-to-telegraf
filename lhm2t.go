package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	sensor_json, err := ioutil.ReadAll(os.Stdin)
	if err != nil { log.Fatal(err) }

	var tree map[string]interface{}
	err = json.Unmarshal(sensor_json, &tree)
	if err != nil { log.Fatal(err) }

	sensors := search_for_sensor(tree, []string{})
	for _, s := range(sensors) {
		badchars := regexp.MustCompile(`[^A-Za-z0-9_.°%#]`)
		tags := ""
		for k, v := range(s.Tag) {
			tags = tags + fmt.Sprintf(",%s=%s", badchars.ReplaceAllLiteralString(k, "_"), badchars.ReplaceAllLiteralString(v, "_"))
		}
		fmt.Printf("%s%s %s_%s=%s\n", badchars.ReplaceAllLiteralString(s.Measurement, "_"), tags, badchars.ReplaceAllLiteralString(s.Field, "_"), badchars.ReplaceAllLiteralString(s.Unit, "_"), s.Value)
	}
}

type Sensor struct {
	Measurement string
	Field string
	Value string
	Unit string
	Tag map[string]string
}

func parse_sensor(node map[string]interface{}, path []string) Sensor {
	v := strings.SplitN(node["Value"].(string), " ", 2)
	tag := make(map[string]string)
	if m := regexp.MustCompile(`^(.*) #(\d+)$`).FindStringSubmatch(path[len(path)-1]); m != nil {
		tag[m[1]] = m[2]
		path = append(path[0:len(path)-1], m[1])
	}
	tag["host"] = path[1]
	return Sensor{
		Tag: tag,
		Value: v[0],
		Unit: v[1],
		Measurement: strings.Join(path[2:len(path)-1], "."),
		Field: path[len(path)-1],
	}
}

func search_for_sensor(node map[string]interface{}, path []string) []Sensor {
	sensors := make([]Sensor, 0)
	path = append(path, node["Text"].(string))

	if _, ok := node["SensorId"]; ok {
		//log.Printf("%s\n", node["SensorId"])
		sensors = append(sensors, parse_sensor(node, path))
	}

	if _, ok := node["Children"]; ok {
		for _, c := range(node["Children"].([]interface{})) {
			sensors = append(sensors, search_for_sensor(c.(map[string]interface{}), path)...)
		}
	}
	return sensors
}

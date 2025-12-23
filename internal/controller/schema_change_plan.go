package controller

import (
	"ariga.io/atlas/atlasexec"
	"encoding/json"
	"log"
	"os"
	"reflect"
)

func init() {
	log.SetOutput(os.Stdout)
}

type SchemaChangePlanner struct {
	Level string
}

var (
	GChangePlanner = &SchemaChangePlanner{Level: "INFO"}
)

func (s *SchemaChangePlanner) Describe(data interface{}) {
	strData := ""
	if reflect.TypeOf(data) == reflect.TypeOf("") {
		strData = data.(string)
	} else {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return
		}
		strData = string(jsonData)
	}
	vData := []byte(strData)
	s.Write(vData)
}

func (s *SchemaChangePlanner) DescribeEx(changes atlasexec.Changes) {
	strData, _ := json.MarshalIndent(changes, "", "  ")
	s.Write(strData)
}

func (s *SchemaChangePlanner) Write(p []byte) (n int, err error) {
	log.Printf("====================SchemaPlan[%s]-Start======================", s.Level)
	log.Printf("%s", string(p))
	log.Printf("====================SchemaPlan[%s]-End========================", s.Level)
	return len(p), nil
}

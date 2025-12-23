// Copyright 2025 The Atlas Operator Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"ariga.io/atlas/atlasexec"
	"encoding/json"
	"log"
	"os"
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

func (s *SchemaChangePlanner) DescribeEx(schemaApply *atlasexec.SchemaApply) {

	type DBChange struct {
		Database string `json:"database"`
		Changes  atlasexec.Changes
	}

	change := DBChange{
		Database: schemaApply.URL.Schema,
		Changes:  schemaApply.Changes,
	}
	strData, _ := json.MarshalIndent(change, "", "  ")
	log.Printf("====================SchemaPlan[%s]-Start======================", s.Level)
	log.Printf("\n%s", string(strData))
	log.Printf("====================SchemaPlan[%s]-End========================", s.Level)
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Rule struct {
	RuleType      string        `json:"rule-type"`
	RuleId        string        `json:"rule-id"`
	RuleName      string        `json:"rule-name"`
	RuleAction    string        `json:"rule-action"`
	ObjectLocator ObjectLocator `json:"object-locator"`
	Filters       []Filter      `json:"filters"`
}

type FilterCondition struct {
	FilterOperator string `json:"filter-operator"`
	Value          string `json:"value"`
}

type Filter struct {
	FilterType       string
	ColumnName       string
	FilterConditions string
}

type ObjectLocator struct {
	SchemaName string
	TableName  string
	TableType  string
}

type Mapping struct {
	Rules []Rule
}

func main() {
	// rules := &Rele{Re: "Frank"}
	start()
}

func start() {
	mappings := createData()
	jsonString, err := format(mappings)
	if err != nil {
		fmt.Println(err)
	} else {
		// showContent(jsonString)
		fullFileName := "./Traze-Mappings-result.json"
		saveFile(jsonString, fullFileName)

	}
}
func createData() *Mapping {
	rules := []Rule{}
	objectLocatior := ObjectLocator{"name", "table", "type"}
	objectLocatior.SchemaName = "schema name"
	objectLocatior.TableName = "table name"

	filters := []Filter{}

	for i := 1; i < 3; i++ {

		filters = append(filters, Filter{"ss", "fff", "bb"})

	}

	rules = append(rules, Rule{"type", "id", "name", "RuleAction", objectLocatior, filters})
	rules = append(rules, Rule{"type2", "id22", "name2", "RuleAction333", objectLocatior, filters})
	mappings := &Mapping{rules}

	return mappings
}

func format(mappings *Mapping) (data string, err error) {
	b, err := json.Marshal(mappings)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

func showContent(data string) {
	fmt.Println(data)
}
func saveFile(data string, fullFileName string) bool {
	f, err := os.Create(fullFileName)
	showContent(data)
	if err == nil {
		f.WriteString(data)
		f.Close()
		return true
	} else {
		fmt.Println(err)
	}
	return false

}

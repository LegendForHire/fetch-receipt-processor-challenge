package configLoader

import (
	"os"
    "sync"
    "encoding/json"
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/rule"
	"github.com/legendforhire/fetch-receipt-processor-challenge/rules/centsRule"
	"github.com/legendforhire/fetch-receipt-processor-challenge/rules/dayModRule"
	"github.com/legendforhire/fetch-receipt-processor-challenge/rules/itemCountRule"
	"github.com/legendforhire/fetch-receipt-processor-challenge/rules/itemDescriptionModAndPriceRule"
	"github.com/legendforhire/fetch-receipt-processor-challenge/rules/retailerAlphaNumericCountRule"
	"github.com/legendforhire/fetch-receipt-processor-challenge/rules/timeRule"
	"fmt"
)

type Config struct {
	Rules []rule.Rule
}

var onlyConfig *Config

var lock = &sync.Mutex{}


func GetConfig() *Config{
	if onlyConfig == nil {
        lock.Lock()
        defer lock.Unlock()
        if onlyConfig == nil {
			file, err := os.Open("config.json")
			if err != nil {
				fmt.Println("Error opening file:", err)
				return nil
			}
			defer file.Close()
			// Decode the JSON data into a map
			decoder := json.NewDecoder(file)
			var config map[string]any
			err = decoder.Decode(&config)
			if err != nil {
				fmt.Println("Error decoding JSON:", err)
				return nil
			}
			constructorMap := map[string]func(map[string]any) rule.Rule{
				"centsRule": func(ruleData map[string]any) rule.Rule{ 
					cents := make([]string, len(ruleData["cents"].([]any)))
					for i := range cents{
						cents[i] = ruleData["cents"].([]any)[i].(string)
					}
					return centsRule.CentsRuleConstructor(cents, int(ruleData["points"].(float64)))
				},
				"dayModRule": func(ruleData map[string]any) rule.Rule{ return dayModRule.DayModRuleConstructor(int(ruleData["modulo"].(float64)), int(ruleData["result"].(float64)), int(ruleData["points"].(float64)))},
				"itemCountRule": func(ruleData map[string]any) rule.Rule{ return itemCountRule.ItemCountRuleConstructor(int(ruleData["threshhold"].(float64)), int(ruleData["points"].(float64)))},
				"itemDescriptionModAndPriceRule": func(ruleData map[string]any) rule.Rule{ return itemDescriptionModAndPriceRule.ItemDescriptionModAndPriceRuleConstructor(int(ruleData["modulo"].(float64)), int(ruleData["result"].(float64)), ruleData["roundUp"].(bool), ruleData["points"].(float64))},
				"retailerAlphaNumericCountRule": func(ruleData map[string]any) rule.Rule{ return retailerAlphaNumericCountRule.RetailerAlphaNumericCountRuleConstructor(int(int(ruleData["points"].(float64))))},
				"timeRule": func(ruleData map[string]any) rule.Rule{ return timeRule.TimeRuleConstructor(ruleData["startTime"].(string), ruleData["endTime"].(string), int(ruleData["points"].(float64)))},
			}
			rules := []rule.Rule{}
			for _, ruleData := range config["rules"].([]any){
				rules = append(rules, constructorMap[ruleData.(map[string]any)["type"].(string)](ruleData.(map[string]any)))
			} 
            onlyConfig = &Config{rules}
        }
    }

    return onlyConfig
}
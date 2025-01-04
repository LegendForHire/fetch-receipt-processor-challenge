package centsRule

import (
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/rule"
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/receipt"
    "strings"
)

type CentsRule struct {
	*rule.AbstractBoolRule
	cents []string
}

func CentsRuleConstructor(cents []string, points int) *CentsRule{
  a:=&rule.AbstractBoolRule{Points: points}
  r:=&CentsRule{a, cents}
  a.Rule=r
  return r
}

func (rule *CentsRule) GetCalcData(receiptData receipt.Receipt) any{
	for _, centAmt := range rule.cents {
		if(strings.HasSuffix(receiptData.Total, centAmt)){
			return true
		}
	}
	return false
}
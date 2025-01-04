package retailerAlphaNumericCountRule

import (
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/rule"
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/receipt"
)

type RetailerAlphaNumericCountRule struct {
	*rule.AbstractCountRule
}

func RetailerAlphaNumericCountRuleConstructor(points int) *RetailerAlphaNumericCountRule{
  a:=&rule.AbstractCountRule{Points: points}
  r:=&RetailerAlphaNumericCountRule{a}
  a.Rule=r
  return r
}

func (rule *RetailerAlphaNumericCountRule) GetCalcData(receiptData receipt.Receipt) any{
	n := 0
	for _, b := range receiptData.Retailer {
		if ('a' <= b && b <= 'z') ||
            ('A' <= b && b <= 'Z') ||
            ('0' <= b && b <= '9'){
				n++
			}
	}
	return n;
}
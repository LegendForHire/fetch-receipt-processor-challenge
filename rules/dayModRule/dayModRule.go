package dayModRule

import (
	  "github.com/legendforhire/fetch-receipt-processor-challenge/model/rule"
	  "github.com/legendforhire/fetch-receipt-processor-challenge/model/receipt"
)

type DayModRule struct {
	*rule.AbstractBoolRule
	modulo int
	result int
}

func DayModRuleConstructor(modulo int, result int, points int) *DayModRule{
  a:=&rule.AbstractBoolRule{Points: points}
  r:=&DayModRule{a, modulo, result}
  a.Rule=r
  return r
}

func (rule *DayModRule) GetCalcData(receiptData receipt.Receipt) any{
  daySuffixNum := int(receiptData.PurchaseDate[len(receiptData.PurchaseDate)-1]-0)
	return (daySuffixNum%rule.modulo == rule.result)
}
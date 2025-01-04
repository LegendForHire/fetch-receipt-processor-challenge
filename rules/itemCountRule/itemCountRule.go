package itemCountRule

import (
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/rule"
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/receipt"
)

type ItemCountRule struct {
	*rule.AbstractCountRule
	threshhold int
}

func ItemCountRuleConstructor(threshhold int, points int) *ItemCountRule{
  a:=&rule.AbstractCountRule{Points: points}
  r:=&ItemCountRule{a, threshhold}
  a.Rule=r
  return r
}

func (rule *ItemCountRule) GetCalcData(receiptData receipt.Receipt) any{
	return len(receiptData.Items)/rule.threshhold
}
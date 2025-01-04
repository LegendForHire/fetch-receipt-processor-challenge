package itemDescriptionModAndPriceRule

import (
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/rule"
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/receipt"
    "strconv"
)

type ItemDescriptionModAndPriceRule struct {
	*rule.AbstractMultRule
	modulo int
	result int
}

func ItemDescriptionModAndPriceRuleConstructor(modulo int, result int, roundUp bool, points float64) *ItemDescriptionModAndPriceRule{
  a:=&rule.AbstractMultRule{RoundUp: roundUp, Points: points}
  r:=&ItemDescriptionModAndPriceRule{a, modulo, result}
  a.Rule=r
  return r
}

func (rule *ItemDescriptionModAndPriceRule) GetCalcData(receiptData receipt.Receipt) any{	
	prices := []float64{}
	for _, item := range receiptData.Items {
		if(len(item.Description)%rule.modulo == rule.result){
			price, _ := strconv.ParseFloat(item.Price, 64)
			prices = append(prices, price)
		}
	}
	return prices
}
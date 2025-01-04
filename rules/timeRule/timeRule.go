package timeRule

import (
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/rule"
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/receipt"
)

type TimeRule struct {
	*rule.AbstractBoolRule
	startTime string
	endTime string
}

func TimeRuleConstructor(startTime string, endTime string, points int) *TimeRule{
  a:=&rule.AbstractBoolRule{Points: points}
  r:=&TimeRule{a, startTime, endTime}
  a.Rule=r
  return r
}

func (rule *TimeRule) GetCalcData(receiptData receipt.Receipt) any{
	afterST := false
	beforeET := false
	for i, b := range []byte(receiptData.PurchaseTime) {
		if(b != ':'){
			
			if((b < rule.startTime[i] && !afterST) || (b > rule.endTime[i]&& !beforeET)){
				return false
			} else{
				afterST= afterST || b > rule.startTime[i]
				beforeET= beforeET || b < rule.endTime[i]
				if(afterST&&beforeET){
					return true;
				}
			}
		}
	}
	return false;
}
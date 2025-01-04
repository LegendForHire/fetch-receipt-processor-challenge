package rule 

import (
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/receipt"
    "math"
)

type Rule interface {
	GetCalcData(receiptData receipt.Receipt) any
	CalcPoints(receiptData receipt.Receipt) int
}

type AbstractBoolRule struct {
	Points int
	Rule
}
type AbstractCountRule struct {
	Points int
	Rule
}
type AbstractMultRule struct {
	RoundUp bool
	Points float64
	Rule
}

func (rule *AbstractBoolRule) CalcPoints(receiptData receipt.Receipt) int{
	var calcData bool = rule.GetCalcData(receiptData).(bool)
	if(calcData) {
		return rule.Points
	}
	return 0
}
func (rule *AbstractCountRule) CalcPoints(receiptData receipt.Receipt) int{
	var calcData int = rule.GetCalcData(receiptData).(int)
	return rule.Points*calcData
}
func (rule *AbstractMultRule) CalcPoints(receiptData receipt.Receipt) int{
	calculatedPoints := 0
	var calcData []float64 = rule.GetCalcData(receiptData).([]float64 )
	for _, element := range calcData  {
		if rule.RoundUp {
			calculatedPoints += int(math.Ceil(element*rule.Points))
		} else {
			calculatedPoints += int(math.Floor(element*rule.Points))
		}
	}
	return calculatedPoints
}
package api

import (
	"net/http"
  	"github.com/gin-gonic/gin"
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/inMemDB"
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/configLoader"
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/receipt"
	"github.com/legendforhire/fetch-receipt-processor-challenge/model/rule"
	"github.com/google/uuid"
	"sync"
)

func ProcessReceipt(c *gin.Context){
	receiptData := receipt.Receipt{} 
	err := c.ShouldBindBodyWithJSON(&receiptData); 
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
    }
	config := configLoader.GetConfig()
	var wg sync.WaitGroup
	wg.Add(len(config.Rules))
	var lock = &sync.Mutex{}
	points := 0
	for _, aRule := range config.Rules{
		go func(ruleData rule.Rule){
			defer wg.Done()
			rulePoints := ruleData.CalcPoints(receiptData)
			lock.Lock()
			defer lock.Unlock()
			points+=rulePoints
		}(aRule)
	}
	wg.Wait()
	id := uuid.New().String()
	database := inMemDB.GetInstance()
	database.PointMap[id] = points
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func GetReceiptPoints(c *gin.Context){
	database := inMemDB.GetInstance()
	id := c.Param("id")
	points, ok := database.PointMap[id]
	if(ok){
		c.JSON(http.StatusOK, gin.H{
			id: points,
		})
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
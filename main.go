package main

import (
  "github.com/gin-gonic/gin"
  "github.com/legendforhire/fetch-receipt-processor-challenge/api"
)

func main(){
	r := gin.Default()
	r.GET("receipts/:id/points", api.GetReceiptPoints)
	r.POST("receipts/process", api.ProcessReceipt)
	r.Run()
}
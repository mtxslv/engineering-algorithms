package main

import (
	"fmt"
	"utils/utils"
)

func main(){
	// load fruits and prices
	fruitsAndPrices := utils.LoadCsv("./fruits-prices.csv")
	for _, item := range(fruitsAndPrices){
		fmt.Printf("%+v\n", item)
	} 

}
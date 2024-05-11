package useFieldTagsInMarshaledStructs

import "encoding/json"

type StockB struct {
	Price int
	Name  string
}

var structInstance = StockB{
	Price: 137,
	Name:  "UBER",
}

func stockUsageExBad() {
	bytes, err := json.Marshal(StockB{
		Price: 137,
		Name:  "UBER",
	})
	print(bytes, err)
}

func stockUsageExBad2() {
	bytes, err := json.Marshal(structInstance)
	print(bytes, err)
}

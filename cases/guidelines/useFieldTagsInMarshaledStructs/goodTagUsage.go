package useFieldTagsInMarshaledStructs

import "encoding/json"

type StockG struct {
	Price int    `json:"price"`
	Name  string `json:"name"`
	// Safe to rename Name to Symbol.
}

var structInstanceGood = StockG{
	Price: 137,
	Name:  "UBER",
}

func stockUsageExGood() {
	bytes, err := json.Marshal(StockG{
		Price: 137,
		Name:  "UBER",
	})
	print(bytes, err)
}

func stockUsageExGood2() {
	bytes, err := json.Marshal(structInstanceGood)
	print(bytes, err)
}

# types

extend the golang types for json/xml Marshal and Unmarshal


### Amount
```go
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/gogap/types"
)

type Goods struct {
	Amount types.Amount
}

func main() {
	fmt.Println("Amount(json)")
	jsonAmount()
	fmt.Println("Amount(xml)")
	xmlAmount()
}

func jsonAmount() {
	goods := Goods{Amount: 123}
	if x, e := json.MarshalIndent(goods, "", " "); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(string(x))
	}

	strJson := `{"Amount":"000000000000456"}`
	goods2 := Goods{}
	if e := json.Unmarshal([]byte(strJson), &goods2); e != nil {
		fmt.Println(goods2, e)
	} else {
		fmt.Println(goods2)
	}
}

func xmlAmount() {
	goods := Goods{Amount: 123}
	if x, e := xml.MarshalIndent(goods, "", " "); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(string(x))
	}

	strJson := `<Goods><Amount>000000000000456</Amount></Goods>`
	goods2 := Goods{}
	if e := xml.Unmarshal([]byte(strJson), &goods2); e != nil {
		fmt.Println(goods2, e)
	} else {
		fmt.Println(goods2)
	}
}

```
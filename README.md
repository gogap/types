# types

extend the golang types for json/xml Marshal and Unmarshal

- Amount:	000000001000001
- Date: 	2015-01-02
- Time: 	13:14:55
- DateTime: 2015-01-02 13:14:55

### Example
```go
package main

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/gogap/logs"
	"github.com/gogap/types"
)

type Goods struct {
	Amount     types.Amount
	CreateDate types.Date
	DateTime   types.DateTime
	Time       types.Time
}

func main() {
	jsonExample()
	xmlExample()

	time.Sleep(time.Second)
}

func jsonExample() {
	goods := Goods{
		Amount:     123,
		CreateDate: types.Date(time.Now()),
		DateTime:   types.DateTime(time.Now()),
		Time:       types.Time(time.Now())}

	if x, e := json.MarshalIndent(goods, "", " "); e != nil {
		logs.Error(e)
	} else {
		logs.Debug(string(x))
	}

	strJson := `{"Amount":"000000000000456", 
				"CreateDate":"2015-02-11",
				"DateTime":"2015-02-11 11:12:13",
				"Time":"21:12:13"}`

	goods2 := Goods{}
	if e := json.Unmarshal([]byte(strJson), &goods2); e != nil {
		logs.Error(e)
	} else {
		logs.Pretty("", goods2)
	}
}

func xmlExample() {
	goods := Goods{
		Amount:     123,
		CreateDate: types.Date(time.Now()),
		DateTime:   types.DateTime(time.Now()),
		Time:       types.Time(time.Now())}

	if x, e := xml.MarshalIndent(goods, "", " "); e != nil {
		logs.Error(e)
	} else {
		logs.Debug(string(x))
	}

	strJson := `<Goods>
				<Amount>000000000000456</Amount>
				<CreateDate>2015-02-11</CreateDate>
				<DateTime>2015-02-11 11:12:13</DateTime>
				<Time>20:12:13</Time>
				</Goods>`

	goods2 := Goods{}
	if e := xml.Unmarshal([]byte(strJson), &goods2); e != nil {
		logs.Error(e)
	} else {
		logs.Pretty("", goods2)
	}
}
```
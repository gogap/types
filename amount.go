package types

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

const (
	TYPES_AMOUNT_MARSHAL_FORMAT = "%015d"
	TYPES_AMOUNT_LENGTH         = 15
)

type Amount int64

func (i Amount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(fmt.Sprintf(TYPES_AMOUNT_MARSHAL_FORMAT, i), start)
	return nil
}

func (i *Amount) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var content string
	if e := d.DecodeElement(&content, &start); e != nil {
		return fmt.Errorf("get the type Amount field of %s error", start.Name.Local)
	}

	if len(content) != TYPES_AMOUNT_LENGTH {
		return fmt.Errorf("the type Amount field of %s length error", start.Name.Local)
	}

	if v, e := strconv.Atoi(content); e != nil {
		return fmt.Errorf("the type Amount field of %s is not a number, value is: %s", start.Name.Local, content)
	} else {
		*i = Amount(v)
	}
	return nil
}

func (p Amount) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("\""+TYPES_AMOUNT_MARSHAL_FORMAT+"\"", p)
	return []byte(str), nil
}

func (p *Amount) UnmarshalJSON(text []byte) (err error) {
	if len(text) != TYPES_AMOUNT_LENGTH+2 {
		err = fmt.Errorf("amount value length is not 15, error value is: %s", string(text))
		return
	}

	strAmount := string(text[1 : TYPES_AMOUNT_LENGTH+1])

	if v, e := strconv.Atoi(strAmount); e != nil {
		return fmt.Errorf("amount should be a number, error value is: %s", strAmount)
	} else {
		*p = Amount(v)
	}
	return nil
}

package types

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type DoubleAmount int64

func (p DoubleAmount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeElement(fmt.Sprintf(TYPES_AMOUNT_MARSHAL_FORMAT, p), start)
	return nil
}

func (p *DoubleAmount) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var content string
	if e := d.DecodeElement(&content, &start); e != nil {
		return fmt.Errorf("get the type DoubleAmount field of %s error", start.Name.Local)
	}

	if len(content) != TYPES_AMOUNT_LENGTH {
		return fmt.Errorf("the type DoubleAmount field of %s length error", start.Name.Local)
	}

	if v, e := strconv.Atoi(content); e != nil {
		return fmt.Errorf("the type DoubleAmount field of %s is not a number, value is: %s", start.Name.Local, content)
	} else {
		*p = DoubleAmount(v)
	}
	return nil
}

func (p DoubleAmount) MarshalJSON() ([]byte, error) {
	str := strconv.FormatFloat(float64(p)/100.00, 'f', 2, 64)
	return []byte("\"" + str + "\""), nil
}

func (p *DoubleAmount) UnmarshalJSON(text []byte) (err error) {
	strText := strings.Trim(string(text), "\"")
	strText = strings.TrimSpace(strText)

	if strText == "" {
		err = fmt.Errorf("amount value is empty")
		return
	}

	lenText := len(strText)
	if strText[lenText-3] != '.' {
		err = fmt.Errorf("amount value should have 2 decimal number, error value is: %s", strText)
		return
	}

	strDoubleAmount := strings.Replace(strText, ".", "", 1)

	if v, e := strconv.Atoi(strDoubleAmount); e != nil {
		return fmt.Errorf("convert double amount value into int64 error, error value is: %s, err: %s", text, e)
	} else {
		*p = DoubleAmount(v)
	}
	return nil
}

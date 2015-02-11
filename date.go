package types

import (
	"encoding/xml"
	"fmt"
	"time"
)

const (
	TYPES_DATE_LAYOUT = "2006-01-02"
	TYPES_DATE_LENGTH = len(TYPES_DATE_LAYOUT)
)

type Date time.Time

func (p Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(p)
	str := t.Format(TYPES_DATE_LAYOUT)
	e.EncodeElement(str, start)
	return nil
}

func (p *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var content string
	if e := d.DecodeElement(&content, &start); e != nil {
		return fmt.Errorf("get the type Date field of %s error", start.Name.Local)
	}

	if len(content) != TYPES_DATE_LENGTH {
		return fmt.Errorf("the type Date field of %s length error", start.Name.Local)
	}

	if v, e := time.Parse(TYPES_DATE_LAYOUT, content); e != nil {
		return fmt.Errorf("the type Date field of %s is not a time, value is: %s", start.Name.Local, content)
	} else {
		*p = Date(v)
	}
	return nil
}

func (p Date) MarshalJSON() ([]byte, error) {
	t := time.Time(p)
	str := "\"" + t.Format(TYPES_DATE_LAYOUT) + "\""

	return []byte(str), nil
}

func (p *Date) UnmarshalJSON(text []byte) (err error) {
	strDate := string(text[1 : TYPES_DATE_LENGTH+1])

	if v, e := time.Parse(TYPES_DATE_LAYOUT, strDate); e != nil {
		return fmt.Errorf("Date should be a time, error value is: %s", strDate)
	} else {
		*p = Date(v)
	}
	return nil
}

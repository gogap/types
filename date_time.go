package types

import (
	"encoding/xml"
	"fmt"
	"time"
)

const (
	TYPES_DATE_TIME_LAYOUT = "2006-01-02 15:04:05"
	TYPES_DATE_TIME_LENGTH = len(TYPES_DATE_TIME_LAYOUT)
)

type DateTime time.Time

func (p DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(p)
	str := t.Format(TYPES_DATE_TIME_LAYOUT)
	e.EncodeElement(str, start)
	return nil
}

func (p *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var content string
	if e := d.DecodeElement(&content, &start); e != nil {
		return fmt.Errorf("get the type DateTime field of %s error", start.Name.Local)
	}

	if len(content) != TYPES_DATE_TIME_LENGTH {
		return fmt.Errorf("the type DateTime field of %s length error", start.Name.Local)
	}

	if v, e := time.Parse(TYPES_DATE_TIME_LAYOUT, content); e != nil {
		return fmt.Errorf("the type DateTime field of %s is not a time, value is: %s", start.Name.Local, content)
	} else {
		*p = DateTime(v)
	}
	return nil
}

func (p DateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(p)
	str := "\"" + t.Format(TYPES_DATE_TIME_LAYOUT) + "\""

	return []byte(str), nil
}

func (p *DateTime) UnmarshalJSON(text []byte) (err error) {
	strDateTime := string(text[1 : TYPES_DATE_TIME_LENGTH+1])

	if v, e := time.Parse(TYPES_DATE_TIME_LAYOUT, strDateTime); e != nil {
		return fmt.Errorf("DateTime should be a time, error value is: %s", strDateTime)
	} else {
		*p = DateTime(v)
	}
	return nil
}

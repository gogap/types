package types

import (
	"encoding/xml"
	"fmt"
	"time"
)

const (
	TYPES_TIME_LAYOUT = "15:04:05"
	TYPES_TIME_LENGTH = len(TYPES_TIME_LAYOUT)
)

type Time time.Time

func (p Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(p)
	str := t.Format(TYPES_TIME_LAYOUT)
	e.EncodeElement(str, start)
	return nil
}

func (p *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var content string
	if e := d.DecodeElement(&content, &start); e != nil {
		return fmt.Errorf("get the type Time field of %s error", start.Name.Local)
	}

	if len(content) != TYPES_TIME_LENGTH {
		return fmt.Errorf("the type Time field of %s length error", start.Name.Local)
	}

	if v, e := time.Parse(TYPES_TIME_LAYOUT, content); e != nil {
		return fmt.Errorf("the type Time field of %s is not a time, value is: %s", start.Name.Local, content)
	} else {
		*p = Time(v)
	}
	return nil
}

func (p Time) MarshalJSON() ([]byte, error) {
	t := time.Time(p)
	str := "\"" + t.Format(TYPES_TIME_LAYOUT) + "\""

	return []byte(str), nil
}

func (p *Time) UnmarshalJSON(text []byte) (err error) {
	strTime := string(text[1 : TYPES_TIME_LENGTH+1])

	if v, e := time.Parse(TYPES_TIME_LAYOUT, strTime); e != nil {
		return fmt.Errorf("Time should be a time, error value is: %s", strTime)
	} else {
		*p = Time(v)
	}
	return nil
}

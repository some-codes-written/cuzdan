package exchangerates

import (
	"encoding/xml"
	"time"
)

func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const shortform = "20060102"
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(shortform, v)
	if err != nil {
		return err
	}
	*c = customTime{parse}
	return nil
}

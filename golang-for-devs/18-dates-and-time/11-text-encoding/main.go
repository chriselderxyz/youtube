package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// Text Encoding/Decoding
/////////////////////////////////////////////////////////////

type Event struct {
	At time.Time `xml:"at"`
}

func main() {

	str := `<Event><at>2037-04-22T18:09:34+02:00</at></Event>`

	// Create Event from XML string
	e := &Event{}
	err := xml.Unmarshal([]byte(str), e)
	if err != nil {
		fmt.Println("err: ", err)
	}

	// Create XML string from Event
	marshaledBytes, err := xml.Marshal(e)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println("e.At: ", e.At)
	fmt.Println("xml.Marshal result: ", string(marshaledBytes))
}

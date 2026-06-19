package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/////////////////////////////////////////////////////////////
// JSON Encoding/Decoding
/////////////////////////////////////////////////////////////

type Event struct {
	At time.Time `json:"at"`
}

func main() {
	str := `{"at":"2037-04-22T18:09:34+02:00"}`

	// Create Event with json string
	e := &Event{}
	err := json.Unmarshal([]byte(str), e)
	if err != nil {
		fmt.Println("err: ", err)
	}

	// Create json string with Event
	bytes, err := json.Marshal(e)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println("e.At: ", e.At)
	fmt.Println("json.Marshal result: ", string(bytes))
}

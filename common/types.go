package common

import (
	"encoding/json"
	"time"
)

type MyTime struct {
	time.Time
}

func (i *MyTime) UnmarshalJSON(data []byte) error {

	var timestamp int64
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err // Handle any initial unmarshaling errors
	}

	sec := timestamp / 1000          // Milliseconds to seconds
	nsec := (timestamp % 1000) * 1e6 // Remaining nanoseconds
	i.Time = time.Unix(sec, nsec)

	return nil
}

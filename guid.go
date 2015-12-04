package nd

import (
	"strconv"
	"time"
)

var Guidv4 func() []byte

func init() {
	ResetGuidv4()
}

func ResetGuidv4() {
	Guidv4 = func() []byte {
		return []byte(strconv.FormatInt(time.Now().UnixNano(), 10))
	}
}

func Guidv4String() string {
	t := Guidv4()
	return string(t)
}

func ForceGuid(guid string) error {
	Guidv4 = func() []byte { return []byte(guid) }
	return nil
}

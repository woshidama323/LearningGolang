package customMarshal

import (
	"strconv"
	"time"
)

type TestJson struct {
	Testitme UnixTimestamp
}

type UnixTimestamp time.Time

func (ut UnixTimestamp) MarshalJSON() ([]byte, error) {
	s := strconv.Itoa(int(time.Time(ut).Unix()))
	return []byte(s), nil
}

func (ut *UnixTimestamp) UnmarshalJSON(data []byte) error {
	unix, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	*ut = UnixTimestamp(time.Unix(int64(unix), 0))
	return nil
}

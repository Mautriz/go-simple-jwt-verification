package helpers

import (
	"time"
)

// CurrTimestampMS ritorna il timestamp corrente in millisecondi (come javascript)
func CurrTimestampMS() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// CurrTimestampSecond ritorna il timestamp corrente in millisecondi (come javascript)
func CurrTimestampSecond() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}
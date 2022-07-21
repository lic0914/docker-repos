package utils
import (
	"time"
	"strconv"
)
func GetTs() string{
	t := time.Now() // .Unix() 秒
    ts := t.UnixNano() / 1e6 //毫秒
	return strconv.FormatInt(ts,10)
}
package helper

import (
	"errors"
	"strings"
	"time"
)

// Str2TimeParse 将字符串转换为时间结构
func (tk *TsTime) Str2TimeParse(str string, format ...string) (time.Time, error) {
	f := ""
	if len(format) > 0 {
		f = strings.Trim(format[0], " ")
	} else {
		f = "2006-01-02 15:04:05"
	}
	if len(str) != len(f) {
		return time.Now(), errors.New("parameter format error")
	}
	return time.Parse(f, str)
}

// Time 获取当前Unix时间戳(秒)
func (tk *TsTime) Time() int64 {
	return time.Now().Unix()
}

// MicroTime 获取当前Unix时间戳(微秒).
func (tk *TsTime) MicroTime() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// ServiceUptime 纳秒
func (tk *TsTime) ServiceUptime() time.Duration {
	return time.Since(NowTime)
}


// Str2Timestamp 将字符串转换为时间戳,秒.
// str 为要转换的字符串;
// format 为该字符串的格式,默认为"2006-01-02 15:04:05" .
func (tk *TsTime) Str2Timestamp(str string, format ...string) (int64, error) {
	tim, err := tk.Str2TimeParse(str, format...)
	if err != nil {
		return 0, err
	}

	return tim.Unix(), nil
}
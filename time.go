package helper

import (
	"errors"
	"strings"
	"time"
)

// Str2TimeParse 将字符串转换为时间结构.
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

// Time 获取当前Unix时间戳(秒).
func (tk *TsTime) Time() int64 {
	return time.Now().Unix()
}

// MicroTime 获取当前Unix时间戳(微秒).
func (tk *TsTime) MicroTime() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// ServiceUptime 纳秒.
func (tk *TsTime) ServiceUptime() time.Duration {
	return time.Since(NowTime)
}


// Str2Timestamp 将字符串转换为时间戳,秒.
// str 为要转换的字符串;
// format 为该字符串的格式,默认为"2006-01-02 15:04:05".
func (tk *TsTime) Str2Timestamp(str string, format ...string) (int64, error) {
	tim, err := tk.Str2TimeParse(str, format...)
	if err != nil {
		return 0, err
	}

	return tim.Unix(), nil
}

// GetMonthDays 获取指定年月的天数, 默认当前年份
func (tk *TsTime) GetMonthDays(month int, years ...int) int {
	months := map[int]int{1: 31, 3: 31, 4: 30, 5: 31, 6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 31}

	if days, ok := months[month]; ok {
		return days
	} else if month < 1 || month > 12 {
		return 0
	}

	var year int
	yLen := len(years)
	if yLen == 0 {
		year = time.Now().Year()
	} else {
		year = years[0]
	}

	if year%100 == 0 {
		if year%400 == 0 {
			return 29
		} else {
			return 28
		}
	} else if year%4 == 0 {
		return 29
	} else {
		return 28
	}
}
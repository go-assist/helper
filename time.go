package helper

import (
	"errors"
	"strings"
	"time"
)

// 这个时间函数的包也不错, github.com/golang-module/carbon

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
func (tk *TsTime) Str2Timestamp(str string, format ...string) (t int64, e error) {
	tim, err := tk.Str2TimeParse(str, format...)
	if err != nil {
		return
	}
	t = tim.Unix()
	return
}

// GetMonthDays 获取指定年月的天数, 默认当前年份.
func (tk *TsTime) GetMonthDays(month int, years ...int) (d int) {
	months := map[int]int{1: 31, 3: 31, 4: 30, 5: 31, 6: 30, 7: 31, 8: 31, 9: 30, 10: 31, 11: 30, 12: 31}

	if days, ok := months[month]; ok {
		d = days
		return
	}

	if month < 1 || month > 12 {
		return
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
			d = 29
			return
		} else {
			d = 28
			return
		}
	} else if year%4 == 0 {
		d = 29
		return
	} else {
		d = 28
		return
	}
}
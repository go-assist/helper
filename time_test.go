package helper

import (
	"fmt"
	"testing"
)

func TestStr2TimeParse(t *testing.T)  {
	timeStr := "2021-08-30 11:27:00"
	_, err := TTime.Str2TimeParse(timeStr)
	if err != nil {
		t.Errorf("The error %v \n", err)
	}

	_, err = TTime.Str2TimeParse(timeStr, "2006-01-02 15:04:05")
	if err != nil {
		t.Errorf("The error %v \n", err)
	}
}

func TestTime(t *testing.T) {
	ti := fmt.Sprintf("%d", TTime.Time())
	if len(ti) != 10 {
		t.Error("Time unit test fail")
		return
	}
}

func TestMicroTime(t *testing.T) {
	ti := fmt.Sprintf("%d", TTime.MicroTime())
	if len(ti) != 16 {
		t.Error("MicroTime unit test fail")
		return
	}
}

func TestServiceUptime(t *testing.T) {
	res := TTime.ServiceUptime()
	if int64(res) <= 0 {
		t.Error("ServiceUptime unit test fail")
		return
	}
}

func TestStr2Timestamp(t *testing.T) {
	ti, err := TTime.Str2Timestamp("2019-07-11 10:11:23")
	if err != nil || ti <= 0 {
		t.Error("Str2Timestamp unit test fail")
		return
	}

	_, err = TTime.Str2Timestamp("02/01/2016 15:04:05")
	if err == nil {
		t.Error("Str2Timestamp unit test fail")
		return
	}

	_, err = TTime.Str2Timestamp("2020-02-01 13:39:36", "2019-07- 11 10: 11:23")
	if err == nil {
		t.Error("Str2Timestamp unit test fail")
		return
	}
}

func TestGetMonthDays(t *testing.T) {
	days28 := TTime.GetMonthDays(2)
	days31 := TTime.GetMonthDays(3)
	days29 := TTime.GetMonthDays(2, 2000)
	if days28 != 28 || days31 != 31 || days29 != 29 {
		t.Error("GetMonthDays unit test fail")
	}
}
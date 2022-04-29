package helper

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestByteToString(t *testing.T) {
	b := []byte("123abc")
	s := TStr.ByteToString(b)
	got := s
	want := `123abc`
	if !reflect.DeepEqual(got, want) {
		t.Errorf("The values of %v is not %v\n", got, want)
	}
}

func TestConvString(t *testing.T) {
	// int
	i := 8
	got8 := strconv.Itoa(i)
	want8 := TStr.ConvString(i)
	if !reflect.DeepEqual(got8, want8) {
		t.Errorf("The values of int %v is not %v\n", got8, want8)
	}

	// unit64
	var n uint64 = 8
	got64u := strconv.Itoa(int(n))
	want64u := TStr.ConvString(n)
	if !reflect.DeepEqual(got64u, want64u) {
		t.Errorf("The values of unit64 %v is not %v\n", got64u, want64u)
	}

	// int64
	var ts int64 = 8
	got64 := strconv.FormatInt(ts, 10)
	want64 := TStr.ConvString(ts)
	if !reflect.DeepEqual(got64, want64) {
		t.Errorf("The values of int64 %v is not %v\n", got64, want64)
	}

	// float
	f := 3.1415
	got64f := strconv.FormatFloat(f, 'f', -1, 64)
	want64f := TStr.ConvString(f)
	if !reflect.DeepEqual(got64f, want64f) {
		t.Errorf("The values of float64 %v is not %v\n", got64f, want64f)
	}

	// string
	s := `is string`
	wants := TStr.ConvString(s)
	if !reflect.DeepEqual(s, wants) {
		t.Errorf("The values of string %v is not %v\n", s, wants)
	}

	// other
	var o bool
	want0 := TStr.ConvString(o)
	if reflect.DeepEqual(o, want0) {
		t.Errorf("The values of bool %v is not %v\n", o, wants)
	}

	ss := ""
	wantSS := TStr.ConvString(ss)
	if !reflect.DeepEqual(ss, wantSS) {
		t.Errorf("The values of empty string %v is not %v\n", ss, wantSS)
	}
}

func TestBase64Encode(t *testing.T)  {
	s := 1234
	s1 := strconv.Itoa(s)
	gotInt := TStr.Base64Encode(s1)
	wantInt := `MTIzNA==`
	if gotInt != wantInt {
		t.Errorf("The int values of %v is not %v\n", gotInt, wantInt)
	}

	s2 := `1234abc`
	gotStr := TStr.Base64Encode(s2)
	wantStr := `MTIzNGFiYw==`
	if gotStr != wantStr {
		if gotStr != wantStr {
			t.Errorf("The str values of %v is not %v\n", gotStr, wantStr)
		}
	}
}

func TestStructName(t *testing.T) {
	e := `Example`
	structName := TStr.StructName(&Example{})
	if e != structName {
		t.Errorf("The struct name values of %v is not %v\n", structName, e)
	}
}

func TestCalculatePercentage(t *testing.T) {
	e1 := `50%`
	p1 := TStr.CalculatePercentage(1, 2)
	if e1 != p1 {
		t.Errorf("The percentage values of %v is not %v\n", e1, p1)
	}

	e2 := `33.33%`
	p2 := TStr.CalculatePercentage(1, 3)

	if e2 != p2 {
		t.Errorf("The percentage values of %v is not %v\n", e2, p2)
	}

	e3 := `0.00%`
	p3 := TStr.CalculatePercentage(1, 0)
	if e3 != p3 {
		t.Errorf("The percentage values of %v is not %v\n", e3, p3)
	}
}

func TestRandStringRunes(t *testing.T) {
	n1 := 1
	s1 := TStr.RandStringRunes(n1)
	if len(s1) != n1 {
		t.Errorf("The round len values of %v is not %v\n", n1, len(s1))
	}
}

func TestRandomString(t *testing.T) {
	var n1 uint8 = 4
	s1 := TStr.RandomString(n1, 1)

	if len(s1) != 4 {
		t.Errorf("The round type values of 1 len of %v is not %v\n", n1, len(s1))
	}

	s2 := TStr.RandomString(4, 2)
	if len(s2) != 4 {
		t.Errorf("The round type values of values2 len of %v is not %v\n", n1, len(s2))
	}

	s3 := TStr.RandomString(4, 3)
	if len(s3) != 4 {
		t.Errorf("The round type values of 3 len of %v is not %v\n", n1, len(s3))
	}

	s4 := TStr.RandomString(4, 4)
	if len(s3) != 4 {
		t.Errorf("The round type values of 4 len of %v is not %v\n", n1, len(s4))
	}

	s5 := TStr.RandomString(4, 5)
	if len(s3) != 4 {
		t.Errorf("The round type values of 5 len of %v is not %v\n", n1, len(s5))
	}
}

func TestGetBetweenStr(t *testing.T) {
	s := "@abc456"
	e := TStr.GetBetweenStr(`@abc456%`, "@", "%")
	if s != e {
		t.Errorf("The GetBetweenStr values of %v is not %v\n", e, s)
	}

	e1 := TStr.GetBetweenStr(`vcx@abc456%aio`, "@", "%")
	if s != e1 {
		t.Errorf("The GetBetweenStr values of %v is not %v\n", e1, s)
	}

	e2 := TStr.GetBetweenStr(`@abc456%aio`, "@", "%")
	if s != e1 {
		t.Errorf("The GetBetweenStr values of %v is not %v\n", e2, s)
	}

	e3 := TStr.GetBetweenStr(`aaaa@abc456%`, "@", "%")
	if s != e3 {
		t.Errorf("The GetBetweenStr values of %v is not %v\n", e3, s)
	}

	e4 := TStr.GetBetweenStr(`@abc456%`, "@", "%")
	if s != e3 {
		t.Errorf("The GetBetweenStr values of %v is not %v\n", e4, s)
	}
}

func TestSubstr(t *testing.T) {
	s := `a`
	e := TStr.Substr(`abc@123&`, 0, 1)
	if s != e {
		t.Errorf("The Substr values of %v is not %v\n", e, s)
	}

	s1 := `3&`
	e1 := TStr.Substr(`abc@123&`, -1, 2)
	if s1 != e1 {
		t.Errorf("The Substr values of %v is not %v\n", e1, s1)
	}

	s2 := `abc@123&`
	e2 := TStr.Substr(`abc@123&`, 0, 9999)
	if s2 != e2 {
		t.Errorf("The Substr values of %v is not %v\n", e2, s2)
	}
}

func TestSubStrLeftOrRight(t *testing.T) {
	s1 := `abc中国`
	t1 := `abc`
	right1, err := TStr.SubStrLeftOrRight(s1, t1, "right", true)
	if err != nil {
		t.Errorf("The SubStrLeftOrRight errors %v\n", err)
	}
	if right1 != s1 {
		t.Errorf("The values of %v is not %v\n", right1, s1)
	}

	right2, err := TStr.SubStrLeftOrRight(s1, t1, "right", false)
	if err != nil {
		t.Errorf("The SubStrLeftOrRight errors %v\n", err)
	}
	s2 := `bc中国`
	if right2 != s2 {
		t.Errorf("The values of %v is not %v\n", right2, s2)
	}

	left1, err := TStr.SubStrLeftOrRight(s1, t1, "left", true)
	if err != nil {
		t.Errorf("The SubStrLeftOrRight errors %v\n", err)
	}
	s3 := `a`
	if left1 != s3 {
		t.Errorf("The values of %v is not %v\n", left1, s3)
	}

	left2, err := TStr.SubStrLeftOrRight(s1, t1, "left", false)
	if err != nil {
		t.Errorf("The SubStrLeftOrRight errors %v\n", err)
	}
	s4 := ``
	if left2 != s4 {
		t.Errorf("The values of %v is not %v\n", left2, s4)
	}
}

func TestMD5(t *testing.T) {
	s1 := `123456`
	m1 := `E10ADC3949BA59ABBE56E057F20F883E`
	md51 := TStr.MD5(s1)
	if m1 != md51 {
		t.Errorf("The values of %v is not %v\n", md51, m1)
	}
	s2 := ``
	m2 := `D41D8CD98F00B204E9800998ECF8427E`
	md52 := TStr.MD5(s2)
	if m2 != md52 {
		t.Errorf("The values of %v is not %v\n", md52, m2)
	}
}

func TestUcFirst(t *testing.T) {
	s1 := ""
	u1 := ""
	uc1 := TStr.UcFirst(s1)
	if u1 != uc1 {
		t.Errorf("The values of %v is not %v\n", uc1, u1)
	}

	s2 := "abc"
	u2 := "Abc"
	uc2 := TStr.UcFirst(s2)
	if u2 != uc2 {
		t.Errorf("The values of %v is not %v\n", uc2, u2)
	}

	s3 := "Abc"
	uc3 := TStr.UcFirst(s3)
	if s3 != uc3 {
		t.Errorf("The values of %v is not %v\n", uc3, s3)
	}
}

func TestLcFirst(t *testing.T) {
	s1 := ""
	l1 := ""
	lc1 := TStr.LcFirst(s1)
	if s1 != lc1 {
		t.Errorf("The values of %v is not %v\n", lc1, l1)
	}

	s2 := "abc"
	l2 := "Abc"
	lc2 := TStr.LcFirst(l2)
	if s2 != lc2 {
		t.Errorf("The values of %v is not %v\n", lc2, s2)
	}

	s3 := "abc"
	lc3 := TStr.LcFirst(s3)
	if s3 != lc3 {
		t.Errorf("The values of %v is not %v\n", lc3, s3)
	}
}

func TestUcWords(t *testing.T) {
	s1 := "hello world"
	u1 := "Hello World"
	uc1 := TStr.UcWords(s1)
	if u1 != uc1 {
		t.Errorf("The values of %v is not %v\n", uc1, u1)
	}

	s2 := "hello_aa"
	u2 := "Hello_aa"
	uc2 := TStr.UcWords(s2)
	if u2 != uc2 {
		t.Errorf("The values of %v is not %v\n", uc2, u2)
	}

	s3 := ""
	u3 := ""
	uc3 := TStr.UcWords(s3)
	if u3 != uc3 {
		t.Errorf("The values of %v is not %v\n", uc3, u3)
	}
}

func TestLcWords(t *testing.T) {
	s1 := ""
	l1 := ""
	lc1 := TStr.LcWords(s1)
	if l1 != lc1 {
		t.Errorf("The values of %v is not %v\n", lc1, l1)
	}

	s2 := "Hello World"
	l2 := "hello world"
	lc2 := TStr.LcWords(s2)
	if l2 != lc2 {
		t.Errorf("The values of %v is not %v\n", lc2, l2)
	}

	s3 := "hello_bb"
	l3 := "hello_bb"
	lc3 := TStr.LcWords(s3)
	if l3 != lc3 {
		t.Errorf("The values of %v is not %v\n", lc3, l3)
	}
}

func TestShuffle(t *testing.T) {
	s := "123abc*&^"
	sfx := TStr.Shuffle(s)
	if len(sfx) != len(s) {
		t.Errorf("The length of %v is not %v\n", len(sfx), len(s))
	}
}

func TestStropsFirstFind(t *testing.T)  {
	s1 := "asd123fgh456"
	f1 := "789"
	pos1 := TStr.StropsFirstFind(s1, f1, 0)
	if pos1 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos1, -1)
	}

	f2 := "123"
	pos2 := TStr.StropsFirstFind(s1, f2, 0)
	if pos2 != 3 {
		t.Errorf("The pos of %v is not %v\n", pos2, 3)
	}

	pos3 := TStr.StropsFirstFind(s1, f2, 9999)
	if pos3 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos3, -1)
	}

	pos4 := TStr.StropsFirstFind(s1, f2, -9999)
	if pos4 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos4, -1)
	}

	f5 := "F"
	pos5 := TStr.StropsFirstFind(s1, f5, 0)
	if pos5 != 6 {
		t.Errorf("The pos of %v is not %v\n", pos5, -1)
	}

	f6 := ""
	pos6 := TStr.StropsFirstFind(s1, f6, 0)
	if pos6 != 0 {
		t.Errorf("The pos of %v is not %v\n", pos6, 0)
	}
}

func TestStropsFirst(t *testing.T)  {
	s1 := "asd123fgh456"
	f1 := "789"
	pos1 := TStr.StropsFirst(s1, f1, 0)
	if pos1 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos1, -1)
	}

	f2 := "123"
	pos2 := TStr.StropsFirst(s1, f2, 0)
	if pos2 != 3 {
		t.Errorf("The pos of %v is not %v\n", pos2, 3)
	}

	pos3 := TStr.StropsFirst(s1, f2, 9999)
	if pos3 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos3, -1)
	}

	pos4 := TStr.StropsFirst(s1, f2, -9999)
	if pos4 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos4, -1)
	}

	f5 := "F"
	pos5 := TStr.StropsFirst(s1, f5, 0)
	if pos5 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos5, -1)
	}

	f6 := ""
	pos6 := TStr.StropsFirst(s1, f6, 0)
	if pos6 != 0 {
		t.Errorf("The pos of %v is not %v\n", pos6, 0)
	}
}

func TestStropsLast(t *testing.T) {
	s1 := "asd123asd"
	f1 := "789"
	pos1 := TStr.StropsLast(s1, f1, 0)
	if pos1 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos1, -1)
	}

	f2 := "asd"
	pos2 := TStr.StropsLast(s1, f2, 0)
	if pos2 != 6 {
		t.Errorf("The pos of %v is not %v\n", pos2, 3)
	}

	pos3 := TStr.StropsLast(s1, f2, 9999)
	if pos3 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos3, -1)
	}

	pos4 := TStr.StropsLast(s1, f2, -9999)
	if pos4 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos4, -1)
	}

	f5 := "A"
	pos5 := TStr.StropsLast(s1, f5, 0)
	if pos5 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos5, -1)
	}

	f6 := ""
	pos6 := TStr.StropsLast(s1, f6, 0)
	if pos6 != 9 {
		t.Errorf("The pos of %v is not %v\n", pos6, 9)
	}
}

func TestStropsLastFind(t *testing.T) {
	s1 := "asd123asd"
	f1 := "789"
	pos1 := TStr.StropsLastFind(s1, f1, 0)
	if pos1 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos1, -1)
	}

	f2 := "asd"
	pos2 := TStr.StropsLastFind(s1, f2, 0)
	if pos2 != 6 {
		t.Errorf("The pos of %v is not %v\n", pos2, 3)
	}

	pos3 := TStr.StropsLastFind(s1, f2, 9999)
	if pos3 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos3, -1)
	}

	pos4 := TStr.StropsLastFind(s1, f2, -9999)
	if pos4 != -1 {
		t.Errorf("The pos of %v is not %v\n", pos4, -1)
	}

	f5 := "A"
	pos5 := TStr.StropsLastFind(s1, f5, 0)
	if pos5 != 6 {
		t.Errorf("The pos of %v is not %v\n", pos5, 6)
	}

	f6 := ""
	pos6 := TStr.StropsLastFind(s1, f6, 0)
	if pos6 != 9 {
		t.Errorf("The pos of %v is not %v\n", pos6, 9)
	}
}

func TestReverse(t *testing.T) {
	s1 := ""
	r1 := TStr.Reverse(s1)
	if s1 != r1 {
		t.Errorf("The values of %v is not %v\n", r1, s1)
	}

	s2 := "abc123"
	want := "321cba"
	r2 := TStr.Reverse(s2)
	if r2 != want {
		t.Errorf("The values of %v is not %v\n", r2, s2)
	}
}

func TestMd5Hex(t *testing.T) {
	b1 := []byte{56,50,55,99,99,98,48,101,101,97,56,97,55,48,54,99,52,99,51,52,97,49,54,56,57,49,102,56,52,101,55,98}
	hex1 := TStr.Md5Hex([]byte("12345"), 0)
	if !bytes.Equal(b1, hex1) {
		t.Errorf("The hex of %v is not %v\n", hex1, b1)
	}

	b2 := []byte{56,50}
	hex2 := TStr.Md5Hex([]byte("12345"), 2)
	if !bytes.Equal(b2, hex2) {
		t.Errorf("The hex of %v is not %v\n", hex2, b2)
	}
}

func TestTrim(t *testing.T) {
	str := " hello world!你好 世界！　"
	res := TStr.Trim(str)
	if res[0] != 'h' {
		t.Error("Trim unit test fail")
		return
	}

	res = TStr.Trim("\v\t 0.0.0\f\n ")
	if res != "0.0.0" {
		t.Error("Trim unit test fail")
		return
	}

	TStr.Trim(str, "\n")
}

func TestLtrim(t *testing.T) {
	str := " hello world!你好 世界！　"
	res := TStr.Ltrim(str)
	if res[0] != 'h' {
		t.Error("Ltrim unit test fail")
		return
	}
	TStr.Ltrim(str, "\n")
}

func TestRtrim(t *testing.T) {
	str := " hello world!你好 世界！　"
	res := TStr.Rtrim(str, "　")
	if strings.HasSuffix(res, "　") {
		t.Error("Rtrim unit test fail")
		return
	}
	TStr.Rtrim(str)
}

func TestDStrPos(t *testing.T) {
	var str string
	var arr []string

	chk1, itm1 := TStr.DStrPos(str, arr, false)
	if chk1 || itm1 != "" {
		t.Error("DStrPos unit test fail")
		return
	}

	str = "Hello 你好, World 世界！"
	arr = []string{"he", "好", "world"}

	chk2, itm2 := TStr.DStrPos(str, arr, false)
	if !chk2 || itm2 == "" {
		t.Error("DStrPos unit test fail")
		return
	}

	chk3, itm3 := TStr.DStrPos(str, arr, true)
	if !chk3 || itm3 != "好" {
		t.Error("DStrPos unit test fail")
		return
	}

	arr = []string{"呵呵", "时间", "gogo"}
	chk4, itm4 := TStr.DStrPos(str, arr, true)
	if chk4 || itm4 != "" {
		t.Error("DStrPos unit test fail")
		return
	}
}

func TestMbSubstr(t *testing.T) {
	TStr.MbSubstr("", 0)
	TStr.MbSubstr("abcdef", 0)
	for _, test := range exampleStrTests {
		actual := TStr.MbSubstr(test.param, test.start, test.length)
		if actual != test.expected {
			t.Errorf("Expected Substr(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func TestDoZlibCompressAndUnCompress(t *testing.T) {
	zipStr := `我是一只小小的goper,写呀写呀写bug.`
	zip, _ := TStr.DoZlibCompress([]byte(zipStr))

	unzip, _ := TStr.DoZlibUnCompress(zip)
	unzipString := string(unzip)

	if zipStr != unzipString {
		t.Errorf("zip string is %v, not %v", zipStr, unzipString)
		return
	}
}
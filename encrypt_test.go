package helper

import (
	"fmt"
	"strings"
	"testing"
)

func TestEncBase64Encode(t *testing.T) {
	str := []byte("This is an string to encode")
	res := TEncrypt.Base64Encode(str)
	if !strings.HasSuffix(res, "=") {
		t.Error("Base64Encode unit test fail")
		return
	}
}

func TestBase64Decode(t *testing.T) {
	str := "VGypsyPacyBibiBlumNvZGVkIHN0camlZw=="
	_, err := TEncrypt.Base64Decode(str)
	if err != nil {
		t.Error("Base64Decode unit test fail")
		return
	}
	_, err = TEncrypt.Base64Decode("#iu3498r")
	if err == nil {
		t.Error("Base64Decode unit test fail")
		return
	}
	_, err = TEncrypt.Base64Decode("VGypsy")
	_, err = TEncrypt.Base64Decode("VGypsyB")
}

func TestBase64UrlEncodeDecode(t *testing.T) {
	str := []byte("This is an string to encode")
	res := TEncrypt.Base64UrlEncode(str)
	if strings.HasSuffix(res, "=") {
		t.Error("Base64UrlEncode unit test fail")
		return
	}

	_, err := TEncrypt.Base64UrlDecode(res)
	if err != nil {
		t.Error("Base64UrlDecode unit test fail")
		return
	}
}

func TestAuthCode(t *testing.T) {
	key := "123456"
	str := "hello world"

	res, _ := TEncrypt.AuthCode(str, key, true, 0)
	if res == "" {
		t.Error("AuthCode Encode unit test fail")
		return
	}

	res2, _ := TEncrypt.AuthCode(res, key, false, 0)
	if res2 == "" {
		t.Error("AuthCode Decode unit test fail")
		return
	}

	res, _ = TEncrypt.AuthCode(str, key, true, -3600)
	TEncrypt.AuthCode(res, key, false, 0)
	TEncrypt.AuthCode("", key, true, 0)
	TEncrypt.AuthCode("", "", true, 0)
	TEncrypt.AuthCode("7caeNfPt/N1zHdj5k/7i7pol6NHsVs0Cji7c15h4by1RYcrBoo7EEw==", key, false, 0)
	TEncrypt.AuthCode("7caeNfPt/N1zHdj5k/7i7pol6N", key, false, 0)
	TEncrypt.AuthCode("123456", "", false, 0)
	TEncrypt.AuthCode("1234#iu3498r", "", false, 0)
}

func TestPasswordHashVerify(t *testing.T) {
	pwd := []byte("123456")
	has, err := TEncrypt.PasswordHash(pwd)
	if err != nil {
		t.Error("PasswordHash unit test fail")
		return
	}

	chk := TEncrypt.PasswordVerify(pwd, has)
	if !chk {
		t.Error("PasswordVerify unit test fail")
		return
	}

	_, _ = TEncrypt.PasswordHash(pwd, 1)
	//慎用20以上,太耗时
	_, _ = TEncrypt.PasswordHash(pwd, 15)
	_, _ = TEncrypt.PasswordHash(pwd, 33)
}

func BenchmarkPasswordHash(b *testing.B) {
	b.ResetTimer()
	pwd := []byte("123456")
	for i := 0; i < b.N; i++ {
		//太耗时,只测试少量的
		if i > 10 {
			break
		}
		_, _ = TEncrypt.PasswordHash(pwd)
	}
}

func BenchmarkPasswordVerify(b *testing.B) {
	b.ResetTimer()
	pwd := []byte("123456")
	has := []byte("$2a$10$kCv6ljsVuTSI54oPkWereEmUNTW/zj0Dgh6qF4Vz0w4C3gVf/w7a")
	for i := 0; i < b.N; i++ {
		//太耗时,只测试少量的
		if i > 10 {
			break
		}
		TEncrypt.PasswordVerify(pwd, has)
	}
}

func TestEasyEncryptDecrypt(t *testing.T) {
	key := "123456"
	str := "hello world你好!hello world你好!hello world你好!hello world你好!"
	enc := TEncrypt.EasyEncrypt(str, key)
	if enc == "" {
		t.Error("EasyEncrypt unit test fail")
		return
	}

	dec := TEncrypt.EasyDecrypt(enc, key)
	if dec != str {
		t.Error("EasyDecrypt unit test fail")
		return
	}

	dec = TEncrypt.EasyDecrypt("你好，世界！", key)
	if dec != "" {
		t.Error("EasyDecrypt unit test fail")
		return
	}

	TEncrypt.EasyEncrypt("", key)
	TEncrypt.EasyEncrypt("", "")
	TEncrypt.EasyDecrypt(enc, "1qwerty")
	TEncrypt.EasyDecrypt("123", key)
	TEncrypt.EasyDecrypt("1234#iu3498r", key)
}

func BenchmarkEasyDecrypt(b *testing.B) {
	b.ResetTimer()
	key := "123456"
	str := "e10azZacODumpY"
	for i := 0; i < b.N; i++ {
		TEncrypt.EasyDecrypt(str, key)
	}
}

func TestHmacShaX(t *testing.T) {
	str := []byte("hello world")
	key := []byte("123456")
	res1 := TEncrypt.HmacShaX(str, key, 1)
	res2 := TEncrypt.HmacShaX(str, key, 256)
	res3 := TEncrypt.HmacShaX(str, key, 512)
	if res1 == "" || res2 == "" || res3 == "" {
		t.Error("HmacShaX unit test fail")
		return
	}
}

func TestHmacShaXPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	str := []byte("hello world")
	key := []byte("123456")
	TEncrypt.HmacShaX(str, key, 4)
}

func BenchmarkHmacShaX(b *testing.B) {
	b.ResetTimer()
	str := []byte("hello world")
	key := []byte("123456")
	for i := 0; i < b.N; i++ {
		TEncrypt.HmacShaX(str, key, 256)
	}
}
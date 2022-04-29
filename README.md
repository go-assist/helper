# Helper
##### 1.仓库迁移至☞ https://github.com/go-assist/helper


##### 2.一些常用的助手函数工具包
包含: 

字符串 (TStr)

数组 (TArr)

整型 (TInt)

浮点型 (TFloat)

文件 (TFile)

url (TUri)

hash (THash)

动态调用方法 (TCallFunc)

debug (TDebug)

os (TOs)

类型转换 (TConv)

时间 (TTime)

定时任务 (TCorn)

uuid (TUuid)

json (TJson)

加解密 (TEncrypt)

以上操作☝☝☝

##### 3.获取 ❤❤❤
go version >= 1.15

go get github.com/go-assist/helper

##### 4. 示例 for example 

```Golang
package helper

import (
	"fmt"
	"github.com/go-assist/helper"
)

func example() {
	// string
	s1 := `123456`
	md5 := helper.TStr.MD5(s1)
	fmt.Println(md5) // E10ADC3949BA59ABBE56E057F20F883E

	s2 := `hello world`
	ucFirst := helper.TStr.UcFirst(s2)
	fmt.Println(ucFirst) // Hello world

	sfx := helper.TStr.Shuffle(s1)
	fmt.Println(sfx) // 125436

	// array
	arr := [5]int{1, 2, 3, 4, 5}
	i := 2
	if helper.TArr.InArray(i, arr) {
		fmt.Printf(" %v in %v \n", i, arr) //  2 in [1 2 3 4 5]
	}

	// debug
	funcName := helper.TDebug.GetFuncName(helper.TArr.ArrayDiff, true) // ...ArrayDiff-fm
	fmt.Println(funcName) // ArrayDiff-fm

	// float
	round := helper.TFloat.Round(4.56)
	fmt.Println(round) // 4

	// json
	jsonArr := `[{"email_address":"test1@email.com"},{"email_address":"test2@email.com"}]`
	m := helper.TJson.JsonToMapArr(jsonArr)
	fmt.Println(m) // [map[email_address:test1@email.com] map[email_address:test2@email.com]]

	// convert
	inter := helper.TConv.Int2Str(s1)
	fmt.Println(inter) // 123456

	// hash, 此方法可以用来获取0-100区间的随机值
	h := "abc123tre"
	hashcode := helper.THash.HashCode(h, 2)
	fmt.Println(hashcode) // 52 

	// os
	endian := helper.TOs.GetEndian()
	fmt.Println(endian) // LittleEndian

	// file
	f := "./testname.txt"
	fileName := helper.TFile.Basename(f)
	fmt.Println(fileName) // testname.txt

	// encrypt
	ek := "Key"
	enc := helper.TEncrypt.EasyEncrypt(s1, ek)
	fmt.Println(enc) // 89735695aWtqZ2ps
	dec := helper.TEncrypt.EasyDecrypt(enc, ek)
	fmt.Println(dec) // 123456

	// url
	uri := `http://localhost/report?Av=5.3.5&Bd=bdtest&Cid=023&CityCode=101030100&Did=70836bc3ae68fddbc78ce5a917ae9e9d60c712df&Imei=`
	qm := helper.TUri.ParseUriQueryToMap(uri)
	fmt.Println(qm) // map[Av:5.3.5 Bd:bdtest Cid:023 CityCode:101030100 Did:70836bc3ae68fddbc78ce5a917ae9e9d60c712df Imei:]
	
	// int
	abs := helper.TInt.Abs(-1)
	fmt.Println(abs) // 1
	
	// time
	days := helper.TTime.GetMonthDays(2)
	fmt.Println(days) // 28
	
	// validate
	isEven := helper.TInt.IsEven(3)
	fmt.Println(isEven) // false
	
	// zlib string
	zip, _ := helper.TStr.DoZlibCompress([]byte(`我是一只小小的goper,写呀写呀写bug`))
	fmt.Println(zip)
	unzip, _ := helper.TStr.DoZlibUnCompress(zip)
	unzipString := string(unzip)
	fmt.Println(unzipString)
	
	// ... other 
}
```

##### 5. 使用过程如有疑问欢迎issue ｡◕‿◕｡

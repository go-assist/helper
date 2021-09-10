package helper

import (
	_ "encoding/json"
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	//数组
	arr := [5]int{1, 2, 3, 4, 5}
	it := 2
	if !TArr.InArray(it, arr) {
		t.Errorf("The values of %v not in %v \n", it, arr)
	}

	//字典
	mp := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	it2 := "a"
	it3 := "bb"
	if TArr.InArray(it2, mp) {
		t.Errorf("The values of %v not in %v \n", it2, mp)
	} else if !TArr.InArray(it3, mp) {
		t.Errorf("The values of %v not in %v \n", it3, mp)
	}

	if TArr.InArray(it2, "abc") {
		t.Errorf("The values of %v not in %v \n", it2, arr)
	}
}

func TestArrayFill(t *testing.T) {
	n1 := 4
	r1 := TArr.ArrayFill("abc", n1)
	if len(r1) != n1 {
		t.Errorf("The values of %v not in %v \n", len(r1), n1)
	}

	n2 := 0
	r2 := TArr.ArrayFill("abc", n2)
	if len(r2) != n2 {
		t.Errorf("The values of %v not in %v \n", len(r2), n2)
	}
}

func TestArrayFlip(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	res := TArr.ArrayFlip(mp)
	if val, ok := res[1]; !ok || fmt.Sprintf("%v", val) != "a" {
		t.Error("ArrayFlip unit test fail")
		return
	}

	var sli = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"
	res = TArr.ArrayFlip(sli)

	TArr.ArrayFlip("hello")
}

func TestArrayKeys(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	res := TArr.ArrayKeys(mp)
	if len(res) != 3 {
		t.Error("ArrayKeys unit test fail")
		return
	}

	var sli = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"
	res = TArr.ArrayKeys(sli)
	if len(res) != 5 {
		t.Error("ArrayKeys unit test fail")
		return
	}

	TArr.ArrayKeys("hello")
}

func TestArrayValues(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	res := TArr.ArrayValues(mp, false)
	if len(res) != 3 {
		t.Error("ArrayValues unit test fail")
		return
	}

	var sli = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"
	res = TArr.ArrayValues(sli, false)
	if len(res) != 5 {
		t.Error("ArrayValues unit test fail")
		return
	}

	TArr.ArrayValues("hello", false)
}

func TestMergeSlice(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	var arr = [10]int{1, 2, 3, 4, 5, 6}
	var sli = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"

	res1 := TArr.MergeSlice(false, arr, sli)
	if len(res1) != 15 {
		t.Error("MergeSlice unit test fail")
		return
	}

	res2 := TArr.MergeSlice(true, arr, sli)
	if len(res2) != 13 {
		t.Error("MergeSlice unit test fail")
		return
	}
	TArr.MergeSlice(true)
	TArr.MergeSlice(false, "hello")
}

func TestMergeMap(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	mp1 := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	mp2 := map[string]int{
		"h": 1,
		"i": 2,
		"j": 3,
	}

	res := TArr.MergeMap(true, mp1, mp2)
	_, err := TJson.JsonEncode(res)
	if err != nil {
		t.Error("MergeMap unit test fail")
		return
	}
	TArr.MergeMap(false)
	TArr.MergeMap(false, mp1, mp2)
	TArr.MergeMap(false, mp1, mp2, "hello")
}

func TestArrayChunk(t *testing.T) {
	size := 3
	var arr = [11]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	res1 := TArr.ArrayChunk(arr, size)
	if len(res1) != 4 {
		t.Error("ArrayChunk unit test fail")
		return
	}

	var mySlice []int
	TArr.ArrayChunk(mySlice, 1)
}

func TestArrayPad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	var sli []int
	var arr = [3]string{"a", "b", "c"}

	res1 := TArr.ArrayPad(sli, 5, 1)
	res2 := TArr.ArrayPad(arr, 6, "d")
	res3 := TArr.ArrayPad(arr, -6, "d")
	res4 := TArr.ArrayPad(arr, 2, "d")
	if len(res1) != 5 || len(res2) != 6 || fmt.Sprintf("%v", res3[0]) != "d" || len(res4) != 3 {
		t.Error("ArrayPad unit test fail")
		return
	}

	TArr.ArrayPad("hello", 2, "d")
}

func TestArraySlice(t *testing.T) {
	var sli []int
	var arr = [6]string{"a", "b", "c", "d", "e", "f"}

	res1 := TArr.ArraySlice(sli, 0, 1)
	res2 := TArr.ArraySlice(arr, 1, 2)
	res3 := TArr.ArraySlice(arr, -3, 2)
	res4 := TArr.ArraySlice(arr, -3, 4)
	if len(res1) != 0 || len(res2) != 2 || len(res3) != 2 || len(res4) != 3 {
		t.Error("ArraySlice unit test fail")
		return
	}
}

func TestArrayRand(t *testing.T) {
	var arr = [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var sli []int

	res1 := TArr.ArrayRand(sli, 1)
	res2 := TArr.ArrayRand(arr, 3)
	res3 := TArr.ArrayRand(arr, 9)

	if len(res1) != 0 || len(res2) != 3 || len(res3) != 8 {
		t.Error("ArraySlice unit test fail")
		return
	}
}

func TestArrayColumn(t *testing.T) {
	//数组切片
	jsonStr := `[{"name":"zhang3","age":23,"sex":1},{"name":"li4","age":30,"sex":1},{"name":"wan5","age":25,"sex":0},{"name":"zhao6","age":50,"sex":0}]`
	var arr interface{}
	err := TJson.JsonDecode([]byte(jsonStr), &arr)
	if err != nil {
		t.Error("JsonDecode unit test fail")
		return
	}

	res := TArr.ArrayColumn(arr, "name")
	if len(res) != 4 {
		t.Error("ArrayColumn unit test fail")
		return
	}

	//字典
	jsonStr = `{"person1":{"name":"zhang3","age":23,"sex":1},"person2":{"name":"li4","age":30,"sex":1},"person3":{"name":"wan5","age":25,"sex":0},"person4":{"name":"zhao6","age":50,"sex":0}}`
	err = TJson.JsonDecode([]byte(jsonStr), &arr)
	if err != nil {
		t.Error("JsonDecode unit test fail")
		return
	}

	res = TArr.ArrayColumn(arr, "name")
	if len(res) != 4 {
		t.Error("ArrayColumn unit test fail")
		return
	}
}

func TestArrayPushPop(t *testing.T) {
	var arr []interface{}
	length := TArr.ArrayPush(&arr, 1, 2, 3, "a", "b", "c")
	if length != 6 {
		t.Error("ArrayPush unit test fail")
		return
	}

	last := TArr.ArrayPop(&arr)
	if fmt.Sprintf("%v", last) != "c" {
		t.Error("ArrayPop unit test fail")
		return
	}
	arr = nil
	TArr.ArrayPop(&arr)
}

func BenchmarkArrayPush(b *testing.B) {
	b.ResetTimer()
	var arr []interface{}
	for i := 0; i < b.N; i++ {
		TArr.ArrayPush(&arr, 1, 2, 3, "a", "b", "c")
	}
}

func BenchmarkArrayPop(b *testing.B) {
	b.ResetTimer()
	var arr = []interface{}{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		TArr.ArrayPop(&arr)
	}
}

func TestArrayShiftUnshift(t *testing.T) {
	var arr []interface{}
	length := TArr.ArrayUnshift(&arr, 1, 2, 3, "a", "b", "c")
	if length != 6 {
		t.Error("ArrayUnshift unit test fail")
		return
	}

	first := TArr.ArrayShift(&arr)
	if fmt.Sprintf("%v", first) != "1" {
		t.Error("ArrayPop unit test fail")
		return
	}
	arr = nil
	TArr.ArrayShift(&arr)
}

func BenchmarkArrayUnshift(b *testing.B) {
	b.ResetTimer()
	var arr []interface{}
	for i := 0; i < b.N; i++ {
		TArr.ArrayUnshift(&arr, 1, 2, 3, "a", "b", "c")
	}
}

func BenchmarkArrayShift(b *testing.B) {
	b.ResetTimer()
	var arr = []interface{}{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		TArr.ArrayShift(&arr)
	}
}

func TestArrayKeyExistsArr(t *testing.T) {
	var arr []interface{}
	TArr.ArrayPush(&arr, 1, 2, 3, "a", "b", "c")

	chk1 := TArr.ArrayKeyExists(1, arr)
	chk2 := TArr.ArrayKeyExists(-1, arr)
	chk3 := TArr.ArrayKeyExists(6, arr)
	if !chk1 || chk2 || chk3 {
		t.Error("ArrayKeyExists unit test fail")
		return
	}

	var key interface{}
	TArr.ArrayKeyExists(key, arr)

	arr = nil
	TArr.ArrayKeyExists(1, arr)
}

func TestArrayReverse(t *testing.T) {
	var arr = []interface{}{"a", "b", "c", "d", "e"}
	res := TArr.ArrayReverse(arr)

	if len(res) != 5 || fmt.Sprintf("%s", res[2]) != "c" {
		t.Error("ArrayReverse unit test fail")
		return
	}

	var mySlice []int
	TArr.ArrayReverse(mySlice)
}

func TestArrayReversePanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	TArr.ArrayReverse("hello")
}

func TestImplode(t *testing.T) {
	var arr = []string{"a", "b", "c", "d", "e"}
	res := TArr.Implode(",", arr)

	arr = nil
	res = TArr.Implode(",", arr)
	if res != "" {
		t.Error("Implode slice unit test fail")
		return
	}

	//字典
	var mp1 = make(map[string]string)
	res = TArr.Implode(",", mp1)
	if res != "" {
		t.Error("Implode map unit test fail")
		return
	}

	mp2 := map[string]string{
		"a": "aa",
		"b": "bb",
		"c": "cc",
		"d": "dd",
	}
	res = TArr.Implode(",", mp2)
	if res == "" {
		t.Error("Implode map unit test fail")
		return
	}
}

func TestImplodePanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	TArr.Implode(",", "hello")
}

func BenchmarkImplode(b *testing.B) {
	b.ResetTimer()
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		TArr.Implode(",", sli)
	}
}

func TestJoinStrings(t *testing.T) {
	var arr []string

	res := TArr.JoinStrings(arr, ",")
	if res != "" {
		t.Error("JoinStrings unit test fail")
		return
	}

	arr = append(arr, "a", "b", "c", "d", "e")
	TArr.JoinStrings(arr, ",")
}

func BenchmarkJoinStrings(b *testing.B) {
	b.ResetTimer()
	var arr = []string{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		TArr.JoinStrings(arr, ",")
	}
}

func TestJoinJoinIntArr(t *testing.T) {
	var arr []int

	res := TArr.JoinIntArr(arr, ",")
	if res != "" {
		t.Error("JoinStrings unit test fail")
		return
	}

	arr = append(arr, 1, 2, 3, 4, 5, 6)
	TArr.JoinIntArr(arr, ",")
}

func BenchmarkJoinIntArr(b *testing.B) {
	b.ResetTimer()
	var arr = []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		TArr.JoinIntArr(arr, ",")
	}
}

func TestUniqueIntArr(t *testing.T) {
	arr := []int{-3, 9, -5, 0, 5, -3, 0, 7}
	res := TArr.UniqueIntArr(arr)
	if len(arr) == len(res) {
		t.Error("UniqueIntArr unit test fail")
		return
	}
}

func BenchmarkUniqueIntArr(b *testing.B) {
	b.ResetTimer()
	arr := []int{-3, 9, -5, 0, 5, -3, 0, 7}
	for i := 0; i < b.N; i++ {
		TArr.UniqueIntArr(arr)
	}
}

func TestUnique64IntArr(t *testing.T) {
	arr := []int64{-3, 9, -5, 0, 5, -3, 0, 7}
	res := TArr.Unique64IntArr(arr)
	if len(arr) == len(res) {
		t.Error("Unique64IntArr unit test fail")
		return
	}
}

func BenchmarkUnique64IntArr(b *testing.B) {
	b.ResetTimer()
	arr := []int64{-3, 9, -5, 0, 5, -3, 0, 7}
	for i := 0; i < b.N; i++ {
		TArr.Unique64IntArr(arr)
	}
}

func TestUniqueStrings(t *testing.T) {
	var arr1 []string
	res1 := TArr.UniqueStringsArr(arr1)
	if len(res1) != 0 {
		t.Error("UniqueStrings unit test fail")
		return
	}

	arr2 := []string{"", "hello", "world", "hello", "你好", "world", "1234"}
	res2 := TArr.UniqueStringsArr(arr2)
	if len(arr2) == len(res2) {
		t.Error("UniqueStrings unit test fail")
		return
	}
}

func BenchmarkUniqueStrings(b *testing.B) {
	b.ResetTimer()
	arr := []string{"", "hello", "world", "hello", "你好", "world", "1234"}
	for i := 0; i < b.N; i++ {
		TArr.UniqueStringsArr(arr)
	}
}

func TestArrayDiff(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	ar1 := []string{"aa", "bb", "cc", "dd", ""}
	ar2 := []string{"bb", "cc", "ff", "gg", ""}
	mp1 := map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": ""}
	mp2 := map[string]string{"a": "0", "b": "2", "c": "4", "g": "4", "h": ""}

	var ar3 []string
	var mp3 = make(map[string]string)

	res1 := TArr.ArrayDiff(ar1, ar2)
	res2 := TArr.ArrayDiff(mp1, mp2)
	if len(res1) != len(res2) {
		t.Error("ArrayDiff unit test fail")
		return
	}

	res5 := TArr.ArrayDiff(ar3, ar1)
	res6 := TArr.ArrayDiff(ar1, ar3)
	if len(res5) != 0 || len(res6) != 4 {
		t.Error("ArrayDiff unit test fail")
		return
	}

	res7 := TArr.ArrayDiff(mp3, mp1)
	res8 := TArr.ArrayDiff(mp1, mp3)
	if len(res7) != 0 || len(res8) != 4 {
		t.Error("ArrayDiff unit test fail")
		return
	}

	res9 := TArr.ArrayDiff(ar3, mp1)
	res10 := TArr.ArrayDiff(ar1, mp3)
	res11 := TArr.ArrayDiff(ar1, mp1)
	if len(res9) != 0 || len(res10) != len(res11) {
		t.Error("ArrayDiff unit test fail")
		return
	}

	res12 := TArr.ArrayDiff(mp3, ar1)
	res13 := TArr.ArrayDiff(mp1, ar3)
	res14 := TArr.ArrayDiff(mp1, ar1)
	if len(res12) != 0 || len(res13) != len(res14) {
		t.Error("ArrayDiff unit test fail")
		return
	}

	TArr.ArrayDiff("hello", ar1)
}

func BenchmarkArrayDiff(b *testing.B) {
	b.ResetTimer()
	ar1 := []string{"aa", "bb", "cc", "dd", ""}
	ar2 := []string{"bb", "cc", "ff", "gg", ""}
	for i := 0; i < b.N; i++ {
		TArr.ArrayDiff(ar1, ar2)
	}
}

func TestArrayUnique(t *testing.T) {
	arr1 := map[string]string{"a": "green", "0": "red", "b": "green", "1": "blue", "2": "red"}
	arr2 := []string{"aa", "bb", "cc", "", "bb", "aa"}
	res1 := TArr.ArrayUnique(arr1)
	res2 := TArr.ArrayUnique(arr2)
	if len(res1) == 0 || len(res2) == 0 {
		t.Error("ArrayUnique fail")
		return
	}
}

func TestArrayUniquePanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	_ = TArr.ArrayUnique("hello")
}

func TestArraySearchMulti(t *testing.T) {
	type item map[string]interface{}

	var list []interface{}
	arr := make(map[string]interface{})
	cond := make(map[string]interface{})

	res1 := TArr.ArraySearchItem(list, cond)
	res2 := TArr.ArraySearchItem(arr, cond)
	if res1 != nil || res2 != nil {
		t.Error("ArraySearchItem fail")
	}

	mul1 := TArr.ArraySearchMulti(list, cond)
	mul2 := TArr.ArraySearchMulti(arr, cond)
	if mul1 != nil || mul2 != nil {
		t.Error("ArraySearchMulti fail")
	}

	item1 := item{"age": 20, "name": "test1", "location": "us", "tel": "13712345678"}
	item2 := item{"age": 21, "name": "test2", "location": "cn", "tel": "13712345679"}
	item3 := item{"age": 22, "name": "test3", "location": "en", "tel": "13712345670"}
	item4 := item{"age": 23, "name": "test4", "location": "fr", "tel": "13712345671"}
	item5 := item{"age": 21, "name": "test5", "location": "cn", "tel": "13712345672"}

	list = append(list, item1, item2, item3, item4, item5, nil, "hello")
	arr["a"] = item1
	arr["b"] = item2
	arr["c"] = item3
	arr["c"] = item4
	arr["d"] = nil
	arr["d"] = "world"
	arr["e"] = item5

	cond1 := map[string]interface{}{"age": 23}
	cond2 := map[string]interface{}{"age": 21, "location": "cn"}
	cond3 := map[string]interface{}{"age": 22, "location": "cn", "tel": "13712345671"}

	res3 := TArr.ArraySearchItem(list, cond1)
	res4 := TArr.ArraySearchItem(arr, cond1)
	if res3 == nil || res4 == nil {
		t.Error("ArraySearchItem unit test fail")
	}

	mul3 := TArr.ArraySearchMulti(list, cond1)
	mul4 := TArr.ArraySearchMulti(arr, cond1)
	if mul3 == nil || mul4 == nil {
		t.Error("ArraySearchMulti unit test fail")
	}

	res5 := TArr.ArraySearchItem(list, cond2)
	res6 := TArr.ArraySearchItem(arr, cond2)
	if res5 == nil || res6 == nil {
		t.Error("ArraySearchItem unit test fail")
	}

	mul5 := TArr.ArraySearchMulti(list, cond2)
	mul6 := TArr.ArraySearchMulti(arr, cond2)
	if mul5 == nil || mul6 == nil {
		t.Error("ArraySearchMulti unit test fail")
	}

	res7 := TArr.ArraySearchItem(list, cond3)
	res8 := TArr.ArraySearchItem(arr, cond3)
	if res7 != nil || res8 != nil {
		t.Error("ArraySearchItem unit test fail")
	}

	mul7 := TArr.ArraySearchMulti(list, cond3)
	mul8 := TArr.ArraySearchMulti(arr, cond3)
	if mul7 != nil || mul8 != nil {
		t.Error("ArraySearchMulti unit test fail")
	}
}

func BenchmarkArraySearchMulti(b *testing.B) {
	b.ResetTimer()
	type item map[string]interface{}
	var list []interface{}

	item1 := item{"age": 20, "name": "test1", "location": "us", "tel": "13712345678"}
	item2 := item{"age": 21, "name": "test2", "location": "cn", "tel": "13712345679"}
	item3 := item{"age": 22, "name": "test3", "location": "en", "tel": "13712345670"}
	item4 := item{"age": 23, "name": "test4", "location": "fr", "tel": "13712345671"}
	item5 := item{"age": 21, "name": "test5", "location": "cn", "tel": "13712345672"}
	list = append(list, item1, item2, item3, item4, nil, "hello", item5)
	cond := map[string]interface{}{"age": 21, "location": "cn"}
	for i := 0; i < b.N; i++ {
		TArr.ArraySearchMulti(list, cond)
	}
}

func TestArraySearchMultiPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	TArr.ArraySearchMulti("hello", map[string]interface{}{"a": 1})
}

func TestArraySearchItemPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	TArr.ArraySearchItem("hello", map[string]interface{}{"a": 1})
}

func BenchmarkArraySearchItem(b *testing.B) {
	b.ResetTimer()
	type item map[string]interface{}
	var list []interface{}

	item1 := item{"age": 20, "name": "test1", "location": "us", "tel": "13712345678"}
	item2 := item{"age": 21, "name": "test2", "location": "cn", "tel": "13712345679"}
	item3 := item{"age": 22, "name": "test3", "location": "en", "tel": "13712345670"}
	item4 := item{"age": 23, "name": "test4", "location": "fr", "tel": "13712345671"}
	list = append(list, item1, item2, item3, item4, nil, "hello")
	cond := map[string]interface{}{"age": 21, "location": "cn"}
	for i := 0; i < b.N; i++ {
		TArr.ArraySearchItem(list, cond)
	}
}

func TestArrayCombine(t *testing.T) {
	keys := []interface{}{"a", "b", "c"}
	values := []interface{}{"aa", "bb", "cc"}
	want := map[interface{}]interface{}{"a":"aa", "b":"bb", "c": "cc"}
	got := TArr.ArrayCombine(keys, values)
	if len(want) != len(got) {
		t.Errorf("the length of %v, does not %v \n", len(got), len(want))
	}

	if _, ok := got["a"]; !ok {
		t.Errorf("ArrayCombine unit test fail \n")
	}
}

func TestIsMapBySprintf(t *testing.T) {
	m := make(map[interface{}]interface{}, 2)
	m["a"] = "aa"
	m["b"] = "b" // map

	array := [5]int {1,2,3,4,5} // array

	i := 12 // int

	s := "abc" // string

	mResult := TArr.IsMapBySprintf(m)
	arrResult := TArr.IsMapBySprintf(array)
	intResult := TArr.IsMapBySprintf(i)
	strResult := TArr.IsMapBySprintf(s)

	if !mResult || arrResult || intResult || strResult {
		t.Errorf("IsMapBySprintf unit test fail \n")
	}
}
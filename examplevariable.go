package helper

type Example struct {
	Examples string
}

var (
	jsonExample = `{"k1":"v1","k2":"v2"}`
	errJson  = `{"k1":"v1","k2":"v2}`
	jsonArr = `[{"email_address":"test1@uber.com"},{"email_address":"test2@uber.com"}]`
	jsonMap = map[string] interface{} {"C":5, "Go":4.5, "Python":4.5, "C++":2 }
	jsonStruct = map[string] interface{} {"Examples":"test"}
	exampleStrTests = []struct {
		param    string
		start    int
		length   int
		expected string
	}{
		{"ab你好世界cdef01", 0, 4, "ab你好"},
		{"ab你好世界cdef02", -2, 4, "02"},
		{"ab你好世界cdef03", 0, -2, "ab你好世界cdef"},
		{"ab你好世界cdef04", -20, 8, ""},
		{"ab你好世界cdef05", 5, 50, "界cdef05"},
	}
	exampleUriStr = `http://localhost/report?Av=5.3.5&Bd=bdtest&Cid=023&CityCode=101030100&Did=70836bc3ae68fddbc78ce5a917ae9e9d60c712df&Imei=`
	exampleUriTests = []struct {
		param    string
		isMain   bool
		expected string
	}{
		{"", false, ""},
		{"hello world", false, ""},
		{"http://login.localhost:3000", false, "login.localhost"},
		{"https://play.golang.com:3000/p/3R1TPyk8qck", false, "play.golang.com"},
		{"https://www.siongui.github.io/pali-chanting/zh/archives.html", true, "github.io"},
		{"http://foobar.中文网/", false, "foobar.中文网"},
		{"foobar.com/abc/efg/h=1", false, ""},
		{"127.0.0.1", false, ""},
	}
)
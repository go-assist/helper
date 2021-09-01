package helper

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/user"
	"strings"
	"testing"
	"time"
)

func TestIsWindows(t *testing.T) {
	res := TOs.IsWindows()
	if res {
		t.Error("IsWindows unit test fail")
		return
	}
}

func BenchmarkIsWindows(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.IsWindows()
	}
}

func TestIsLinux(t *testing.T) {
	res := TOs.IsLinux()
	if !res {
		t.Error("IsLinux unit test fail")
		return
	}
}

func BenchmarkIsLinux(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.IsLinux()
	}
}

func TestIsMac(t *testing.T) {
	res := TOs.IsMac()
	if res {
		t.Error("IsMac unit test fail")
		return
	}
}

func BenchmarkIsMac(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.IsMac()
	}
}

func TestPwd(t *testing.T) {
	res := TOs.Pwd()
	if res == "" {
		t.Error("Pwd unit test fail")
		return
	}
}

func BenchmarkPwd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.Pwd()
	}
}

func TestGetWorkList(t *testing.T) {
	res, err := TOs.GetWorkList()
	if err != nil || res == "" {
		t.Error("Getcwd unit test fail")
		return
	}
}

func BenchmarkGetWorkList(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TOs.GetWorkList()
	}
}

func TestChdir(t *testing.T) {
	err := TOs.Chdir("./testdata")
	if err != nil {
		println(err.Error())
		t.Error("Chdir unit test fail")
		return
	}

	err = TOs.Chdir("../")
	if err != nil {
		println(err.Error())
		t.Error("Chdir unit test fail")
		return
	}
	_ = TOs.Chdir("")
}

func BenchmarkChdir(b *testing.B) {
	b.ResetTimer()
	dir := TOs.Pwd()
	for i := 0; i < b.N; i++ {
		_ = TOs.Chdir(dir)
	}
}

func TestHomeDir(t *testing.T) {
	_, err := TOs.HomeDir()
	if err != nil {
		t.Error("Pwd unit test fail")
		return
	}
}

func BenchmarkHomeDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TOs.HomeDir()
	}
}

func TestLocalIP(t *testing.T) {
	_, err := TOs.LocalIP()
	if err != nil {
		t.Error("LocalIP unit test fail")
		return
	}
}

func BenchmarkLocalIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TOs.LocalIP()
	}
}

func TestOutboundIP(t *testing.T) {
	_, err := TOs.OutboundIP()
	if err != nil {
		t.Error("OutboundIP unit test fail")
		return
	}
}

func BenchmarkOutboundIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TOs.OutboundIP()
	}
}

func TestIsPublicIP(t *testing.T) {
	ipStr, _ := TOs.LocalIP()
	ipAddr := net.ParseIP(ipStr)
	if TOs.IsPublicIP(ipAddr) {
		t.Error("IsPublicIP unit test fail")
		return
	}
	TOs.IsPublicIP(net.ParseIP("127.0.0.1"))
	TOs.IsPublicIP(net.ParseIP("172.16.0.1"))
	TOs.IsPublicIP(net.ParseIP("192.168.0.1"))
	//google
	TOs.IsPublicIP(net.ParseIP("172.217.26.142"))
	//google IPv6
	TOs.IsPublicIP(net.ParseIP("2404:6800:4005:80f::200e"))
}

func BenchmarkIsPublicIP(b *testing.B) {
	b.ResetTimer()
	ipStr, _ := TOs.LocalIP()
	ipAddr := net.ParseIP(ipStr)
	for i := 0; i < b.N; i++ {
		TOs.IsPublicIP(ipAddr)
	}
}

func TestGetIPs(t *testing.T) {
	ips := TOs.GetIPs()
	if len(ips) == 0 {
		t.Error("GetIPs unit test fail")
		return
	}
}

func BenchmarkGetIPs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.GetIPs()
	}
}

func TestGetMacAddrArr(t *testing.T) {
	macs := TOs.GetMacAddrArr()
	if len(macs) == 0 {
		t.Error("GetMacAddrArr unit test fail")
		return
	}
}

func BenchmarkGetMacAddrs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.GetMacAddrArr()
	}
}

func TestHostname(t *testing.T) {
	res, err := TOs.Hostname()
	if err != nil || res == "" {
		t.Error("Hostname unit test fail")
		return
	}
}

func BenchmarkHostname(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TOs.Hostname()
	}
}

func TestGetIpByHostname(t *testing.T) {
	name := "localhost"
	ip, err := TOs.GetIpByHostname(name)
	if err != nil || ip != "127.0.0.1" {
		t.Error("GetIpByHostname unit test fail")
		return
	}

	_, _ = TOs.GetIpByHostname("::1")
	_, _ = TOs.GetIpByHostname("hello")
}

func BenchmarkGetIpByHostname(b *testing.B) {
	b.ResetTimer()
	name := "localhost"
	for i := 0; i < b.N; i++ {
		_, _ = TOs.GetIpByHostname(name)
	}
}

func TestGetIpsByDomain(t *testing.T) {
	name := "google.com"
	ips, err := TOs.GetIpsByDomain(name)
	if err != nil || len(ips) == 0 {
		t.Error("GetIpsByDomain unit test fail")
		return
	}

	ips, err = TOs.GetIpsByDomain("hello")
	if err == nil || len(ips) > 0 {
		t.Error("GetIpsByDomain unit test fail")
		return
	}
}

func BenchmarkGetIpsByDomain(b *testing.B) {
	b.ResetTimer()
	name := "google.com"
	for i := 0; i < b.N; i++ {
		_, _ = TOs.GetIpsByDomain(name)
	}
}

func TestGetHostByIp(t *testing.T) {
	ip := "127.0.0.1"
	host, err := TOs.GetHostByIp(ip)
	if err != nil || host == "" {
		t.Error("GetHostByIp unit test fail")
		return
	}

	_, _ = TOs.GetHostByIp("192.168.1.1")
}

func BenchmarkGetHostByIp(b *testing.B) {
	b.ResetTimer()
	ip := "127.0.0.1"
	for i := 0; i < b.N; i++ {
		_, _ = TOs.GetHostByIp(ip)
	}
}

func TestGoMemory(t *testing.T) {
	mem := TOs.GoMemory()
	if mem == 0 {
		t.Error("GoMemory unit test fail")
		return
	}
}

func BenchmarkGoMemory(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.GoMemory()
	}
}

//func TestMemoryUsage(t *testing.T) {
//	// 虚拟内存
//	used1, free1, total1 := TOs.MemoryUsage(true)
//	//usedRate1 := float64(used1) / float64(total1)
//	if used1 <= 0 || free1 <= 0 || total1 <= 0 {
//		t.Error("MemoryUsage(true) unit test fail")
//		return
//	}
//
//	// 真实物理内存
//	used2, free2, total2 := TOs.MemoryUsage(false)
//	//usedRate2 := float64(used2) / float64(total2)
//	if used2 <= 0 || free2 <= 0 || total2 <= 0 {
//		t.Error("MemoryUsage(false) unit test fail")
//		return
//	}
//}

//func BenchmarkMemoryUsageVirtual(b *testing.B) {
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		TOs.MemoryUsage(true)
//	}
//}

//func BenchmarkMemoryUsagePhysic(b *testing.B) {
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		TOs.MemoryUsage(false)
//	}
//}

func TestCpuUsage(t *testing.T) {
	usr, idle, total := TOs.CpuUsage()

	// 注意: usr/total,两个整数相除，结果是整数，取整数部分为0
	usedRate := float64(usr) / float64(total)
	freeRate := float64(idle) / float64(total)

	if usedRate == 0 || freeRate == 0 {
		t.Error("CpuUsage unit test fail")
		return
	}
}

func BenchmarkCpuUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.CpuUsage()
	}
}

//func TestDiskUsage(t *testing.T) {
//	used, free, total := TOs.DiskUsage("/")
//	if used <= 0 || free <= 0 || total <= 0 {
//		t.Error("DiskUsage unit test fail")
//		return
//	}
//}

//func BenchmarkDiskUsage(b *testing.B) {
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		TOs.DiskUsage("/")
//	}
//}

func TestSetenvGetenv(t *testing.T) {
	name1 := "HELLO"
	name2 := "HOME"

	err := TOs.Setenv(name1, "world")
	if err != nil {
		t.Error("Setenv unit test fail")
		return
	}

	val1 := TOs.Getenv(name1)
	val2 := TOs.Getenv(name2)
	if val1 != "world" || val2 == "" {
		t.Error("GetENV unit test fail")
		return
	}
}

func BenchmarkSetenv(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = TOs.Setenv("HELLO", "world")
	}
}

func BenchmarkGetENV(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.Getenv("HELLO")
	}
}

func TestGetEndian_IsLittleEndian(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	endi := TOs.GetEndian()
	isLit := TOs.IsLittleEndian()

	if fmt.Sprintf("%v", endi) == "" {
		t.Error("GetEndian unit test fail")
		return
	} else if isLit && fmt.Sprintf("%v", endi) != "LittleEndian" {
		t.Error("IsLittleEndian unit test fail")
		return
	}
}

func BenchmarkGetEndian(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.GetEndian()
	}
}

func BenchmarkIsLittleEndian(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.IsLittleEndian()
	}
}

func TestExec(t *testing.T) {
	cmd := " ls -a -h"
	ret, _, _ := TOs.Exec(cmd)
	if ret == 1 {
		t.Error("Exec unit test fail")
		return
	}

	cmd = " ls -a\"\" -h 'hehe'"
	_, _, _ = TOs.Exec(cmd)
}

func BenchmarkExec(b *testing.B) {
	b.ResetTimer()
	cmd := " ls -a -h"
	for i := 0; i < b.N; i++ {
		_, _, _ = TOs.Exec(cmd)
	}
}

func TestSystem(t *testing.T) {
	cmd := " ls -a -h"
	ret, _, _ := TOs.System(cmd)
	if ret == 1 {
		t.Error("System unit test fail")
		return
	}

	cmd = "123"
	_, _, _ = TOs.System(cmd)

	cmd = " ls -a\"\" -h 'hehe'"
	_, _, _ = TOs.System(cmd)

	cmd = "ls -a /root/"
	_, _, _ = TOs.System(cmd)

	filename := ""
	for i := 0; i < 10000; i++ {
		filename = fmt.Sprintf("./testdata/empty/zero_%d", i)
		TFile.Touch(filename, 0)
	}

	cmd = "ls -a -h ./testdata/empty"
	_, _, _ = TOs.System(cmd)
	_, _, _ = TOs.System(cmd)
	_, _, _ = TOs.System(cmd)

	cmd = "touch /root/hello"
	_, _, _ = TOs.System(cmd)
	_ = TFile.DelDir("./testdata/empty", false)
}

func BenchmarkSystem(b *testing.B) {
	b.ResetTimer()
	cmd := " ls -a -h"
	for i := 0; i < b.N; i++ {
		_, _, _ = TOs.System(cmd)
	}
}

func TestChmodChown(t *testing.T) {
	file := "./testdata"
	res1 := TOs.Chmod(file, 0777)

	usr, _ := user.Current()
	uid := TConv.Str2Int(usr.Uid)
	guid := TConv.Str2Int(usr.Gid)

	res2 := TOs.Chown(file, uid, guid)

	if !res1 || !res2 {
		t.Error("Chmod unit test fail")
		return
	}
}

func BenchmarkChmod(b *testing.B) {
	b.ResetTimer()
	file := "./testdata"
	for i := 0; i < b.N; i++ {
		TOs.Chmod(file, 0777)
	}
}

func BenchmarkChown(b *testing.B) {
	b.ResetTimer()
	file := "./testdata"
	usr, _ := user.Current()
	uid := TConv.Str2Int(usr.Uid)
	guid := TConv.Str2Int(usr.Gid)
	for i := 0; i < b.N; i++ {
		TOs.Chown(file, uid, guid)
	}
}

func TestGetTempDir(t *testing.T) {
	res := TOs.GetTempDir()
	if res == "" {
		t.Error("GetTempDir unit test fail")
		return
	}
}

func BenchmarkGetTempDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.GetTempDir()
	}
}

func TestPrivateCIDR(t *testing.T) {
	res := TOs.PrivateCIDR()
	if len(res) == 0 {
		t.Error("PrivateCIDR unit test fail")
		return
	}
}

func BenchmarkPrivateCIDR(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.PrivateCIDR()
	}
}

func TestIsPrivateIp(t *testing.T) {
	// 无效Ip
	res, err := TOs.IsPrivateIp("hello")
	if res || err == nil {
		t.Error("IsPrivateIp unit test fail")
		return
	}

	// TPrivyCiders未初始化数据
	if len(TPrivyCiders) != 0 {
		t.Error("IsPrivateIp unit test fail")
		return
	}

	// docker ip
	res, err = TOs.IsPrivateIp("172.17.0.1")
	if !res || err != nil {
		t.Error("IsPrivateIp unit test fail")
		return
	}

	//外网ip
	res, err = TOs.IsPrivateIp("220.181.38.148")
	if res || err != nil {
		t.Error("IsPrivateIp unit test fail")
		return
	}

	// TPrivyCiders 已初始化数据
	if len(TPrivyCiders) == 0 {
		t.Error("IsPrivateIp unit test fail")
		return
	}
}

func BenchmarkIsPrivateIp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = TOs.IsPrivateIp("172.17.0.1")
	}
}

func TestClientIp(t *testing.T) {
	// Create type and function for testing
	type testIP struct {
		name     string
		request  *http.Request
		expected string
	}

	newRequest := func(remoteAddr, xRealIP string, xForwardedFor ...string) *http.Request {
		h := http.Header{}
		h.Set("X-Real-IP", xRealIP)
		for _, address := range xForwardedFor {
			h.Set("X-Forwarded-For", address)
		}

		return &http.Request{
			RemoteAddr: remoteAddr,
			Header:     h,
		}
	}

	// Create test data
	publicAddr1 := "144.12.54.87"
	publicAddr2 := "119.14.55.11"
	publicAddr3 := "8.8.8.8:8080"
	localAddr1 := "127.0.0.0"
	localAddr2 := "::1"

	testData := []testIP{
		{
			name:     "No header,no port",
			request:  newRequest(publicAddr1, ""),
			expected: publicAddr1,
		}, {
			name:     "No header,has port",
			request:  newRequest(publicAddr3, ""),
			expected: publicAddr3,
		}, {
			name:     "Has X-Forwarded-For",
			request:  newRequest("", "", publicAddr1),
			expected: publicAddr1,
		}, {
			name:     "Has multiple X-Forwarded-For",
			request:  newRequest("", "", localAddr1, publicAddr1, publicAddr2),
			expected: publicAddr2,
		}, {
			name:     "Has X-Real-IP",
			request:  newRequest("", publicAddr1),
			expected: publicAddr1,
		}, {
			name:     "Local ip",
			request:  newRequest("", localAddr2),
			expected: localAddr2,
		},
	}

	// Run test
	var actual string
	for _, v := range testData {
		actual = TOs.ClientIp(v.request)
		if v.expected == "::1" {
			if actual != "127.0.0.1" {
				t.Errorf("%s: expected %s but get %s", v.name, v.expected, actual)
			}
		} else {
			if strings.Contains(v.expected, ":") {
				ip, _, _ := net.SplitHostPort(v.expected)
				if ip != actual {
					t.Errorf("%s: expected %s but get %s", v.name, v.expected, actual)
				}
			} else {
				if v.expected != actual {
					t.Errorf("%s: expected %s but get %s", v.name, v.expected, actual)
				}
			}
		}
	}
}

func BenchmarkClientIp(b *testing.B) {
	b.ResetTimer()
	req := &http.Request{
		RemoteAddr: "216.58.199.14",
	}
	for i := 0; i < b.N; i++ {
		TOs.ClientIp(req)
	}
}

func TestGetSystemInfo(t *testing.T) {
	info := TOs.GetSystemInfo()
	fmt.Printf("%+v\n", info)
}

func BenchmarkGetSystemInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.GetSystemInfo()
	}
}

func TestIsPortOpen(t *testing.T) {
	var tests = []struct {
		host     string
		port     interface{}
		protocol string
		expected bool
	}{
		{"", 23, "", false},
		{"localhost", 0, "", false},
		{"127.0.0.1", 23, "", false},
		{"golang.org", 80, "udp", true},
		{"golang.org", 80, "tcp", true},
		{"www.google.com", "443", "tcp", true},
	}
	for _, test := range tests {
		actual := TOs.IsPortOpen(test.host, test.port, test.protocol)
		if actual != test.expected {
			t.Errorf("Expected IsChineseName(%q, %v, %q) to be %v, got %v", test.host, test.port, test.protocol, test.expected, actual)
		}
	}

	TOs.IsPortOpen("127.0.0.1", 80, "tcp")
	TOs.IsPortOpen("::", 80, "tcp")
	TOs.IsPortOpen("::", 80, "")
	TOs.IsPortOpen("::", 80)
}

func BenchmarkIsPortOpen(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.IsPortOpen("127.0.0.1", 80, "tcp")
	}
}

func TestGetPidByPortGetProcessExeByPid(t *testing.T) {
	message := "Hi there!\n"

	time.AfterFunc(time.Millisecond*200, func() {
		getPidByInode("1234", nil)
		TOs.GetPidByPort(22)
		TOs.GetPidByPort(25)
		TOs.GetPidByPort(1999)
		res := TOs.GetPidByPort(2020)
		exePath := TOs.GetProcessExeByPid(res)
		if res == 0 {
			t.Error("GetPidByPort unit test fail")
			return
		}
		if exePath == "" {
			t.Error("getProcessExeByPid unit test fail")
			return
		}
	})

	time.AfterFunc(time.Millisecond*500, func() {
		conn, err := net.Dial("tcp", ":2020")
		if err != nil {
			t.Fatal(err)
		}
		defer func(conn net.Conn) {
			err = conn.Close()
			if err != nil {

			}
		}(conn)

		if _, err = fmt.Fprintf(conn, message); err != nil {
			t.Fatal(err)
		}
	})

	l, err := net.Listen("tcp", ":2020")
	if err != nil {
		t.Fatal(err)
	}
	defer func(l net.Listener) {
		_ = l.Close()
	}(l)
	for {
		conn, err1 := l.Accept()
		if err1 != nil {
			return
		}

		buf, err2 := ioutil.ReadAll(conn)
		if err2 != nil {
			t.Fatal(err2)
		}

		if msg := string(buf[:]); msg != message {
			t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, message)
		}
		_ = conn.Close()
		return // Done
	}
}

func BenchmarkGetPidByPort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.GetPidByPort(2020)
	}
}

func BenchmarkGetProcessExeByPid(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.GetProcessExeByPid(2020)
	}
}

func TestForceGC(t *testing.T) {
	TOs.ForceGC()
}

func TestTriggerGC(t *testing.T) {
	TOs.TriggerGC()
}

func BenchmarkForceGC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.ForceGC()
	}
}

func BenchmarkTriggerGC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TOs.TriggerGC()
	}
}
package helper

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"

	//"golang.org/x/sys/windows"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	//"syscall"
	"unicode"
	//"unsafe"
)

// SystemInfo 系统信息
type SystemInfo struct {
	ServerName   string  `json:"server_name"`    //服务器名称
	Runtime      int64   `json:"run_time"`       //服务运行时间,纳秒
	GoroutineNum int     `json:"goroutine_num"`  //goroutine数量
	CpuNum       int     `json:"cpu_num"`        //cpu核数
	CpuUser      float64 `json:"cpu_user"`       //cpu用户态比率
	CpuFree      float64 `json:"cpu_free"`       //cpu空闲比率
	DiskUsed     uint64  `json:"disk_used"`      //已用磁盘空间,字节数
	DiskFree     uint64  `json:"disk_free"`      //可用磁盘空间,字节数
	DiskTotal    uint64  `json:"disk_total"`     //总磁盘空间,字节数
	MemUsed      uint64  `json:"mem_used"`       //已用内存,字节数
	MemSys       uint64  `json:"mem_sys"`        //系统内存占用量,字节数
	MemFree      uint64  `json:"mem_free"`       //剩余内存,字节数
	MemTotal     uint64  `json:"mem_total"`      //总内存,字节数
	AllocGolang  uint64  `json:"alloc_golang"`   //golang内存使用量,字节数
	AllocTotal   uint64  `json:"alloc_total"`    //总分配的内存,字节数
	Lookups      uint64  `json:"lookups"`        //指针查找次数
	Mallocs      uint64  `json:"mallocs"`        //内存分配次数
	Frees        uint64  `json:"frees"`          //内存释放次数
	LastGCTime   uint64  `json:"last_gc_time"`   //上次GC时间,纳秒
	NextGC       uint64  `json:"next_gc"`        //下次GC内存回收量,字节数
	PauseTotalNs uint64  `json:"pause_total_ns"` //GC暂停时间总量,纳秒
	PauseNs      uint64  `json:"pause_ns"`       //上次GC暂停时间,纳秒
}

// Echo 输出一个或多个字符串.
func (to *TsOs) Echo(args ...interface{}) {
	log.Print(args...)
}

// Exit 输出一条消息,并退出当前脚本.
func (to *TsOs) Exit(msg interface{}, status int) {
	log.Println(msg)
	os.Exit(status)
}

// Die 函数输出一条消息,并退出当前脚本.
func (to *TsOs) Die(msg interface{}, status int) {
	log.Println(msg)
	os.Exit(status)
}

// IsWindows 当前操作系统是否Windows.
func (to *TsOs) IsWindows() bool {
	return "windows" == runtime.GOOS
}

// IsLinux 当前操作系统是否Linux.
func (to *TsOs) IsLinux() bool {
	return "linux" == runtime.GOOS
}

// IsMac 当前操作系统是否Mac OS/X.
func (to *TsOs) IsMac() bool {
	return "darwin" == runtime.GOOS
}

// Pwd 获取当前程序运行所在的路径,注意和GetWorkList有所不同.
func (to *TsOs) Pwd() (dir string) {
	var ex string
	var err error
	ex, err = os.Executable()
	if err == nil {
		exReal, _ := filepath.EvalSymlinks(ex)
		exReal, _ = filepath.Abs(exReal)
		dir = filepath.Dir(exReal)
	}
	return
}

// GetWorkList 取得当前工作目录(程序可能在任务中进行多次目录切换).
func (to *TsOs) GetWorkList() (dir string, err error) {
	dir, err = os.Getwd()
	return
}

// Chdir 改变/进入新的工作目录.
func (to *TsOs) Chdir(dir string) error {
	return os.Chdir(dir)
}

// HomeDir 获取当前用户的主目录(仅支持Unix-like system).
func (to *TsOs) HomeDir() (string, error) {
	// Unix-like system, so just assume Unix
	home := os.Getenv("HOME")

	usr, err := user.Current()
	if nil == err {
		home = usr.HomeDir
	}

	return home, err
}

// LocalIP 获取本机第一个NIC's IP.
func (to *TsOs) LocalIP() (string, error) {
	ip := ""
	addrArr, err := net.InterfaceAddrs()
	if len(addrArr) > 0 {
		for _, addr := range addrArr {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if nil != ipnet.IP.To4() {
					ip = ipnet.IP.String()
					break
				}
			}
		}
	}

	return ip, err
}

// OutboundIP 获取本机的出口IP.
func (to *TsOs) OutboundIP() (string, error) {
	out := ""
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if conn != nil {
		addr := conn.LocalAddr().(*net.UDPAddr)
		out = addr.IP.String()
		_ = conn.Close()
	}

	return out, err
}

// IsPublicIP 是否公网IP.
func (to *TsOs) IsPublicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := ip.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// GetIPs 获取本机的IP列表.
func (to *TsOs) GetIPs() (ips []string) {
	interfaceAddrArr, _ := net.InterfaceAddrs()
	if len(interfaceAddrArr) > 0 {
		for _, addr := range interfaceAddrArr {
			ipNet, isValidIpNet := addr.(*net.IPNet)
			if isValidIpNet && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					ips = append(ips, ipNet.IP.String())
				}
			}
		}
	}

	return
}

// GetMacAddrArr 获取本机的Mac网卡地址列表.
func (to *TsOs) GetMacAddrArr() (macAddrArr []string) {
	netInterfaces, _ := net.Interfaces()
	if len(netInterfaces) > 0 {
		for _, netInterface := range netInterfaces {
			macAddr := netInterface.HardwareAddr.String()
			if len(macAddr) == 0 {
				continue
			}
			macAddrArr = append(macAddrArr, macAddr)
		}
	}

	return
}

// Hostname 获取主机名.
func (to *TsOs) Hostname() (string, error) {
	return os.Hostname()
}

// GetIpByHostname 返回主机名对应的 IPv4地址.
func (to *TsOs) GetIpByHostname(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		for _, v := range ips {
			if v.To4() != nil {
				return v.String(), nil
			}
		}
		return "", nil
	}
	return "", err
}

// GetIpsByDomain 获取互联网域名/主机名对应的 IPv4 地址列表.
func (to *TsOs) GetIpsByDomain(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if ips != nil {
		var ipStrArr []string
		for _, v := range ips {
			if v.To4() != nil {
				ipStrArr = append(ipStrArr, v.String())
			}
		}
		return ipStrArr, nil
	}
	return nil, err
}

// GetHostByIp 获取指定的IP地址对应的主机名.
func (to *TsOs) GetHostByIp(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if names != nil {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}

// GoMemory MemoryGetUsage 获取当前go程序的内存使用,返回字节数.
func (to *TsOs) GoMemory() uint64 {
	stat := new(runtime.MemStats)
	runtime.ReadMemStats(stat)
	return stat.Alloc
}

// MemoryUsage 获取内存使用率(仅支持linux),单位字节.
// 参数 virtual,是否取虚拟内存.
// used为已用,
// free为空闲,
// total为总数.
// win 与linux 不通用 废弃.
//func (to *TsOs) MemoryUsage(virtual bool) (used, free, total uint64) {
//	if virtual {
//		// 虚拟机的内存
//		contents, err := ioutil.ReadFile("/proc/memento")
//		if err == nil {
//			lines := strings.Split(string(contents), "\n")
//			for _, line := range lines {
//				fields := strings.Fields(line)
//				if len(fields) == 3 {
//					val, _ := strconv.ParseUint(fields[1], 10, 64) // kB
//
//					if strings.HasPrefix(fields[0], "MemTotal") {
//						total = val * 1024
//					} else if strings.HasPrefix(fields[0], "MemFree") {
//						free = val * 1024
//					}
//				}
//			}
//
//			//计算已用内存
//			used = total - free
//		}
//	} else {
//		// 真实物理机内存
//		sys := &syscall.Sysinfo_t{}
//		err := syscall.Sysinfo(sys)
//		if err == nil {
//			total = sys.Totalram * uint64(syscall.Getpagesize()/1024)
//			free = sys.Freeram * uint64(syscall.Getpagesize()/1024)
//			used = total - free
//		}
//	}
//
//	return
//}

// CpuUsage 获取CPU使用率(仅支持linux),单位jiffies(节拍数).
// user为用户态(用户进程)的运行时间,
// idle为空闲时间,
// total为累计时间.
func (to *TsOs) CpuUsage() (user, idle, total uint64) {
	contents, _ := ioutil.ReadFile("/proc/stat")
	if len(contents) > 0 {
		lines := strings.Split(string(contents), "\n")
		for _, line := range lines {
			fields := strings.Fields(line)
			if fields[0] == "cpu" {
				//CPU指标：user，nice, system, idle, iowait, irq, softirq
				// cpu  130216 19944 162525 1491240 3784 24749 17773 0 0 0

				numFields := len(fields)
				for i := 1; i < numFields; i++ {
					val, _ := strconv.ParseUint(fields[i], 10, 64)
					total += val // tally up all the numbers to get total ticks
					if i == 1 {
						user = val
					} else if i == 4 { // idle is the 5th field in the cpu line
						idle = val
					}
				}
				break
			}
		}
	}

	return
}

// DiskUsage 获取磁盘/目录使用情况,单位字节.参数path为目录.
// used为已用, free为空闲, total为总数.
// win 与 linux mac不通用 已废弃.
//func (to *TsOs) DiskUsage(path string) (used, free, total uint64) {
//	//if runtime.GOOS == "linux" {
//	//	fs := &syscall.Statfs_t{}
//	//	err := syscall.Statfs(path, fs)
//	//	if err == nil {
//	//		total = fs.Blocks * uint64(fs.Bsize)
//	//		free = fs.Bfree * uint64(fs.Bsize)
//	//		used = total - free
//	//	}
//	//}
//	if runtime.GOOS == "windows" {
//		h := windows.MustLoadDLL("kernel32.dll")
//		c := h.MustFindProc("GetDiskFreeSpaceExW")
//		lpFreeBytesAvailable := uint64(0)
//		lpTotalNumberOfBytes := uint64(0)
//		lpTotalNumberOfFreeBytes := uint64(0)
//		_, _, err := c.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("C:"))),
//			uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
//			uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
//			uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
//		if err == nil {
//			total = lpTotalNumberOfBytes
//			free = lpTotalNumberOfFreeBytes
//			used = lpFreeBytesAvailable
//			return
//		}
//	}
//	return
//}

// Setenv 设置一个环境变量的值.
func (to *TsOs) Setenv(varName, data string) error {
	return os.Setenv(varName, data)
}

// Getenv 获取一个环境变量的值.
func (to *TsOs) Getenv(varName string) string {
	return os.Getenv(varName)
}

// GetEndian 获取系统字节序类型,小端返回binary.LittleEndian,大端返回binary.BigEndian .
func (to *TsOs) GetEndian() binary.ByteOrder {
	return getEndian()
}

// IsLittleEndian 系统字节序类型是否小端存储.
func (to *TsOs) IsLittleEndian() bool {
	return isLittleEndian()
}

// Exec 执行一个外部命令.
// retInt为1时失败,为0时成功;outStr为执行命令的输出;errStr为错误输出.
// 命令如
// "ls -a"
// "/bin/bash -c \"ls -a\""
func (to *TsOs) Exec(command string) (retInt int, outStr, errStr []byte) {
	// split command
	q := rune(0)
	parts := strings.FieldsFunc(command, func(r rune) bool {
		switch {
		case r == q:
			q = rune(0)
			return false
		case q != rune(0):
			return false
		case unicode.In(r, unicode.Quotation_Mark):
			q = r
			return false
		default:
			return unicode.IsSpace(r)
		}
	})

	// remove the " and ' on both sides
	for i, v := range parts {
		f, l := v[0], len(v)
		if l >= 2 && (f == '"' || f == '\'') {
			parts[i] = v[1 : l-1]
		}
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		retInt = 1 //失败
		stderr.WriteString(err.Error())
		errStr = stderr.Bytes()
	} else {
		retInt = 0 //成功
		outStr, errStr = stdout.Bytes(), stderr.Bytes()
	}

	return
}

// System 与Exec相同,但会同时打印标准输出和标准错误.
func (to *TsOs) System(command string) (retInt int, outStr, errStr []byte) {
	// split command
	q := rune(0)
	parts := strings.FieldsFunc(command, func(r rune) bool {
		switch {
		case r == q:
			q = rune(0)
			return false
		case q != rune(0):
			return false
		case unicode.In(r, unicode.Quotation_Mark):
			q = r
			return false
		default:
			return unicode.IsSpace(r)
		}
	})

	// remove the " and ' on both sides
	for i, v := range parts {
		f, l := v[0], len(v)
		if l >= 2 && (f == '"' || f == '\'') {
			parts[i] = v[1 : l-1]
		}
	}

	var stdout, stderr bytes.Buffer
	var err error

	cmd := exec.Command(parts[0], parts[1:]...)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	outWr := io.MultiWriter(os.Stdout, &stdout)
	errWr := io.MultiWriter(os.Stderr, &stderr)

	err = cmd.Start()
	if err != nil {
		retInt = 1 //失败
		stderr.WriteString(err.Error())
		fmt.Printf("%s\n", stderr.Bytes())
		return
	}

	go func() {
		_, _ = io.Copy(outWr, stdoutIn)
	}()
	go func() {
		_, _ = io.Copy(errWr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		stderr.WriteString(err.Error())
		fmt.Println(stderr.Bytes())
		retInt = 1 //失败
	} else {
		retInt = 0 //成功
	}
	outStr, errStr = stdout.Bytes(), stderr.Bytes()

	return
}

// Chmod 改变文件模式.
func (to *TsOs) Chmod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}

// Chown 改变文件的所有者.
func (to *TsOs) Chown(filename string, uid, gid int) bool {
	return os.Chown(filename, uid, gid) == nil
}

// GetTempDir 返回用于临时文件的目录.
func (to *TsOs) GetTempDir() string {
	return os.TempDir()
}

// PrivateCIDR 获取私有网段的CIDR(无类别域间路由).
func (to *TsOs) PrivateCIDR() []*net.IPNet {
	maxCidrBlocks := []string{
		"127.0.0.1/8",    // localhost
		"10.0.0.0/8",     // 24-bit block
		"172.16.0.0/12",  // 20-bit block
		"192.168.0.0/16", // 16-bit block
		"169.254.0.0/16", // link local address
		"::1/128",        // localhost IPv6
		"fc00::/7",       // unique local address IPv6
		"fe80::/10",      // link local address IPv6
	}

	ipNets := make([]*net.IPNet, len(maxCidrBlocks))
	for i, maxCidrBlock := range maxCidrBlocks {
		_, cidr, _ := net.ParseCIDR(maxCidrBlock)
		ipNets[i] = cidr
	}

	return ipNets
}

// IsPrivateIp 是否私有IP地址(ipv4/ipv6).
func (to *TsOs) IsPrivateIp(address string) (bool, error) {
	ip := net.ParseIP(address)
	if ip == nil {
		return false, errors.New("address is not valid ip")
	}

	if TPrivyCiders == nil {
		TPrivyCiders = to.PrivateCIDR()
	}
	for i := range TPrivyCiders {
		if TPrivyCiders[i].Contains(ip) {
			return true, nil
		}
	}

	return false, nil
}

// ClientIp 获取客户端真实IP,req为http请求.
func (to *TsOs) ClientIp(req *http.Request) string {
	// 获取头部信息,有可能是代理
	xRealIP := req.Header.Get("X-Real-Ip")
	xForwardedFor := req.Header.Get("X-Forwarded-For")

	// If both empty, return IP from remote address
	if xRealIP == "" && xForwardedFor == "" {
		var remoteIP string

		// If there are colon in remote address, remove the port number
		// otherwise, return remote address as is
		if strings.ContainsRune(req.RemoteAddr, ':') {
			remoteIP, _, _ = net.SplitHostPort(req.RemoteAddr)
		} else {
			remoteIP = req.RemoteAddr
		}

		return remoteIP
	}

	// Check list of IP in X-Forwarded-For and return the first global address
	// X-Forwarded-For是逗号分隔的IP地址列表,如"10.0.0.1, 10.0.0.2, 10.0.0.3"
	for _, address := range strings.Split(xForwardedFor, ",") {
		address = strings.TrimSpace(address)
		isPrivate, err := to.IsPrivateIp(address)
		if !isPrivate && err == nil {
			return address
		}
	}

	if xRealIP == "::1" {
		xRealIP = "127.0.0.1"
	}

	// If nothing succeed, return X-Real-IP
	return xRealIP
}

// GetSystemInfo 获取系统运行信息.
// 内存等win/linux/mac不通用,已废弃.
func (to *TsOs) GetSystemInfo() *SystemInfo {
	//运行时信息
	mstat := &runtime.MemStats{}
	runtime.ReadMemStats(mstat)

	//CPU信息
	cpuUser, cpuIdel, cpuTotal := to.CpuUsage()
	cpuUserRate := float64(cpuUser) / float64(cpuTotal)
	cpuFreeRate := float64(cpuIdel) / float64(cpuTotal)

	//磁盘空间信息
	//diskUsed, diskFree, diskTotal := to.DiskUsage("/")

	//内存使用信息
	//memUsed, memFree, memTotal := to.MemoryUsage(true)

	serverName, _ := os.Hostname()

	return &SystemInfo{
		ServerName:   serverName,
		Runtime:      int64(TTime.ServiceUptime()),
		GoroutineNum: runtime.NumGoroutine(),
		CpuNum:       runtime.NumCPU(),
		CpuUser:      cpuUserRate,
		CpuFree:      cpuFreeRate,
		//DiskUsed:     diskUsed,
		DiskUsed: 0,
		//DiskFree:     diskFree,
		DiskFree: 0,
		//DiskTotal:    diskTotal,
		DiskTotal: 0,
		//MemUsed:      memUsed,
		MemUsed: 0,
		MemSys:  mstat.Sys,
		//MemFree:      memFree,
		MemFree: 0,
		//MemTotal:     memTotal,
		MemTotal:     0,
		AllocGolang:  mstat.Alloc,
		AllocTotal:   mstat.TotalAlloc,
		Lookups:      mstat.Lookups,
		Mallocs:      mstat.Mallocs,
		Frees:        mstat.Frees,
		LastGCTime:   mstat.LastGC,
		NextGC:       mstat.NextGC,
		PauseTotalNs: mstat.PauseTotalNs,
		PauseNs:      mstat.PauseNs[(mstat.NumGC+255)%256],
	}
}

// IsPortOpen 检查主机端口是否开放.protocols为协议名称,可选,默认tcp.
func (to *TsOs) IsPortOpen(host string, port interface{}, protocols ...string) bool {
	if TValidate.IsHost(host) && TValidate.IsPort(port) {
		// 默认tcp协议
		protocol := "tcp"
		if len(protocols) > 0 && len(protocols[0]) > 0 {
			protocol = strings.ToLower(protocols[0])
		}

		conn, _ := net.DialTimeout(protocol, net.JoinHostPort(host, TConv.ToStr(port)), CheckConnectTimeout)
		if conn != nil {
			_ = conn.Close()
			return true
		}
	}

	return false
}

//GetPidByPort 根据端口号获取监听的进程PID.
func (to *TsOs) GetPidByPort(port int) (pid int) {
	files := []string{
		"/proc/net/tcp",
		"/proc/net/udp",
		"/proc/net/tcp6",
		"/proc/net/udp6",
	}

	procDirs, _ := filepath.Glob("/proc/[0-9]*/fd/[0-9]*")
	for _, filePath := range files {
		lines, _ := TFile.ReadInArray(filePath)
		for _, line := range lines[1:] {
			fields := strings.Fields(line)
			if len(fields) < 10 {
				continue
			}

			//非 LISTEN 监听状态
			if fields[3] != "0A" {
				continue
			}

			//本地ip和端口
			ipport := strings.Split(fields[1], ":")
			locPort, _ := TConv.Hex2Dec(ipport[1])

			// 非该端口
			if int(locPort) != port {
				continue
			}
			pid = getPidByInode(fields[9], procDirs)
			if pid > 0 {
				return
			}
		}
	}
	return
}

// GetProcessExeByPid 根据PID获取进程的执行路径.
func (to *TsOs) GetProcessExeByPid(pid int) string {
	return getProcessExeByPid(pid)
}

// ForceGC 强制手动GC垃圾回收(阻塞).
func (to *TsOs) ForceGC() {
	runtime.GC()
	debug.FreeOSMemory()
}

// TriggerGC 触发GC(非阻塞).
func (to *TsOs) TriggerGC() {
	go func() {
		to.ForceGC()
	}()
}

package Helper

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"regexp"
	"unicode"

	"github.com/forgoer/openssl"
	"github.com/gin-gonic/gin"
)

const (
	//密码验证选项 只能含有
	PWD_OPT_Number  uint16 = 1 << iota //数字 	 0001
	PWD_OPT_Lower                      //小写 	 0010
	PWD_OPT_Upper                      //大写 	 0100
	PWD_OPT_Special                    //特殊符号 1000
)

//加密过程：
//  1、处理数据，对数据进行填充，采用PKCS7（当密钥长度不够时，缺几位补几个几）的方式。
//  2、对数据进行加密，采用AES加密方法中CBC加密模式
//  3、对得到的加密数据，进行base64加密，得到字符串
// 解密过程相反

// 16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法
// key不能泄露
var PwdKey = []byte("@8KPDy5OFwauml&owZRksreESAONkZ*y")

// AES解密 部分post接口传参用
func AESDecrypt(c *gin.Context) (result string, err error) {
	receiveData, _ := ioutil.ReadAll(c.Request.Body)
	dst, err := openssl.AesECBDecrypt(receiveData, PwdKey, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	} else {
		return string(dst), nil
	}

}

func AESDecrypt_New(data string) (result string, err error) {
	dataByte, _ := base64.StdEncoding.DecodeString(data)
	dst, err := openssl.AesECBDecrypt(dataByte, PwdKey, openssl.PKCS7_PADDING)
	if err != nil {
		return "", err
	} else {
		return string(dst), nil
	}
}

func padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func AesEncryptECB(origData []byte) (string, error) {

	//key只能是 16 24 32长度
	block, err := aes.NewCipher(PwdKey)
	if err != nil {
		return "", err
	}
	//padding
	origData = padding(origData, block.BlockSize())
	//存储每次加密的数据

	//分组分块加密
	buffer := bytes.NewBufferString("")
	tmpData := make([]byte, block.BlockSize()) //存储每次加密的数据
	for index := 0; index < len(origData); index += block.BlockSize() {
		block.Encrypt(tmpData, origData[index:index+block.BlockSize()])
		buffer.Write(tmpData)
	}

	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

// 密码强度必须为字⺟⼤⼩写+数字+符号，密码长度为6-16位
func CheckPasswordLever(ps string) error {
	if len(ps) < 6 || len(ps) > 16 {
		return fmt.Errorf("密码长度为6-16位")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err != nil {
		return fmt.Errorf("password need num :%v", err)
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err != nil {
		return fmt.Errorf("password need a_z :%v", err)
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err != nil {
		return fmt.Errorf("password need A_Z :%v", err)
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err != nil {
		return fmt.Errorf("password need symbol :%v", err)
	}
	return nil
}

// 密码长度为6-16位，密码不能为纯数字或纯英文
func VerifyPwd(pwd string, options uint16) bool {
	if pwd == "" {
		return false
	}
	if len(pwd) < 6 || len(pwd) > 16 {
		return false
	}
	// 用于记录验证结果
	var result uint16
	for _, r := range pwd {
		switch {
		case unicode.IsNumber(r):
			result = result | PWD_OPT_Number
		case unicode.IsLower(r):
			result = result | PWD_OPT_Lower
		case unicode.IsUpper(r):
			result = result | PWD_OPT_Upper
		case unicode.IsPunct(r) || unicode.IsSymbol(r): //标点符号 和 字符
			result = result | PWD_OPT_Special
		default:
			return false
		}
		// 比较结果和设置项
		// 当 cp.options&cp.result != cp.result 表示密码字符串超出 options 范围
		if options&result != result {
			return false
		}
	}
	return true
}

// func GetPhysicalID() string {
// 	var ids []string
// 	if guid, err := GetMachineGuid(); err != nil {
// 		panic(err.Error())
// 	} else {
// 		ids = append(ids, guid)
// 	}
// 	if cpuinfo, err := GetCPUInfo(); err != nil && len(cpuinfo) > 0 {
// 		panic(err.Error())
// 	} else {
// 		ids = append(ids, cpuinfo[0].VendorID+cpuinfo[0].PhysicalID)
// 	}
// 	if mac, err := GetMACAddress(); err != nil {
// 		panic(err.Error())
// 	} else {
// 		ids = append(ids, mac)
// 	}
// 	sort.Strings(ids)
// 	idsstr := strings.Join(ids, "|/|")
// 	return GetMd5String(idsstr, true, true)
// }

// //获取MAC地址
// func GetMACAddress() (string, error) {
// 	netInterfaces, err := net.Interfaces()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	mac, macerr := "", errors.New("无法获取到正确的MAC地址")
// 	for i := 0; i < len(netInterfaces); i++ {
// 		//fmt.Println(netInterfaces[i])
// 		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
// 			addrs, _ := netInterfaces[i].Addrs()
// 			for _, address := range addrs {
// 				ipnet, ok := address.(*net.IPNet)
// 				//fmt.Println(ipnet.IP)
// 				if ok && ipnet.IP.IsGlobalUnicast() {
// 					// 如果IP是全局单拨地址，则返回MAC地址
// 					mac = netInterfaces[i].HardwareAddr.String()
// 					return mac, nil
// 				}
// 			}
// 		}
// 	}
// 	return mac, macerr
// }

// type cpuInfo struct {
// 	CPU        int32  `json:"cpu"`
// 	VendorID   string `json:"vendorId"`
// 	PhysicalID string `json:"physicalId"`
// }

// type win32_Processor struct {
// 	Manufacturer string
// 	ProcessorID  *string
// }

// func GetCPUInfo() ([]cpuInfo, error) {
// 	var ret []cpuInfo
// 	var dst []win32_Processor
// 	q := wmi.CreateQuery(&dst, "")
// 	fmt.Println(q)
// 	if err := WmiQuery(q, &dst); err != nil {
// 		return ret, err
// 	}

// 	var procID string
// 	for i, l := range dst {
// 		procID = ""
// 		if l.ProcessorID != nil {
// 			procID = *l.ProcessorID
// 		}

// 		cpu := cpuInfo{
// 			CPU:        int32(i),
// 			VendorID:   l.Manufacturer,
// 			PhysicalID: procID,
// 		}
// 		ret = append(ret, cpu)
// 	}

// 	return ret, nil
// }

// // WMIQueryWithContext - wraps wmi.Query with a timed-out context to avoid hanging
// func WmiQuery(query string, dst interface{}, connectServerArgs ...interface{}) error {
// 	ctx := context.Background()
// 	if _, ok := ctx.Deadline(); !ok {
// 		ctxTimeout, cancel := context.WithTimeout(ctx, 3000000000) //超时时间3s
// 		defer cancel()
// 		ctx = ctxTimeout
// 	}

// 	errChan := make(chan error, 1)
// 	go func() {
// 		errChan <- wmi.Query(query, dst, connectServerArgs...)
// 	}()

// 	select {
// 	case <-ctx.Done():
// 		return ctx.Err()
// 	case err := <-errChan:
// 		return err
// 	}
// }

// func GetMachineGuid() (string, error) {
// 	// there has been reports of issues on 32bit using golang.org/x/sys/windows/registry, see https://github.com/shirou/gopsutil/pull/312#issuecomment-277422612
// 	// for rationale of using windows.RegOpenKeyEx/RegQueryValueEx instead of registry.OpenKey/GetStringValue
// 	var h windows.Handle
// 	err := windows.RegOpenKeyEx(windows.HKEY_LOCAL_MACHINE, windows.StringToUTF16Ptr(`SOFTWARE\Microsoft\Cryptography`), 0, windows.KEY_READ|windows.KEY_WOW64_64KEY, &h)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer windows.RegCloseKey(h)

// 	const windowsRegBufLen = 74 // len(`{`) + len(`abcdefgh-1234-456789012-123345456671` * 2) + len(`}`) // 2 == bytes/UTF16
// 	const uuidLen = 36

// 	var regBuf [windowsRegBufLen]uint16
// 	bufLen := uint32(windowsRegBufLen)
// 	var valType uint32
// 	err = windows.RegQueryValueEx(h, windows.StringToUTF16Ptr(`MachineGuid`), nil, &valType, (*byte)(unsafe.Pointer(&regBuf[0])), &bufLen)
// 	if err != nil {
// 		return "", err
// 	}

// 	hostID := windows.UTF16ToString(regBuf[:])
// 	hostIDLen := len(hostID)
// 	if hostIDLen != uuidLen {
// 		return "", fmt.Errorf("HostID incorrect: %q\n", hostID)
// 	}

// 	return hostID, nil
// }

// //生成32位md5字串
// func GetMd5String(s string, upper bool, half bool) string {
// 	h := md5.New()
// 	h.Write([]byte(s))
// 	result := hex.EncodeToString(h.Sum(nil))
// 	if upper == true {
// 		result = strings.ToUpper(result)
// 	}
// 	if half == true {
// 		result = result[8:24]
// 	}
// 	return result
// }

// //利用随机数生成Guid字串
// func UniqueId() string {
// 	b := make([]byte, 48)
// 	if _, err := rand.Read(b); err != nil {
// 		return ""
// 	}
// 	return GetMd5String(base64.URLEncoding.EncodeToString(b), true, false)
// }

// Md5加密
func Md5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

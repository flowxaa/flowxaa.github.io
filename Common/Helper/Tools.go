package Helper

import (
	"math"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// IsContain 判断string 数组是否包含 字符串 返回true 为包含
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func IsIntContain(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// GetPageCount 分页返回总页数
func GetPageCount(totalcount int64, pagesize int64) int64 {
	var pagecount int64 = 0
	if totalcount%pagesize == 0 {
		pagecount = totalcount / pagesize
	} else {
		pagecount = totalcount/pagesize + 1
	}
	return pagecount
}

func GetProcName(areaId string, procName string) string {

	if areaId != "0" {
		return "exec " + procName + "_tw"
	} else {
		return "exec " + procName
	}
}

func GetProcNameChoice(areaId string, procName string, procName2 string) string {
	if areaId == "0" {
		return "exec " + procName
	} else {
		return "exec " + procName2
	}
}

// ConverModelList []map[string]interface{} 转 []T
func ConverModelList[T any](mapdata []map[string]interface{}) (s []T) {
	for _, k := range mapdata {
		var t T
		mapstructure.Decode(k, &t)
		s = append(s, t)
	}
	return s
}

// GetQueryIntParam 获取int类型参数
func GetQueryIntParam(ctx *gin.Context, key string) (result int) {

	result, err := strconv.Atoi(ctx.Query(key))
	if err != nil {
		result = 0
	}
	return result
}

// GetQueryIntParamD 获取int类型参数,加默认值
func GetQueryIntParamD(ctx *gin.Context, key string, defaultval int) (result int) {
	result, err := strconv.Atoi(ctx.Query(key))
	if err != nil {
		result = defaultval
	}
	return result
}

// GetQueryInt64Param 获取int64类型参数
func GetQueryInt64Param(ctx *gin.Context, key string) (result int64) {
	result, err := strconv.ParseInt(ctx.Query(key), 10, 64)
	if err != nil {
		result = 0
	}
	return result
}

// GetQueryInt64ParamD 获取int64类型参数,加默认值
func GetQueryInt64ParamD(ctx *gin.Context, key string, defaultval int64) (result int64) {
	result, err := strconv.ParseInt(ctx.Query(key), 10, 64)
	if err != nil {
		result = defaultval
	}
	return result
}

// IpToInt 将IP地址转为整数形
func IpToInt(ip string) int64 {
	ipArr := strings.Split(ip, ".")
	if len(ipArr) == 3 {
		ip = ip + ".0"
		ipArr = strings.Split(ip, ".")
	}

	var ip_int int64 = 0
	p1, _ := strconv.ParseInt(ipArr[0], 10, 64)
	p1 = p1 * 256 * 256 * 256
	p2, _ := strconv.ParseInt(ipArr[1], 10, 64)
	p2 = p2 * 256 * 256
	p3, _ := strconv.ParseInt(ipArr[2], 10, 64)
	p3 = p3 * 256
	p4, _ := strconv.ParseInt(ipArr[3], 10, 64)

	ip_int = p1 + p2 + p3 + p4

	return ip_int
}

// Wrap 将float64转成精确的int64
func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

// Unwrap 将int64恢复成正常的float64
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

// WrapToFloat64 精准float64
func WrapToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}

// UnwrapToInt64 精准int64
func UnwrapToInt64(num int64, retain int) int64 {
	return int64(Unwrap(num, retain))
}

// GetHourDiffer 计算两个时间差 （返回小时）
func GetHourDiffer(start_time, end_time time.Time) int64 {
	var hour int64
	t1 := start_time
	t2 := end_time
	if t1.Before(t2) {
		diff := t2.Unix() - t1.Unix() //
		hour = diff / 3600
		return hour
	} else {
		return hour
	}
}

// GetFirstDateOfMonth 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// GetLastDateOfMonth 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// GetZeroTime 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// int64 转 int
func Int64ToInt(val int64) int {
	strval := strconv.FormatInt(val, 10)
	retval, err := strconv.Atoi(strval)
	if err != nil {
		return 0
	}
	return retval
}

func ChangeApiPath(Path string) (result string) {
	if strings.Contains(Path, "ng/uptHeadimg") {
		result = strings.Replace(Path, "ng/uptHeadimg", "UserInfo/Upload", 1)
	}
	if strings.Contains(Path, "ng/uplHeadimg") {
		result = strings.Replace(Path, "ng/uplHeadimg", "Upload/UserPhoto", -1)
	}
	if strings.Contains(Path, "ng/regFirst") {
		result = strings.Replace(Path, "ng/regFirst", "Account/CheckMobileCG", -1)
	}
	if strings.Contains(Path, "ng/regSecond") {
		result = strings.Replace(Path, "ng/regSecond", "Account/RegisterPhoneCG", -1)
	}
	if strings.Contains(Path, "ng/apLogin") {
		result = strings.Replace(Path, "ng/apLogin", "Account/AppleLoginCG", -1)
	}
	if strings.Contains(Path, "ng/LTaskState") {
		result = strings.Replace(Path, "ng/LTaskState", "Room/GetLuxuryTaskState", -1)
	}
	if strings.Contains(Path, "ng/jyyz") {
		result = strings.Replace(Path, "ng/jyyz", "Account/GetCaptcha", -1)
	}
	if strings.Contains(Path, "ng/logout") {
		result = strings.Replace(Path, "ng/logout", "Account/CheckLogOut", -1)
	}
	if strings.Contains(Path, "ng/authSubmit") {
		result = strings.Replace(Path, "ng/authSubmit", "Passport/HumanAuth", -1)
	}
	if strings.Contains(Path, "ng/wcLogin") {
		result = strings.Replace(Path, "ng/wcLogin", "Account/WeiXinLoginQM", -1)
	}
	if strings.Contains(Path, "user/quickLogin") {
		result = strings.Replace(Path, "user/quickLogin", "Account/PhoneQuickLogin", -1)
	}
	if strings.Contains(Path, "user/gquickLogin") {
		result = strings.Replace(Path, "user/gquickLogin", "Account/PhoneQuickLogin", -1)
	}
	if strings.Contains(Path, "ng/smsLogin") {
		result = strings.Replace(Path, "ng/smsLogin", "Account/SMSLogin", -1)
	}
	if strings.Contains(Path, "mxgj/hidState") {
		result = strings.Replace(Path, "mxgj/hidState", "About/isHidden", -1)
	}
	if strings.Contains(Path, "mxzj/checkCode") {
		result = strings.Replace(Path, "mxzj/checkCode", "Account/VerifySmsCode", -1)
	}
	if result == "" {
		result = Path
	}
	return result
}

func NewProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.Host = target.Host
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = ChangeApiPath(req.URL.Path)
		req.URL.Path, req.URL.RawPath = newJoinURLPath(target, req.URL)
		//req.RemoteAddr = strings.Replace(req.RemoteAddr, "127.0.0.1", "10.199.4.140", 1)
		//req.RemoteAddr = strings.Replace(req.RemoteAddr, "[::1]", "10.199.4.140", 1)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
	return &httputil.ReverseProxy{Director: director}
}
func newJoinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return newSingleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

func newSingleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

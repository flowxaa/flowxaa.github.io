package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	//Helper "songerDailyWage/Common/Helper"
	logger "WebSocketDemo/Common/log"
	//"songerDailyWage/SqlServerHelper"
)

func PostData() {

	params := "signature + dddd + sign"
	//logger.LogOfInfo("url=", "http://cps.leyoubuji.com/caiyu-cps/cashout/balance?"+params)
	resp, err := http.Post("http://127.0.0.1:8083/index", "application/x-www-form-urlencoded", strings.NewReader(params))
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		logger.LogOfInfo("调用API接口失败:", err)
	}

	//logger.LogOfInfo("resp= ", resp.Body)

	//获取返回值
	contents, _ := ioutil.ReadAll(resp.Body)
	//[]byte转*json 返回*Json,err json 是个结构体
	//result, _ := simplejson.NewJson(contents)
	fmt.Println(string(contents))
}

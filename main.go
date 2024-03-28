package main

import "WebSocketDemo/route"

func main() {
	//api.PostData()

	//	获取路由
	ro := route.GetRoute()
	//	监听端口
	_ = ro.Run("0.0.0.0:8083")

	// ticker := time.NewTicker(10 * time.Second)
	// defer ticker.Stop()

	// for range ticker.C {
	// 	wk := time.Now().Weekday().String()
	// 	fmt.Println("定时执行  当前时间 " + time.Now().Format("15:04:05") + " " + wk)

	// 	iclock := "16:11:11" //pub.ReadLinString("config", "clock.ctime")
	// 	fmt.Println("接口执行时间iclock = " + iclock)

	// 	//周一至周五执行 周六日不执行
	// 	if wk == "Saturday" || wk == "Sunday" {
	// 		continue
	// 	}

	// 	nowStr := time.Now().Format("15:04:05")

	// 	//iclock = "11111"

	// 	//if nowStr == iclock {
	// 	fmt.Println("执行接口nowStr = ", nowStr)
	// 	//continue
	// 	api.PostData()

	// 	// resp, err := http.Get(apiUrl)
	// 	// if resp != nil {
	// 	// 	defer resp.Body.Close()
	// 	// }
	// 	// if err != nil {
	// 	// 	fmt.Printf("调用API接口失败: %v \n", err)
	// 	// }

	// 	// fmt.Printf("resp= %v\n", resp)
	// 	fmt.Println("执行接口完成")
	// 	//}

	// }

}

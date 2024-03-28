package Helper

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ChangeHost(hostString string, ctx *gin.Context) string {
	host := ctx.Request.Host

	nhost := strings.Replace(host, strings.Split(host, ".")[0]+".", "", 1)

	var ImgHostStr string
	var MallHostStr string
	if strings.Contains(nhost, "qingyajiu.com") {
		ImgHostStr = "res.qingyajiu.com"
		MallHostStr = "mall.qingyajiu.com"
	} else if strings.Contains(nhost, "lxjb.com") {
		ImgHostStr = "res.lxjb.com"
		MallHostStr = "mall.lxjb.com"
	} else if strings.Contains(nhost, "dizhou.net") {
		ImgHostStr = "res.dizhou.net"
		MallHostStr = "mall.dizhou.net"
	} else if strings.Contains(nhost, "szchhkj.cn") {
		// ImgHostStr = "res.szchhkj.cn"
		// MallHostStr = "mall.szchhkj.cn"
		ImgHostStr = "res.lxjb.com"
		MallHostStr = "mall.lxjb.com"
	} else {
		ImgHostStr = "res.qingyajiu.com"
		MallHostStr = "mall.qingyajiu.com"
	}

	hostString = strings.Replace(hostString, "liveimg.miaobolive.com", ImgHostStr, 1)
	hostString = strings.Replace(hostString, "liveimg.9158.com", ImgHostStr, 1)
	hostString = strings.Replace(hostString, "liveimg.maozhuazb.com", ImgHostStr, 1)
	hostString = strings.Replace(hostString, "live.maozhuazb.com", MallHostStr, 1)
	hostString = strings.Replace(hostString, "mc.maozhuazb.com", MallHostStr, 1)

	return hostString
}

var tool = {
    getQueryStr: function (name) {
        var reg = new RegExp("(^|\\?|&)" + name + "=([^&]*)(\\s|&|$)", "i");
        if (reg.test(location.href)) return decodeURIComponent(RegExp.$2.replace(/\+/g, " "));
        return "";
    },
    setalter: function (content) {
        layer.open({
            content: content, skin: 'msg', time: 2
        });
    },
    client: function () {
        var u = navigator.userAgent;
        var isAndroid = u.toLowerCase().indexOf('android') > -1 || u.indexOf('Adr') > -1; //android终端
        var isiOS = !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/); //ios终端
        if (isiOS || u.indexOf('iPad') > -1 || !isAndroid) {
            phone = "ios";
        } else {
            phone = "android";
        }
        return phone;
    },
    setAppHeight(type){
        //设置APP显示高度 type 1:4分之1屏 2:半屏 3:4分之3屏 4:全屏 5:自适应
        try{
            if (tool.client() == "ios") {
                //dsBridge.call('SetHeight',type);
                window.webkit.messageHandlers.SetHeight.postMessage({"type":type});
                //window.location.href = "SetHeight:type:"+type;
            } else {
                return android.SetHeight(type);
            }
        }catch(e){
            //console.log('出现错误, 如果在非android环境下访问, 出现该警告是正常的.')
        }
    },
    UrlReplaceRandom(param){
        if(param != "" && param.lastIndexOf("&r=") > -1){
            var len =param.lastIndexOf("&r=");
            var b = param.substr(len + 3,param.length); 
            param = param.replace(b,Math.random());
        }
        return param;
    }
};
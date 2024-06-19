<%@ WebHandler Language="C#" Class="goldwords" %>

using System;
using System.Web;
using System.Net;


using System.IO;
using System.Web.Security;
using System.Net.Sockets;


using System.Security.Authentication;
using System.Security.Cryptography.X509Certificates;

using LitJson;
using System.Text;

public class goldwords : IHttpHandler {

    public void ProcessRequest (HttpContext context) {
        //context.Response.ContentType = "text/plain";
        //context.Response.Write("Hello World");

        string words = GetHttpWebResponse("https://api.lovelive.tools/api/SweetNothings","");//GetJsonData("http://www.flowx.cn/api/getGoldWords.aspx","","");
        if (words != "")
        {
            context.Response.Write("{\"code\":\"ok\",\"formJson\":\"" + words + "\"}");
            return;
        }

    }

    public bool IsReusable {
        get {
            return false;
        }
    }


    public string GetJson() {
        return "";
    }
    public static string GetJsonData(string tmpUrlI, string tmpKeyName, string tmpKeyValue)
    {
        string str = "";
        //获取请求
        HttpWebRequest request = WebRequest.Create(tmpUrlI) as HttpWebRequest;
        //加入请求头
       // request.Headers.Add(tmpKeyName, tmpKeyValue);
        using (HttpWebResponse response = request.GetResponse() as HttpWebResponse)
        {
            //获取响应数据流
            StreamReader reader = new StreamReader(response.GetResponseStream(), Encoding.UTF8);
            //获得json字符串
            string jsonstr = reader.ReadLine();


            JsonData jsonData = JsonMapper.ToObject(jsonstr);//JsonMapper.ToObject(File.ReadAllText("F://2019/api.flowxa.com/www.flowx.cn/data.json"));//
            if (jsonData != null)
            {
                //取library下键值字典
                JsonData jsonDicList = jsonData["returnObj"];
                str = jsonDicList[0].ToString();
                //foreach (JsonData item in jsonDicList)
                //{
                //    GwordItem JO = new GwordItem();
                //    JO.contents = item["contents"].ToString();
                //    JO.id = item["id"].ToString();
                //    JO.idx = item["idx"].ToString();
                //    GwordItem.wordList.Add(JO);
                //}
            }
        }

        return str;
    }



    public string GetHttpWebResponse(string url, string postData, string code = "utf-8")
    {

        ServicePointManager.ServerCertificateValidationCallback = (sender, certificate, chain, sslPolicyErrors) => true;
        ServicePointManager.SecurityProtocol = SecurityProtocolType.Ssl3
| SecurityProtocolType.Tls
| (SecurityProtocolType)0x300 //Tls11
| (SecurityProtocolType)0xC00; //Tls12
        ServicePointManager.ServerCertificateValidationCallback = new System.Net.Security.RemoteCertificateValidationCallback(CheckValidationResult);
        HttpWebRequest request = (HttpWebRequest)HttpWebRequest.Create(url);
        HttpWebResponse response = null;
        try
        {
            response = (HttpWebResponse)request.GetResponse();
            using (StreamReader reader = new StreamReader(response.GetResponseStream(), System.Text.Encoding.GetEncoding(code)))
            {
                return reader.ReadToEnd();
            }
        }
        catch (WebException e)
        {
            return e.Message;
        }
        finally
        {
            if (response != null) response.Close();
        }
    }

    public static bool CheckValidationResult(object sender, X509Certificate certificate, X509Chain chain, System.Net.Security.SslPolicyErrors errors)
    {
        return true;
    }

}
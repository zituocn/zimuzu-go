package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io"
	"strings"
	"time"
)

//api数据配置
var (
	cid       = "###"
	accesskey = "###"
	client    = "###"
)

//HttpGet方法
//返回*httplib.BeegoHTTPRequest
//请或取和引用github.com/astaxie/beego/httplib
func HttpGet(url string) (req *httplib.BeegoHTTPRequest) {
	t := time.Now().Unix()
	str := fmt.Sprintf("%s$$%s&&%d", cid, accesskey, t)
	key := Md5(str)
	req = httplib.Get("http://api.ousns.net" + url)
	req.SetTimeout(15*time.Second, 15*time.Second)
	req.Param("cid", cid)
	req.Param("accesskey", key)
	req.Param("timestamp", fmt.Sprintf("%d", t))
	req.Param("client", client)

	return req
}

//md5编码
//输出小写
func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return strings.ToLower(fmt.Sprintf("%x", m.Sum(nil)))
}

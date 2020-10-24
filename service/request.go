package service

import (
	uuid "github.com/iris-contrib/go.uuid"
	"net/http"
	"strings"
)

var m = map[string]string{
	"uuid":           "",
	"currentResd":    "",
	"fromHbToZjDate": "",
	"fromHbToZj":     "C",
	"fromWtToHzDate": "",
	"fromWtToHz":     "B",
	"meetDate":       "",
	"meetCase":       "C",
	"travelDate":     "",
	"travelCase":     "D",
	"medObsvReason":  "",
	"medObsv":        "B",
	"belowCaseDesc":  "",
	"belowCase":      "D",
	"temperature":    "",
	"notApplyReason": "",
	"hzQRCode":       "A",
	"specialDesc":    "",
}

//获取cookie值,生成uuid并赋值,home赋值
func GetCookie(user, pwd, userAgent, home string) string {
	var cookie string
	cookies, _ := Request(Req{
		Url:    "https://nco.zjgsu.edu.cn/login",
		Method: http.MethodPost,
		Header: map[string]string{
			"User-Agent": userAgent,
		},
		FormData: map[string]string{
			"name":  user,
			"psswd": pwd,
		},
	})
	if cookies != nil {
		u, err := uuid.NewV4()
		if err != nil {
			panic("genUUidErr")
		}
		uid := u.String()
		m["uuid"] = uid
		m["currentResd"] = home
		cookie = cookies.Name + "=" + cookies.Value + ";" + " _ncov_uuid=" + uid + "; _ncov_username=" + user + "; _ncov_psswd=" + pwd
	}
	if cookie == "" {
		panic("cookieEmpty")
	}
	return cookie
}

//post 报送表单json
func PostInfo(cookie, userAgent string) bool {
	_, body := Request(Req{
		Url:    "https://nco.zjgsu.edu.cn/",
		Method: http.MethodPost,
		Header: map[string]string{
			"User-Agent": userAgent,
			"Cookie":     cookie,
		},
		JsonData: m,
	})
	return strings.Contains(body, "报送成功")
}

// 微信推送
func Push2WeChat(pushKey, id, name string, result bool) {
	if result {
		Request(Req{
			Url:    "https://sc.ftqq.com/" + pushKey + ".send",
			Method: http.MethodPost,
			FormData: map[string]string{
				"text": "打卡成功" + id + name,
			},
		})
	} else {
		Request(Req{
			Url:    "https://sc.ftqq.com/" + pushKey + ".send",
			Method: http.MethodPost,
			FormData: map[string]string{
				"text": "打卡失败" + id + name,
			},
		})
	}
}
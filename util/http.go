package util

import (
  "os"
  "fmt"
  "time"
  "net/http"
  "net/url"
//   "log"
)

func GetTestVideoUrl(w http.ResponseWriter) {
  str_time := time.Now().Format("2006-01-02 15:04:05")
  fmt.Fprintln(w, "#EXTM3U")
  fmt.Fprintln(w, "#EXTINF:-1 tvg-name=\""+str_time+"\" tvg-logo=\"https://cdn.jsdelivr.net/gh/youshandefeiyang/IPTV/logo/tg.jpg\" group-title=\"列表更新时间\","+str_time)
  fmt.Fprintln(w, "https://cdn.jsdelivr.net/gh/youshandefeiyang/testvideo/time/time.mp4")
}

func GetLivePrefix(r *http.Request) string {
	// 尝试从环境变量读取url
	envUrl := os.Getenv("LIVE_URL")
	// log.Println("env url:", envUrl)
	if envUrl == "" {
		// 默认url
		envUrl = "https://www.goodiptv.club"
	  }
    firstUrl := DefaultQuery(r, "url", envUrl)
    realUrl, _ := url.QueryUnescape(firstUrl)
    return realUrl
}

func DefaultQuery(r *http.Request, name string, defaultValue string) string {
  param := r.URL.Query().Get(name)
  if param == "" {
    return defaultValue
  }
  return param
}
  
func Duanyan(adurl string, realurl any) string {
  var liveurl string
  if str, ok := realurl.(string); ok {
    liveurl = str
  } else {
	liveurl = adurl
  }
  // log.Println("Redirect url:", liveurl)
 return liveurl
}

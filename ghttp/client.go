package ghttp

import "regexp"

// IsMobile 根据请求携带的userAgent判断请求客户端是否是移动端
func IsMobile(userAgent string) bool {
	var isMobile = false
	reg := regexp.MustCompile(`(?i:(blackberry|configuration\/cldc|hp |hp-|htc |htc_|htc-|iemobile|kindle|midp|mmp|motorola|mobile|nokia|opera mini|opera |Googlebot-Mobile|YahooSeeker\/M1A1-R2D2|android|iphone|ipod|mobi|palm|palmos|pocket|portalmmm|ppc;|smartphone|sonyericsson|sqh|spv|symbian|treo|up.browser|up.link|vodafone|windows ce|xda |xda_|MicroMessenger))`)
	if len(reg.FindAllString(userAgent, -1)) > 0 {
		isMobile = true
	}
	return isMobile
}

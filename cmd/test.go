package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/vislee/uasurfer"
)

func test() {
	// s := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36; 360Spider)"
	// s := "DadaStaff/10.29.0 (iPhone; iOS 13.6; Scale/3.00)"
	// s := "curl/7.19.7 (x86_64-redhat-linux-gnu) libcurl/7.19.7 NSS/3.27.1 zlib/1.2.3 libidn/1.18 libssh2/1.4.2"
	// s := "Mozilla/5.0 (compatible; Baiduspider-render/2.0; +http://www.baidu.com/search/spider.html)"
	// s := "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)'"
	// s := "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 YisouSpider/5.0 Safari/537.36"
	// s := "Sogou web spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)"
	// s := "Sogou web spider/4.0(+http://www.sogou.com/docs/help/webmasters.htm#07)"
	// s := "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm"
	// s := "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
	// s := "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0); 360Spider"
	// s := "Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)"
	s := "Mozilla/5.0 (compatible; AhrefsBot/7.0; +http://ahrefs.com/robot/))"
	// s := "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1 (compatible; Baiduspider-render/2.0; +http://www.baidu.com/search/spider.html)"
	ua := &uasurfer.UserAgent{}
	ua.Reset()
	uasurfer.ParseUserAgent(s, ua)
	fmt.Println(ua)
}

func main() {
	bufin := bufio.NewReader(os.Stdin)
	for {
		str, err := bufin.ReadString('\n')
		if err != nil {
			fmt.Println("err:", err.Error())
			return
		}

		ua := &uasurfer.UserAgent{}
		ua.Reset()
		uasurfer.ParseUserAgent(str, ua)
		fmt.Println(str, ua)
	}
}

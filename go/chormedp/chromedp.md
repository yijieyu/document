# 什么是cdp
    最近抓取网站的播放地址，由于视频的播放地址是php使用混编的方式写入html中，并且播放地址会经过kt_player.js加密后才是最终的播放地址
    最开始想的方式是通过下载kt_player.js，然后把原始的播放地址通过go的robertkrimen/otto运行后拿到最终地址，
    由于kt_player.js使用了浏览器中的document，Windows等对象，robertkrimen/otto并不支持这些原始对象，此路不通
    后面通过Google/百度查找到无头浏览器，此工具可以模拟浏览器的所有操作，最终通过这个工具获取到播放地址
    
    搜了下相关概念，无头浏览器的话python里就是selenium驱动的，广泛使用的headless browser解决方案PhantomJS已经宣布不再继续维护，转而推荐使用headless chrome。Headless Chrome 是 Chrome 浏览器的无界面形态，可以在不打开浏览器的gui前提下，使用所有 Chrome 支持的特性运行你的程序
    
    反爬措施的目的就是保证正常用户的访问，拒绝爬虫的访问。这个时候，我们就在思索一件事，不管他步骤怎样复杂化，他还是要对正常的浏览器提供业务支持，换而言之，他再复杂的请求步骤也会被浏览器完美执行。使用浏览器自己当爬虫，加大了资源消耗，爬取速度明显变慢，但是简化了开发步骤，缩短了开发周期，在某些情况下，这个技术还是非常有利可图的。
    
    golang里驱动headless chrome有着开源库chromedp(在2017年的gopher大会上有展示过)，它是使用Chrome Debugging Protocol(简称cdp) 并且没有外部依赖 (如Selenium, PhantomJS等)。
    
    浏览器本身其实还充当着一个服务端的角色，大家应该都用过chrome浏览器的F12，也就是devtools，其实这是一个web应用，当你使用devtools的时候，而你看到的浏览器调试工具界面，其实只是一个前端应用，在这中间通信的，就是cdp，他是基于websocket的，一个让devtools和浏览器内核交换数据的通道。[cdp的官方文档地址](https://chromedevtools.github.io/devtools-protocol/) 可以点击查阅。
    
# chromedp能做什么
   - 反爬虫js，例如有的网页后台js自动发送心跳包，浏览器里会自动运行，不需要我们自动处理
   - 针对于前端页面的自动化测试
   - 解决类似VueJS和SPA之类的渲染
   - 解决网页的懒加载
   - 网页截图和pdf导出，而不需要额外的去学习其他的库实现
   - seo训练和刷点击量
   - 执行javascript 代码
   - 设置dom的标签属性
#### 使用前提
    类似Jquery，懂一点html和css以及js，因为操作html的dom元素需要用到xpath和css选择器之类的，如果F12的element里会右击复制selector也行，但是复杂的选择器还得需要xpath或者css选择器。不会使用的话简单教下：
    chrome打开网页F12后下面的调试工具出来后点击Elements,然后点击elements右边的那个框框里的鼠标箭头，点击后变蓝色，然后放到网页上选中区域点击一下，下面的内容就跳到对应地方，然后下面右击html的标签->Copy->COpy selector或者xpath，就能复制选择器了。
    
#### 安装
    go get -u github.com/chromedp/chromedp@master
    
## 使用一
   - 打开腾讯视频搜索页面 https://v.qq.com/x/search/
   - 输入哪吒
   - 点击全网搜
   - 获取第一个结果
   - 搜索截图
#### 代码
```go
package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)
func main() {

	options := []chromedp.ExecAllocatorOption{
		// 开启GUI来debug
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		// 设置UA
		chromedp.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36`),
	}

	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	ctx := context.Background()
	c, cc := chromedp.NewExecAllocator(ctx, options...)
	defer cc()
	// create context
	ctx, cancel := chromedp.NewContext(c)
	defer cancel()
	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://v.qq.com/x/search/?q=&stag=&smartbox_ab=`),
		chromedp.WaitVisible(`#searchForm`),
		chromedp.SendKeys(`#keywords`, `哪吒`, chromedp.ByID),
		chromedp.Click(`.search_btn`, chromedp.NodeVisible),
		chromedp.WaitVisible(`.wrapper_main .result_item `),
		chromedp.CaptureScreenshot(&buf),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("/tmp/fullScreenshot.png", buf, 0644); err != nil {
		log.Fatal(err)
	}
}
```
![image](https://raw.githubusercontent.com/yijieyu/document/master/go/chormedp/fullScreenshot.png)

## 使用二
   - 抓取页面的运行js后的结果
   
#### 被抓取页面的js代码
```js
    var flashvars = {
        video_url: 'http://domain.com/kt_player/demo/demo_video.mp4',
        preview_url: 'http://domain.com/kt_player/demo/demo_preview.jpg'
    };
    var params = {allowfullscreen: 'true', allowscriptaccess: 'always'};
    var player_obj = kt_player('kt_player', '/kt_player/kt_player.swf', '600', '400', flashvars, params);
```
#### go 代码
```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/davecgh/go-spew/spew"
)

type playerObj struct {
	VideoURL string
	Rnd      string
}

func main() {
	// create context

	t1 := time.Now()

	defer func() {
		t2 := time.Now()
		spew.Dump(t2.Sub(t1))
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	p := &playerObj{}
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.sexvid.xxx/eva-lovia-and-van-wylde-have-sex-in-porn-parody-game-of-balls.html`),
		chromedp.WaitVisible("#kt_player .fp-player", chromedp.NodeVisible),
		chromedp.EvaluateAsDevTools(scriptVideo, &p.VideoURL),
		chromedp.EvaluateAsDevTools(scriptRnd, &p.Rnd),
	)

	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(p)
	spew.Dump(p.VideoURL + "&" + p.Rnd)
}

const (
	scriptVideo = `player_obj.conf.video_url`
	scriptRnd   = `player_obj.conf.rnd`
)
```
运行结果：
```text
(*main.playerObj)(0xc000160100)({
  VideoURL: (string) (len=107) "https://www.sexvid.xxx/get_file/7/83919b5498e71fef8aeab85ba011ec677858523f54/51000/51656/51656.mp4/?br=2538",
  Rnd: (string) (len=10) "1571041408"
 })
 
 (string) (len=118) "https://www.sexvid.xxx/get_file/7/83919b5498e71fef8aeab85ba011ec677858523f54/51000/51656/51656.mp4/?br=2538&1571041408"
 (time.Duration) 28.979338921s
 
 Process finished with exit code 0
```

# chromedp官方 
   - demo https://github.com/chromedp/examples
   - api https://godoc.org/github.com/chromedp/chromedp
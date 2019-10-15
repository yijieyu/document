# 什么是goquery
    最近在做关于爬虫的工作，经常解析页面内容。
    使用到goquery这个库比较多，尤其是对爬取到的HTML进行选择和查找匹配的内容时，goquery的选择器使用尤其多，而且还有很多不常用但又很有用的选择器
    如果大家以前做过前端开发，对jquery不会陌生，goquery类似jquery，它是jquery的go版本实现。使用它，可以很方便的对HTML进行处理

# 常用的选择器和数据
```go
package main

import (
	"github.com/PuerkitoBio/goquery"
)

func main() {
 
    doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
    
    // ID选择器
    doc.Find("#id")
    // class选择器
    doc.Find(".class")
    
   // 列表
   doc.Find("#id .class").Each(func(i int, s *goquery.Selection){
   	// 属性获取
   	href,ok:=s.Find("a").Attr("href")
   	// 第一次属性获取失败，则再次尝试获取
   	if !ok {
    			for i := 0; i < 3; i++ {
    				href,ok=s.Find("a").Attr("href")
    				if ok {
    					break
    				}
    			}
    		}
   	// 文本获取
   	text :=s.Find("a").Text()
   })
    
    // 选择列表的第n个
    doc.Find("#id .class").Eq(n).Text()
    
    // 下面的节点
    doc.NextAllFiltered(".class/#id")
}

```
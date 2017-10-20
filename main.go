
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
	"github.com/axgle/mahonia"
	//"bytes"
	//"github.com/PuerkitoBio/goquery"
	//"golang.org/x/text/encoding/simplifiedchinese"
	//"golang.org/x/text/transform"
)

/*
 GO 爬虫
 1、获取 => net/http
    a、使用 io/ioutil.ReadAll() 读取返回内容
    b、用 string() 转 []uint8 为可读字符串
	c、编码 => encoding / golang.org/x/text/encoding
 2、解析 => regexp / github.com/PuerkitoBio/goquery
 3、并发 =>

 */

var (
	contentLinkReg = regexp.MustCompile(`<dd><a href="(.+)">(.+)</a></dd>`)
)

type Chapter struct {
	link string
	title string
	content string
}

/*
 请求指定url，返回页面内容
 */
func request(url string) (content string) {
	res, errGet := http.Get(url)
	if errGet != nil {
		fmt.Println(errGet)
		return
	}
	defer res.Body.Close()
	reader, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		fmt.Println(errRead)
		return
	}
	content = string(reader)
	return
}

/*func transcode(str string) (tostr string) {
	//str = "中国"
	byteStr := []byte(str)
	fmt.Printf("%T",byteStr)
	reader := transform.NewReader(bytes.NewReader(byteStr), simplifiedchinese.GBK.NewEncoder())
	toReader, errRead := ioutil.ReadAll(reader)
	if errRead != nil {
		fmt.Println(errRead)
		return
	}
	tostr = string(toReader)
	fmt.Println(tostr)
	return
}*/

func pickChaptersByReg (content string) (chapters []Chapter) {
	cLabels := contentLinkReg.FindAllStringSubmatch(content, 10000) // 10000是最大长度
	chapters = make([]Chapter, 0)
	// 有时不得不声明一个unused的变量，可以用_代表
	// _ 可以在同一个作用域内声明多个而不报错
	for _,v := range cLabels{
		var chapter Chapter
		chapter.link = v[1]
		chapter.title = v[2]
		chapters = append(chapters, chapter)
	}
	return
}

func pickChaptersByQuery (content string) (chapters []Chapter) {
	//doc, err := goquery.NewDocumentFromReader(content)
	// ...
	// goquery不能接受html[string]作为参数，不适应于此
	// 可以接受:
	// url [string]
	// reader [io.reader]
	// response [http.response]
	// *html.Node
	// *Document
	return
}

func main() {
	fmt.Println("cawler start")

	const url = "http://www.biquge.tv/0_1/"
	content := request(url)

	//content = transcode(content)
	fmt.Println(content)
	enc := mahonia.NewEncoder("gbk")
	fmt.Println(enc.ConvertString(content))


	chapters := make([]Chapter, 0)

	// 通过正则获取章节列表
	chapters = pickChaptersByReg(content)

	// 通过Goquery获取章节列表
	// chapters = pickChaptersByQuery(content)

	fmt.Println(chapters)
}


package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"io/ioutil"
	"errors"
)

type RetBody struct {
	Status int
	Raw string
	Code int
	Msg string
	Data interface{}
}

// TODO 预览图应从后端中转，并做Refer等反反盗链处理
// TODO 页面美化，移动端适配
// TODO 实测以支持各种图片格式
func main() {
	server()
}

func server() {
	//全局设置为产品环境
	gin.SetMode(gin.ReleaseMode)
	// 端口
	port := 666
	//获得路由实例
	router := gin.Default()
	// 静态资源
	router.StaticFile("/", "./public/index.html")
	// 注册接口
	//router.GET("/", IndexHandler)
	router.GET("/pick", PickHandler)
	// 启动服务
	fmt.Println(`Server start @`, port)
	router.Run(fmt.Sprintf(":%d", port))
}

func CReturn(c *gin.Context, ret RetBody) {
	if ret.Status == 0 {
		ret.Status = http.StatusOK
	}
	if ret.Raw != "" {
		c.Data(ret.Status, "text/plain", []byte(ret.Raw))
	} else {
		type JSONRet struct {
			Code int `json:"Code"`
			Msg string `json:"Msg"`
			Data interface{}
		}
		if ret.Msg == "" {ret.Msg = "操作成功"}
		if ret.Data == "" {ret.Data = gin.H{}}
		jsonRet := JSONRet{Code: ret.Code, Msg: ret.Msg, Data: ret.Data}
		c.JSON(ret.Status, jsonRet)
	}
}

func IndexHandler(c *gin.Context) {
	CReturn(c, RetBody{Raw: "Index"})
}

func PickHandler(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		CReturn(c, RetBody{Code: 1, Msg: "未提交Key"})
		return
	}
	// 析出av号
	avid := pickAVID(value)
	if avid == "" {
		CReturn(c, RetBody{Code: 2, Msg: "提交的Key中不包含av号"})
		return
	}
	fmt.Println("检测到AV号:", avid)
	// 根据avid获取封面地址
	coverUrl, err := pickCoverUrl(avid)
	if err != nil {
		fmt.Println("尝试利用AV号", avid, "获取视频封面时错误：", err)
		CReturn(c, RetBody{Code: 3, Msg: "获取视频封面时错误", Data: gin.H{"reason": fmt.Sprintf("%s", err)}})
		return
	}
	CReturn(c, RetBody{Data: gin.H{"coverUrl": coverUrl, "avid": avid}})
}

func pickAVID(value string) (avid string) {
	var regAVID = regexp.MustCompile(`^((http|https)://www.bilibili.com/video/){0,1}(av){0,1}(\d{8})`)
	matches := regAVID.FindStringSubmatch(value)
	if len(matches) > 0 {
		avid = matches[len(matches) - 1]
	}
	return
}

func pickCoverUrl(avid string) (coverUrl string, err error) {
	content, _, err := request(fmt.Sprintf("http://www.bilibili.com/video/av%s/", avid))
	if err != nil {
		return
	}
	var regCover = regexp.MustCompile(`<img src="(//i\d{1}\.hdslb\.com/bfs/archive/\w+\.(jpg|png))" `)
	matches := regCover.FindStringSubmatch(content)
	if len(matches) > 0 {
		coverUrl = fmt.Sprintf("http:%s", matches[1])
	} else {
		err = errors.New("no cover img label matched")
	}
	return
}

func request(url string) (content string, reader []byte, err error) {
	fmt.Println(url)
	res, errGet := http.Get(url)
	if errGet != nil {
		fmt.Println(errGet)
		err = errors.New("http get failed")
		return
	}
	defer res.Body.Close()
	reader, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		fmt.Println(errRead)
		err = errors.New("res body read failed")
		return
	}
	content = string(reader)
	return
}
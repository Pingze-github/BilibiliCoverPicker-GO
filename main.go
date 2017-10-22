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
	String string
	Raw []byte
	ContentType string
	Code int
	Msg string
	Data interface{}
}

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
	router.GET("/imgpipe", ImgPiperHandler)
	// 启动服务
	fmt.Println(`Server start @`, port)
	router.Run(fmt.Sprintf(":%d", port))
}

func CReturn(c *gin.Context, ret RetBody) {
	if ret.Status == 0 {
		ret.Status = http.StatusOK
	}
	if ret.String != "" {
		c.Data(ret.Status, "text/plain", []byte(ret.String))
	} else if ret.Raw != nil {
		if ret.ContentType == "" {
			ret.ContentType = "text/plain"
		}
		c.Data(ret.Status, ret.ContentType, []byte(ret.Raw))
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
	CReturn(c, RetBody{String: "Index"})
}

func PickHandler(c *gin.Context) {
	key, exist := c.GetQuery("key")
	if !exist {
		CReturn(c, RetBody{Code: 1, Msg: "未提交Key"})
		return
	}
	// 析出av号
	avid := pickAVID(key)
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

func ImgPiperHandler(c *gin.Context) {
	src, exist := c.GetQuery("src")
	if !exist {
		CReturn(c, RetBody{Code: 1, Msg: "未提交图片路径src"})
		return
	}
	_, bytes, err := request(src)
	if err != nil {
		return
	}
	CReturn(c, RetBody{Raw: bytes, ContentType: "image/jpeg"})
}

func pickAVID(value string) (avid string) {
	var regAVID = regexp.MustCompile(`^((http|https)://www.bilibili.com/video/){0,1}(av){0,1}(\d+)`)
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

func request(url string) (content string, bytes []byte, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil);
	req.Header.Set("Referer", "http://www.bilibili.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
	res, errRes := client.Do(req)
	if errRes != nil {
		fmt.Println(errRes)
		err = errors.New("http get failed")
		return
	}
	defer res.Body.Close()
	bytes, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		fmt.Println(errRead)
		err = errors.New("res body read failed")
		return
	}
	content = string(bytes)
	return
}
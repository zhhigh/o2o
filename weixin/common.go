package weixin

import (
    "time"
)

var (
	Token string
)

const (
    Text = "text"
    Location = "location"
    Image = "image"
    News = "news"
	WeixinHost        = "https://api.weixin.qq.com/cgi-bin"
	WeixinQRScene     = "https://api.weixin.qq.com/cgi-bin/qrcode"
	WeixinShowQRScene = "https://mp.weixin.qq.com/cgi-bin/showqrcode"
	WeixinFileURL     = "http://file.api.weixin.qq.com/cgi-bin/media"
)

type MsgBase struct {
    ToUserName string
    FromUserName string
    CreateTime time.Duration
    MsgType string
    Content string
}

type PicBase struct {
	ToUserName string
	FromUserName string
	CreateTime time.Duration
	MsgType string
}


/*
resp := weixin.NewResponse()

resp.ArticleCount = 2
resp.MsgBase = weixin.MsgBase{
ToUserName: "toUser",
FromUserName: "fromUser",
CreateTime: 12345678000000000,
MsgType: weixin.News,
Content: "content",
}

var item weixin.Item
item.Title= "title"
item.Description="desc"
item.PicUrl="http://a.36krcnd.com/photo/fccc728b86eccc843ac48444c0aa1ca2.jpg"
item.Url="http://www.36kr.com/p/207965.html"

resp.Articles = append(resp.Articles,&item)

var a weixin.Item

a.Title= "title"
a.Description="desc"
a.PicUrl="http://a.36krcnd.com/photo/fccc728b86eccc843ac48444c0aa1ca2.jpg"
a.Url="http://www.36kr.com/p/207965.html"

resp.Articles = append(resp.Articles,&a)
data,_ := resp.Encode()
fmt.Println("%s",string(data))

		*/




/*mon = mongo.NewMongoDBConn()
	defer mon.Stop()
	mon.Connect("mongodb://redis:redis@192.241.225.170:27017/rrest")
	mon.SetDBName("rrest")
	mon.SetTableName("msg")

	var r interface{}
	mon.FindOne(bson.M{"serial": "0901"},&r)

	serial := r.(bson.M)["serial"]
	title  := r.(bson.M)["title"]
	picurl := r.(bson.M)["picurl"]
	url    := r.(bson.M)["url"]
	fmt.Println(serial,title,picurl,url) */
//fmt.Println(&q)
//fmt.Println(reflect.TypeOf(r))



/* 图文消息格式
<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>12345678</CreateTime>
<MsgType><![CDATA[news]]></MsgType>
<ArticleCount>2</ArticleCount>
<Articles>
<item>
<Title><![CDATA[title1]]></Title>
<Description><![CDATA[description1]]></Description>
<PicUrl><![CDATA[picurl]]></PicUrl>
<Url><![CDATA[url]]></Url>
</item>
<item>
<Title><![CDATA[title]]></Title>
<Description><![CDATA[description]]></Description>
<PicUrl><![CDATA[picurl]]></PicUrl>
<Url><![CDATA[url]]></Url>
</item>
</Articles>
</xml>
*/



/*文字
<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[fromUser]]></FromUserName>
<CreateTime>12345678</CreateTime>
<MsgType><![CDATA[text]]></MsgType>
<Content><![CDATA[你好]]></Content>
</xml>
*/



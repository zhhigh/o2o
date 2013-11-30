package weixin

import (
    "time"
    "encoding/xml"
)


type Request struct {
    XMLName xml.Name `xml:"xml"`
    MsgBase // base struct
    Location_X, Location_Y float32
    Scale int
    Label string
    PicUrl string
    MsgId int
}

func DecodeRequest(data []byte) (req *Request, err error) {
    req = &Request{}
    if err = xml.Unmarshal(data, req); err != nil {
        return
    }
    req.CreateTime *= time.Second
    return
}

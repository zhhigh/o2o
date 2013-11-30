package weixin

import (
    "fmt"
    "sort"
    "net/http"
    "io/ioutil"
    "crypto/sha1"
    "github.com/mikespook/golib/log"
)

var (
    Token = "thisiswechattoken"
)

type HandlerFunc func(*Request)(*Response, error)

func Handle(w http.ResponseWriter, r *http.Request, h HandlerFunc) {
    defer r.Body.Close()
    if r.Method == "POST" || r.Method == "post" {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            log.Error(err)
            w.WriteHeader(500)
            return
        }
        var wreq *Request
        if wreq, err = DecodeRequest(body); err != nil {
            log.Error(err)
            w.WriteHeader(500)
            return
        }
        wresp, err := h(wreq)
        if err != nil {
            log.Error(err)
            w.WriteHeader(500)
            return
        }
        data, err := wresp.Encode()
        if _, err := w.Write(data); err != nil {
            log.Error(err)
            w.WriteHeader(500)
        }
		fmt.Println("handleer")
		fmt.Println(string(data))
		fmt.Println("handleer")
        return
    } else {
        if Signature(Token, r.FormValue("timestamp"),
            r.FormValue("nonce")) == r.FormValue("signature") {
            w.Write([]byte(r.FormValue("echostr")))
        } else {
            w.WriteHeader(403)
        }
    }
}

func Signature(token, timestamp, nonce string) string {
    strs := sort.StringSlice{token, timestamp, nonce}
    sort.Strings(strs)
    str := ""
    for _, s := range strs {
        str += s
    }
    h := sha1.New()
    h.Write([]byte(str))
    return fmt.Sprintf("%x", h.Sum(nil))
}

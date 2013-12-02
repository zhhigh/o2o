/*
Created by zhhigh
Date :2013-11-11
*/
package o2o

import (
    "crypto/md5"
    "crypto/sha1"
    "encoding/hex"
    "io"
    "encoding/base64"
    "fmt"
    "time"
    "hash"
)

type Alg struct{
    Cipher string
    md5      hash.Hash
}
//var cipher = "密鑰"
//var h = md5.New()


func NewAlg() *Alg{
     return &Alg{}
}

/*set keyt
*/
func (a *Alg)SetKeyt(key string){
    a.md5     = md5.New()
    a.Cipher = key
}


func (a *Alg)Md5(data string)(string){
	h := md5.New()
	h.Write([]byte(data))
	return  hex.EncodeToString(h.Sum(nil))
}


func (a *Alg)Sha1(data string)(string){
	h := sha1.New()
    io.WriteString(h,data)
    return  hex.EncodeToString(h.Sum(nil))
}



func (a *Alg)CipherEncode(sourceText string) string {
    a.md5.Write([]byte(a.Cipher))
    cipherHash := fmt.Sprintf("%x", a.md5.Sum(nil))
    a.md5.Reset()
    inputData := []byte(sourceText)
    loopCount := len(inputData)
    outData := make([]byte,loopCount)
    for i:= 0; i < loopCount ; i++ {
        outData[i] = inputData[i] ^ cipherHash[i%32]
    }   
    return fmt.Sprintf("%s", outData)
}

func (a *Alg)Eencode(sourceText string) string { 
    a.md5.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
    noise := fmt.Sprintf("%x", a.md5.Sum(nil))
    a.md5.Reset()
    inputData := []byte(sourceText)
    loopCount := len(inputData)
    outData := make([]byte,loopCount*2)
    
    for i, j := 0,0; i < loopCount ; i,j = i+1,j+1 {        
        outData[j] = noise[i%32]
        j++
        outData[j] = inputData[i] ^ noise[i%32]
    }
    
    return base64.StdEncoding.EncodeToString([]byte(a.CipherEncode(fmt.Sprintf("%s", outData))))
}

func (a *Alg)Ddecode(sourceText string) string {
    buf, err := base64.StdEncoding.DecodeString(sourceText)
    if err != nil {
        fmt.Println("Decode(%q) failed: %v", sourceText, err)
        return ""
    }
    inputData := []byte(a.CipherEncode(fmt.Sprintf("%s", buf)))
    loopCount := len(inputData)
    outData := make([]byte,loopCount)
    for i, j := 0,0; i < loopCount ; i,j = i+2,j+1 {        
        outData[j] = inputData[i] ^ inputData[i+1]
    }
    return fmt.Sprintf("%s", outData)
}

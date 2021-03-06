package logs

import (
	"encoding/json"
	"log"
	"os"
	//"fmt"
)

type ConsoleWriter struct {
	lg    *log.Logger
	level int
}

func NewConsole() LoggerInterface {
	cw := new(ConsoleWriter)
	cw.lg = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	cw.level = LevelTrace
	return cw
}

func (c *ConsoleWriter) Init(jsonconfig string) error {
	//fmt.Println("Init",jsonconfig)
	var m map[string]interface{}
	err := json.Unmarshal([]byte(jsonconfig), &m)
	//fmt.Println("console Init",err)
	if err != nil {
		return err
	}
	if lv, ok := m["level"]; ok {
		c.level = int(lv.(float64))
	}
	//fmt.Println("ConsoleWriter",c)
	//fmt.Println("c.level",c.level)
	return nil
}

func (c *ConsoleWriter) WriteMsg(msg string, level int) error {
    //fmt.Println("--------------write console msg")
	//fmt.Println("level",level)
	//fmt.Println("c.level",c.level)
	//fmt.Println("--------------write console msg")
	if level < c.level {
		return nil
	}
	c.lg.Println(msg)
	//fmt.Println("console WriteMsg",msg)
	return nil
}

func (c *ConsoleWriter) Destroy() {

}

func (c *ConsoleWriter) Flush() {

}

func init() {
	Register("console", NewConsole)
}

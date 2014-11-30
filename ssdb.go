package o2o

import(
	"fmt"
	"github.com/eryx/lessgo/data/hissdb"
)

type ssdb struct{
	Host string
	Port string
	TimeOut int
	MaxConn int
	Conn    *hissdb.Connector
}

func New() *ssdb{
	return &ssdb
}

func (s *ssdb)InitCfg(host string,port string,timeout int,maxconn int){
	s.Host = host
	s.Port = port
	s.TimeOut = timeout
	s.MaxConn = maxconn
}


func (s *ssdb)Connect(){
	conn, err := hissdb.NewConnector(hissdb.Config{
	Host:    s.Host,
	Port:    s.Port,
	Timeout: 3,  // timeout in second, default to 10
	MaxConn: 10,  // max connection number, default to 1
	})
	if err != nil {
		fmt.Println("Connect Error:", err)
		return
	}
	s.Conn = conn
	defer conn.Close()
}



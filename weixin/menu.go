package weixin

/*20141029*/
import (
	"encoding/json"
	"net/http"
	"bytes"
	"errors"
)


type Menu struct {
	Buttons []MenuButton `json:"button,omitempty"`
}

type MenuButton struct {
	Name       string       `json:"name"`
	Type       string       `json:"type,omitempty"`
	Key        string       `json:"key,omitempty"`
	Url        string       `json:"url,omitempty"`
	SubButtons []MenuButton `json:"sub_button,omitempty"`
}

func CreateMenu(menu *Menu) error {
	data, err := json.Marshal(menu)
	if err != nil {
		return err
	} else {
	_, err := postRequest(WeixinHost+"/menu/create?access_token=", Token, data)
	return err
	}
}


func postRequest(reqURL string,token string,data []byte)([]byte,error){
	r,err := http.Post(reqURL,token,"application/json; charset=utf-8",bytes.NewReader(data))
	if err != nil{
		return nil,err
	}
	defer r.Body.Close()

	reply,err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var result response
	if err := json.Unmarshal(reply, &result); err != nil {
		return nil, err
	} else {
		switch result.ErrorCode {
		case 0:
			return reply, nil
		case 42001: // access_token timeout and retry
			continue
		default:
			return nil, errors.New(fmt.Sprintf("WeiXin send post request reply[%d]: %s", result.ErrorCode, result.ErrorMessage))
		}
	}

	return nil, errors.New("WeiXin post request too many times:" + reqURL)

}
/*
func postRequest(reqURL string, c chan accessToken, data []byte) ([]byte, error) {
	for i := 0; i < retryMaxN; i++ {
		token := <-c
		if time.Since(token.expires).Seconds() < 0 {
			r, err := http.Post(reqURL+token.token, "application/json; charset=utf-8", bytes.NewReader(data))
			if err != nil {
				return nil, err
			}
			defer r.Body.Close()
			reply, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			var result response
			if err := json.Unmarshal(reply, &result); err != nil {
				return nil, err
			} else {
				switch result.ErrorCode {
				case 0:
					return reply, nil
				case 42001: // access_token timeout and retry
					continue
				default:
					return nil, errors.New(fmt.Sprintf("WeiXin send post request reply[%d]: %s", result.ErrorCode, result.ErrorMessage))
				}
			}
		}
	}
	return nil, errors.New("WeiXin post request too many times:" + reqURL)
}*/

package weixin

/*20141029*/

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
	_, err := postRequest(weixinHost+"/menu/create?access_token=", wx.tokenChan, data)
	return err
	}
}

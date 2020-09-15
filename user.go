package oauth4thunder

import (
	"context"
	"encoding/json"
	"io/ioutil"
)

type UserInfo struct {
	Name string `json:"name"`
	UserName string `json:"username"`
	Phone string `json:"phone"`
	Avatar string `json:"avatar"`
	Email string `json:"email"`
	Gender int8 `json:"gender"`
	IsSuperuser bool `json:"is_superuser"`
}

func (p *provider) UserInfo (ctx context.Context, code string) (*UserInfo, error){
	token, err := p.ExchangeToken(ctx, code)
	if err != nil {
		return nil, err
	}

	client := p.oauthConf.Client(ctx, token)
	resp, err := client.Get(p.conf.Endpoint + "v1/api/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user UserInfo
	err = json.Unmarshal(bodyByte, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
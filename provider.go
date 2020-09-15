package oauth4thunder

import (
	"context"
	"golang.org/x/oauth2"
)


type Provider interface {
	AuthorizeURL(state string, opts ...oauth2.AuthCodeOption) string
	ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error)
	UserInfo (ctx context.Context, code string) (*UserInfo, error)
}

type Config struct {
	Endpoint string `toml: "endpoint"`
	ClientId string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
	RedirectUrl string `toml:"redirect_url"`
	Scope []string	`toml:"scope"`
}

type provider struct {
	conf Config
	oauthConf *oauth2.Config
}


func New (conf Config) Provider {
	oauthConf := &oauth2.Config{
		ClientID:     conf.ClientId,
		ClientSecret: conf.ClientSecret,
		Endpoint:     oauth2.Endpoint{
			AuthURL: conf.Endpoint + "v1/oauth2/authorize",
			TokenURL: conf.Endpoint + "v1/oauth2/token",
		},
		RedirectURL:  conf.RedirectUrl,
		Scopes:       conf.Scope,
	}
	return &provider{
		conf: conf,
		oauthConf: oauthConf,
	}
}

func (p *provider) AuthorizeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return p.oauthConf.AuthCodeURL(state, opts...)
}

func (p *provider) ExchangeToken(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := p.oauthConf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}



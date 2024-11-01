package config

type Config struct {
	Oauth OauthProvider
}

// 597e71c35284c81e072f2a85038b3c419d67d5b7
type OauthProvider struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	TokenUrl     string `yaml:"token_url"`
	ProfileURL   string `yaml:"profile_url"`
}

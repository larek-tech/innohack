package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/larek-tech/innohack/backend/internal/auth/config"
)

type OauthProvider struct {
	cfg    *config.OauthProvider
	client http.Client
}

func NewOauthProvider(cfg *config.OauthProvider) *OauthProvider {
	return &OauthProvider{
		cfg: cfg,
	}
}

type providerAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func (gp *OauthProvider) getProviderAccessToken(ctx context.Context, code string) (*providerAccessTokenResponse, error) {
	requestBodyMap := map[string]string{
		"client_id":     gp.cfg.ClientID,
		"client_secret": gp.cfg.ClientSecret,
		"code":          code,
		"redirect_uri":  gp.cfg.TokenUrl,
	}

	requestJSON, _ := json.Marshal(requestBodyMap)
	r, err := http.NewRequest(
		"POST",
		gp.cfg.TokenUrl,
		bytes.NewBuffer(requestJSON),
	)
	if err != nil {
		return nil, err
	}

	response, err := gp.client.Do(r)
	if err != nil {
		return nil, err
	}
	bytesResponse, err := io.ReadAll(response.Body)

	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	accessToken := providerAccessTokenResponse{}

	json.Unmarshal(bytesResponse, &accessToken)

	return &accessToken, err
}

func (gp *OauthProvider) getProviderUserData(providerAccessToken *providerAccessTokenResponse) (map[string]interface{}, error) {
	r, err := http.NewRequest("GET", gp.cfg.ProfileURL, nil)

	if err != nil {
		return nil, err
	}

	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", providerAccessToken.AccessToken))

	response, err := http.DefaultClient.Do(r)

	if err != nil {
		return nil, err
	}

	bytesResponse, err := io.ReadAll(response.Body)

	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	newEntity := make(map[string]interface{})

	err = json.Unmarshal(bytesResponse, &newEntity)
	if err != nil {
		return nil, err
	}

	return newEntity, err
}

func (s *Service) OauthGetAccessToken(ctx context.Context, code string) (string, error) {
	resp, err := s.oauth.getProviderAccessToken(ctx, code)
	if err != nil {
		return "", err
	}
	return resp.AccessToken, nil
}

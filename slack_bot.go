package slackbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type SlackBot struct {
	webHookUrl *url.URL
	httpClient *http.Client
	header     http.Header
}

func (s *SlackBot) GetWebHookUrl() *url.URL {
	if s == nil {
		return nil
	}
	return s.webHookUrl
}

func (s *SlackBot) SendMessageStr(msg string) error {
	if s == nil {
		return nil
	}
	slackBody, err := json.Marshal(&slackMsg{msg})
	if err != nil {
		return err
	}
	return s.SendMessage(slackBody)
}

func (s *SlackBot) SendMessage(msg []byte) error {
	if s == nil {
		return nil
	}
	if len(msg) == 0 {
		return cEmptyMessageErr
	}
	req := &http.Request{
		Method:        http.MethodPost,
		URL:           s.webHookUrl,
		Header:        s.header,
		Body:          ioutil.NopCloser(bytes.NewReader(msg)),
		ContentLength: int64(len(msg)),
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got %d response from Slack", resp.StatusCode)
	}
	return nil
}

func NewSlackBot(webHookUrl string) (*SlackBot, error) {
	httpClient := &http.Client{}
	httpClient.Transport = &http.Transport{MaxIdleConnsPerHost: 256, MaxIdleConns: 256, IdleConnTimeout: time.Minute * 2}
	header := http.Header{"Content-Type": []string{"application/json"}}
	url, err := url.Parse(webHookUrl)
	if err != nil {
		return nil, err
	}
	return &SlackBot{url, httpClient, header}, nil
}

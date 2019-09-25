package slackbot

import (
	"errors"
)

var (
	cEmptyMessageErr = errors.New("can't send empty message")
)

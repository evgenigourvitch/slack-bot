# pxGoSlackBot
golang slack package

Usage example
```
package main

import (
	"fmt"
	"slack-bot"
)

func main() {
	bot, err := slackbot.NewSlackBot("https://hooks.slack.com/services/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")
	if err != nil {
		panic(err)
	}
	if err := bot.SendMessage("as package"); err != nil {
		fmt.Printf("failed to send meessage, err: %+v\n", err)
	}
}
```

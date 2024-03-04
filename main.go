package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println("")
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6740187676067-6742818374548-P5pqdSE4Nm80Y3lucrFrM6PF")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06MBMXNCR5-6740329421218-627a653d6ec56e1637f4992ee796ed857aaaf84602d7bf2891424cc4be310e5e")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	// Actual Function
	bot.Command("My Year of Birth is <year>", &slacker.CommandDefinition{
		Description: "YOB Calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				response.Reply("Invalid year format. Please provide a valid year.")
				return
			}
			age := 2021 - yob
			r := fmt.Sprintf("Age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

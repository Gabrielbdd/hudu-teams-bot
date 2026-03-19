package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/infracloudio/msbotbuilder-go/core"
	"github.com/infracloudio/msbotbuilder-go/core/activity"
	"github.com/infracloudio/msbotbuilder-go/schema"
)

var handler = activity.HandlerFuncs{
	OnMessageFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
		text := turn.Activity.Text
		reply := fmt.Sprintf("Hello! You said: %s", text)
		return turn.SendActivity(activity.MsgOptionText(reply))
	},
}

type Bot struct {
	adapter core.Adapter
}

func (b *Bot) messagesHandler(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	act, err := b.adapter.ParseRequest(ctx, req)
	if err != nil {
		log.Printf("Failed to parse request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := b.adapter.ProcessActivity(ctx, act, handler); err != nil {
		log.Printf("Failed to process activity: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3978"
	}

	adapter, err := core.NewBotAdapter(core.AdapterSetting{
		AppID:       os.Getenv("TEAMS_APP_ID"),
		AppPassword: os.Getenv("TEAMS_APP_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Failed to create adapter: %v", err)
	}

	bot := &Bot{adapter: adapter}

	http.HandleFunc("/api/messages", bot.messagesHandler)

	addr := ":" + port
	log.Printf("Bot listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

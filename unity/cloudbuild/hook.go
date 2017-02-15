package cloudbuild

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net/http"
)

const (
	StatusQueued        = "queued"
	StatusSentToBuilder = "sentToBuilder"
	StatusStarted       = "started"
	StatusRestarted     = "restarted"
	StatusSuccess       = "success"
	StatusFailure       = "failure"
	StatusCanceled      = "canceled"
	StatusUnknown       = "unknown "
)

var Statuses = map[string]struct {
	Color int
}{
	StatusQueued:        {0x838b8b},
	StatusSentToBuilder: {0xffa500},
}

type Hook struct {
}

func (hook *Hook) Handle(req *http.Request, body []byte) (*discordgo.WebhookParams, error) {
	var payload Payload
	if err := json.Unmarshal(body, &payload); err != nil {
		return nil, err
	}

	embed := &discordgo.MessageEmbed{
		URL:         payload.Links.DashboardURL.Href + payload.Links.DashboardSummary.Href,
		Description: fmt.Sprintf("%s (%s)", payload.ProjectName, payload.BuildTargetName),
	}
	switch payload.BuildStatus {
	case StatusQueued:
		embed.Title = "Build Queued"
		embed.Color = 0x838b8b
	case StatusSentToBuilder:
		embed.Title = "Build Sent to Builder"
		embed.Color = 0xffb90f
	case StatusStarted:
		embed.Title = "Build Started"
		embed.Color = 0xbf3eff
	case StatusRestarted:
		embed.Title = "Build Restarted"
		embed.Color = 0xbf3eff
	case StatusSuccess:
		embed.Title = "Build Succeeded"
		embed.Color = 0xcaff70
	case StatusFailure:
		embed.Title = "Build Failed"
		embed.Color = 0xff3030
	case StatusCanceled:
		embed.Title = "Build Canceled"
		embed.Color = 0xdcdcdc
	case StatusUnknown:
		embed.Title = "Unknown Build Status"
	}

	return &discordgo.WebhookParams{
		Embeds: []*discordgo.MessageEmbed{embed},
	}, nil
}

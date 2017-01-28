package discourse

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
	"net/http"
)

const (
	EventTypeTopic = "topic"
	EventTypePost  = "post"
	EventTypeUser  = "user"

	EventTopicCreated   = "topic_created"
	EventTopicDestroyed = "topic_destroyed"
	EventTopicRecovered = "topic_recovered"

	EventPostCreated   = "post_created"
	EventPostDestroyed = "post_destroyed"
	EventPostRecovered = "post_recovered"
	EventPostEdited    = "post_edited"

	EventUserCreated  = "user_created"
	EventUserApproved = "user_approved"
	EventUserUpdated  = "user_updated"
)

type HookFields struct {
	Filter []string `json:"filter"`
}

type Hook struct {
	HookFields
}

func (hook *Hook) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &hook.HookFields); err != nil {
		return err
	}

	for _, event := range hook.Filter {
		switch event {
		case EventTopicCreated:
		case EventTopicDestroyed:
		case EventTopicRecovered:
		case EventPostCreated:
		case EventPostDestroyed:
		case EventPostRecovered:
		case EventPostEdited:
		case EventUserCreated:
		case EventUserApproved:
		case EventUserUpdated:
		default:
			return errors.Errorf("unknown discourse event: %s", event)
		}
	}

	return nil
}

func (hook *Hook) Handle(req *http.Request, body []byte) (*discordgo.WebhookParams, error) {
	event := req.Header.Get("X-Discourse-Event")
	eventType := req.Header.Get("X-Discourse-Event-Type")
	instance := req.Header.Get("X-Discourse-Instance")
	log.WithFields(log.Fields{
		"event":     event,
		"eventType": eventType,
		"instance":  instance,
	}).Info("Discord: Event Received")

	if !hook.CaresAbout(event) {
		return nil, nil
	}

	var envelope Envelope
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, err
	}

	params := &discordgo.WebhookParams{}
	switch event {
	case EventTopicCreated:
		params.Content = "New topic!"
	case EventTopicDestroyed:
		params.Content = "Topic deleted."
	case EventTopicRecovered:
		params.Content = "Deleted topic restored!"
	case EventPostCreated:
		params.Content = "New post!"
	case EventPostDestroyed:
		params.Content = "Post deleted."
	case EventPostRecovered:
		params.Content = "Deleted post restored!"
	case EventPostEdited:
		params.Content = "Post edited."
	case EventUserCreated:
		params.Content = "New user!"
	case EventUserApproved:
		params.Content = "User approved!"
	case EventUserUpdated:
		params.Content = "User updated!"
	default:
	}

	embed := &discordgo.MessageEmbed{}
	switch eventType {
	case EventTypeTopic:
		firstPost := envelope.Topic.PostStream.Posts[0]
		embed.URL = fmt.Sprintf("%s/t/%s/%d", instance, envelope.Topic.Slug, envelope.Topic.ID)
		embed.Title = envelope.Topic.Title
		embed.Description = StripHTML(firstPost.Cooked)
		embed.Author = &discordgo.MessageEmbedAuthor{
			Name:    firstPost.Username,
			URL:     fmt.Sprintf("%s/users/%s", instance, firstPost.Username),
			IconURL: fmt.Sprintf("%s/%s", instance, AvatarURL(firstPost.AvatarTemplate, 1000)),
		}
	}
	params.Embeds = []*discordgo.MessageEmbed{embed}

	return params, nil
}

func (hook *Hook) CaresAbout(event string) bool {
	if len(hook.Filter) == 0 {
		return true
	}

	for _, e := range hook.Filter {
		if e == event {
			return true
		}
	}
	return false
}

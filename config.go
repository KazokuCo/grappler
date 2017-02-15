package main

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/kazokuco/grappler/discourse"
	"github.com/kazokuco/grappler/unity/cloudbuild"
	"github.com/pkg/errors"
	"net/http"
)

type Impl interface {
	Handle(req *http.Request, body []byte) (*discordgo.WebhookParams, error)
}

type Outbound struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}

type HookFields struct {
	Key    string     `json:"key"`
	Type   string     `json:"type"`
	Invoke []Outbound `json:"invoke"`
}

type Hook struct {
	HookFields

	// Created by deserializing the payload.
	Impl Impl `json:"-"`
}

func (hc *Hook) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &hc.HookFields); err != nil {
		return nil
	}

	if hc.HookFields.Key == "" {
		return errors.New("missing key for hook")
	}

	switch hc.HookFields.Type {
	case "discourse":
		var impl discourse.Hook
		if err := json.Unmarshal(data, &impl); err != nil {
			return err
		}
		hc.Impl = Impl(&impl)
	case "unity.cloudbuild":
		var impl cloudbuild.Hook
		if err := json.Unmarshal(data, &impl); err != nil {
			return err
		}
		hc.Impl = Impl(&impl)
	default:
		return errors.Errorf("unknown hook type: %s", hc.HookFields.Type)
	}

	return nil
}

type Config struct {
	Hooks []Hook `json:"hooks"`
}

package main

import (
	"bytes"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func HandleInvokeHook(hook Hook) httprouter.Handle {
	return func(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			// Connection closed prematurely.
			return
		}
		_ = req.Body.Close()

		hookParams, err := hook.Impl.Handle(req, body)
		if err != nil {
			body, err := json.Marshal(Response{false, err.Error()})
			if err != nil {
				panic(err)
			}
			rw.WriteHeader(http.StatusBadRequest)
			_, _ = rw.Write(body)
			return
		}

		if hookParams != nil {
			body, err := json.Marshal(hookParams)
			if err != nil {
				panic(err)
			}

			for _, out := range hook.Invoke {
				res, err := http.Post(
					discordgo.EndpointWebhookToken(strconv.FormatInt(out.ID, 10), out.Token),
					"application/json",
					bytes.NewReader(body),
				)
				if err != nil {
					body, err := json.Marshal(Response{false, err.Error()})
					if err != nil {
						panic(err)
					}
					rw.WriteHeader(res.StatusCode)
					_, _ = rw.Write(body)
					return
				}
			}
		}

		resp, err := json.Marshal(Response{true, ""})
		if err != nil {
			panic(err)
		}
		rw.Write(resp)
	}
}

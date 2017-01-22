package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ghodss/yaml"
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/tylerb/graceful"
	"gopkg.in/urfave/cli.v2"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func actionRun(cc *cli.Context) error {
	if cc.Bool("help") {
		return cli.ShowAppHelp(cc)
	}

	args := cc.Args()
	if args.Len() == 0 {
		return errors.New("missing argument: config.yml")
	}

	configData, err := ioutil.ReadFile(args.First())
	if err != nil {
		return err
	}

	var config Config
	if err := yaml.Unmarshal(configData, &config); err != nil {
		return err
	}

	router := httprouter.New()
	for _, hook := range config.Hooks {
		path := fmt.Sprintf("/%s/%s", hook.Type, hook.Key)
		router.POST(path, HandleInvokeHook(hook))
		log.WithField("path", path).Info("Registered Hook")
	}

	return graceful.ListenAndServe(
		&http.Server{
			Addr:    "0.0.0.0:8000",
			Handler: router,
		},
		10*time.Second,
	)
}

func main() {
	app := &cli.App{
		Name:      "grappler",
		Usage:     "discord webhook adapter",
		ArgsUsage: "config.yml",
		Action:    actionRun,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "help",
				Aliases: []string{"h"},
				Usage:   "print help and exit",
			},
		},
		HideHelp:    true,
		HideVersion: true,
	}
	app.Run(os.Args)
}

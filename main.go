package main

import (
	"encoding/base64"
	"github.com/akamensky/argparse"
	"github.com/google/logger"
	"github.com/imroc/req"
	"io/ioutil"
	"os"
)

func main() {
	// parse arguments
	parser := argparse.NewParser(os.Args[0], "")
	msg := parser.String("m", "message", &argparse.Options{Required: true, Help: "message to send"})
	title := parser.String("t", "title", &argparse.Options{Required: false, Help: "title for message", Default: ""})
	verbose := parser.Flag("v", "verbose", &argparse.Options{Required: false, Help: "print logs to stdout", Default: false})
	err := parser.Parse(os.Args)
	log := logger.Init("main", *verbose, true, ioutil.Discard)
	log.Info("starting")
	if err != nil {
		log.Fatal(parser.Usage(err))
	}
	// get home api credentials from environment variables
	home_api_user, user_is_set := os.LookupEnv("HOME_API_USER")
	home_api_auth, auth_is_set := os.LookupEnv("HOME_API_AUTH")
	home_api_domain, domain_is_set := os.LookupEnv("HOME_API_DOMAIN")

	if !user_is_set || !auth_is_set || !domain_is_set {
		log.Fatal("missing $HOME_API_USER $HOME_API_AUTH $HOME_API_DOMAIN")
	}
	// send pushover message via home-api
	log.Info("going to send a push message.")
	user_auth := base64.StdEncoding.EncodeToString([]byte(home_api_user + ":" + home_api_auth))
	header := req.Header{
		"Accept":        "application/json",
		"Authorization": "Basic " + user_auth,
	}
	param := req.Param{
		"message": *msg,
		"title":   *title,
	}
	// only url is required, others are optional.
	var resp, req_err = req.Post("https://"+home_api_domain+"/api/pushover", header, param)
	if req_err != nil {
		log.Fatal(req_err)
	}
	log.Info("response: ", resp)
}

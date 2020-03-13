package main

import (
	"github.com/akamensky/argparse"
    log "github.com/sirupsen/logrus"
    "github.com/gregdel/pushover"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
	"context"
	"strconv"
	"fmt"
)
type Event struct {
	Token string `json:"token"`
	User string `json:"user"`
	Message string `json:"message"`
	Title string `json:"title"`
}
type Response struct {
	success bool
	message string
}
var is_lambda bool
var verbose bool
var pushover_api_token string
var pushover_user_key string

func init() {
	// setup optional environment variables
	is_lambda = false
	ISLAMBDA := strings.ToLower(os.Getenv("ISLAMBDA"))
	if len(ISLAMBDA) > 0 {
		is_lambda, _ = strconv.ParseBool(ISLAMBDA)
	}
	verbose = false
	VERBOSE := strings.ToLower(os.Getenv("VERBOSE"))
	if len(VERBOSE) > 0 {
		verbose, _ = strconv.ParseBool(VERBOSE)
	}
	// setup required env vars
	pushover_api_token = strings.ToLower(os.Getenv("PUSHOVER_API_TOKEN"))
	pushover_user_key = strings.ToLower(os.Getenv("PUSHOVER_USER_KEY"))
	if len(pushover_api_token) <= 0 || len(pushover_user_key) <= 0 {
		fmt.Println("environment variables \"PUSHOVER_API_TOKEN\" and \"PUSHOVER_USER_KEY\" are required")
		os.Exit(1)
	}
	// setup logging
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	if verbose {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}
}
func sendMessage(message string, title string) (pushover.Response, error) {
	log.Trace("in sendMessage")
    app := pushover.New(pushover_api_token)
    recipient := pushover.NewRecipient(pushover_user_key)
    msg := pushover.NewMessageWithTitle(message, title)
    response, err := app.SendMessage(msg, recipient)
    if err != nil {
        log.Fatal("error while sending message: \""+err.Error()+"\"")
    } else {
    	log.Info("message sent successfully")
    }
	log.Info("response: \""+response.String()+"\"")
    return *response, err
}
func LambdaHandler(ctx context.Context, req Event) (Response, error) {
	log.Trace("in LambdaHandler")
	resp, err := sendMessage(req.Message, req.Title)
	if err != nil {
		return Response{success: false, message: err.Error()}, err
	} else {
		return Response{success: true, message: resp.String()}, nil
	}
}
func CommandLine() {
	log.Trace("in CommandLine")
	// parse arguments
	parser := argparse.NewParser(os.Args[0], "")
	msg := parser.String("m", "message", &argparse.Options{Required: true, Help: "message to send"})
	title := parser.String("t", "title", &argparse.Options{Required: true, Help: "title for message", Default: ""})
	verbose := parser.Flag("v", "verbose", &argparse.Options{Required: false, Help: "print logs to stdout", Default: false})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(parser.Usage(err))
		os.Exit(1)
	}
	if *verbose {
		log.SetLevel(log.TraceLevel)
	}
	_, err = sendMessage(*msg, *title)
	if err != nil {
		log.Error(err)
	}
}
func main() {
	log.Info("starting")
	if is_lambda {
		lambda.Start(LambdaHandler)
	} else if is_lambda == false {
		CommandLine()
	}
	log.Info("done.")
}

# -*- coding: utf-8 -*-

import json
import logging
import os
from chump import Application
import argparse
import sys

formatter = logging.Formatter('[%(asctime)s] - %(filename)s - %(levelname)s: %(message)s')
log = logging.getLogger(__file__)
log.setLevel(logging.DEBUG)
ch = logging.StreamHandler()

# optional env variables
is_lambda = os.getenv("ISLAMBDA", "false")
verbose = os.getenv("VERBOSE", "false")
# required env variables
pushover_api_token = os.getenv("PUSHOVER_API_TOKEN")
pushover_user_key = os.getenv("PUSHOVER_USER_KEY")

if not pushover_api_token or not pushover_user_key:
    print("environment variables \"PUSHOVER_API_TOKEN\" and \"PUSHOVER_USER_KEY\" are required")
    sys.exit(1)

def send_message(message, title):
    log.debug("in send_message")
    app = Application(pushover_api_token)
    user = app.get_user(pushover_user_key)
    msg = user.create_message(
        title=title,
        message=message
    )
    msg.send()
    results = {
        "success": msg.is_sent,
        "err": msg.error
    }
    log.info(results)
    return results


def lambda_handler(event, context):
    if verbose.upper() == "TRUE":
        ch.setLevel(logging.DEBUG)
    else:
        ch.setLevel(logging.INFO)
    ch.setFormatter(formatter)
    log.addHandler(ch)      
    log.debug("in lambda_handler")
    log.debug("event: " + json.dumps(event, indent=2))
    log.debug("context: " + str(context))
    if "Records" in event and "Sns" in event['Records'][0]:
        message = event['Records'][0]['Sns']['Message']
        title = event['Records'][0]['Sns']['Subject']
    else:
        message = event["message"]
        title = event["title"]
    return send_message(message, title)

def commandline():
    parser = argparse.ArgumentParser()
    parser.add_argument('--message', '-m', help="message to send", required=True)
    parser.add_argument('--title', '-t', help="title for the message", required=True)
    parser.add_argument('--verbose', '-v', action="store_true", help="verbose logging", required=False)
    args = parser.parse_args()
    if args.verbose:
        ch.setLevel(logging.DEBUG)
    else:
        ch.setLevel(logging.WARNING)
    ch.setFormatter(formatter)
    log.addHandler(ch)
    log.debug("in commandline")
    send_message(args.message, args.title)


if __name__ == '__main__':
    commandline()

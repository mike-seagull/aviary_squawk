aviary_squawk
===
send a pushover notification via my home-api
___
[![Build Status](https://travis-ci.com/mike-seagull/aviary_squawk.svg?branch=master)](https://travis-ci.com/mike-seagull/aviary_squawk)
![alt text](mine.png "squawk")

#### Environment Variables
* HOME_API_USER
* HOME_API_AUTH
* HOME_API_DOMAIN
#### To check a domain
```aviary_squawk -m $MESSAGE -t $TITLE```
#### Good use case
```/slow/process && aviary_sqawk -m "process successful" || aviary_squawk -m "process failed"```

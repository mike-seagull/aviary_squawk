pushover_squawk
===
sends a pushover notifications
___
![alt text](mine.png "squawk")

#### Environment Variables
##### Required:
* PUSHOVER_API_TOKEN
* PUSHOVER_USER_KEY
##### Optional:
* ISLAMBDA
* VERBOSE
#### To check a domain
```pushover_squawk -m $MESSAGE -t $TITLE```
#### Good use case
```/painfully/slow/process && pushover_squawk -m "process successful" || pushover_squawk -m "process failed"```

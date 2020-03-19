pushover_squawk
===
sends a pushover notifications
___
![alt text](mine.png "squawk")

#### Required Environment Variables
* PUSHOVER_API_TOKEN
* PUSHOVER_USER_KEY
#### To push a message

`squawk -m $MESSAGE -t $TITLE`
#### Lambda
It can also be used as a lambda function by providing an environment variable:
`ISLAMBDA=true`
##### To package for Lambda:
`invoke lambda.build`
##### To run as a Lambda function locally:
`invoke lambda.run`
#### Good use case
`/painfully/slow/process && pushover_squawk -m "process successful" || pushover_squawk -m "process failed"`
#### To bundle to a binary:
`invoke build --aws-lambda`

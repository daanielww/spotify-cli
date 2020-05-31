# Spotify-Cli

This is a service that aggregates data from the spotify API and stores it in AWS S3, the service is hosted in an EC2 instance and on AWS Lambda. An AWS Cloudwatch cron event is scheduled to trigger the lambda to gather data on a daily basis automatically.

## lambda
The lambda configuration cannot be run locally. If you wish to run this application locally, use the `server-local` configuration.

## server
Both a development and a production server is provided under `cmd`

While the Lambda will automatically run daily, it is nice to have it running on EC2, so the api can easily be invoked at anytime.

use `go run cmd/server/main.go` to run the production server
- Note: if you try to run this locally it will fail as it cannot contact AWS S3

use `go run cmd/server-local/main.go -id <ID> -secret <secret>` to run the development server
- provide the spotify api key id with the `id` flag
- provide the spotify api key secret with the `secret` flag
- the development server uses a mock AWS S3 server stored in local memory

## client
Both a development and production client is provided

use `go run client/main.go` to run the client
- provide the `development` flag to run the development client instead of the production one. The development client will interact with the development server
- provide the `tracks` flag to hit the tracks endpoint instead
## purpose
The aggregated data will be used for an AI project relating to song recommendations.
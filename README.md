# Spotify-Cli

service that aggregates data from the spotify API and stores it in AWS S3, the service is hosted in an EC2 instance.

## server
Both a development a production server is provided under `cmd`

use `go run cmd/server/main.go` to run the production server
- Note: if you try to run this locally it will fail as it cannot contact AWS S3


use `go run cmd/server-local/main.go -id <ID> -secret <secret>` to run the development server
- provide the spotify api key id with the `id` flag
- provide the spotify api key secret with the `secret` flag
- the development server uses a mock AWS S3 server

## client
Both a development and production client interface is provided `client`

use `go run client/main.go` to run the client
- provide the `development` flag to run the development client instead of the production one
- the development client will interact with the development server

## updates
Still a work-in-progress the data will be used for an AI project for song recommendations.
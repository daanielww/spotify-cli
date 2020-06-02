# Spotify-Cli

This is a service that aggregates data from the spotify API and stores it in AWS S3.

## Production Configuration (AWS lambda)
The production configuration is in `cmd/lambda-aws` and runs on AWS Lambda. An AWS Cloudwatch cron event is scheduled to trigger the lambda to gather data on a daily basis automatically.

The lambda configuration cannot be run locally. If you wish to run this application locally, use the `server-local` development configuration.

## Staging Server Configuration
A staging server is provided in `cmd/server-staging`. This server is hosted in an AWS EC2 container on the cloud. It will store data in a staging S3 bucket, so there is data isolation between development and production data.

The staging server provides endpoints that can be accessed via the cmd line client below

use `go run cmd/server-staging/main.go` to run the staging server on EC2
- Note: if you try to run this locally it will fail as it cannot contact AWS S3

## Local Development Server Configuration
A local development server is provided in `cmd/server-local`. This can be run locally on your machine and will use a mock AWS S3 server to store data.

use `go run cmd/server-local/main.go -id <ID> -secret <secret>` to run the development server
- provide the spotify api key id with the `id` flag
- provide the spotify api key secret with the `secret` flag
- the development server uses a mock AWS S3 server stored in local memory

## Client
A client is provided to access the local server and staging server

use `go run client/main.go` to run the client. By default this will hit the staging server endpoint
- by default it will hit the `playlist and albums` endpoint. If you wish to get the tracks provide the `tracks` flag to hit the tracks endpoint instead
- provide the `development` flag to hit the local server instead of staging.
## Purpose
The aggregated data will be used for an AI project relating to song recommendations.
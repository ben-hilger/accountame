# Accountame

[![Go - 22.0](https://img.shields.io/badge/Go-22.0-2ea44f?logo=go&logoColor=%2300ADD8)](https://tip.golang.org/doc/go1.22)
[![postgresql - 16](https://img.shields.io/static/v1?label=postgresql&message=16&color=%234169E1&logo=postgresql&logoColor=%234169E1)](https://tip.golang.org/doc/go1.22)


Accountame is a project that seeks to provide a tool to turn your phone  

* Hold yourself acoountable to daily/weekly/monthly actions
* Create an accountability group to keep you and your friends accountable
* Specify rewards for making a positive step forward, or a penalty for not following through

# Downloading the App 

The app is not live yet, but the plan is to make it available first on the iOS app store, and later the Google Play Store

# Contributing

Contributing is highly encouraged! Feel free to open a pull request and/or create an issue. Keep in mind there's currently only myself working on it in my free time, so it may take a couple of days to respond.

We want to maximize the developer experience to make contributing as easy as possible. So, we've minimized the dependencies, and compiled all of the required commands into a (hopefully easy to use) script.

All necessary commands to run the local environment are designed to be able to run in the root directory, limiting the need to navigate into specific project folders

## Dependencies

* [Docker](https://docs.docker.com/engine/install/)
* Bash 
* Xcode 
* [AtlasGo](https://atlasgo.io/getting-started)
  * **Note:** There will be work to replace this with the official docker container, removing this dependency

## Project Structure

* `api/`: The Go backend API
  * `Dockerfile`: The production Docker file without hot reload
* `database/`: Stores the database schema, and scripts to manage database migrations
  * `schema.sql`: Contains the working-in-progress schema
    * When making database changes, make the change in this file
  * `generated_schema.sql`: The generated sql schema from `atlas inspect`
    * DO NOT EDIT, this is regenerated when migrating, and used as the main schema when starting the stack
  * `local/`:
    * `migrate_local.sh`: Migrates the local stack database, setting all the required environment variables
  * `migrate.sh`: Migrates the specified database with the schema in `schema.sql`
    * For local development, run the `local/migrate-local.sh` script
  * `dump-db.sh`: Updates the `generated_schema.sql` with the latest schema
    * For local development, this is already executed when running the `local/migrate-local.sh` script
* `cdk/`: Stores the infrastructure as code, using AWS CDK
* `ios/`: Stores the Xcode ios project
* `docker-compose.yaml`: The local docker stack
* `run.sh`: Script with quality-of-life commands to make using this stack easier

## Running Locally

### Startup

There are two options to start the stack locally:
* `sh run.sh start`
* `docker-compose up`

All local environment variables are defined in the docker-compose.yaml. DO NOT USE THEM IN A PRODUCTION ENVIRONMENT

### DB Migrations

When updating the database, you can edit the schema in the `database/schema.sql` file

When you want to apply those updates, run `sh run.sh migrate` and follow the prompts to accept/deny the changes

### Hot Reload

* To remove the need to restart the stack while developing the Golang API, [Air](https://github.com/cosmtrek/air) is utilized for hot reloading

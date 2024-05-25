#!/bin/bash

ACTION=$1

if [[ "$ACTION" == "start" ]]; then
  docker-compose up --remove-orphans
elif [[ "$ACTION" == "migrate" ]]; then
  sh database/local/migrate-local.sh
elif [[ "$ACTION" == "ngrok" ]]; then
  docker run --net=host -it -e NGROK_AUTHTOKEN="$NGROK_AUTHTOKEN" ngrok/ngrok http 8080
else
  echo "Unsupported action $ACTION"
fi
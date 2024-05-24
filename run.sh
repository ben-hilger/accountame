#!/bin/bash

ACTION=$1

if [[ "$ACTION" == "start" ]]; then
  docker-compose up --remove-orphans
elif [[ "$ACTION" == "migrate" ]]; then
  sh database/local/migrate-local.sh
else
  echo "Unsupported action $ACTION"
fi
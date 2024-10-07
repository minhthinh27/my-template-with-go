#!/bin/sh

if [ -n "$APP_ENV" ]; then
    CONFIG_FILE="./configs/config-$APP_ENV.yaml"
else
    CONFIG_FILE="./configs/config.yaml"
fi

exec "$@" -conf "$CONFIG_FILE"
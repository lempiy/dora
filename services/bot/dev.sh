#!/usr/bin/env bash

declare NAME="dora_bot"

echo "|----------------------------|"
echo "|  BUILD DORA BOT DEV IMAGE  |"
echo "|----------------------------|"

echo "Start building image of ${NAME}..."
cp -r cert services/bot
cp pb/bot.proto services/bot
docker build -t dev-dora-${NAME}:latest services/bot
if [ $? -eq 0 ]; then
    echo "Successfully built ${NAME}!"
else
    exit 1
fi

# sed -i -e "s/${NAME}:v.*/${NAME}:${1}/g" ./docker/docker-compose-build.yml
echo "Remove temporal files ..."
rm -rf services/bot/cert
rm services/bot/bot.proto

echo "Build $1 | Local Time: $(date)"

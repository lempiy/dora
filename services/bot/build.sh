#!/usr/bin/env bash

declare NAME="dora_bot"

echo "|------------------------|"
echo "|  BUILD DORA BOT IMAGE  |"
echo "|------------------------|"

if [ -z "$1" ]
    then
        echo 'Build version should be supplied.'
        exit 1
fi

echo "Start building image of ${NAME}..."
cp -r cert services/bot
cp pb/bot.proto services/bot
docker build -t lempiy/${NAME}:$1 services/bot
if [ $? -eq 0 ]; then
    echo "Successfully built ${NAME}!"
else
    exit 1
fi

echo "Pushing ${NAME}:$1 image to docker-hub..."
docker push lempiy/${NAME}:$1
if [ $? -eq 0 ]; then
    echo "Image ${NAME}:$1 pushed successfully!"
else
    exit 1
fi

# sed -i -e "s/${NAME}:v.*/${NAME}:${1}/g" ./docker/docker-compose-build.yml
echo "Remove temporal files ..."
rm -rf services/bot/cert
rm services/bot/bot.proto

echo "Build $1 | Local Time: $(date)"

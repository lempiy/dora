#!/usr/bin/env bash

declare NAME="dora_api"

echo "|------------------------|"
echo "|  BUILD DORA API IMAGE  |"
echo "|------------------------|"

if [ -z "$1" ]
    then
        echo 'Build version should be supplied.'
        exit 1
fi

echo "Building binary..."
go build -o ${GOPATH}/src/github.com/lempiy/dora/services/api/${NAME} \
    github.com/lempiy/dora/services/api
if [ $? -eq 0 ]; then
    echo "Build for ${NAME} is done!"
else
    exit 1
fi

echo "Start building image of ${NAME}..."
cp -r cert services/api
docker build -t lempiy/${NAME}:$1 ${GOPATH}/src/github.com/lempiy/dora/services/api
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
rm services/api/${NAME}
rm -rf services/api/cert

echo "Build $1 | Local Time: $(date)"

FROM node:10.9
RUN mkdir /app
COPY . /app/
RUN cd app && \
    npm install && \
    cp servers.js ./node_modules/steam/lib/
WORKDIR /app
CMD [ "npm", "start" ]

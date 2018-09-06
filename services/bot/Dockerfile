FROM node:10.9
RUN mkdir /app
COPY . /app/
RUN cd app && \
    npm install
WORKDIR /app
CMD [ "npm", "start" ]

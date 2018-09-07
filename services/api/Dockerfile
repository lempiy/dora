FROM debian:9.5
RUN mkdir /app
COPY dora_api* /app/
COPY ./cert /app/cert
WORKDIR /app
CMD [ "/app/dora_api" ]

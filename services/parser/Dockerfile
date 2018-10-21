FROM ubuntu:16.04
RUN mkdir /app
COPY dora_parser* /app/
COPY ./cert /app/cert
WORKDIR /app
RUN mkdir .temp
CMD [ "/app/dora_parser" ]

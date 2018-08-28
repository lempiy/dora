FROM debian:9.5
RUN mkdir /app
COPY dora_parser* /app/
COPY ../../cert /app/cert
WORKDIR /app
ENTRYPOINT /app/dora_parser

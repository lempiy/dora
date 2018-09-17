'use strict'

const PORT = 6000 // TODO: ENV var
const PATH_TO_PROTO = './bot.proto' // TODO: Docker COPY + ENV var
const PATH_TO_CERT = './cert/ca.crt'
const PATH_TO_SERVER_CERT = './cert/server.crt' 
const PATH_TO_SERVER_KEY = './cert/server.key' 

const fs = require('fs')
const grpc = require('grpc')
const handlers = require('./handlers')
const service = grpc.load(PATH_TO_PROTO)

const cacert = fs.readFileSync(PATH_TO_CERT),
    cert = fs.readFileSync(PATH_TO_SERVER_CERT),
    key = fs.readFileSync(PATH_TO_SERVER_KEY),
    authority = {
        private_key: key,
        cert_chain: cert,
    }
const credentials = grpc.ServerCredentials.createSsl(null, [authority])
const server = new grpc.Server()

server.addProtoService(service.BotService.service, handlers)
server.bind(`0.0.0.0:${PORT}`, credentials)
console.log(`Starting server on port ${PORT}`)
server.start()

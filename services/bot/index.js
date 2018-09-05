'use strict'

const PORT = 7000 // TODO: ENV var
const PATH_TO_PROTO = './bot.proto' // TODO: Docker COPY + ENV var
const PATH_TO_CERT = './cert/ca.crt'
const PATH_TO_SERVER_CERT = './cert/server.crt' 
const PATH_TO_SERVER_KEY = './cert/server.key' 

const fs = require('fs')
const grpc = require('grpc')
const service = grpc.load(PATH_TO_PROTO)

const cacert = fs.readFileSync(PATH_TO_CERT),
    cert = fs.readFileSync(PATH_TO_SERVER_CERT),
    key = fs.readFileSync(PATH_TO_SERVER_KEY),
    authority = {
        private_key: key,
        cert_chain: cert,
    }
const credentials = grpc.ServerCredentials.createSsl(cacert, [authority])
const server = new grpc.Server()

server.addProtoService(service.BotService.service, {
    getPlayerCard: null,
    getMatchesHistory: null,
    getMatchDetails: null,
})
server.bind(`127.0.0.1:${PORT}`, credentials)
console.log(`Starting server on port ${PORT}`)
server.start()

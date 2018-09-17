'use strict';

const config = require("../config")
const DotaBot = require("dota2-bot")

const loginDetails = {
    account_name: config.steam_name,
    password: config.steam_pass
}

const bot = new DotaBot(loginDetails, true, false)
let connected = false
const _requestQueue = new Map()

// Watch common events
bot.Dota2.on('profileCardData', function (accId, data) {
    const resolve = _requestQueue.get(`profileCardData_${accId}`)
    if (!resolve) {
        throw new Error(`unexpected event with ID 'profileCardData_${accId}'`)
    }
    resolve(accId, data)
    _requestQueue.delete(`profileCardData_${accId}`)
})
bot.Dota2.on('playerMatchHistoryData', function (err, matchHistoryResponse) {
    const requestId = matchHistoryResponse.request_id
    const resolve = _requestQueue.get(`playerMatchHistoryData_${requestId}`)
    if (!resolve) {
        throw new Error(`unexpected event with ID 'playerMatchHistoryData_${requestId}'`)
    }
    resolve(matchHistoryResponse.matches || [])
    _requestQueue.delete(`playerMatchHistoryData_${requestId}`)
})
bot.Dota2.on('matchDetailsData', function(matchID, matchDetailsData) {
    const resolve = _requestQueue.get(`matchDetailsData_${matchID}`)
    if (!resolve) {
        throw new Error(`unexpected event with ID 'matchDetailsData_${matchID}'`)
    }
    resolve(matchID, matchDetailsData)
    _requestQueue.delete(`matchDetailsData_${matchID}`)
})

// TODO: request timeouts
module.exports = {
    get isConnected() {
        return connected
    },
    set isConnected(value) {
        throw new Error('cannot set isConnected directly')
    },
    connect() {
        bot.connect()
        connected = true
    },
    async getProfileData(profileID) {
        return new Promise((resolve, reject) => {
            _requestQueue.set(`profileCardData_${profileID}`, resolve)
            bot.schedule(() => bot.Dota2.requestProfileCard(+profileID))
        })
    },
    async getPlayerMatchHistory(profileID, options) {
        return new Promise((resolve, reject) => {
            const request_id = Math.floor(Math.random() * 1e7)
            const defaultOptions = {
                request_id,
                matches_requested: 10
            }
            _requestQueue.set(`playerMatchHistoryData_${request_id}`, resolve)
            bot.schedule(() => {
                bot.Dota2.requestPlayerMatchHistory(+profileID, {...defaultOptions, ...options})
            })
        })
    },
    async getMatchDetails(matchId) {
        return new Promise((resolve, reject) => {
            _requestQueue.set(`matchDetailsData_${matchID}`, resolve)
            bot.schedule(() => {
                bot.Dota2.requestMatchDetails(+matchId)
            })
        })
    }
}

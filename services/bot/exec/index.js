const config = require("../config")
const DotaBot = require("dota2-bot")

const loginDetails = {
    account_name: config.steam_name,
    password: config.steam_pass
}
const bot = new DotaBot(loginDetails, true, false)
const connected = false
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
bot.Dota2.on('playerMatchHistoryData', function (requestId, matchHistoryResponse) {
    const resolve = _requestQueue.get(`playerMatchHistoryData_${requestId}`)
    if (!resolve) {
        throw new Error(`unexpected event with ID 'playerMatchHistoryData_${requestId}'`)
    }
    resolve(requestId, matchHistoryResponse)
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
    set isConnected() {
        throw new Error('cannot set isConnected directly')
    },
    connect() {
        bot.connect()
        connected = true
    },
    getProfileData(profileID) {
        return new Promise((resolve, reject) => {
            _requestQueue.set(`profileCardData_${profileID}`, resolve)
            bot.schedule(() => bot.Dota2.requestProfileCard(profileID))
        })
    },
    getPlayerMatchHistory(profileID, options) {
        return new Promise((resolve, reject) => {
            const request_id = `${profileID}_${Math.floor(Math.random() * 10000)}`
            const defaultOptions = {
                request_id,
                matches_requested: 10
            }
            _requestQueue.set(`playerMatchHistoryData_${request_id}`, resolve)
            bot.schedule(() => {
                _requestQueue.set(`playerMatchHistoryData_${request_id}`, resolve)
                bot.Dota2.requestPlayerMatchHistory(profileID, {...defaultOptions, ...options})
            })
        })
    },
    getMatchDetails(matchId) {
        return new Promise((resolve, reject) => {
            _requestQueue.set(`matchDetailsData_${matchID}`, resolve)
            bot.schedule(() => {
                bot.Dota2.requestMatchDetails(matchId)
            })
        })
    }
}

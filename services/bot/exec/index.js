'use strict';

const config = require("../config")
const DotaBot = require("dota2-bot")

const loginDetails = {
    account_name: config.steam_name,
    password: config.steam_pass
}

const bot = new DotaBot(loginDetails, true, false)
let connected = false

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
            bot.schedule(() => bot.Dota2.requestProfileCard(+profileID, (err, card) => {
                if (err) {
                    reject(err)
                    return
                }
                resolve(card)
            }))
        })
    },
    async getPlayerMatchHistory(profileID, options) {
        return new Promise((resolve, reject) => {
            const defaultOptions = {
                matches_requested: 10
            }
            bot.schedule(() => {
                bot.Dota2.requestPlayerMatchHistory(+profileID, {...defaultOptions, ...options}, (err, matchHistoryResponse) => {
                    if (err) {
                        reject(err)
                        return
                    }
                    resolve(matchHistoryResponse.matches || [])
                })
            })
        })
    },
    async getMatchDetails(matchId) {
        return new Promise((resolve, reject) => {
            bot.schedule(() => {
                bot.Dota2.requestMatchDetails(+matchId, (err, matchDetailsData) => {
                    if (err) {
                        reject(err)
                        return
                    }
                    resolve(matchDetailsData)
                })
            })
        })
    }
}

'use strict';

const bot = require("../exec")

module.exports = {
    async getPlayerCard(call, callback) {
        const playerID = call.request.player_id
        if (!bot.isConnected) bot.connect()
        const card = await bot.getProfileData(playerID)
        if (!card) {
            callback('cannot get user card')
        }
        callback(null, card)
    },
    async getMatchesHistory(call, callback) {
        const playerID = call.request.player_id;
        if (!bot.isConnected) bot.connect()
        const matches = await bot.getPlayerMatchHistory(playerID)
        if (!card) {
            callback('cannot user matches history')
        }
        callback(null, matches)
    },
    async getMatchDetails(call, callback) {
        const matchID = call.request.match_id
        if (!bot.isConnected) bot.connect()
        const match = await bot.getMatchDetails(matchID)
        if (!card) {
            callback('cannot get match details')
        }
        callback(null, match)
    }
}

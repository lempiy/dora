'use strict';

const bot = require("../exec")
const long = require("long")

function deLong(obj) {
    if (Array.isArray(obj)) {
        return obj.map(v => deLong(v))
    } else if (long.isLong(obj)) {
        return obj.toString()
    } else if (obj instanceof Object) {
        var newObj = {};
        Object.keys(obj).forEach(d => {
            if (obj.hasOwnProperty(d)) {
                if (obj[d] instanceof Object) {
                    newObj[d] = deLong(obj[d])
                } else {
                    newObj[d] = obj[d]
                }
            }
        })
        return newObj
    } else {
        return obj
    }
}

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
        let matches = await bot.getPlayerMatchHistory(playerID, {})
        if (!matches) {
            callback('cannot user matches history')
        }
        callback(null, {
            matches: deLong(matches)
        })
    },
    async getMatchDetails(call, callback) {
        const matchID = call.request.match_id
        if (!bot.isConnected) bot.connect()
        const match = await bot.getMatchDetails(matchID)
        if (!match) {
            callback('cannot get match details')
        }
        callback(null, match)
    }
}
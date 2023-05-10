import Axios from "@/http";

export function createAlarmLevel(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/alarm-level', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveAlarmLevel(alarmLevelId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/alarm-level/${ alarmLevelId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateAlarmLevel(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/alarm-level', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteAlarmLevel(alarmLevelId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/alarm-level/${ alarmLevelId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteAlarmLevelWithFlag(alarmLevelId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/alarm-level/${ alarmLevelId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllAlarmLevels() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/alarm-level/all').then(res => {
            resolve(res)
        })
    })
}

export function queryAlarmLevel (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/alarm-level/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
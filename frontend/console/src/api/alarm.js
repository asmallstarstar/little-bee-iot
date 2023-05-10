import Axios from "@/http";

export function createAlarm(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/alarm', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveAlarm(alarmId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/alarm/${ alarmId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateAlarm(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/alarm', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteAlarm(alarmId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/alarm/${ alarmId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteAlarmWithFlag(alarmId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/alarm/${ alarmId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllAlarms() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/alarm/all').then(res => {
            resolve(res)
        })
    })
}

export function queryAlarm (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/alarm/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function ackAlarm(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/alarm/ack', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
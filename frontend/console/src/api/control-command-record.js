import Axios from "@/http";

export function createControlCommandRecord(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/control-command-record', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveControlCommandRecord(controlCommandRecordId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/control-command-record/${ controlCommandRecordId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateControlCommandRecord(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/control-command-record', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteControlCommandRecord(controlCommandRecordId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/control-command-record/${ controlCommandRecordId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteControlCommandRecordWithFlag(controlCommandRecordId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/control-command-record/${ controlCommandRecordId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllControlCommandRecords() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/control-command-record/all').then(res => {
            resolve(res)
        })
    })
}

export function queryControlCommandRecord (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/control-command-record/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
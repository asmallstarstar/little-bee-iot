import Axios from "@/http";

export function createSignalDi(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-di', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveSignalDi(signalDiId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/signal-di/${ signalDiId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateSignalDi(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/signal-di', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalDi(signalDiId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/signal-di/${ signalDiId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalDiWithFlag(signalDiId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/signal-di/${ signalDiId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllSignalDis() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/signal-di/all').then(res => {
            resolve(res)
        })
    })
}

export function querySignalDi (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-di/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
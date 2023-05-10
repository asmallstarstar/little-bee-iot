import Axios from "@/http";

export function createSignal(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveSignal(signalId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/signal/${ signalId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateSignal(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/signal', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignal(signalId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/signal/${ signalId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalWithFlag(signalId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/signal/${ signalId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllSignals() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/signal/all').then(res => {
            resolve(res)
        })
    })
}

export function querySignal (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
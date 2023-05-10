import Axios from "@/http";

export function createSignalSi(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-si', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveSignalSi(signalSiId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/signal-si/${ signalSiId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateSignalSi(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/signal-si', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalSi(signalSiId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/signal-si/${ signalSiId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalSiWithFlag(signalSiId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/signal-si/${ signalSiId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllSignalSis() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/signal-si/all').then(res => {
            resolve(res)
        })
    })
}

export function querySignalSi (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-si/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
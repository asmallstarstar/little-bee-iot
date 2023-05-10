import Axios from "@/http";

export function createSignalThreshhold(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-threshhold', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveSignalThreshhold(signalThreshholdId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/signal-threshhold/${ signalThreshholdId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateSignalThreshhold(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/signal-threshhold', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalThreshhold(signalThreshholdId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/signal-threshhold/${ signalThreshholdId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalThreshholdWithFlag(signalThreshholdId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/signal-threshhold/${ signalThreshholdId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllSignalThreshholds() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/signal-threshhold/all').then(res => {
            resolve(res)
        })
    })
}

export function querySignalThreshhold (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-threshhold/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
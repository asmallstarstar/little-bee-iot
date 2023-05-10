import Axios from "@/http";

export function createSignalOutput(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-output', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveSignalOutput(signalOutputId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/signal-output/${ signalOutputId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateSignalOutput(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/signal-output', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalOutput(signalOutputId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/signal-output/${ signalOutputId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalOutputWithFlag(signalOutputId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/signal-output/${ signalOutputId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllSignalOutputs() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/signal-output/all').then(res => {
            resolve(res)
        })
    })
}

export function querySignalOutput (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-output/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
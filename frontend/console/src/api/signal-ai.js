import Axios from "@/http";

export function createSignalAi(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-ai', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveSignalAi(signalAiId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/signal-ai/${ signalAiId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateSignalAi(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/signal-ai', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalAi(signalAiId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/signal-ai/${ signalAiId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalAiWithFlag(signalAiId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/signal-ai/${ signalAiId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllSignalAis() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/signal-ai/all').then(res => {
            resolve(res)
        })
    })
}

export function querySignalAi (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-ai/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
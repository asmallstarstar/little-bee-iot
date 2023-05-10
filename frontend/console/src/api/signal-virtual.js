import Axios from "@/http";

export function createSignalVirtual(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-virtual', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveSignalVirtual(signalVirtualId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/signal-virtual/${ signalVirtualId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateSignalVirtual(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/signal-virtual', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalVirtual(signalVirtualId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/signal-virtual/${ signalVirtualId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalVirtualWithFlag(signalVirtualId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/signal-virtual/${ signalVirtualId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllSignalVirtuals() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/signal-virtual/all').then(res => {
            resolve(res)
        })
    })
}

export function querySignalVirtual (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-virtual/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
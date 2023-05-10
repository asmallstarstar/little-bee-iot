import Axios from "@/http";

export function createSignalAcquisition(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-acquisition', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveSignalAcquisition(signalAcquisitionId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/signal-acquisition/${ signalAcquisitionId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateSignalAcquisition(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/signal-acquisition', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalAcquisition(signalAcquisitionId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/signal-acquisition/${ signalAcquisitionId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalAcquisitionWithFlag(signalAcquisitionId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/signal-acquisition/${ signalAcquisitionId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllSignalAcquisitions() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/signal-acquisition/all').then(res => {
            resolve(res)
        })
    })
}

export function querySignalAcquisition (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-acquisition/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
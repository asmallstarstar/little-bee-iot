import Axios from "@/http";

export function createSignalUnification(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-unification', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveSignalUnification(signalUnificationId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/signal-unification/${ signalUnificationId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateSignalUnification(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/signal-unification', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalUnification(signalUnificationId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/signal-unification/${ signalUnificationId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalUnificationWithFlag(signalUnificationId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/signal-unification/${ signalUnificationId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllSignalUnifications() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/signal-unification/all').then(res => {
            resolve(res)
        })
    })
}

export function querySignalUnification (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-unification/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
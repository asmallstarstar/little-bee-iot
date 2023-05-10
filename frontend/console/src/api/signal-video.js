import Axios from "@/http";

export function createSignalVideo(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-video', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveSignalVideo(signalVideoId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/signal-video/${ signalVideoId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateSignalVideo(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/signal-video', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalVideo(signalVideoId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/signal-video/${ signalVideoId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteSignalVideoWithFlag(signalVideoId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/signal-video/${ signalVideoId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllSignalVideos() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/signal-video/all').then(res => {
            resolve(res)
        })
    })
}

export function querySignalVideo (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/signal-video/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
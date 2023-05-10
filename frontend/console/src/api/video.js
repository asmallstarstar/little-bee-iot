import Axios from "@/http";

export function createVideo(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/video', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveVideo(videoId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/video/${ videoId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateVideo(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/video', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteVideo(videoId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/video/${ videoId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteVideoWithFlag(videoId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/video/${ videoId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllVideos() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/video/all').then(res => {
            resolve(res)
        })
    })
}

export function queryVideo (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/video/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
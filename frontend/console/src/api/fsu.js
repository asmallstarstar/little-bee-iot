import Axios from "@/http";

export function createFsu(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/fsu', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveFsu(fsuId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/fsu/${ fsuId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateFsu(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/fsu', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteFsu(fsuId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/fsu/${ fsuId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteFsuWithFlag(fsuId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/fsu/${ fsuId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllFsus() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/fsu/all').then(res => {
            resolve(res)
        })
    })
}

export function queryFsu (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/fsu/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
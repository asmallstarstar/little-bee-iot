import Axios from "@/http";

export function createSi(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/si', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveSi(siId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/si/${ siId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateSi(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/si', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteSi(siId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/si/${ siId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteSiWithFlag(siId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/si/${ siId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllSis() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/si/all').then(res => {
            resolve(res)
        })
    })
}

export function querySi (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/si/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
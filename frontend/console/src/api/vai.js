import Axios from "@/http";

export function createVai(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/vai', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveVai(vaiId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/vai/${ vaiId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateVai(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/vai', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteVai(vaiId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/vai/${ vaiId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteVaiWithFlag(vaiId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/vai/${ vaiId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllVais() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/vai/all').then(res => {
            resolve(res)
        })
    })
}

export function queryVai (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/vai/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
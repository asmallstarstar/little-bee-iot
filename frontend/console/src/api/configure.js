import Axios from "@/http";

export function createConfigure(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/configure', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveConfigure(configureId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/configure/${ configureId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateConfigure(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/configure', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteConfigure(configureId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/configure/${ configureId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteConfigureWithFlag(configureId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/configure/${ configureId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllConfigures() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/configure/all').then(res => {
            resolve(res)
        })
    })
}

export function queryConfigure (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/configure/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
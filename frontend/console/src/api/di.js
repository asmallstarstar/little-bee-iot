import Axios from "@/http";

export function createDi(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/di', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveDi(diId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/di/${ diId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateDi(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/di', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteDi(diId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/di/${ diId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteDiWithFlag(diId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/di/${ diId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllDis() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/di/all').then(res => {
            resolve(res)
        })
    })
}

export function queryDi (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/di/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
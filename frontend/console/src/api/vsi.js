import Axios from "@/http";

export function createVsi(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/vsi', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveVsi(vsiId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/vsi/${ vsiId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateVsi(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/vsi', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteVsi(vsiId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/vsi/${ vsiId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteVsiWithFlag(vsiId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/vsi/${ vsiId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllVsis() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/vsi/all').then(res => {
            resolve(res)
        })
    })
}

export function queryVsi (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/vsi/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
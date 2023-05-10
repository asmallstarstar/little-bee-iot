import Axios from "@/http";

export function createVdi(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/vdi', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveVdi(vdiId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/vdi/${ vdiId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateVdi(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/vdi', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteVdi(vdiId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/vdi/${ vdiId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteVdiWithFlag(vdiId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/vdi/${ vdiId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllVdis() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/vdi/all').then(res => {
            resolve(res)
        })
    })
}

export function queryVdi (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/vdi/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
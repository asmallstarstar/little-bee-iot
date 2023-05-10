import Axios from "@/http";

export function createFsuType(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/fsu-type', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveFsuType(fsuTypeId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/fsu-type/${ fsuTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateFsuType(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/fsu-type', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteFsuType(fsuTypeId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/fsu-type/${ fsuTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteFsuTypeWithFlag(fsuTypeId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/fsu-type/${ fsuTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllFsuTypes() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/fsu-type/all').then(res => {
            resolve(res)
        })
    })
}

export function queryFsuType (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/fsu-type/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
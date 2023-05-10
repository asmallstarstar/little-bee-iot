import Axios from "@/http";

export function createAreaType(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/area-type', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveAreaType(areaTypeId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/area-type/${ areaTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateAreaType(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/area-type', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteAreaType(areaTypeId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/area-type/${ areaTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteAreaTypeWithFlag(areaTypeId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/area-type/${ areaTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllAreaTypes() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/area-type/all').then(res => {
            resolve(res)
        })
    })
}

export function queryAreaType (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/area-type/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
import Axios from "@/http";

export function createDriverType(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/driver-type', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveDriverType(driverTypeId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/driver-type/${ driverTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateDriverType(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/driver-type', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteDriverType(driverTypeId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/driver-type/${ driverTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteDriverTypeWithFlag(driverTypeId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/driver-type/${ driverTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllDriverTypes() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/driver-type/all').then(res => {
            resolve(res)
        })
    })
}

export function queryDriverType (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/driver-type/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
import Axios from "@/http";

export function createDeviceType(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/device-type', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveDeviceType(deviceTypeId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/device-type/${ deviceTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateDeviceType(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/device-type', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteDeviceType(deviceTypeId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/device-type/${ deviceTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteDeviceTypeWithFlag(deviceTypeId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/device-type/${ deviceTypeId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllDeviceTypes() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/device-type/all').then(res => {
            resolve(res)
        })
    })
}

export function queryDeviceType (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/device-type/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
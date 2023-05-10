import Axios from "@/http";

export function createDriver(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/driver', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveDriver(driverId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/driver/${ driverId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateDriver(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/driver', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteDriver(driverId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/driver/${ driverId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteDriverWithFlag(driverId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/driver/${ driverId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllDrivers() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/driver/all').then(res => {
            resolve(res)
        })
    })
}

export function queryDriver (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/driver/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
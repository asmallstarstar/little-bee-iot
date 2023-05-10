import Axios from "@/http";

export function createLogicDevice(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/logic-device', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveLogicDevice(logicDeviceId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/logic-device/${ logicDeviceId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateLogicDevice(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/logic-device', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteLogicDevice(logicDeviceId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/logic-device/${ logicDeviceId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteLogicDeviceWithFlag(logicDeviceId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/logic-device/${ logicDeviceId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllLogicDevices() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/logic-device/all').then(res => {
            resolve(res)
        })
    })
}

export function queryLogicDevice (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/logic-device/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
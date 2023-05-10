import Axios from "@/http";

export function createLogicArea(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/logic-area', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveLogicArea(logicAreaId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/logic-area/${ logicAreaId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateLogicArea(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/logic-area', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteLogicArea(logicAreaId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/logic-area/${ logicAreaId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteLogicAreaWithFlag(logicAreaId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/logic-area/${ logicAreaId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllLogicAreas() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/logic-area/all').then(res => {
            resolve(res)
        })
    })
}

export function queryLogicArea (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/logic-area/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
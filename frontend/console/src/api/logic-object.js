import Axios from "@/http";

export function createLogicObject(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/logic-object', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveLogicObject(logicObjectId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/logic-object/${ logicObjectId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateLogicObject(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/logic-object', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteLogicObject(logicObjectId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/logic-object/${ logicObjectId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteLogicObjectWithFlag(logicObjectId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/logic-object/${ logicObjectId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllLogicObjects() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/logic-object/all').then(res => {
            resolve(res)
        })
    })
}

export function queryLogicObject (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/logic-object/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function getLogicObjectPath(logicObjectId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/path/logic-object/${ logicObjectId }`).then(res => {
            resolve(res)
        })
    })
}

export function getSignalPath(signalId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/path/signal/${ signalId }`).then(res => {
            resolve(res)
        })
    })
}
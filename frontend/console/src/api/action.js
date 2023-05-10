import Axios from "@/http";

export function createAction(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/action', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveAction(actionId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/action/${ actionId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateAction(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/action', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteAction(actionId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/action/${ actionId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteActionWithFlag(actionId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/action/${ actionId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllActions() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/action/all').then(res => {
            resolve(res)
        })
    })
}

export function queryAction (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/action/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
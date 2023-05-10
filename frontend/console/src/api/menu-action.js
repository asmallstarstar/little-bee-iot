import Axios from "@/http";

export function createMenuAction(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/menu-action', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveMenuAction(menuActionId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/menu-action/${ menuActionId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateMenuAction(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/menu-action', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteMenuAction(menuActionId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/menu-action/${ menuActionId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteMenuActionWithFlag(menuActionId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/menu-action/${ menuActionId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllMenuActions() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/menu-action/all').then(res => {
            resolve(res)
        })
    })
}

export function queryMenuAction (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/menu-action/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
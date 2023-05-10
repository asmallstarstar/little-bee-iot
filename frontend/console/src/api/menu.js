import Axios from "@/http";

export function createMenu(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/menu', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveMenu(menuId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/menu/${ menuId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateMenu(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/menu', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteMenu(menuId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/menu/${ menuId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteMenuWithFlag(menuId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/menu/${ menuId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllMenus() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/menu/all').then(res => {
            resolve(res)
        })
    })
}

export function queryMenu (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/menu/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function batchCreateMenuAction (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/menu-action/batch', params, { isJSON: true }).then(res => {
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

export function assembledMenus() {
    return new Promise((resolve, reject) => {
        Axios.get('/user/assembled-menus').then(res => {
            resolve(res)
        })
    })
}

export function getUserAllActions() {
    return new Promise((resolve, reject) => {
        Axios.get('/user/all-actions').then(res => {
            resolve(res)
        })
    })
}
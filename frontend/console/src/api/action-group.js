import Axios from "@/http";

export function createActionGroup(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/action-group', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveActionGroup(actionGroupId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/action-group/${ actionGroupId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateActionGroup(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/action-group', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteActionGroup(actionGroupId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/action-group/${ actionGroupId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteActionGroupWithFlag(actionGroupId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/action-group/${ actionGroupId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllActionGroups() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/action-group/all').then(res => {
            resolve(res)
        })
    })
}

export function queryActionGroup (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/action-group/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
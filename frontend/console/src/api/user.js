import Axios from "@/http";

export function login (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/public/user/login', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function createUser(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/user', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveUser(userId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/user/${ userId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateUser(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/user', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteUser(userId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/user/${ userId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteUserWithFlag(userId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/user/${ userId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllUsers() {
    return new Promise((resolve, reject) => {
        Axios.get('/user/all').then(res => {
            resolve(res)
        })
    })
}

export function getAllUserAlias() {
    return new Promise((resolve, reject) => {
        Axios.get('/user/alias-all').then(res => {
            resolve(res)
        })
    })
}

export function queryUser (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/user/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function joinGroup (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/user/join-group', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function removeGroup (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/user/remove-group', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function getUserGroup(userId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/user/join-group/${ userId }`).then(res => {
            resolve(res)
        })
    })
}

export function batchGrantUser (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/user/grant-user/batch', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function getUseGrant(userId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/user/grant-user/${ userId }`).then(res => {
            resolve(res)
        })
    })
}

export function modifyPasword(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/user/password', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function verifyUser (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/user/verify', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
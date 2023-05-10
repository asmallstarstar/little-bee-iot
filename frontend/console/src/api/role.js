import Axios from "@/http";

export function createRole(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/role', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveRole(roleId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/role/${ roleId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateRole(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/role', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteRole(roleId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/role/${ roleId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteRoleWithFlag(roleId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/role/${ roleId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllRoles() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/role/all').then(res => {
            resolve(res)
        })
    })
}

export function queryRole (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/role/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function grantBatchRole(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/user/grant-group/batch', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function getRoleGranted(roleId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/user/grant-group/${roleId}`).then(res => {
            resolve(res)
        })
    })
}

export function getRolePermissByUserId(userId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/user/grant-user-with-roles/${userId}`).then(res => {
            resolve(res)
        })
    })
}

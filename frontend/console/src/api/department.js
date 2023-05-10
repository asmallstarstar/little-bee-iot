import Axios from "@/http";

export function createDepartment(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/department', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveDepartment(departmentId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/department/${ departmentId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateDepartment(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/department', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteDepartment(departmentId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/department/${ departmentId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteDepartmentWithFlag(departmentId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/department/${ departmentId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllDepartments() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/department/all').then(res => {
            resolve(res)
        })
    })
}

export function queryDepartment (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/department/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
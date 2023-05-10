import Axios from "@/http";

export function createOutput(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/output', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveOutput(outputId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/output/${ outputId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateOutput(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/output', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteOutput(outputId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/output/${ outputId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteOutputWithFlag(outputId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/output/${ outputId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllOutputs() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/output/all').then(res => {
            resolve(res)
        })
    })
}

export function queryOutput (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/output/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
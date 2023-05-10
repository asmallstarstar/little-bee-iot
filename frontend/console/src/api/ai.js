import Axios from "@/http";

export function createAi(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/ai', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveAi(aiId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/ai/${ aiId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateAi(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/ai', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteAi(aiId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/ai/${ aiId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteAiWithFlag(aiId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/ai/${ aiId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllAis() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/ai/all').then(res => {
            resolve(res)
        })
    })
}

export function queryAi (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/ai/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
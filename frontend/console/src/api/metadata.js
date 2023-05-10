import Axios from "@/http";

export function createMetadata(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/metadata', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveMetadata(metadataId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/metadata/${ metadataId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateMetadata(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/metadata', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteMetadata(metadataId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/metadata/${ metadataId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteMetadataWithFlag(metadataId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/metadata/${ metadataId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllMetadatas() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/metadata/all').then(res => {
            resolve(res)
        })
    })
}

export function queryMetadata (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/metadata/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
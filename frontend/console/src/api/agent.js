import Axios from "@/http";

export function createAgent(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/agent', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function retrieveAgent(agentId) {
    return new Promise((resolve, reject) => {
        Axios.get(`/config/agent/${ agentId }`).then(res => {
            resolve(res)
        })
    })
}

export function updateAgent(params) {
    return new Promise((resolve, reject) => {
        Axios.put('/config/agent', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function deleteAgent(agentId) {
    return new Promise((resolve, reject) => {
        Axios.delete(`/config/agent/${ agentId }`).then(res => {
            resolve(res)
        })
    })
}

export function deleteAgentWithFlag(agentId) {
    return new Promise((resolve, reject) => {
        Axios.put(`/config/agent/${ agentId }`).then(res => {
            resolve(res)
        })
    })
}

export function getAllAgents() {
    return new Promise((resolve, reject) => {
        Axios.get('/config/agent/all').then(res => {
            resolve(res)
        })
    })
}

export function queryAgent (params) {
    return new Promise((resolve, reject) => {
        Axios.post('/config/agent/query', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
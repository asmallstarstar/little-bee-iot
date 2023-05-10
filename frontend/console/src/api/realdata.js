import Axios from "@/http";

export function getRealdata(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/realdata', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}

export function controlCommand(params) {
    return new Promise((resolve, reject) => {
        Axios.post('/control', params, { isJSON: true }).then(res => {
            resolve(res)
        })
    })
}
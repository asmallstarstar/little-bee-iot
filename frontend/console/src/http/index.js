import axios from 'axios'
import router from '../router'
import { useUsersStore } from '@/stores/user'
import { ElMessage } from 'element-plus'
import i18n from '@/lang/index'

const tip = (msg,cb=null) => {
    ElMessage({
        showClose: false,
        message: msg,
        type: 'error',
        onClose:cb
    })
}

const login = () => {
    router.replace({
        path: '/login',
        query: {
            redirect: router.currentRoute.fullPath
        }
    })
}

const handleError = (status, other) => {
    const t = i18n.global.t
    switch (status) {
        case 400:
            tip(t('language.http.BAD_REQUEST'))
            break
        case 401:
            tip(t('language.http.LOGIN_EXPIRED'),login())
            break
        case 403:
            const userStore = useUsersStore()
            userStore.resetUserInfo()
            tip(t('language.http.UNAUTHORIZED_OPERATION'),login())
            break
        case 404:
            tip(t('language.http.RESOURCE_NOT_FOUND'))
            break
        case 500:
            tip(t('language.http.INTERNAL_SERVER_ERROR'))
            break
        default:
            console.log(other);
    }
}


const instance = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL,
    timeout: 8000
})

instance.defaults.headers.post['Content-Type'] = 'application/json';
instance.defaults.headers.put['Content-Type'] = 'application/json';

instance.interceptors.request.use(
    config => {
        const usersStore = useUsersStore()
        const token = usersStore.token || "";
        token && (config.headers.token = token);
        return config;
    },
    error => Promise.error(error)
)

instance.interceptors.response.use(
    res => res.status === 200 ? Promise.resolve(res) : Promise.reject(res),
    error => {
        const { response } = error
        if (response) {
            handleError(response.status, response.data.message)
            return Promise.reject(response)
        } else {
            return Promise.reject(error)
        }
    });

export default instance;

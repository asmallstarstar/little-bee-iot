import { defineStore } from 'pinia'

export const useUsersStore = defineStore('users', {
    state: () => {
        return {
            userId: 0,
            departmentId: 0,
            userName: "",
            userNick: "",
            phone: "",
            email: "",
            sex: 0,
            avatar: "",
            status: 0,
            remark: "",
            token:""
        }
    },
    actions: {
        setUserInfo(userInfo,token){
            if(token){
                this.token = token
            }
            this.userId = userInfo.userId
            this.departmentId = userInfo.departmentId
            this.userName = userInfo.userName
            this.userNick = userInfo.userNick
            this.phone = userInfo.phone
            this.email = userInfo.email
            this.sex = userInfo.sex
            this.avatar = userInfo.avatar
            this.status = userInfo.status
            this.remark = userInfo.remark
        },
        resetUserInfo(){
            this.userId = 0
            this.departmentId = 0
            this.userName = ""
            this.userNick = ""
            this.phone = ""
            this.email = ""
            this.sex = ""
            this.avatar = ""
            this.status = 0
            this.remark = ""
            this.token = ""
        }
    },
    persist: {
        enabled: true,
        strategies: [
            { storage: sessionStorage, paths: ['userId','departmentId','userName','userNick','phone','email','sex','status','remark'] },
            { storage: localStorage, paths: ['token'] },
        ],
    }
})

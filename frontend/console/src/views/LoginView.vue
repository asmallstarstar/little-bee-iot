<template>
  <div class="login-wrap">
    <el-card class="login-card">
      <div class="login-title">LITTLE BEE CONSOLE</div>
      <el-form :model="loginParam" :rules="rules" ref="loginForm"  class="login-content">
        <el-form-item prop="userName">
          <el-input v-model="loginParam.userName" :placeholder="$t('language.login.USERNAME')">
            <template #prepend>
              <el-icon>
                <User></User>
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
              type="password"
              :placeholder="$t('language.login.PASSWORD')"
              v-model="loginParam.password"
              @keyup.enter="submitForm(loginForm)"
          >
            <template #prepend>
              <el-icon>
                <Lock></Lock>
              </el-icon>
            </template>
          </el-input>
        </el-form-item>
        <div class="login-btn">
          <el-button type="primary" @click="submitForm(loginForm)">{{$t('language.login.LOGIN')}}</el-button>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import {useUsersStore} from "../stores/user"
import i18n from '@/lang/index'
import { ref, reactive } from 'vue'
import { Lock, User } from '@element-plus/icons-vue'
import router from "@/router";
import {useRoute} from "vue-router";
import {login} from "../api/user";
import {useMenusStore} from "../stores/menu";
import {useTabsStore} from "../stores/tab";
import {ElMessage} from "element-plus";
import {assembledMenus, getUserAllActions} from "@/api/menu";
import {useActionsStore} from "@/stores/action";

const t = i18n.global.t
const loginParam = reactive({
  userName: 'admin',
  password: 'admin'
})
const loginForm = ref();
const rules = {
  userName: [{required: true, message: t('language.login.USERNAME'), trigger: 'blur'}],
  password: [{required: true, message: t('language.login.PASSWORD'), trigger: 'blur'}]
}
const route = useRoute()

const submitForm = (formEl) => {
  if (!formEl) return
  formEl.validate((valid) => {
    if (valid) {
      const params = {
        userName:loginParam.userName,
        password:loginParam.password
      }
      login(params).then(res=>{
        if(res.data.isSuccess){
          const userStore = useUsersStore()
          userStore.setUserInfo(res.data.result.user,res.data.result.token)
          const path = route.query.redirect
          if(path!=undefined){
            router.push(path).then(()=>{
              const menusStore = useMenusStore()
              const paths = menusStore.menus.filter(x=>x.path==path)
              if(paths.length>0){
                const tabsStore = useTabsStore()
                tabsStore.addTab(paths[0].menuName,paths[0].title,paths[0].path)
              }
            })
          }else{
            Promise.all([assembledMenus(), getUserAllActions()]).then(function (results) {
              const actionStore = useActionsStore()
              const menusStore = useMenusStore()
              const menus = results[0].data.result.items
              const userOwnedActions = results[1].data.result.actionName

              menusStore.setMenus(menus)
              menusStore.assembledRoutes()
              actionStore.setUserActions(userOwnedActions)

              router.push('/home')

            })
          }
        }else{
          ElMessage({
            showClose: true,
            message: res.data.errorDesc,
            type: 'error',
          })
        }
      })
    }
  })
}

</script>

<style scoped>
.login-wrap {
  position: relative;
  width: 100%;
  height: 100%;
  background-image: url(../assets/img/login-bg.jpg);
  background-size: 100%;
}

.login-card {
  position: absolute;
  left: 50%;
  top: 50%;
  width: 350px;
  margin: -190px 0 0 -175px;
  border-radius: 5px;
  background: rgba(255, 255, 255, 1);
  overflow: hidden;
}

.login-title {
  width: 100%;
  line-height: 50px;
  text-align: center;
  font-size: 20px;
}

.login-content {
  padding: 10px 10px;
}

.login-btn {
  text-align: center;
}

.login-btn button {
  width: 100%;
  height: 36px;
  margin-bottom: 5px;
}

</style>

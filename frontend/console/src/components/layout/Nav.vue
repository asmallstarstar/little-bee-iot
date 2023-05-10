<template>
  <div>
  <el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal"
           @select="handleSelect" :ellipsis="false" >
    <el-menu-item index="/logo">
      <template #title>
        <el-avatar shape="square" size="default" src="/logo.png" /><span>{{$t('language.common.TITLE')}}</span>
      </template>
    </el-menu-item>
    <template v-for="item in menusStore.getMenusByLevel(1)">
      <el-sub-menu :index="item.url" v-if="menusStore.haveChildMenu(item.menuId) && item.isShow">
        <template #title>
          <el-icon><component :is="item.icon" /></el-icon><span>{{ item.menuAlias }}</span>
        </template>
        <template v-for="childItem in menusStore.getMenusByParentId(item.menuId)">
          <el-menu-item :index="childItem.url" v-if="childItem.isShow">
            <template #title>
              <el-icon><component :is="childItem.icon" /></el-icon><span>{{ childItem.menuAlias }}</span>
            </template>
          </el-menu-item>
        </template>
      </el-sub-menu>
      <el-menu-item :index="item.url" v-else-if="item.isShow">
        <template #title>
          <el-icon><component :is="item.icon" /></el-icon><span>{{ item.menuAlias }}</span>
        </template>
      </el-menu-item>
    </template>
    <div class="flex-grow" />
    <el-sub-menu index="/language">
      <template #title><el-icon><Location /></el-icon>{{languageLabel}}</template>
      <el-menu-item v-for="item in i18n.languages" :index="item.locale" >
        <template #title><span>{{ item.label }}</span></template>
      </el-menu-item>
    </el-sub-menu>
    <el-sub-menu index="/user">
      <template #title><el-icon><User /></el-icon>{{usersStore.userNick}}</template>
      <el-menu-item index="/user-info" @click="showUserInfo"><el-icon><InfoFilled /></el-icon>{{$t('language.user.MODIFY_USER_INFO')}}</el-menu-item>
      <el-menu-item index="/password" @click="showUserPassword"><el-icon><Lock /></el-icon>{{$t('language.user.MODIFY_USER_PASSWORD')}}</el-menu-item>
      <el-menu-item index="/quit"><el-icon><SwitchButton /></el-icon>{{$t('language.common.QUIT')}}</el-menu-item>
    </el-sub-menu>
  </el-menu>
    <user-info ref="userInfoInstanceRef"></user-info>
    <user-password ref="userPasswordInstanceRef"></user-password>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import {useMenusStore} from "@/stores/menu";
import {useUsersStore} from "../../stores/user";
import {useTabsStore} from "@/stores/tab";
import { useI18n } from 'vue-i18n'
import i18n from "@/lang";
import router from "@/router";
import UserInfo from "@/components/user/UserInfo.vue";
import UserPassword from "@/components/user/UserPassword.vue";

const t = i18n.global.t
const activeIndex =ref('/home')
const menusStore = useMenusStore()
const usersStore = useUsersStore()
const tabsStore = useTabsStore()
const userInfoInstanceRef = ref()
const userPasswordInstanceRef = ref()

const { locale } = useI18n()
const languageLabel = ref(t('language.common.LANGUAGE'))
languageLabel.value = i18n.getCurrentLanguageLabel()

function handleSelect(key, keyPath) {
  //language
  const locales = i18n.languages.map(x=>x.locale)
  const localesSet = new Set(locales)
  if( localesSet.has(key) ){
    localStorage.setItem("language",key)
    locale.value = key
    languageLabel.value=i18n.getCurrentLanguageLabel()
    return
  }

  const menus = tabsStore.getSignalTreeMenu()
  const keys = menus.filter(x=>x.url===key)
  if(keys.length>0){//depends on signal tree
    const res = tabsStore.findTabByPath(keys[0].url)
    if(res==null){
      router.push(keys[0].url)
    }else{
      router.push(res.path)
    }
  }else{
    const links = menusStore.getMenusByLevel(2)
    const linksSet = new Set(links.map(x=>x.url))
    if( linksSet.has(key) || key==='/home' ){
      router.push(key)
    }
  }

  //quit
  if(key==='/quit'){
    usersStore.resetUserInfo()
    tabsStore.removeAll()
    router.push('/login')
  }
}

function showUserInfo() {
  userInfoInstanceRef.value.showUserDialog()
}

function showUserPassword() {
  userPasswordInstanceRef.value.showModifyPasswordDialog()
}

</script>
<style scoped>
.flex-grow {
  flex-grow: 1;
}
</style>

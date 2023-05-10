<template>
  <div class="tabs">
    <ul>
      <li class="tabs-li" v-for="(item, index) in tabsStore.tabs" :class="{ active: isActive(item.path) }" :key="index">
        <router-link :to="item.path" class="tabs-li-title">{{ item.title }}</router-link>
        <el-icon @click="closeTab(item.path)" v-if="item.name!=='home'"><Close /></el-icon>
      </li>
    </ul>
    <div class="tabs-close-box">
      <el-dropdown @command="handleTabs">
        <el-button size="small" type="primary">
          {{$t('language.tabs.TABS_OPTION')}}
          <el-icon class="el-icon--right"><arrow-down /></el-icon>
        </el-button>
        <template #dropdown>
          <el-dropdown-menu size="small">
            <el-dropdown-item command="other">{{$t('language.tabs.CLOSE_OTHERS')}}</el-dropdown-item>
            <el-dropdown-item command="all">{{$t('language.tabs.CLOSE_ALL')}}</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
</template>

<script setup>
import { Close } from '@element-plus/icons-vue'
import {useTabsStore} from "../../stores/tab";
import {useRoute} from "vue-router";
import i18n from "@/lang";
import {useMenusStore} from "@/stores/menu";

const t = i18n.global.t
const tabsStore = useTabsStore()
const menusStore = useMenusStore()

function homeMenu() {
  const homes = menusStore.menus.filter(x=>x.menuName=='home')
  if(homes.length>0){
    tabsStore.addTab(homes[0].menuName,homes[0].title,homes[0].path)
  }
}
homeMenu()

const route = useRoute();
const isActive = (path) => {
  return path === route.path;
}

const closeTab = (tabPath) => {
  tabsStore.removeTab(tabPath)
}

const closeAll = () => {
  tabsStore.removeAll()
}

const closeOther = () => {
  tabsStore.removeOthers(route.path)
}

const handleTabs = (command) => {
  command === 'other' ? closeOther() : closeAll();
}

</script>

<style scoped>

.tabs {
  position: relative;
  height: 30px;
  overflow: hidden;
  background: #fff;
  padding-right: 120px;
  box-shadow: 0 5px 10px #ddd;
}

.tabs ul {
  box-sizing: border-box;
  width: 100%;
  height: 100%;
}

.tabs-li {
  display: flex;
  align-items: center;
  float: left;
  margin: 3px 5px 2px 3px;
  border-radius: 3px;
  font-size: 12px;
  overflow: hidden;
  cursor: pointer;
  height: 23px;
  border: 1px solid #e9eaec;
  background: #fff;
  padding: 0 5px 0 12px;
  color: #666;
  -webkit-transition: all 0.3s ease-in;
  -moz-transition: all 0.3s ease-in;
  transition: all 0.3s ease-in;
}

.tabs-li:not(.active):hover {
  background: #f8f8f8;
}

.tabs-li.active {
  border: 1px solid #409EFF;
  background-color: #409eff;
}

.tabs-li-title {
  float: left;
  max-width: 80px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  margin-right: 5px;
  color: #666;
}

.tabs-li.active .tabs-li-title {
  color: #fff;
}

.tabs-close-box {
  position: absolute;
  right: 0;
  top: 0;
  box-sizing: border-box;
  padding-top: 1px;
  text-align: center;
  width: 110px;
  height: 30px;
  background: #fff;
  box-shadow: -3px 0 15px 3px rgba(0, 0, 0, 0.1);
  z-index: 10;
}

</style>

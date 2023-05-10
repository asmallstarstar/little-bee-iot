<template>
  <el-tree-select
      v-model="currentMenu"
      :data="menusStore.menuTree"
      :render-after-expand="false"
      node-key="id"
      default-expand-all
      :placeholder="$t('language.menuAction.SELECT_MENU')"
      @change="$emit('selectMenu',currentMenu)"
  />
</template>

<script setup>
import {useMenusStore} from "@/stores/menu";
import {ref} from "vue";
import i18n from "@/lang";

const menusStore =  useMenusStore()
const currentMenu = ref()
const t = i18n.global.t


const getCurrentMenuId = () =>{
  return currentMenu.value
}

defineExpose({
  getCurrentMenuId
})

</script>

<style scoped>

</style>

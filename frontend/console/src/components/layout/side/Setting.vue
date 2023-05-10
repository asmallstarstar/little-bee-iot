<template>
  <el-scrollbar :max-height="maxHeight">
  <el-menu :default-active="activeIndex" class="el-menu-demo" mode="vertical"
           @select="handleSelect" :ellipsis="false" router>
    <template v-for="item in menusStore.getMenusByLevel(3)">
      <el-menu-item  :index="item.path" v-if="item.isShow">
        <template #title>
          <el-icon><component :is="item.icon" /></el-icon><span>{{ item.menuAlias }}</span>
        </template>
      </el-menu-item>
    </template>
  </el-menu>
</el-scrollbar>
</template>

<script setup>
import {useMenusStore} from "@/stores/menu";
import i18n from "@/lang";
import {ref} from "vue";
let maxHeight = ref(window.innerHeight - 60)

const menusStore = useMenusStore()
const t = i18n.global.t
const activeIndex =ref('')

function handleSelect(key, keyPath) {

}

</script>

<style scoped>

</style>
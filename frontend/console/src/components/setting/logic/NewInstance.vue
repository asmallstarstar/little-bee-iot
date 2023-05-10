<template>
  <el-button-group class="ml-4">
    <el-button @click="createLogic('AREA')" plain size="small" :icon="OfficeBuilding"
               v-if="parentObjectType===signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_CENTER ||
               parentObjectType===signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_AREA" >
      {{$t('language.newInstance.NEW_AREA')}}
    </el-button>
    <el-button @click="createLogic('DEVICE')" plain size="small" :icon="Discount"
               v-if="parentObjectType===signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_CENTER ||
               parentObjectType === signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_AREA ||
               parentObjectType ===signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_DEVICE" >
      {{$t('language.newInstance.NEW_DEVICE')}}
    </el-button>
    <el-button @click="createSignal('AI')" plain size="small" :icon="Drizzling" >{{$t('language.newInstance.NEW_AI')}}</el-button>
    <el-button @click="createSignal('DI')" plain size="small" :icon="Open" >{{$t('language.newInstance.NEW_DI')}}</el-button>
    <el-button @click="createSignal('SI')" plain size="small" :icon="Document" >{{$t('language.newInstance.NEW_SI')}}</el-button>
    <el-button @click="createSignal('AO')" plain size="small" :icon="Notebook" >{{$t('language.newInstance.NEW_AO')}}</el-button>
    <el-button @click="createSignal('DO')" plain size="small" :icon="Notebook" >{{$t('language.newInstance.NEW_DO')}}</el-button>
    <el-button @click="createSignal('SO')" plain size="small" :icon="Notebook" >{{$t('language.newInstance.NEW_SO')}}</el-button>
    <el-button @click="createSignal('VAI')" plain size="small" :icon="Collection" >{{$t('language.newInstance.NEW_VAI')}}</el-button>
    <el-button @click="createSignal('VDI')" plain size="small" :icon="ScaleToOriginal" >{{$t('language.newInstance.NEW_VDI')}}</el-button>
    <el-button @click="createSignal('VSI')" plain size="small" :icon="CopyDocument" >{{$t('language.newInstance.NEW_VSI')}}</el-button>
    <el-button @click="createSignal('VIDEO')" plain size="small" :icon="VideoCamera" >{{$t('language.newInstance.NEW_VIDEO')}}</el-button>
  </el-button-group>
</template>

<script setup>
import { OfficeBuilding, Discount, Drizzling, Open, Document, Notebook,
  Collection, ScaleToOriginal, CopyDocument,VideoCamera} from '@element-plus/icons-vue'
import {
  signalTreeObjectTypeEnum,
  signalTypeEnum, operatedTypeEnum
} from "@/utils/common";
import router from "@/router";
import {useSignalsStore} from "@/stores/signal";
const props = defineProps(['parentObjectType','parentLogicObjectId'])
const signalsStore = useSignalsStore()

function createLogic(objType) {
  let t = signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_UNKNOWN
  if(objType=='AREA'){
    t = signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_AREA
  }
  if(objType=='DEVICE'){
    t = signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_DEVICE
  }
  router.push({ name: 'logicObject',
    params: {
      logicObjectId: props.parentLogicObjectId,
      operatedType:operatedTypeEnum.OPERATED_TYPE_CREATE,
      objectType:t
    }
  })
}

function createSignal(newIdName) {
  router.push({ name: 'signal',
    params: {
      signalId: 0,
      operatedType:operatedTypeEnum.OPERATED_TYPE_CREATE,
      objectType: createNewObjectParam(newIdName),
      parentLogicObjectId:props.parentLogicObjectId
    }
  })
}

function createNewObjectParam(newIdName) {
  if(newIdName==='AI'){
    return signalTypeEnum.SIGNAL_TYPE_AI
  }
  if(newIdName==='DI') {
    return signalTypeEnum.SIGNAL_TYPE_DI
  }
  if(newIdName==='SI'){
    return signalTypeEnum.SIGNAL_TYPE_SI
  }
  if(newIdName==='VAI'){
    return signalTypeEnum.SIGNAL_TYPE_VAI
  }
  if(newIdName==='VDI') {
    return signalTypeEnum.SIGNAL_TYPE_VDI
  }
  if(newIdName==='VSI'){
    return signalTypeEnum.SIGNAL_TYPE_VSI
  }
  if(newIdName==='AO') {
    return signalTypeEnum.SIGNAL_TYPE_AO
  }
  if(newIdName==='DO') {
    return signalTypeEnum.SIGNAL_TYPE_DO
  }
  if(newIdName==='SO') {
    return signalTypeEnum.SIGNAL_TYPE_SO
  }
  if(newIdName==='VIDEO'){
    return signalTypeEnum.SIGNAL_TYPE_VIDEO
  }
  return signalTypeEnum.SIGNAL_TYPE_UNKNOWN
}

</script>

<style scoped>

</style>
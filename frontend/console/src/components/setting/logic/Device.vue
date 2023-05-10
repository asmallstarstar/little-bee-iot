<template>
  <el-form :model="device">
    <el-form-item :label="$t('language.logicObject.DEVICE_NAME')" label-position="left" label-width="120px" prop="deviceAlias">
      <el-input v-model="device.deviceAlias" />
    </el-form-item>
    <el-form-item :label="$t('language.deviceType.DEVICE_TYPE_NAME')" label-position="left" label-width="120px" prop="areaType">
      <el-select v-model="device.deviceType" class="m-2" :placeholder="$t('language.common.SELECT')">
        <el-option v-for="item in deviceTypes" :key="item.deviceTypeId" :label="item.deviceTypeName" :value="item.deviceTypeId"/>
      </el-select>
    </el-form-item>
  </el-form>
</template>

<script setup>
import {ref, toRaw, watch} from "vue";
import {operatedTypeEnum} from "@/utils/common";
import {retrieveLogicDevice} from "@/api/logic-device";
import {getAllDeviceTypes} from "@/api/device-type";

const deviceTypes = ref()
const props = defineProps(['logicObjectId','operatedType'])
const  device = ref({
  logicObjectId:0,
  deviceAlias:"",
  deviceType:null
})

watch(props,(newX)=>{
  if(newX.operatedType!==operatedTypeEnum.OPERATED_TYPE_CREATE){
    retrieveLogicDevice(newX.logicObjectId).then(x=>{
      delete x.data.result['@type']
      device.value = x.data.result
    })
  }else{
    device.value = {
      logicObjectId:0,
      deviceAlias:"",
      deviceType:null
    }
  }
},{immediate: true, deep: true})

getAllDeviceTypes().then(function (x) {
  deviceTypes.value = x.data.result.items
})

const getDevice = () => {
  const {...rawData} = toRaw(device.value)
  return rawData
}

defineExpose({
  getDevice
})

</script>

<style scoped>
.el-input ,.el-select {
  width: 500px;
}
</style>
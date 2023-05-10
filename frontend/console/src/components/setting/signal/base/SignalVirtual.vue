<template>
  <el-form :model="signalVS">
    <el-form-item :label="$t('language.signal.VIRTUAL_SIGNAL_EXPRESS')" label-position="left" label-width="120px" prop="signalExpress">
      <el-input v-model="signalVS.signalExpress"/>
    </el-form-item>
  </el-form>
</template>

<script setup>
import {ref, toRaw, watch} from "vue";
import {operatedTypeEnum} from "@/utils/common";
import {retrieveSignalVirtual} from "@/api/signal-virtual";

const props = defineProps(['signalId','operatedType'])
const  signalVS = ref({
  signalId:-1,
  signalExpress:""
})

watch(props,(newX)=>{
  if(newX.operatedType!==operatedTypeEnum.OPERATED_TYPE_CREATE){
    retrieveSignalVirtual(newX.signalId).then(x=>{
      delete x.data.result['@type']
      signalVS.value = x.data.result
    })
  }else{
    signalVS.value = {
      signalId:-1,
      signalExpress:""
    }
  }
},{immediate: true, deep: true})

const getSignalVirtual = () => {
  const {...rawData} = toRaw(signalVS.value)
  return rawData
}

defineExpose({
  getSignalVirtual
})

</script>

<style scoped>
.el-input ,.el-select {
  width: 500px;
}
</style>
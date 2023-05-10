<template>
  <el-form :model="signalSi">
    <el-form-item :label="$t('language.signal.RECORD_PERIOD')" label-position="left" label-width="120px" prop="recordPeriod">
      <el-input v-model="signalSi.recordPeriod"/>
    </el-form-item>
  </el-form>
</template>

<script setup>
import {ref, toRaw, watch} from "vue";
import {operatedTypeEnum} from "@/utils/common";
import {retrieveSignalSi} from "@/api/signal-si";

const props = defineProps(['signalId','operatedType'])
const  signalSi = ref({
  signalId:-1,
  recordPeriod:60
})

watch(props,(newX)=>{
  if(newX.operatedType!==operatedTypeEnum.OPERATED_TYPE_CREATE){
    retrieveSignalSi(newX.signalId).then(x=>{
      delete x.data.result['@type']
      signalSi.value = x.data.result
    })
  }else{
    signalSi.value = {
      signalId:-1,
      recordPeriod:60
    }
  }
},{immediate: true, deep: true})

const getSignalSi = () => {
  const {...rawData} = toRaw(signalSi.value)
  return rawData
}

defineExpose({
  getSignalSi
})

</script>

<style scoped>
.el-input ,.el-select {
  width: 500px;
}
</style>
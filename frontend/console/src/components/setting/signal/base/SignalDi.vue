<template>
  <el-form :model="signalDi">
    <el-form-item :label="$t('language.signalDi.SIGNAL_DI_VALUE_FLIP')" label-position="left" label-width="120px" prop="multipleRate">
      <el-checkbox v-model="signalDi.isFlip" :label="$t('language.signalDi.SIGNAL_DI_IS_FLIP')" />
    </el-form-item>
    <el-form-item :label="$t('language.signal.RECORD_PERIOD')" label-position="left" label-width="120px" prop="recordPeriod">
      <el-input v-model="signalDi.recordPeriod"/>
    </el-form-item>
    <el-form-item :label="$t('language.signalDi.SIGNAL_DI_HIGH_DESC')" label-position="left" label-width="120px" prop="highDesc">
      <el-input v-model="signalDi.highDesc"/>
    </el-form-item>
    <el-form-item :label="$t('language.signalDi.SIGNAL_DI_LOW_DESC')" label-position="left" label-width="120px" prop="lowDesc">
      <el-input v-model="signalDi.lowDesc"/>
    </el-form-item>
  </el-form>
</template>

<script setup>
import {ref, toRaw, watch} from "vue";
import {operatedTypeEnum} from "@/utils/common";
import {retrieveSignalDi} from "@/api/signal-di";

const props = defineProps(['signalId','operatedType'])
const  signalDi = ref({
  signalId:-1,
  isFlip:false,
  recordPeriod:60,
  highDesc:"",
  lowDesc:"",
})

watch(props,(newX)=>{
  if(newX.operatedType!==operatedTypeEnum.OPERATED_TYPE_CREATE){
    retrieveSignalDi(newX.signalId).then(x=>{
      delete x.data.result['@type']
      signalDi.value = x.data.result
    })
  }else{
    signalDi.value = {
      signalId:-1,
      isFlip:false,
      recordPeriod:60,
      highDesc:"",
      lowDesc:"",
    }
  }
},{immediate: true, deep: true})

const getSignalDi = () => {
  const {...rawData} = toRaw(signalDi.value)
  return rawData
}

defineExpose({
  getSignalDi
})

</script>

<style scoped>
.el-input ,.el-select {
  width: 500px;
}
</style>
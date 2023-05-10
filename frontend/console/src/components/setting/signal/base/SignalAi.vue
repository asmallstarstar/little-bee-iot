<template>
  <el-form :model="signalAi">
    <el-form-item :label="$t('language.signalAi.SIGNAL_AI_MULTIPLE_RATE')" label-position="left" label-width="120px" prop="multipleRate">
      <el-input v-model="signalAi.multipleRate"/>
    </el-form-item>
    <el-form-item :label="$t('language.signalAi.SIGNAL_AI_OFFSET')" label-position="left" label-width="120px" prop="offsetValue">
      <el-input v-model="signalAi.offsetValue"/>
    </el-form-item>
    <el-form-item :label="$t('language.signal.RECORD_PERIOD')" label-position="left" label-width="120px" prop="recordPeriod">
      <el-input v-model="signalAi.recordPeriod"/>
    </el-form-item>
    <el-form-item :label="$t('language.signalAi.SIGNAL_AI_DECIMAL_PRECISION')" label-position="left" label-width="120px" prop="decimalPrecision">
      <el-input v-model="signalAi.decimalPrecision"/>
    </el-form-item>
    <el-form-item :label="$t('language.signalAi.SIGNAL_AI_VALUE_UNIT')" label-position="left" label-width="120px" prop="valueUnit">
      <el-input v-model="signalAi.valueUnit"/>
    </el-form-item>
  </el-form>
</template>

<script setup>
import {ref, toRaw, watch} from "vue";
import {operatedTypeEnum} from "@/utils/common";
import {retrieveSignalAi} from "@/api/signal-ai";

const props = defineProps(['signalId','operatedType'])
const  signalAi = ref({
  signalId:-1,
  multipleRate:1.0,
  offsetValue:0.0,
  recordPeriod:60,
  decimalPrecision:2,
  valueUnit:""
})

watch(props,(newX)=>{
  if(newX.operatedType!==operatedTypeEnum.OPERATED_TYPE_CREATE){
    retrieveSignalAi(newX.signalId).then(x=>{
      delete x.data.result['@type']
      signalAi.value = x.data.result
    })
  }else{
    signalAi.value = {
      signalId:-1,
      multipleRate:1.0,
      offsetValue:0.0,
      recordPeriod:60,
      decimalPrecision:2,
      valueUnit:""
    }
  }
},{immediate: true, deep: true})

const getSignalAi = () => {
  const {...rawData} = toRaw(signalAi.value)
  return rawData
}

defineExpose({
  getSignalAi
})

</script>

<style scoped>
.el-input ,.el-select {
  width: 500px;
}
</style>
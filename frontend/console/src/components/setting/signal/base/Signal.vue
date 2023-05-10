<template>
  <el-form :model="signal" :rules="rules">
    <el-form-item :label="$t('language.signal.SIGNAL_ID')" label-position="left" label-width="120px" prop="signalId" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_EDIT" required >
      <el-input v-model="signal.signalId" readonly />
    </el-form-item>
    <el-form-item :label="$t('language.signal.SIGNAL_NAME')" label-position="left" label-width="120px" prop="signalName" required>
      <el-input v-model="signal.signalName"/>
    </el-form-item>
    <el-form-item :label="$t('language.common.ALIAS')" label-position="left" label-width="120px" prop="signalAlias">
      <el-input v-model="signal.signalAlias"/>
    </el-form-item>
    <el-form-item :label="$t('language.common.INDUSTRIAL_INTERNET_ID')" label-position="left" label-width="120px" prop="industrialInternetId">
      <el-input v-model="signal.industrialInternetId"/>
    </el-form-item>
    <el-form-item :label="$t('language.signal.SIGNAL_UNIFICATION')" label-position="left" label-width="120px" prop="signalUnificationName">
      <el-select v-model="signal.signalUnificationId" class="m-2" :placeholder="$t('language.common.SELECT')" >
        <el-option v-for="item in signalUnifications" :key="item.signalUnificationId" :label="item.signalUnificationName" :value="item.signalUnificationId"/>
      </el-select>
    </el-form-item>
    <el-form-item :label="$t('language.signal.SIGNAL_TYPE')" label-position="left" label-width="120px" prop="signalType" required>
      <el-select v-model="signal.signalType" class="m-2" :placeholder="$t('language.common.SELECT')" disabled>
        <el-option v-for="item in signalTypeText" :key="item.index" :label="item.name" :value="item.index"/>
      </el-select>
    </el-form-item>
    <el-form-item :label="$t('language.common.PARENT_OBJECT')" label-position="left" label-width="120px" prop="parentLogicObjectPath" required>
      <el-input v-model="signal.parentLogicObjectPath" readonly/>
    </el-form-item>
    <el-form-item :label="$t('language.common.POSITION_AMONG_BROTHERS')" label-position="left" label-width="120px" prop="positionAmongBrothers">
      <el-input v-model="signal.positionAmongBrothers"/>
    </el-form-item>
    <el-form-item :label="$t('language.signal.APPLIED')" label-position="left" label-width="120px" prop="applied">
      <el-checkbox v-model="signal.applied" :label="$t('language.signal.IS_APPLIED')" />
    </el-form-item>
  </el-form>
</template>

<script setup>
import { Search} from '@element-plus/icons-vue'
import {ref, toRaw, watch} from "vue";
import {getParentObjectPath, operatedTypeEnum, signalTypeEnum, signalTypeText} from "@/utils/common";
import {retrieveSignal} from "@/api/signal";
import {getAllSignalUnifications} from "@/api/signal-unification";

const signalUnifications = ref([])
const props = defineProps(['signalId','operatedType','signalObjectType','parentLogicObjectId'])
const signal = ref({
  signalId:null,
  signalName:"",
  signalAlias:"",
  industrialInternetId:"",
  signalUnificationId:null,
  signalUnificationName:'',
  signalType:props.signalObjectType,
  parentLogicId:props.parentLogicObjectId,
  parentLogicObjectPath:"",
  positionAmongBrothers:0,
  applied:true,
  createdAt:0,
  createdBy:0,
  updatedAt:0,
  updatedBy:0
})

getAllSignalUnifications().then(x=>{
  if(x.data.isSuccess){
    signalUnifications.value = x.data.result.items
  }
})

watch(props,(newX)=>{
  if(newX.operatedType!==operatedTypeEnum.OPERATED_TYPE_CREATE){
    retrieveSignal(newX.signalId).then(x => {
      delete x.data.result['@type']
      signal.value = x.data.result
      signal.value.parentLogicObjectPath = getParentObjectPath(signal.value.signalId,false)
    })
  }else{
    signal.value = {
      signalId:-1,
      signalName:"",
      signalUnificationId:null,
      signalUnificationName:'',
      signalType:props.signalObjectType,
      parentLogicId:props.parentLogicObjectId,
      parentLogicObjectPath:"",
      positionAmongBrothers:0,
      applied:true,
      createdAt:0,
      createdBy:0,
      updatedAt:0,
      updatedBy:0
    }
    signal.value.parentLogicObjectPath = getParentObjectPath(props.parentLogicObjectId,true)
  }

},{immediate: true, deep: true})

const getSignal = () => {
  const {signalUnificationName,parentLogicObjectPath,createdAt,updatedAt,createdBy,updatedBy,...rawData} = toRaw(signal.value)
  return rawData
}

defineExpose({
  getSignal
})

const rules = {
  signalName: [{ required: true,trigger: 'blur' }]
}

</script>

<style scoped>
.el-input ,.el-select {
  width: 500px;
}
</style>
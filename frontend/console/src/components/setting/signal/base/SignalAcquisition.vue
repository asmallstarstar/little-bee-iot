<template>
  <el-form :model="signalAcquisition" :rules="rules">
    <el-form-item :label="$t('language.agent.AGENT_NAME')" label-position="left" label-width="120px" prop="agentId" required>
      <el-select v-model="signalAcquisition.agentId" class="m-2" :placeholder="$t('language.common.SELECT')" @change="changeAgent($event)">
        <el-option v-for="item in agents" :key="item.agentId" :label="item.agentName" :value="item.agentId"/>
      </el-select>
    </el-form-item>
    <el-form-item :label="$t('language.fsu.FSU_NAME')" label-position="left" label-width="120px" prop="fsuId" required>
      <el-select v-model="signalAcquisition.fsuId" class="m-2" :placeholder="$t('language.common.SELECT')" @change="changeFsu($event)">
        <el-option v-for="item in fsus.filter(x=>x.agentId==signalAcquisition.agentId)" :key="item.fsuId" :label="item.fsuName" :value="item.fsuId"/>
      </el-select>
    </el-form-item>
    <template v-if="getFsuDataParsingType(signalAcquisition.fsuId)===fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_CENTER">
      <el-form-item :label="$t('language.driver.DRIVER_NAME')" label-position="left" label-width="120px" prop="driverId" required>
        <el-select v-model="signalAcquisition.driverId" class="m-2" :placeholder="$t('language.common.SELECT')">
          <el-option v-for="item in drivers.filter(x=>x.fsuId==signalAcquisition.fsuId)" :key="item.driverId" :label="item.driverName" :value="item.driverId"/>
        </el-select>
      </el-form-item>
      <el-form-item label="X" label-position="left" label-width="120px" prop="x">
        <el-input v-model="signalAcquisition.x"/>
      </el-form-item>
      <el-form-item label="Y" label-position="left" label-width="120px" prop="y">
        <el-input v-model="signalAcquisition.y"/>
      </el-form-item>
      <el-form-item label="Z" label-position="left" label-width="120px" prop="z">
        <el-input v-model="signalAcquisition.z"/>
      </el-form-item>
    </template>
    <template v-else-if="getFsuDataParsingType(signalAcquisition.fsuId)===fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_BOTTOM">
      <el-form-item :label="$t('language.common.METADATA')" label-position="left" label-width="120px" prop="metadataId" >
        <el-select v-model="signalAcquisition.signalBindMetadataId" class="m-2" :placeholder="$t('language.common.SELECT')" >
          <el-option v-for="item in metadatas" :key="item.metadataId" :label="item.metadataAlias" :value="item.metadataId"/>
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('language.common.METADATA_INSTANCE')" label-position="left" label-width="120px" prop="configureId" >
        <MetaInstance :metadata-id="signalAcquisition.signalBindMetadataId" :configure-id="0" :read-only="false" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_CREATE" ref="metaInstanceRef"></MetaInstance>
        <MetaInstance :metadata-id="signalAcquisition.signalBindMetadataId" :configure-id="signalAcquisition.signalBindConfigureId"  :read-only="false" v-else ref="metaInstanceRef"></MetaInstance>
      </el-form-item>
    </template>
  </el-form>
</template>

<script setup>
import MetaInstance from "@/components/setting/metadata/MetaInstance.vue"

import {ref, toRaw, watch} from "vue";
import {fsuDataParsingTypeEnum, operatedTypeEnum} from "@/utils/common";
import {retrieveSignalAcquisition} from "@/api/signal-acquisition";
import {getAllAgents} from "@/api/agent";
import {getAllFsus, retrieveFsu} from "@/api/fsu";
import {getAllDrivers} from "@/api/driver";
import {getAllMetadatas} from "@/api/metadata";
import {getAllConfigures} from "@/api/configure";
import {getAllFsuTypes} from "@/api/fsu-type";

const agents = ref()
const fsus = ref([])
const drivers = ref([])
const metadatas = ref()
const configures = ref()
const fsuTypes = ref()
const metaInstanceRef = ref()
const props = defineProps(['signalId','operatedType'])
const signalAcquisition = ref({
  signalId:-1,
  agentId:null,
  fsuId:null,
  driverId:null,
  x:0,
  y:0,
  z:0,
  signalBindMetadataId:null,
  signalBindConfigureId:0,
})

Promise.all([getAllAgents(),getAllFsus(),getAllDrivers(),getAllMetadatas(),getAllConfigures(),getAllFsuTypes()]).then(function (x) {
  agents.value = x[0].data.result.items
  fsus.value = x[1].data.result.items
  drivers.value = x[2].data.result.items
  metadatas.value = x[3].data.result.items
  configures.value = x[4].data.result.items
  fsuTypes.value = x[5].data.result.items
})

watch(props,(newX)=>{
  if(newX.operatedType!==operatedTypeEnum.OPERATED_TYPE_CREATE){
    retrieveSignalAcquisition(newX.signalId).then(x=>{
      delete x.data.result['@type']
      signalAcquisition.value = x.data.result
      retrieveFsu(signalAcquisition.value.fsuId).then(k=>{
       signalAcquisition.value.agentId = k.data.result.agentId
      })
    })
  }else{
    signalAcquisition.value = {
      signalId:-1,
      agentId:null,
      fsuId:null,
      driverId:null,
      x:0,
      y:0,
      z:0,
      signalBindMetadataId:null,
      signalBindConfigureId:0,
    }
  }
},{immediate: true, deep: true})

function getFsuDataParsingType(fsuId) {
  if(fsus.value.length>0 && fsuTypes.value.length>0){
    const l = fsus.value.filter(x=>x.fsuId===fsuId)
    if(l.length>0){
      const m = fsuTypes.value.filter(x=>x.fsuTypeId===l[0].fsuTypeId)
      if(m.length>0){
        return m[0].dataAnalysisMode
      }
    }
  }
  return fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_UNKNOWN
}

function changeAgent(value){
  signalAcquisition.value.fsuId = null
  signalAcquisition.value.driverId = null
}

function changeFsu(value){
  signalAcquisition.value.driverId = null
  const l = fsus.value.filter(x=>x.fsuId==value)
  if(l.length==1){
    signalAcquisition.value.signalBindMetadataId = l[0].metadataId
  }
}

const getSignalAcquisition = () => {
  const {...rawData} = toRaw(signalAcquisition.value)
  return rawData
}

const getMetadataAttribute = () =>{
  if(signalAcquisition.value.signalBindMetadataId && metaInstanceRef.value){
    const metadataAttributeJson = JSON.stringify(metaInstanceRef.value.getMetadataAttribute())
    return {
      configureId:signalAcquisition.value.signalBindConfigureId,
      metadataId:signalAcquisition.value.signalBindMetadataId,
      configureJson:metadataAttributeJson
    }
  }
  return null
}

defineExpose({
  getSignalAcquisition,getMetadataAttribute
})

const rules = {
  agentId: [{ required: true,trigger: 'blur' }],
  fsuId: [{ required: true,trigger: 'blur' }]
}

</script>

<style scoped>
.el-input ,.el-select {
  width: 500px;
}
</style>
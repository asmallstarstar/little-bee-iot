<template>
  <el-form :model="signalVideo" :rules="rules">
    <el-form-item :label="$t('language.signal.VIDEO_PLAY_TYPE')" label-position="left" label-width="120px" prop="videoPlayType" required>
      <el-select v-model="signalVideo.videoPlayType" class="m-2" :placeholder="$t('language.common.SELECT')">
        <el-option v-for="item in videoPlayTypeText" :key="item.index" :label="item.name" :value="item.index"/>
      </el-select>
    </el-form-item>
    <el-form-item :label="$t('language.common.METADATA')" label-position="left" label-width="120px" prop="metadataId" >
      <el-select v-model="signalVideo.videoBindMetadataId" class="m-2" :placeholder="$t('language.common.SELECT')" >
        <el-option v-for="item in metadatas" :key="item.metadataId" :label="item.metadataAlias" :value="item.metadataId"/>
      </el-select>
    </el-form-item>
    <el-form-item :label="$t('language.common.METADATA_INSTANCE')" label-position="left" label-width="120px" prop="configureId" >
      <MetaInstance :metadata-id="signalVideo.videoBindMetadataId" :configure-id="0" :read-only="false" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_CREATE" ref="metaInstanceRef"></MetaInstance>
      <MetaInstance :metadata-id="signalVideo.videoBindMetadataId" :configure-id="signalVideo.videoBindConfigureId"  :read-only="false" v-else ref="metaInstanceRef"></MetaInstance>
    </el-form-item>
  </el-form>
</template>

<script setup>
import MetaInstance from "@/components/setting/metadata/MetaInstance.vue"

import {ref, toRaw, watch} from "vue";
import {operatedTypeEnum, videoPlayTypeText} from "@/utils/common";
import {getAllMetadatas} from "@/api/metadata";
import {getAllConfigures} from "@/api/configure";
import {retrieveSignalVideo} from "@/api/signal-video";

const metadatas = ref()
const configures = ref()
const metaInstanceRef = ref()
const props = defineProps(['signalId','operatedType'])
const signalVideo = ref({
  signalId:-1,
  videoPlayType:null,
  videoBindMetadataId:null,
  videoBindConfigureId:0,
})

Promise.all([getAllMetadatas(),getAllConfigures()]).then(function (x) {
  metadatas.value = x[0].data.result.items
  configures.value = x[1].data.result.items
})

watch(props,(newX)=>{
  if(newX.operatedType!==operatedTypeEnum.OPERATED_TYPE_CREATE){
    retrieveSignalVideo(newX.signalId).then(x=>{
      delete x.data.result['@type']
      signalVideo.value = x.data.result
    })
  }else{
    signalVideo.value = {
      signalId:-1,
      videoPlayType:null,
      videoBindMetadataId:null,
      videoBindConfigureId:0,
    }
  }
},{immediate: true, deep: true})

const getVideoAcquisition = () => {
  const {...rawData} = toRaw(signalVideo.value)
  return rawData
}

const getMetadataAttribute = () =>{
  if(signalVideo.value.videoBindMetadataId && metaInstanceRef.value){
    const metadataAttributeJson = JSON.stringify(metaInstanceRef.value.getMetadataAttribute())
    return {
      configureId:signalVideo.value.videoBindConfigureId,
      metadataId:signalVideo.value.videoBindMetadataId,
      configureJson:metadataAttributeJson
    }
  }
  return null
}

defineExpose({
  getVideoAcquisition,getMetadataAttribute
})

const rules = {
  videoPlayType: [{ required: true,trigger: 'blur' }]
}

</script>

<style scoped>
.el-input ,.el-select {
  width: 500px;
}
</style>
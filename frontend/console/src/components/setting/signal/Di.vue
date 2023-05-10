<template>
  <el-scrollbar :max-height="maxHeight">
    <el-collapse v-model="activeNames" >
      <el-collapse-item :title="$t('language.signal.SIGNAL_BASE_PROPERTY')" name="signalProperty">
        <signal :signal-id="signalId" :operated-type="operatedType" :signal-object-type="signalObjectType" :parent-logic-object-id="parentLogicObjectId" ref="signalRef"></signal>
      </el-collapse-item>
      <el-collapse-item :title="$t('language.signal.SIGNAL_ACQUISITION_PROPERTY')" name="acquisitionProperty">
        <signal-acquisition :signal-id="signalId" :operated-type="operatedType" ref="signalAcquisitionRef"></signal-acquisition>
      </el-collapse-item>
      <el-collapse-item :title="$t('language.signal.SIGNAL_DI_PROPERTY')" name="diProperty">
        <signal-di :signal-id="signalId" :operated-type="operatedType" ref="signalDiRef"></signal-di>
      </el-collapse-item>
      <el-collapse-item :title="$t('language.signal.SIGNAL_THRESHOLD_PROPERTY')" name="thresholdProperty">
        <signal-threshhold :signal-id="signalId" :operated-type="operatedType" :signal-object-type="signalObjectType" ref="signalThresholdRef"></signal-threshhold>
      </el-collapse-item>
    </el-collapse>
    <div class="operate-group">
      <el-button type="danger" size="small" @click="handleDelete" v-if="operatedType!=operatedTypeEnum.OPERATED_TYPE_CREATE && actionsStore.isPermission(['delete-signal-di'])">{{$t('language.common.DELETE')}}</el-button>
      <el-button type="primary" size="small" @click="handleConfirm" v-if="actionsStore.isPermission(['create-signal-di','update-signal-di'])">{{$t('language.common.CONFIRM')}}</el-button>
    </div>
  </el-scrollbar>
</template>

<script setup>
import {ref, watch} from "vue";
import {operatedTypeEnum, signalTypeEnum} from "@/utils/common";
import {useRoute} from "vue-router";
import Signal from "@/components/setting/signal/base/Signal.vue";
import SignalAcquisition from "@/components/setting/signal/base/SignalAcquisition.vue";
import SignalDi from "@/components/setting/signal/base/SignalDi.vue";
import SignalThreshhold from "@/components/setting/signal/base/SignalThreshhold.vue";
import {useActionsStore} from "@/stores/action";
import {ElMessage, ElMessageBox} from "element-plus";
import {useTabsStore} from "@/stores/tab";
import router from "@/router";
import {useSignalsStore} from "@/stores/signal";
import i18n from "@/lang";
import {createDi, deleteDi, updateDi} from "@/api/di";

const t = i18n.global.t
const route = useRoute()
const actionsStore = useActionsStore()
const maxHeight = ref(window.innerHeight-120)
const activeNames = ref("signalProperty")
const signalId = ref(-1)
const operatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
const signalObjectType = ref(signalTypeEnum.SIGNAL_TYPE_UNKNOWN)
const parentLogicObjectId = ref(-1)
const signalRef = ref()
const signalAcquisitionRef = ref()
const signalDiRef = ref()
const signalThresholdRef = ref()
const signalsStore = useSignalsStore()

async function onRouteUpdate(params) {
  signalId.value = Number(params.signalId)
  operatedType.value = Number(params.operatedType)
  signalObjectType.value = Number(params.objectType)
  parentLogicObjectId.value = Number(params.parentLogicObjectId)
}

watch(() => route.params, (newValue, oldValue) => {
  if(newValue.signalId){
    onRouteUpdate(newValue)
  }
}, {immediate: true, deep: true})

function handleDelete() {
  const signal = signalRef.value.getSignal()
  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.common.SIGNAL')+": "+signal.signalName, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteDi(signalId.value).then(x=>{
          if(x.data.isSuccess){
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            signalsStore.signals = signalsStore.signals.filter(x=>x.signalId!=signalId.value)
            const tabsStore = useTabsStore()
            tabsStore.removeTab(route.path)
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

function handleConfirm() {
  let obj = {}
  obj.signal = signalRef.value.getSignal()
  obj.signalAcquisition = signalAcquisitionRef.value.getSignalAcquisition()
  obj.signalDi = signalDiRef.value.getSignalDi()
  obj.signalThreshold = signalThresholdRef.value.getSignalThreshold()
  const configure = signalAcquisitionRef.value.getMetadataAttribute()
  if(configure){
    obj.signalBindConfigure = configure
  }
  let fn = createDi
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateDi
  }
  //console.log(JSON.stringify(obj))
  fn(obj).then(x=>{
    if(x.data.isSuccess){
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
        for(let i in signalsStore.signals){
          if(signalsStore.signals[i].signalId==obj.signal.signalId){
            signalsStore.signals[i] = obj.signal
            break
          }
        }
      }else if(operatedType.value==operatedTypeEnum.OPERATED_TYPE_CREATE){
        const res = x.data.result
        delete res['@type']
        signalsStore.signals.push(res.signal)

        const tabsStore = useTabsStore()
        tabsStore.removeTab(route.path)
        router.replace({ name: 'signal',
          params: {
            signalId: x.data.result.signal.signalId,
            operatedType:operatedTypeEnum.OPERATED_TYPE_EDIT,
            objectType:x.data.result.signal.signalType,
            parentLogicObjectId:x.data.result.signal.parentLogicId,
          }
        })
      }
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  }).catch(()=>{
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  })
}

</script>

<style scoped>
.operate-group{
  margin: 10px 10px;
}
</style>
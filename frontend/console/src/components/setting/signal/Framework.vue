<template>
  <div>
    <component :is="currentComponent"></component>
  </div>
</template>

<script setup>
import Ai from "@/components/setting/signal/Ai.vue"
import Di from "@/components/setting/signal/Di.vue"
import Si from "@/components/setting/signal/Si.vue"
import VAi from "@/components/setting/signal/VAi.vue"
import VDi from "@/components/setting/signal/VDi.vue"
import VSi from "@/components/setting/signal/VSi.vue"
import Output from "@/components/setting/signal/Output.vue"
import Video from "@/components/setting/signal/Video.vue"
import {ref, shallowRef, watch} from "vue";
import {signalTypeEnum} from "@/utils/common";
import {useRoute, useRouter} from "vue-router";

const route = useRoute()
const router = useRouter()
const currentComponent = shallowRef(null)
const signal = ref(null)

watch(() => route.params, (newValue, oldValue) => {
  if(newValue.signalId){
    onRouteUpdate(newValue)
  }
}, {immediate: true, deep: true})

function onRouteUpdate(params) {
  currentComponent.value = getCurrentComponentBySignalType(Number(params.objectType))
}

function getCurrentComponentBySignalType(signalType) {
  switch (signalType) {
    case signalTypeEnum.SIGNAL_TYPE_AI:
      return Ai
    case signalTypeEnum.SIGNAL_TYPE_DI:
      return Di
    case signalTypeEnum.SIGNAL_TYPE_SI:
      return Si
    case signalTypeEnum.SIGNAL_TYPE_VAI:
      return VAi
    case signalTypeEnum.SIGNAL_TYPE_VDI:
      return VDi
    case signalTypeEnum.SIGNAL_TYPE_VSI:
      return VSi
    case signalTypeEnum.SIGNAL_TYPE_AO:
    case signalTypeEnum.SIGNAL_TYPE_DO:
    case signalTypeEnum.SIGNAL_TYPE_SO:
      return Output
    case signalTypeEnum.SIGNAL_TYPE_VIDEO:
      return Video
  }
  return null
}

</script>

<style scoped>

</style>
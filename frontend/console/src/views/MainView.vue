<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <Nav></Nav>
      </el-header>
      <el-container>
        <el-aside :width="SIDE_BAR_WIDTH">
          <router-view name="sidebar" />
        </el-aside>
        <el-main>
          <div class="content">
            <TabButtons class="tab-button"></TabButtons>
            <router-view v-slot="{ Component }">
              <transition name="move" mode="out-in">
                <keep-alive>
                  <component :is="Component"></component>
                </keep-alive>
              </transition>
            </router-view>
          </div>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import Nav from '@/components/layout/Nav.vue'
import FsuTree from '@/components/layout/side/FsuTree.vue'
import SignalTree from '@/components/layout/side/SignalTree.vue'
import Setting from '@/components/layout/side/Setting.vue'
import AlarmQuery from '@/components/query/AlarmQuery.vue'
import CommandQuery from '@/components/query/CommandQuery.vue'
import HistoryMonitorQuery from '@/components/query/HistoryMonitorQuery.vue'
import LogQuery from '@/components/query/LogQuery.vue'
import RealData from '@/components/real/RealData.vue'
import RealAlarm from '@/components/real/RealAlarm.vue'
import ActionConfig from '@/components/setting/ActionConfig.vue'
import ActionGroup from "@/components/setting/ActionGroup.vue";
import ActionTree from "@/components/setting/ActionTree.vue";
import AlarmLevel from '@/components/setting/AlarmLevel.vue'
import AreaType from "@/components/setting/AreaType.vue";
import DepartmentConfig from '@/components/setting/DepartmentConfig.vue'
import DeviceType from "@/components/setting/DeviceType.vue";
import MenuAction from "@/components/setting/MenuAction.vue";
import MenuConfig from '@/components/setting/MenuConfig.vue'
import MenuTree from "@/components/setting/MenuTree.vue";
import MetadataConfig from "@/components/setting/MetadataConfig.vue"
import RoleConfig from '@/components/setting/RoleConfig.vue'
import SignalUnification from "@/components/setting/SignalUnification.vue";
import AgentConfig from "@/components/system/AgentConfig.vue"
import ConfigurationSetting from "@/components/system/ConfigurationSetting.vue"
import MonitorObject from "@/components/system/MonitorObject.vue"
import SystemSetting from "@/components/system/SystemSetting.vue"
import UserConfig from "@/components/user/UserConfig.vue"
import Home from "@/components/Home.vue"
import Agent from "@/components/setting/agent/Agent.vue";
import Driver from "@/components/setting/agent/Driver.vue";
import Fsu from "@/components/setting/agent/Fsu.vue";
import Framework from "@/components/setting/signal/Framework.vue"
import LogicObject from "@/components/setting/logic/LogicObject.vue"
import { ref } from 'vue'
import {useTabsStore} from "@/stores/tab";
import {onBeforeRouteUpdate} from "vue-router";
import TabButtons from "@/components/layout/TabButtons.vue";
import {useAgentsStore} from "@/stores/agent";
import {useSignalsStore} from "@/stores/signal";
import i18n from '@/lang/index'
import {
  operatedTypeEnum, signalTreeObjectTypeEnum, signalTypeEnum
} from "@/utils/common";
const t = i18n.global.t
const SIDE_BAR_WIDTH = ref("200")
const tabsStore = useTabsStore()

onBeforeRouteUpdate(to => {
  let title = to.meta.title
  if(to.name=='agent'){
    const agentsStore = useAgentsStore()
    const res = agentsStore.findAgentById(to.params.agentId)
    if(res!=null){
      title = res.agentName
    }else{
      title = ""
    }
  }
  if(to.name=='fsu'){
    const agentsStore = useAgentsStore()
    const res = agentsStore.findFsuById(to.params.fsuId)
    if(res!=null){
      title = res.fsuName
    }else{
      title = ""
    }
  }
  if(to.name=='logicObject'){
    const signalsStore = useSignalsStore()
    const operate = Number(to.params.operatedType)
    const logicType = Number(to.params.objectType)
    if(operate === operatedTypeEnum.OPERATED_TYPE_CREATE){
      if(logicType===signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_AREA){
        title = t("language.newInstance.NEW_AREA")
      }
      if(logicType===signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_DEVICE){
        title = t("language.newInstance.NEW_DEVICE")
      }
    }else{
      const logicObjectId = Number(to.params.logicObjectId)
      const res = signalsStore.findLogicObjectById(logicObjectId)
      if(res!=null){
        title = res.logicObjectName
      }else if(logicObjectId==0){
        title = i18n.global.t('language.common.SURVEILLANCE_CENTER_SIGNAL')
      }else{
        title = "not found title"
      }
    }
  }
  if(to.name=='signal'){
    const signalsStore = useSignalsStore()
    const operate = Number(to.params.operatedType)
    const signalType = Number(to.params.objectType)
    if(operate === operatedTypeEnum.OPERATED_TYPE_CREATE){
      title = getTitleBySignalType(signalType)
    }else{
      const signalId = Number(to.params.signalId)
      const res = signalsStore.findSignalById(signalId)
      if(res!=null){
        title = res.signalName
      }else{
        title = "not found title"
      }
    }
  }
  if(to.name=='configuration' || to.name=='realdata'){
    let prefix = ""
    if(to.name=='configuration'){
      prefix = i18n.global.t('language.common.CONFIGURATION')
    }
    if(to.name=='realdata'){
      prefix = i18n.global.t('language.common.REALDATA')
    }
    if(to.params.nodeType == signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_CENTER ||
        to.params.nodeType == signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_AREA||
        to.params.nodeType == signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_DEVICE){
      const signalsStore = useSignalsStore()
      const logicObjectId = Number(to.params.nodeId)
      const res = signalsStore.findLogicObjectById(logicObjectId)
      if(res!=null){
        title = prefix+":"+res.logicObjectName
      }else if(logicObjectId==0){
        title = prefix+":"+i18n.global.t('language.common.SURVEILLANCE_CENTER_SIGNAL')
      }else{
        title = "not found title"
      }
    }else if(to.params.nodeType == signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_SIGNAL){
      const signalsStore = useSignalsStore()
      const signalId = Number(to.params.nodeId)
      const res = signalsStore.findSignalById(signalId)
      if(res!=null){
        title = prefix+":"+res.signalName
      }else{
        title = "not found title"
      }
    }
  }
  tabsStore.addTab(to.name,title,to.path)
});

function getTitleBySignalType(signalType){
  if(signalType===signalTypeEnum.SIGNAL_TYPE_AI){
    return t("language.newInstance.NEW_AI")
  }
  if(signalType===signalTypeEnum.SIGNAL_TYPE_DI){
    return t("language.newInstance.NEW_DI")
  }
  if(signalType===signalTypeEnum.SIGNAL_TYPE_SI){
    return t("language.newInstance.NEW_SI")
  }
  if(signalType===signalTypeEnum.SIGNAL_TYPE_VAI){
    return t("language.newInstance.NEW_VAI")
  }
  if(signalType===signalTypeEnum.SIGNAL_TYPE_VDI){
    return t("language.newInstance.NEW_VDI")
  }
  if(signalType===signalTypeEnum.SIGNAL_TYPE_VSI){
    return t("language.newInstance.NEW_VSI")
  }
  if(signalType===signalTypeEnum.SIGNAL_TYPE_AO){
    return t("language.newInstance.NEW_AO")
  }
  if(signalType===signalTypeEnum.SIGNAL_TYPE_DO){
    return t("language.newInstance.NEW_DO")
  }
  if(signalType===signalTypeEnum.SIGNAL_TYPE_SO){
    return t("language.newInstance.NEW_SO")
  }
  if(signalType===signalTypeEnum.SIGNAL_TYPE_VIDEO){
    return t("language.newInstance.NEW_VIDEO")
  }
  return "unknown title"
}

</script>

<style scoped>

.el-main {
  --el-main-padding:0
}

.content {
  width: auto;
  /*height: 100%;*/
  padding: 5px;
  /*overflow-y: scroll;*/
  box-sizing: border-box;
}

.tab-button{
  margin-bottom: 5px;
}

</style>

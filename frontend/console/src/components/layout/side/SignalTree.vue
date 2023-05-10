<template>
  <el-scrollbar :max-height="maxHeight">
    <el-input v-model="filterText" :placeholder="$t('language.common.INPUT')" class="filter-input"/>
    <el-tree ref="treeRef" :data="signalTree" node-key="id" @node-click="handleNodeClick"
             default-expand-all :filter-node-method="filterNode" :expand-on-click-node="false" >
      <template #default="{ node, data }">
            <span class="custom-tree-node">
              <el-icon size="16" v-if="data.type === signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_CENTER" class="icon"><Warning/></el-icon>
              <el-icon size="16" v-if="data.type === signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_AREA" class="icon"><OfficeBuilding/></el-icon>
              <el-icon size="16" v-if="data.type === signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_DEVICE" class="icon"><Discount/></el-icon>
              <el-icon size="16" v-if="data.signalType === signalTypeEnum.SIGNAL_TYPE_AI" class="icon"><Drizzling/></el-icon>
              <el-icon size="16" v-if="data.signalType === signalTypeEnum.SIGNAL_TYPE_DI" class="icon"><Open/></el-icon>
              <el-icon size="16" v-if="data.signalType === signalTypeEnum.SIGNAL_TYPE_SI" class="icon"><Document/></el-icon>
              <el-icon size="16" v-if="data.signalType === signalTypeEnum.SIGNAL_TYPE_AO" class="icon"><Notebook/></el-icon>
              <el-icon size="16" v-if="data.signalType === signalTypeEnum.SIGNAL_TYPE_DO" class="icon"><Tickets/></el-icon>
              <el-icon size="16" v-if="data.signalType === signalTypeEnum.SIGNAL_TYPE_SO" class="icon"><Memo/></el-icon>
              <el-icon size="16" v-if="data.signalType === signalTypeEnum.SIGNAL_TYPE_VAI" class="icon"><Collection/></el-icon>
              <el-icon size="16" v-if="data.signalType === signalTypeEnum.SIGNAL_TYPE_VDI" class="icon"><ScaleToOriginal/></el-icon>
              <el-icon size="16" v-if="data.signalType === signalTypeEnum.SIGNAL_TYPE_VSI" class="icon"><CopyDocument/></el-icon>
              <el-icon size="16" v-if="data.signalType === signalTypeEnum.SIGNAL_TYPE_VIDEO" class="icon"><VideoCamera/></el-icon>
              <span  :class="[node.childNodes.length ? 'bold' : '', node.isCurrent ? 'orange' : '']">{{ node.label }}</span>
            </span>
      </template>
    </el-tree>
  </el-scrollbar>
</template>

<script setup>
import {computed, ref, watch} from "vue";
import {useSignalsStore} from "@/stores/signal";
import router from "@/router";
import {operatedTypeEnum, signalTreeObjectTypeEnum, signalTypeEnum} from "@/utils/common";
import {getAllLogicObjects} from "@/api/logic-object";
import {getAllSignals} from "@/api/signal";
import {useRoute} from "vue-router";

const treeRef = ref()
const filterText = ref('')
const signalsStore = useSignalsStore()
let maxHeight = ref(window.innerHeight - 60)
const route = useRoute()
let operator = ""

watch(() => route.path, (newPath, oldPath) => {
      if(newPath.indexOf("monitor-object")>0 ||
          newPath.indexOf("logic")>0 ||
          newPath.indexOf("signal")>0){
        operator = "CONFIG"
      }
      if(newPath.indexOf("configuration-setting")>0){
        operator = "CONFIGURATION"
      }
      if(newPath.indexOf("real-data")>0){
        operator = "REALDATA"
      }
    },
    { immediate: true });

Promise.all([getAllLogicObjects(),getAllSignals()]).then(function (x) {
  signalsStore.setLogicObjects(x[0].data.result.items)
  signalsStore.setSignals(x[1].data.result.items)
})

const signalTree = computed(() => {
  let signalTreeData = []
  signalsStore.resetSurveillanceCenter()
  signalTreeData.push(signalsStore.surveillanceCenter)
  const sorts = signalsStore.logicObjects.sort((a,b)=>{return a.logicObjectId-b.logicObjectId})
  for(let i in sorts){
    const node = createLogicObjectNode(sorts[i])
    const parentNode = findParentNode(signalTreeData,sorts[i].parentLogicObjectId)
    if(parentNode!=null){
      parentNode.children.push(node)
      parentNode.children = parentNode.children.sort((a,b)=>{return a.data.positionAmongBrothers-b.data.positionAmongBrothers})
    }
  }
  for(let i in signalsStore.signals){
    const node = createSignalNode(signalsStore.signals[i])
    const parentNode = findParentNode(signalTreeData,signalsStore.signals[i].parentLogicId)
    if(parentNode!=null){
      parentNode.children.push(node)
      parentNode.children = parentNode.children.sort((a,b)=>{return a.data.positionAmongBrothers-b.data.positionAmongBrothers})
    }
  }
  return signalTreeData
})

function createLogicObjectNode(logicObject) {
  let node = {}
  node.id = logicObject.logicObjectId
  node.label = logicObject.logicObjectName
  node.type = logicObject.logicObjectTypeId
  node.signalType = null
  node.children = []
  node.data = logicObject
  return node
}

function createSignalNode(signal) {
  let node = {}
  node.id = signal.signalId
  node.label = signal.signalName
  node.type = signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_SIGNAL
  node.signalType = signal.signalType
  node.children = []
  node.data = signal
  return node
}

function findParentNode(treeData,id) {
  for(let i in treeData){
    if(treeData[i].id==id && treeData[i].signalType==null){
      return treeData[i]
    }else if(treeData[i].children.length>0){
      const res = findParentNode(treeData[i].children,id)
      if(res!=null)
        return res
    }
  }
  return null
}

function handleNodeClick(nodeData) {
  if(nodeData.type==signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_CENTER ||
      nodeData.type==signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_AREA  ||
      nodeData.type==signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_DEVICE){
    if(operator=="CONFIG"){
      router.push({ name: 'logicObject',
        params: {
          logicObjectId: nodeData.id,
          operatedType:operatedTypeEnum.OPERATED_TYPE_EDIT,
          objectType:nodeData.type
        }
      })
    }else if(operator=="CONFIGURATION"){
      router.push("/configuration-setting/"+nodeData.type+"/"+nodeData.id)
    }else if(operator=="REALDATA"){
      router.push("/real-data/"+nodeData.type+"/"+nodeData.id)
    }
  }
  if(nodeData.type==signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_SIGNAL){
    if(operator=="CONFIG"){
      let pid = null
      const parentLogicObject = signalsStore.findParentLogicObjectBySignalId(nodeData.id)
      if(parentLogicObject){
        pid = parentLogicObject.logicObjectId
      }
      router.push({ name: 'signal',
        params: {
          signalId: nodeData.id,
          operatedType:operatedTypeEnum.OPERATED_TYPE_EDIT,
          objectType: nodeData.signalType,
          parentLogicObjectId:pid
        }
      })
    }else if(operator=="CONFIGURATION"){
      router.push("/configuration-setting/"+nodeData.type+"/"+nodeData.id)
    }else if(operator=="REALDATA"){
      router.push("/real-data/"+nodeData.type+"/"+nodeData.id)
    }
  }
}

watch(filterText, (val) => {
  treeRef.value.filter(val)
})

const filterNode = (value, data) => {
  if (!value) return true
  return data.label.includes(value)
}

</script>

<style lang="scss" scoped>

.filter-input{
  padding: 5px 5px;
}

.custom-tree-node {
  display: flex;
  align-items: center;
  .icon {
    margin-right: 10px;
    flex: 0 0 auto;
  }
  .bold {
    //font-weight: bold;
  }
  .orange {
    color: #f1a428;
  }
}


</style>
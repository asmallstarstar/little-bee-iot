<template>
  <el-scrollbar :max-height="maxHeight">
    <el-input v-model="filterText" :placeholder="$t('language.common.INPUT')" class="filter-input"/>
    <el-tree ref="treeRef" :data="agentTree" node-key="id" @node-click="handleNodeClick"
             default-expand-all :filter-node-method="filterNode" :expand-on-click-node="false" >
      <template #default="{ node, data }">
            <span class="custom-tree-node">
              <el-icon size="16" v-if="node.level === 1" class="icon"><Warning /></el-icon>
              <el-icon size="16" v-if="node.level === 2" class="icon"><Service /></el-icon>
              <el-icon size="16" v-if="node.level === 3" class="icon"><Link /></el-icon>
              <el-icon size="16" v-if="node.level === 4" class="icon"><Switch /></el-icon>
              <span  :class="[node.childNodes.length ? 'bold' : '', node.isCurrent ? 'orange' : '']">{{ node.label }}</span>
            </span>
      </template>
    </el-tree>
  </el-scrollbar>
</template>

<script setup>
import {useAgentsStore} from "@/stores/agent";
import {getAllAgents} from "@/api/agent";
import {getAllFsus} from "@/api/fsu";
import {getAllDrivers} from "@/api/driver";
import {computed, ref, watch} from "vue";
import {agentObjectTypeEnum} from "@/utils/agent";
import router from "@/router";
const treeRef = ref()
const filterText = ref('')
const agentsStore = useAgentsStore()
let maxHeight = ref(window.innerHeight - 60)

Promise.all([getAllAgents(),getAllFsus(),getAllDrivers()]).then(function (x) {
  agentsStore.setAgents(x[0].data.result.items)
  agentsStore.setFsus(x[1].data.result.items)
  agentsStore.setDrivers(x[2].data.result.items)
})

function handleNodeClick(nodeData) {
  if(nodeData.type==agentObjectTypeEnum.AGENT_OBJECT_TYPE_SURVEILLANCE_CENTER){
    router.push({ name: 'center' })
  }
  if(nodeData.type==agentObjectTypeEnum.AGENT_OBJECT_TYPE_AGENT){
    router.push({ name: 'agent', params: { agentId: nodeData.id } })
  }
  if(nodeData.type==agentObjectTypeEnum.AGENT_OBJECT_TYPE_FSU){
    router.push({ name: 'fsu', params: { fsuId: nodeData.id } })
  }
}

watch(filterText, (val) => {
  treeRef.value.filter(val)
})

const filterNode = (value, data) => {
  if (!value) return true
  return data.label.includes(value)
}

const agentTree = computed(() => {
  let agentTreeData = []
  agentsStore.resetSurveillanceCenter()
  agentTreeData.push(agentsStore.surveillanceCenter)

  for(let i in agentsStore.agents){
    agentsStore.surveillanceCenter.children.push(createAgentNode(agentsStore.agents[i]))
  }
  for(let i in agentsStore.fsus){
    const node = createFSUNode(agentsStore.fsus[i])
    const parentNode = findNode(agentTreeData,agentsStore.fsus[i].agentId,agentObjectTypeEnum.AGENT_OBJECT_TYPE_AGENT)
    if(parentNode!=null){
      parentNode.children.push(node)
    }
  }
  for(let i in agentsStore.drivers){
    const node = createDriverNode(agentsStore.drivers[i])
    const parentNode = findNode(agentTreeData,agentsStore.drivers[i].fsuId,agentObjectTypeEnum.AGENT_OBJECT_TYPE_FSU)
    if(parentNode!=null){
      parentNode.children.push(node)
    }
  }
  return agentTreeData
})

function createAgentNode(agent) {
  let node = {}
  node.id = agent.agentId
  node.label = agent.agentName
  node.type = agentObjectTypeEnum.AGENT_OBJECT_TYPE_AGENT
  node.children = []
  node.data = agent
  return node
}

function createFSUNode(fsu) {
  let node = {}
  node.id = fsu.fsuId
  node.label = fsu.fsuName
  node.type = agentObjectTypeEnum.AGENT_OBJECT_TYPE_FSU
  node.children = []
  node.data = fsu
  return node
}

function createDriverNode(driver) {
  let node = {}
  node.id = driver.driverId
  node.label = driver.driverName
  node.type = agentObjectTypeEnum.AGENT_OBJECT_TYPE_DRIVER
  node.children = []
  node.data = driver
  return node
}

function findNode(agentTreeData,id,agentObjectType) {
  for(let i=0;i<agentTreeData.length;i++){
    if(agentTreeData[i].id==id && agentTreeData[i].type==agentObjectType){
      return agentTreeData[i]
    }else if(agentTreeData[i].children.length>0){
      const res = findNode(agentTreeData[i].children,id,agentObjectType)
      if(res!=null)
        return res
    }
  }
  return null
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
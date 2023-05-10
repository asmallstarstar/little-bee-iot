<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-agent'])">{{$t('language.common.CREATE')}}</el-button>
    <div >
      <el-table :data="agentsStore.agents" :max-height="maxHeight">
        <el-table-column prop="agentId" :label="$t('language.agent.AGENT_ID')" width="80" />
        <el-table-column prop="agentName" :label="$t('language.agent.AGENT_NAME')" width="150" />
        <el-table-column prop="agentNote" :label="$t('language.agent.AGENT_NOTE')" width="150" />
        <el-table-column prop="metadataId" :label="$t('language.common.METADATA')" width="150" >
          <template #default="scope">
            {{getMetadataNameById(scope.row.metadataId)}}
          </template>
        </el-table-column>
        <el-table-column prop="configureId" :label="$t('language.common.METADATA_INSTANCE')" width="200" >
          <template #default="scope">
            <MetaInstance :metadata-id="scope.row.metadataId" :configure-id="scope.row.configureId" :read-only="true" :key="keyNum"></MetaInstance>
          </template>
        </el-table-column>
        <el-table-column :label="$t('language.common.TIME')" width="60">
          <template #default="scope">
            <el-popover effect="light" trigger="hover" placement="top" width="auto">
              <template #default>
                <div>{{$t('language.common.CREATE_TIME')}}:
                  {{ scope.row.createdAt?(new Date(scope.row.createdAt)).toLocaleString():""}}</div>
                <div>{{$t('language.common.UPDATE_TIME')}}:
                  {{ scope.row.updatedAt?(new Date(scope.row.updatedAt)).toLocaleString():"" }}</div>
                <div>{{$t('language.common.CREATOR')}}:
                  {{ scope.row.createdBy!=0?getUserNickByUserId(scope.row.createdBy):"" }}</div>
                <div>{{$t('language.common.MODIFIER')}}:
                  {{ scope.row.updatedBy!=0?getUserNickByUserId(scope.row.updatedBy):"" }}</div>
              </template>
              <template #reference>
                <el-tag>...</el-tag>
              </template>
            </el-popover>
          </template>
        </el-table-column>
        <el-table-column :label="$t('language.common.OPERATE')">
          <template #default="scope">
            <el-button size="small" @click="handleView(scope.$index, scope.row)">{{$t('language.common.VIEW')}}</el-button>
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-agent','update-agent'])">{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-agent'])">{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="rowData" :rules="rules" ref="dialogForm">
        <el-form-item :label="$t('language.agent.AGENT_ID')" label-position="left" label-width="120px" prop="agentId" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_VIEW"  >
          <el-input v-model="rowData.agentId" readonly />
        </el-form-item>
        <el-form-item :label="$t('language.agent.AGENT_NAME')" label-position="left" label-width="120px" prop="agentName" required>
          <el-input v-model="rowData.agentName"/>
        </el-form-item>
        <el-form-item :label="$t('language.agent.AGENT_NOTE')" label-position="left" label-width="120px" prop="agentNote" >
          <el-input v-model="rowData.agentNote"/>
        </el-form-item>
        <el-form-item :label="$t('language.common.METADATA')" label-position="left" label-width="120px" prop="metadataId" >
          <el-select v-model="rowData.metadataId" class="m-2" :placeholder="$t('language.common.SELECT')" >
            <el-option v-for="item in metadatas" :key="item.metadataId" :label="item.metadataAlias" :value="item.metadataId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.common.METADATA_INSTANCE')" label-position="left" label-width="120px" prop="configureId" >
          <MetaInstance :metadata-id="rowData.metadataId" :configure-id="0" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_CREATE" ref="metaInstanceRef" :read-only="false"></MetaInstance>
          <MetaInstance :metadata-id="rowData.metadataId" :configure-id="rowData.configureId" v-else ref="metaInstanceRef" :read-only="false"></MetaInstance>
        </el-form-item>
      </el-form>
      <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">{{$t('language.common.CANCEL')}}</el-button>
            <el-button type="primary" @click="handleConfirm" v-if="operatedType!=operatedTypeEnum.OPERATED_TYPE_VIEW">{{$t('language.common.CONFIRM')}}</el-button>
          </span>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>
import {nextTick, ref, toRaw} from "vue";
import {ElMessage, ElMessageBox} from "element-plus";
import {operatedTypeEnum} from "@/utils/common";
import i18n from "@/lang"
import {getAllUserAlias} from "@/api/user";
import {useActionsStore} from "@/stores/action";
import {useAgentsStore} from "@/stores/agent";
import {getAllMetadatas} from "@/api/metadata";
import {createConfigure, deleteConfigure, getAllConfigures, retrieveConfigure, updateConfigure} from "@/api/configure";
import {createAgent, deleteAgent, updateAgent} from "@/api/agent";
import {initMetadataAttributeValue, metadataAttributeTypeText} from "@/utils/metadata"
import MetaInstance from "@/components/setting/metadata/MetaInstance.vue";

const t = i18n.global.t
let keyNum = ref(0)
let dialogTitle = ref(t('language.common.CREATE'))
let dialogVisible = ref(false)
let operatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
let maxHeight = ref(window.innerHeight - 165)
const dialogForm = ref()
const users = ref()
const metadatas = ref()
const configures = ref()
const metaInstanceRef = ref()
const actionsStore = useActionsStore()
const agentsStore = useAgentsStore()

Promise.all([getAllUserAlias(),getAllMetadatas(),getAllConfigures()]).then(function (x) {
  users.value = x[0].data.result.items
  metadatas.value = x[1].data.result.items
  configures.value = x[2].data.result.items
})

let rowData = ref({
  agentId:null,
  agentName:"",
  agentNote:"",
  metadataId:null,
  configureId:0,
  createdAt:0,
  updatedAt:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  rowData.value = {
    agentId:null,
    agentName:"",
    agentNote:"",
    metadataId:null,
    configureId:0,
    createdAt:0,
    updatedAt:0,
  }
  dialogVisible.value = true
}

function handleView(index,row) {
  dialogTitle.value = t('language.common.VIEW')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_VIEW
  const {...raw} = row
  rowData.value = raw
  if(rowData.value.metadataId==0){
    rowData.value.metadataId = null
  }
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleEdit(index,row) {
  dialogTitle.value = t('language.common.EDIT')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_EDIT
  const {...raw} = row
  rowData.value = raw
  if(rowData.value.metadataId==0){
    rowData.value.metadataId = null
  }
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleDelete(index,row) {
  const {...raw} = row
  rowData.value = raw
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_DELETE

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.agent.AGENT')+": "+rowData.value.agentName, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteAgent(rowData.value.agentId).then(x=>{
          if(x.data.isSuccess){
            if(rowData.value.configureId!=0){
              deleteConfigure(rowData.value.configureId)
            }
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            agentsStore.agents = agentsStore.agents.filter(x=>x.agentId!=rowData.value.agentId)
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

async function handleConfirm() {
  let configureId = 0
  //insert or update metadata instance
  if (rowData.value.metadataId != 0 && rowData.value.metadataId != null) {
    const metadataAttributeJson = JSON.stringify(metaInstanceRef.value.getMetadataAttribute())
    let fn = createConfigure
    let json = {}
    if(rowData.value.configureId==0){
      fn = createConfigure
      json = {
        metadataId: rowData.value.metadataId,
        configureJson: metadataAttributeJson
      }
    }
    else {
      fn = updateConfigure
      json = {
        metadataId: rowData.value.metadataId,
        configureJson: metadataAttributeJson,
        configureId: rowData.value.configureId
      }
      configureId = rowData.value.configureId
    }
    const res = await fn(json)
    if(res.data.isSuccess){
      if(rowData.value.configureId==0){
        let configure = res.data.result
        configureId = configure.configureId
        delete configure['@type']
        configures.value.push(configure)
      }else{
        const l = configures.value.filter(x=>x.configureId==configureId)
        l[0].configureJson = metadataAttributeJson
        l[0].metadataId = rowData.value.metadataId
      }
      keyNum.value++
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
      return
    }
  }else {
    if (rowData.value.metadataId == null) {
      rowData.value.metadataId = 0
    }
  }
  rowData.value.configureId = configureId

  //insert or update agent
  const {createdAt, updatedAt, ...newTable} = rowData.value;
  let fn = createAgent
  if (operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT) {
    fn = updateAgent
  }
  fn(newTable).then(x => {
    if (x.data.isSuccess) {
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      if (operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT) {
        const {...newStoreData} = rowData.value
        for (let i in agentsStore.agents) {
          if (agentsStore.agents[i].agentId == newTable.agentId) {
            agentsStore.agents[i] = newStoreData
            break
          }
        }
      } else {
        const res = x.data.result
        delete res['@type']
        agentsStore.agents.push(res)
      }
      dialogVisible.value = false
    } else {
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  }).catch(() => {
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  })
}

const getUserNickByUserId =  (userId) => {
  if(users && users.value){
    const l = users.value.filter(x=>x.userId == userId)
    if(l.length>0)
      return l[0].userNick
  }
  return ""
}

const getMetadataNameById = (metadataId) => {
  if(metadatas && metadatas.value){
    const l = metadatas.value.filter(x=>x.metadataId == metadataId)
    if(l.length>0)
      return l[0].metadataAlias
  }
  return ""
}

const rules = {
  agentName: [{ required: true,trigger: 'blur' }]
}

</script>

<style scoped>

</style>

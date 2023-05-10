<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-fsu'])">{{$t('language.common.CREATE')}}</el-button>
    <div >
      <el-table :data="getFsusByAgentId($route.params['agentId'])" :max-height="maxHeight">
        <el-table-column prop="fsuId" :label="$t('language.fsu.FSU_ID')" width="80" />
        <el-table-column prop="fsuName" :label="$t('language.fsu.FSU_NAME')" width="150" />
        <el-table-column prop="agentId" :label="$t('language.fsu.FSU_AGENT')" width="150" >
          <template #default="scope">
            {{getAgentNameByAgentId(scope.row.agentId)}}
          </template>
        </el-table-column>
        <el-table-column prop="fsuTypeId" :label="$t('language.fsu.FSU_TYPE')" width="150" >
          <template #default="scope">
            {{getFsuTypeNameByFsuTypeId(scope.row.fsuTypeId)}}
          </template>
        </el-table-column>
        <el-table-column prop="relatedMetadataId" :label="$t('language.fsu.DRIVER_RELATED_FLAG_METADATA')" width="150" >
          <template #default="scope">
            {{getMetadataNameById(scope.row.relatedMetadataId)}}
          </template>
        </el-table-column>
        <el-table-column prop="relatedMetadataId" :label="$t('language.fsu.DRIVER_RELATED_FLAG')" width="200" >
          <template #default="scope">
            <MetaInstance :metadata-id="scope.row.relatedMetadataId" :configure-id="0" :read-only="true" :key="keyNum"></MetaInstance>
          </template>
        </el-table-column>
        <el-table-column prop="metadataId" :label="$t('language.common.METADATA')" width="150" >
          <template #default="scope">
            {{getMetadataNameById(scope.row.metadataId)}}
          </template>
        </el-table-column>
        <el-table-column prop="configureId" :label="$t('language.common.METADATA_INSTANCE')" width="200" >
          <template #default="scope">
            <MetaInstance :metadata-id="scope.row.metadataId" :configure-id="scope.row.configureId" :key="keyNum" :read-only="true"></MetaInstance>
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
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-fsu','update-fsu'])">{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-fsu'])">{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="rowData" :rules="rules" ref="dialogForm">
        <el-form-item :label="$t('language.fsu.FSU_ID')" label-position="left" label-width="150px" prop="fsuId" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_VIEW"  >
          <el-input v-model="rowData.fsuId" readonly />
        </el-form-item>
        <el-form-item :label="$t('language.fsu.FSU_NAME')" label-position="left" label-width="150px" prop="fsuName" required>
          <el-input v-model="rowData.fsuName"/>
        </el-form-item>
        <el-form-item :label="$t('language.fsu.FSU_AGENT')" label-position="left" label-width="150px" prop="agentId"  required>
          <el-select v-model="rowData.agentId" class="m-2" :placeholder="$t('language.common.SELECT')" >
            <el-option v-for="item in agents" :key="item.agentId" :label="item.agentName" :value="item.agentId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.fsu.FSU_TYPE')" label-position="left" label-width="150px" prop="fsuTypeId" required>
          <el-select v-model="rowData.fsuTypeId" class="m-2" :placeholder="$t('language.common.SELECT')" @change="changeFsuType">
            <el-option v-for="item in fsuTypes" :key="item.fsuTypeId" :label="item.fsuTypeName" :value="item.fsuTypeId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.fsu.DRIVER_RELATED_FLAG_METADATA')" label-position="left" label-width="150px" prop="relatedMetadataId">
          <el-select v-model="rowData.relatedMetadataId" class="m-2" :placeholder="$t('language.common.SELECT')">
            <el-option v-for="item in metadatas" :key="item.metadataId" :label="item.metadataAlias" :value="item.metadataId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.fsu.DRIVER_RELATED_FLAG')" label-position="left" label-width="120px" prop="relatedMetadataId" >
          <MetaInstance :metadata-id="rowData.relatedMetadataId" :configure-id="0"  :read-only="true"></MetaInstance>
        </el-form-item>
        <el-form-item :label="$t('language.common.METADATA')" label-position="left" label-width="150px" prop="metadataId" >
          <el-select v-model="rowData.metadataId" class="m-2" :placeholder="$t('language.common.SELECT')" >
            <el-option v-for="item in metadatas" :key="item.metadataId" :label="item.metadataAlias" :value="item.metadataId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.common.METADATA_INSTANCE')" label-position="left" label-width="150px" prop="configureId" >
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
import {computed, nextTick, ref, toRaw} from "vue";
import {ElMessage, ElMessageBox} from "element-plus";
import {operatedTypeEnum} from "@/utils/common";
import i18n from "@/lang"
import {getAllUserAlias} from "@/api/user";
import {useActionsStore} from "@/stores/action";
import {useAgentsStore} from "@/stores/agent";
import {getAllMetadatas} from "@/api/metadata";
import {createConfigure, deleteConfigure, getAllConfigures, retrieveConfigure, updateConfigure} from "@/api/configure";
import {getAllAgents} from "@/api/agent";
import {initMetadataAttributeValue} from "@/utils/metadata"
import MetaInstance from "@/components/setting/metadata/MetaInstance.vue";
import {getAllFsuTypes} from "@/api/fsu-type";
import {createFsu, deleteFsu, updateFsu} from "@/api/fsu";
import {useRoute} from "vue-router";

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
const agents = ref()
const fsuTypes = ref()
const metaInstanceRef = ref()
const actionsStore = useActionsStore()
const agentsStore = useAgentsStore()
const route = useRoute()

Promise.all([getAllUserAlias(),getAllMetadatas(),getAllConfigures(),getAllAgents(),getAllFsuTypes()]).then(function (x) {
  users.value = x[0].data.result.items
  metadatas.value = x[1].data.result.items
  configures.value = x[2].data.result.items
  agents.value = x[3].data.result.items
  fsuTypes.value = x[4].data.result.items
})

const getFsusByAgentId = computed(()=>(agentId) => {
  return agentsStore.fsus.filter(x=>x.agentId==agentId)
})

let rowData = ref({
  fsuId:null,
  fsuName:"",
  agentId:null,
  fsuTypeId:null,
  relatedMetadataId:null,
  metadataId:null,
  configureId:0,
  createdAt:0,
  updatedAt:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  rowData.value = {
    fsuId:null,
    fsuName:"",
    agentId:Number(route.params['agentId']),
    fsuTypeId:null,
    metadataId:null,
    configureId:0,
    createdAt:0,
    updatedAt:0,
  }
  nextTick(()=>{
    dialogVisible.value = true
  })
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

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.fsu.FSU')+": "+rowData.value.fsuName, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteFsu(rowData.value.fsuId).then(x=>{
          if(x.data.isSuccess){
            if(rowData.value.configureId!=0){
              deleteConfigure(rowData.value.configureId)
            }
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            agentsStore.fsus = agentsStore.fsus.filter(x=>x.fsuId!=rowData.value.fsuId)
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
    }else{
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
      }
      else{
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
  let fn = createFsu
  if (operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT) {
    fn = updateFsu
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
        for (let i in agentsStore.fsus) {
          if (agentsStore.fsus[i].fsuId == newTable.fsuId) {
            agentsStore.fsus[i] = newStoreData
            break
          }
        }
      } else {
        const res = x.data.result
        delete res['@type']
        agentsStore.fsus.push(res)
      }
      dialogVisible.value = false
    } else {
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  }).catch(() => {
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  })
}

const changeFsuType = (fsuTypeId) =>{
  const l = fsuTypes.value.filter(x=>x.fsuTypeId == fsuTypeId)
  if(l.length>0)
    rowData.value.metadataId = l[0].metadataId
}

const getMetadataNameById = (metadataId) => {
  if(metadatas && metadatas.value){
    const l = metadatas.value.filter(x=>x.metadataId == metadataId)
    if(l.length>0)
      return l[0].metadataAlias
  }
  return ""
}

const getAgentNameByAgentId = (agentId) => {
  if(agents && agents.value){
    const l = agents.value.filter(x=>x.agentId == agentId)
    if(l.length>0)
      return l[0].agentName
  }
  return ""
}

const getFsuTypeNameByFsuTypeId = (fsuTypeId) =>{
  if(fsuTypes && fsuTypes.value){
    const l = fsuTypes.value.filter(x=>x.fsuTypeId == fsuTypeId)
    if(l.length>0)
      return l[0].fsuTypeName
  }
  return ""
}

const getUserNickByUserId =  (userId) => {
  if(users && users.value){
    const l = users.value.filter(x=>x.userId == userId)
    if(l.length>0)
      return l[0].userNick
  }
  return ""
}

const rules = {
  fsuName: [{ required: true,trigger: 'blur' }],
  fsuTypeId: [{ required: true,trigger: 'blur' }],
  agentId: [{ required: true,trigger: 'blur' }],
}

</script>

<style scoped>

</style>

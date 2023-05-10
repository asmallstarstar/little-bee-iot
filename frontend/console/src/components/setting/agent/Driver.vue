<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-driver'])">{{$t('language.common.CREATE')}}</el-button>
    <div >
      <el-table :data="getDriversByFsuId($route.params['fsuId'])" :max-height="maxHeight">
        <el-table-column prop="driverId" :label="$t('language.driver.DRIVER_ID')" width="80" />
        <el-table-column prop="driverName" :label="$t('language.driver.DRIVER_NAME')" width="150" />
        <el-table-column prop="driverTypeId" :label="$t('language.driverType.DRIVER_TYPE')" width="150">
          <template #default="scope">
            {{getDriverTypeNameByTypeId(scope.row.driverTypeId)}}
          </template>
        </el-table-column>
        <el-table-column prop="fsuId" :label="$t('language.fsu.FSU')" width="150">
          <template #default="scope">
            {{getFsuNameByFsuId(scope.row.fsuId)}}
          </template>
        </el-table-column>
        <el-table-column prop="relatedMetadataId" :label="$t('language.fsu.DRIVER_RELATED_FLAG')" width="200" >
          <template #default="scope">
            <MetaInstance :metadata-id="getFsuRelatedMetadata(scope.row.fsuId)" :configure-id="scope.row.relatedFlagConfigureId" :read-only="true" :key="keyNum"></MetaInstance>
          </template>
        </el-table-column>
        <el-table-column prop="driverAcquisitePeriod" :label="$t('language.driverType.DRIVER_TYPE_ACQUISITE_PERIOD')" width="120" />
        <el-table-column prop="driverTimeoutCount" :label="$t('language.driverType.DRIVER_TYPE_TIMEOUT_COUNT')" width="120" />
        <el-table-column prop="driverResetCount" :label="$t('language.driverType.DRIVER_TYPE_RESET_COUNT')" width="120" />
        <el-table-column prop="isWriteLog" :label="$t('language.driver.IS_WRITE_LOG')" width="120" >
          <template #default="scope">
            {{scope.row.isWriteLog?t("language.common.YES"):t("language.common.NO")}}
          </template>
        </el-table-column>
        <el-table-column prop="driverParameterMetadataId" :label="$t('language.common.METADATA')" width="150" >
          <template #default="scope">
            {{getMetadataNameById(scope.row.driverParameterMetadataId)}}
          </template>
        </el-table-column>
        <el-table-column prop="configureId" :label="$t('language.common.METADATA_INSTANCE')" width="200" >
          <template #default="scope">
            <MetaInstance :metadata-id="scope.row.driverParameterMetadataId" :configure-id="scope.row.driverParameterConfigureId" :key="keyNum" :read-only="true"></MetaInstance>
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
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-driver','update-driver'])">{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-driver'])">{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="rowData" :rules="rules" ref="dialogForm">
        <el-form-item :label="$t('language.driver.DRIVER_ID')" label-position="left" label-width="120px" prop="driverId" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_VIEW"  required>
          <el-input v-model="rowData.driverId" readonly />
        </el-form-item>
        <el-form-item :label="$t('language.driver.DRIVER_NAME')" label-position="left" label-width="120px" prop="driverName" required>
          <el-input v-model="rowData.driverName"/>
        </el-form-item>
        <el-form-item :label="$t('language.driverType.DRIVER_TYPE')" label-position="left" label-width="120px" prop="driverTypeId"  required>
          <el-select v-model="rowData.driverTypeId" class="m-2" :placeholder="$t('language.common.SELECT')" @change="changeDriverType">
            <el-option v-for="item in driverTypes" :key="item.driverTypeId" :label="item.driverTypeName" :value="item.driverTypeId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.fsu.FSU')" label-position="left" label-width="120px" prop="fsuId" required>
          <el-select v-model="rowData.fsuId" class="m-2" :placeholder="$t('language.common.SELECT')" >
            <el-option v-for="item in fsus" :key="item.fsuId" :label="item.fsuName" :value="item.fsuId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.fsu.DRIVER_RELATED_FLAG')" label-position="left" label-width="120px" prop="configureId" v-if="getFsuRelatedMetadata(rowData.fsuId)">
          <MetaInstance :metadata-id="getFsuRelatedMetadata(rowData.fsuId)" :configure-id="0" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_CREATE" ref="relatedFlagInstanceRef" :read-only="false"></MetaInstance>
          <MetaInstance :metadata-id="getFsuRelatedMetadata(rowData.fsuId)" :configure-id="rowData.relatedFlagConfigureId" v-else ref="relatedFlagInstanceRef" :read-only="false"></MetaInstance>
        </el-form-item>
        <el-form-item :label="$t('language.driverType.DRIVER_TYPE_ACQUISITE_PERIOD')" label-position="left" label-width="120px" prop="driverAcquisitePeriod" required>
          <el-input v-model="rowData.driverAcquisitePeriod"/>
        </el-form-item>
        <el-form-item :label="$t('language.driverType.DRIVER_TYPE_TIMEOUT_COUNT')" label-position="left" label-width="120px" prop="driverTimeoutCount" required>
          <el-input v-model="rowData.driverTimeoutCount"/>
        </el-form-item>
        <el-form-item :label="$t('language.driverType.DRIVER_TYPE_RESET_COUNT')" label-position="left" label-width="120px" prop="driverResetCount" required>
          <el-input v-model="rowData.driverResetCount"/>
        </el-form-item>
        <el-form-item :label="$t('language.driver.IS_WRITE_LOG')" label-position="left" label-width="120px" prop="isWriteLog">
          <el-checkbox v-model="rowData.isWriteLog" :label="$t('language.driver.IS_WRITE_LOG')" />
        </el-form-item>
        <el-form-item :label="$t('language.common.METADATA')" label-position="left" label-width="120px" prop="metadataId" >
          <el-select v-model="rowData.driverParameterMetadataId" class="m-2" :placeholder="$t('language.common.SELECT')" >
            <el-option v-for="item in metadatas" :key="item.metadataId" :label="item.metadataAlias" :value="item.metadataId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.common.METADATA_INSTANCE')" label-position="left" label-width="120px" prop="configureId" >
          <MetaInstance :metadata-id="rowData.driverParameterMetadataId" :configure-id="0" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_CREATE" ref="metaInstanceRef" :read-only="false"></MetaInstance>
          <MetaInstance :metadata-id="rowData.driverParameterMetadataId" :configure-id="rowData.driverParameterConfigureId" v-else ref="metaInstanceRef" :read-only="false"></MetaInstance>
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
import {initMetadataAttributeValue} from "@/utils/metadata"
import MetaInstance from "@/components/setting/metadata/MetaInstance.vue";
import {createFsu, deleteFsu, getAllFsus, updateFsu} from "@/api/fsu";
import {useRoute} from "vue-router";
import {getAllDriverTypes} from "@/api/driver-type";
import {createDriver, deleteDriver, updateDriver} from "@/api/driver";

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
const fsus = ref()
const driverTypes = ref()
const metaInstanceRef = ref()
const relatedFlagInstanceRef = ref()
const actionsStore = useActionsStore()
const agentsStore = useAgentsStore()
const route = useRoute()

Promise.all([getAllUserAlias(),getAllMetadatas(),getAllConfigures(),getAllFsus(),getAllDriverTypes()]).then(function (x) {
  users.value = x[0].data.result.items
  metadatas.value = x[1].data.result.items
  configures.value = x[2].data.result.items
  fsus.value = x[3].data.result.items
  driverTypes.value = x[4].data.result.items
})

const getDriversByFsuId = computed(()=>(fsuId) => {
  return agentsStore.drivers.filter(x=>x.fsuId==fsuId)
})

const getFsuRelatedMetadata = computed(()=>(fsuId) => {
  if(fsus && fsus.value){
    const l = fsus.value.filter(x=>x.fsuId==fsuId)
    if(l.length>0){
      return l[0].relatedMetadataId
    }
  }
  return 0
})

let rowData = ref({
  driverId:null,
  driverTypeId:null,
  fsuId:null,
  relatedFlagConfigureId:0,
  driverName:"",
  driverParameterMetadataId:null,
  driverParameterConfigureId:0,
  driverAcquisitePeriod:5,
  driverTimeoutCount:3,
  driverResetCount:3,
  isWriteLog:0,
  createdAt:0,
  updatedAt:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  rowData.value = {
    driverId:null,
    driverTypeId:null,
    fsuId:Number(route.params['fsuId']),
    relatedFlagConfigureId:0,
    driverName:"",
    driverParameterMetadataId:null,
    driverParameterConfigureId:0,
    driverAcquisitePeriod:5,
    driverTimeoutCount:3,
    driverResetCount:3,
    isWriteLog:0,
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
  if(rowData.value.driverParameterMetadataId==0){
    rowData.value.driverParameterMetadataId = null
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
  if(rowData.value.driverParameterMetadataId==0){
    rowData.value.driverParameterMetadataId = null
  }
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleDelete(index,row) {
  const {...raw} = row
  rowData.value = raw
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_DELETE

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.fsu.FSU')+": "+rowData.value.driverName, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteDriver(rowData.value.driverId).then(x=>{
          if(x.data.isSuccess){
            if(rowData.value.driverParameterConfigureId){
              deleteConfigure(rowData.value.driverParameterConfigureId)
            }
            if(rowData.value.relatedFlagConfigureId){
              deleteConfigure(rowData.value.relatedFlagConfigureId)
            }
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            agentsStore.drivers = agentsStore.drivers.filter(x=>x.driverId!=rowData.value.driverId)
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

async function handleConfirm() {
  let configureId = 0
  //insert or update metadata instance
  if (rowData.value.driverParameterMetadataId) {
    const metadataAttributeJson = JSON.stringify(metaInstanceRef.value.getMetadataAttribute())
    let fn = createConfigure
    let json = {}
    if(rowData.value.driverParameterConfigureId==0){
      fn = createConfigure
      json = {
        metadataId: rowData.value.driverParameterMetadataId,
        configureJson: metadataAttributeJson
      }
    }else{
      fn = updateConfigure
      json = {
        metadataId: rowData.value.driverParameterMetadataId,
        configureJson: metadataAttributeJson,
        configureId: rowData.value.driverParameterConfigureId
      }
      configureId = rowData.value.driverParameterConfigureId
    }
    const res = await fn(json)
    if(res.data.isSuccess){
      if(rowData.value.driverParameterConfigureId==0){
        let configure = res.data.result
        configureId = configure.configureId
        delete configure['@type']
        configures.value.push(configure)
      }
      else{
        const l = configures.value.filter(x=>x.configureId==configureId)
        l[0].configureJson = metadataAttributeJson
        l[0].metadataId = rowData.value.driverParameterMetadataId
      }
      keyNum.value++
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
      return
    }
  }else {
    if (rowData.value.driverParameterMetadataId == null) {
      rowData.value.driverParameterMetadataId = 0
    }
  }
  rowData.value.driverParameterConfigureId = configureId

  //insert or update driver related configure
  let relatedConfigureId = 0;
  const f = fsus.value.filter(x=>x.fsuId==rowData.value.fsuId)
  if (f.length>0 && f[0].relatedMetadataId) {
    const metadataAttributeJson = JSON.stringify(relatedFlagInstanceRef.value.getMetadataAttribute())
    let fn = createConfigure
    let json = {}
    if(rowData.value.relatedFlagConfigureId==0){
      fn = createConfigure
      json = {
        metadataId: f[0].relatedMetadataId,
        configureJson: metadataAttributeJson
      }
    }else{
      fn = updateConfigure
      json = {
        metadataId: f[0].relatedMetadataId,
        configureJson: metadataAttributeJson,
        configureId: rowData.value.relatedFlagConfigureId
      }
      relatedConfigureId = rowData.value.relatedFlagConfigureId
    }
    const res = await fn(json)
    if(res.data.isSuccess){
      if(rowData.value.relatedFlagConfigureId==0){
        let configure = res.data.result
        relatedConfigureId = configure.configureId
        delete configure['@type']
        configures.value.push(configure)
      }
      else{
        const l = configures.value.filter(x=>x.configureId==relatedConfigureId)
        l[0].configureJson = metadataAttributeJson
        l[0].metadataId = f[0].relatedMetadataId
      }
      keyNum.value++
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
      return
    }
  }
  rowData.value.relatedFlagConfigureId = relatedConfigureId

  //insert or update agent
  const {createdAt, updatedAt, ...newTable} = rowData.value;
  let fn = createDriver
  if (operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT) {
    fn = updateDriver
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
        for (let i in agentsStore.drivers) {
          if (agentsStore.drivers[i].driverId == newTable.driverId) {
            agentsStore.drivers[i] = newStoreData
            break
          }
        }
      } else {
        const res = x.data.result
        delete res['@type']
        agentsStore.drivers.push(res)
      }
      dialogVisible.value = false
    } else {
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  }).catch(() => {
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  })
}

const changeDriverType = (driverTypeId) =>{
  const l = driverTypes.value.filter(x=>x.driverTypeId == driverTypeId)
  if(l.length>0){
    rowData.value.driverParameterMetadataId = l[0].driverTypeParameterMetadataId
    rowData.value.driverAcquisitePeriod = l[0].driverTypeAcquisitePeriod
    rowData.value.driverTimeoutCount = l[0].driverTypeTimeoutCount
    rowData.value.driverResetCount = l[0].driverTypeResetCount
  }
}

const getMetadataNameById = (metadataId) => {
  if(metadatas && metadatas.value){
    const l = metadatas.value.filter(x=>x.metadataId == metadataId)
    if(l.length>0)
      return l[0].metadataAlias
  }
  return ""
}

const getFsuNameByFsuId = (fsuId) => {
  if(fsus && fsus.value){
    const l = fsus.value.filter(x=>x.fsuId == fsuId)
    if(l.length>0)
      return l[0].fsuName
  }
  return ""
}

const getDriverTypeNameByTypeId = (driverTypeId) =>{
  if(driverTypes && driverTypes.value){
    const l = driverTypes.value.filter(x=>x.driverTypeId == driverTypeId)
    if(l.length>0)
      return l[0].driverTypeName
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
  driverId: [{ required: true,trigger: 'blur' }],
  driverTypeId: [{ required: true,trigger: 'blur' }],
  fsuId: [{ required: true,trigger: 'blur' }],
  driverName: [{ required: true,trigger: 'blur' }],
  driverAcquisitePeriod: [{ required: true,trigger: 'blur' }],
  driverTimeoutCount: [{ required: true,trigger: 'blur' }],
  driverResetCount: [{ required: true,trigger: 'blur' }],
}

</script>

<style scoped>

</style>

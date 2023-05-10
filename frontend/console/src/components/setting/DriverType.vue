<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-driver-type'])">{{$t('language.common.CREATE')}}</el-button>
    <div >
      <el-table :data="tableDataSets" :max-height="maxHeight">
        <el-table-column prop="driverTypeId" :label="$t('language.driverType.DRIVER_TYPE_ID')" width="100" />
        <el-table-column prop="driverTypeName" :label="$t('language.driverType.DRIVER_TYPE_NAME')" width="150" />
        <el-table-column prop="driverTypeFileName" :label="$t('language.driverType.DRIVER_FILENAME')" width="200" />
        <el-table-column prop="driverTypeNote" :label="$t('language.driverType.DRIVER_TYPE_NOTE')" width="200" />
        <el-table-column prop="driverTypeAcquisitePeriod" :label="$t('language.driverType.DRIVER_TYPE_ACQUISITE_PERIOD')" width="100" />
        <el-table-column prop="driverTypeTimeoutCount" :label="$t('language.driverType.DRIVER_TYPE_TIMEOUT_COUNT')" width="100" />
        <el-table-column prop="driverTypeResetCount" :label="$t('language.driverType.DRIVER_TYPE_RESET_COUNT')" width="100" />
        <el-table-column prop="metadataId" :label="$t('language.common.METADATA')" width="150" >
          <template #default="scope">
            {{getMetadataNameById(scope.row.driverTypeParameterMetadataId)}}
          </template>
        </el-table-column>
        <el-table-column prop="metadataId" :label="$t('language.common.METADATA_INSTANCE')" width="200" >
          <template #default="scope">
            <MetaInstance :metadata-id="scope.row.driverTypeParameterMetadataId" :configure-id="0" :read-only="true" :key="keyNum"></MetaInstance>
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
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-driver-type','update-driver-type'])">{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-driver-type'])">{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="rowData" :rules="rules" ref="dialogForm">
        <el-form-item :label="$t('language.driverType.DRIVER_TYPE_ID')" label-position="left" label-width="120px" prop="driverTypeId" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_VIEW" required >
          <el-input v-model="rowData.driverTypeId" readonly />
        </el-form-item>
        <el-form-item :label="$t('language.driverType.DRIVER_TYPE_NAME')" label-position="left" label-width="120px" prop="driverTypeName" required>
          <el-input v-model="rowData.driverTypeName"/>
        </el-form-item>
        <el-form-item :label="$t('language.driverType.DRIVER_FILENAME')" label-position="left" label-width="120px" prop="driverTypeFileName" required>
          <el-input v-model="rowData.driverTypeFileName"/>
        </el-form-item>
        <el-form-item :label="$t('language.driverType.DRIVER_TYPE_NOTE')" label-position="left" label-width="120px" prop="driverTypeNote">
          <el-input v-model="rowData.driverTypeNote"/>
        </el-form-item>
        <el-form-item :label="$t('language.driverType.DRIVER_TYPE_ACQUISITE_PERIOD')" label-position="left" label-width="120px" prop="driverTypeAcquisitePeriod" required>
          <el-input v-model="rowData.driverTypeAcquisitePeriod"/>
        </el-form-item>
        <el-form-item :label="$t('language.driverType.DRIVER_TYPE_TIMEOUT_COUNT')" label-position="left" label-width="120px" prop="driverTypeTimeoutCount" required>
          <el-input v-model="rowData.driverTypeTimeoutCount"/>
        </el-form-item>
        <el-form-item :label="$t('language.driverType.DRIVER_TYPE_TIMEOUT_COUNT')" label-position="left" label-width="120px" prop="driverTypeResetCount" required>
          <el-input v-model="rowData.driverTypeResetCount"/>
        </el-form-item>
        <el-form-item :label="$t('language.common.METADATA')" label-position="left" label-width="120px" prop="metadataId" >
          <el-select v-model="rowData.driverTypeParameterMetadataId" class="m-2" :placeholder="$t('language.common.SELECT')" >
            <el-option v-for="item in metadatas" :key="item.metadataId" :label="item.metadataAlias" :value="item.metadataId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.common.METADATA_INSTANCE')" label-position="left" label-width="120px" prop="metadataId" >
          <MetaInstance :metadata-id="rowData.driverTypeParameterMetadataId" :configure-id="0"  :read-only="false"></MetaInstance>
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
import {nextTick, ref} from "vue";
import {ElMessage, ElMessageBox} from "element-plus";
import {operatedTypeEnum} from "../../utils/common";
import i18n from "@/lang"
import {getAllUserAlias} from "@/api/user";
import {useActionsStore} from "@/stores/action";
import {getAllMetadatas} from "@/api/metadata";
import MetaInstance from "@/components/setting/metadata/MetaInstance.vue";
import {createDriverType, deleteDriverType, getAllDriverTypes, updateDriverType} from "@/api/driver-type";

const t = i18n.global.t
let dialogTitle = ref(t('language.common.CREATE'))
let dialogVisible = ref(false)
let operatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
let maxHeight = ref(window.innerHeight - 165)
const keyNum = ref(0)
const dialogForm = ref()
const users = ref()
const tableDataSets = ref([])
const metadatas = ref([])
const actionsStore = useActionsStore()

Promise.all([getAllUserAlias(),getAllDriverTypes(),getAllMetadatas()]).then(function (x) {
  users.value = x[0].data.result.items
  tableDataSets.value = x[1].data.result.items
  metadatas.value = x[2].data.result.items
})

let rowData = ref({
  driverTypeId:null,
  driverTypeName:"",
  driverTypeFileName:"",
  driverTypeNote:"",
  driverTypeAcquisitePeriod:5,
  driverTypeTimeoutCount:3,
  driverTypeResetCount:3,
  driverTypeParameterMetadataId:null,
  createdAt:0,
  updatedAt:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  rowData.value = {
    driverTypeId:null,
    driverTypeName:"",
    driverTypeFileName:"",
    driverTypeNote:"",
    driverTypeAcquisitePeriod:5,
    driverTypeTimeoutCount:3,
    driverTypeResetCount:3,
    driverTypeParameterMetadataId:null,
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
  if(rowData.value.driverTypeParameterMetadataId==0){
    rowData.value.driverTypeParameterMetadataId = null
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
  if(rowData.value.driverTypeParameterMetadataId==0){
    rowData.value.driverTypeParameterMetadataId = null
  }
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleDelete(index,row) {
  const {...raw} = row
  rowData.value = raw
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_DELETE

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.driverType.DRIVER_TYPE')+": "+rowData.value.driverTypeName, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteDriverType(rowData.value.driverTypeId).then(x=>{
          if(x.data.isSuccess){
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            getDriverTypes()
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

function handleConfirm() {
  const { createdAt,updatedAt, ...newTable } = rowData.value;
  let fn = createDriverType
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateDriverType
  }
  fn(newTable).then(x=>{
    if(x.data.isSuccess){
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
        const { ...newStoreData } = rowData.value
        for(let i in tableDataSets.value){
          if(tableDataSets.value[i].driverTypeId==newTable.driverTypeId){
            tableDataSets.value[i] = newStoreData
            break
          }
        }
      }else{
        const res = x.data.result
        delete res['@type']
        tableDataSets.value.push(res)
      }
      dialogVisible.value=false
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  }).catch(()=>{
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

function getDriverTypes() {
  getAllDriverTypes().then(x=>{
    if(x.data.isSuccess){
      tableDataSets.value = x.data.result.items
    }
  })
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
  driverTypeId: [{ required: true,trigger: 'blur' }],
  driverTypeName: [{ required: true,trigger: 'blur' }],
  driverTypeFileName: [{ required: true,trigger: 'blur' }],
  driverTypeAcquisitePeriod: [{ required: true,trigger: 'blur' }],
  driverTypeTimeoutCount: [{ required: true,trigger: 'blur' }],
  driverTypeResetCount: [{ required: true,trigger: 'blur' }],
}

</script>

<style scoped>

</style>

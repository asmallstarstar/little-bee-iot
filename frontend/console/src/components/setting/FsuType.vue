<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-fsu-type'])">{{$t('language.common.CREATE')}}</el-button>
    <div >
      <el-table :data="tableDataSets" :max-height="maxHeight">
        <el-table-column prop="fsuTypeId" :label="$t('language.fsuType.FSU_TYPE_ID')" width="100" />
        <el-table-column prop="fsuTypeName" :label="$t('language.fsuType.FSU_TYPE_NAME')" width="150" />
        <el-table-column prop="dataAnalysisMode" :label="$t('language.fsuType.DATA_ANALYSIS_MODE')" width="150" >
          <template #default="scope">
            {{getDataAnalysisModeNameByModeId(scope.row.dataAnalysisMode)}}
          </template>
        </el-table-column>
        <el-table-column prop="newInstanceName" :label="$t('language.fsuType.NEW_INSTANCE_NAME')" width="300" />
        <el-table-column prop="metadataId" :label="$t('language.common.METADATA')" width="150" >
          <template #default="scope">
            {{getMetadataNameById(scope.row.metadataId)}}
          </template>
        </el-table-column>
        <el-table-column prop="metadataId" :label="$t('language.common.METADATA_INSTANCE')" width="200" >
          <template #default="scope">
            <MetaInstance :metadata-id="scope.row.metadataId" :configure-id="0" :read-only="true" :key="keyNum"></MetaInstance>
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
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-fsu-type','update-fsu-type'])">{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-fsu-type'])">{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="rowData" :rules="rules" ref="dialogForm">
        <el-form-item :label="$t('language.fsuType.FSU_TYPE_ID')" label-position="left" label-width="120px" prop="fsuTypeId" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_VIEW" required >
          <el-input v-model="rowData.fsuTypeId" readonly />
        </el-form-item>
        <el-form-item :label="$t('language.fsuType.FSU_TYPE_NAME')" label-position="left" label-width="120px" prop="fsuTypeName" required>
          <el-input v-model="rowData.fsuTypeName"/>
        </el-form-item>
        <el-form-item :label="$t('language.fsuType.DATA_ANALYSIS_MODE')" label-position="left" label-width="120px" prop="dataAnalysisMode" required>
          <el-select v-model="rowData.dataAnalysisMode" class="m-2" :placeholder="$t('language.common.SELECT')" >
            <el-option  :key="fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_UNKNOWN" :label="$t('language.common.UNKNOWN')" :value="fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_UNKNOWN"/>
            <el-option  :key="fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_BOTTOM" :label="$t('language.fsuType.ANALYSIS_MODE_BOTTOM')" :value="fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_BOTTOM"/>
            <el-option  :key="fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_CENTER" :label="$t('language.fsuType.ANALYSIS_MODE_CENTER')" :value="fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_CENTER"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.fsuType.NEW_INSTANCE_NAME')" label-position="left" label-width="120px" prop="newInstanceName" required>
          <el-input v-model="rowData.newInstanceName"/>
        </el-form-item>
        <el-form-item :label="$t('language.common.METADATA')" label-position="left" label-width="120px" prop="metadataId" >
          <el-select v-model="rowData.metadataId" class="m-2" :placeholder="$t('language.common.SELECT')" >
            <el-option v-for="item in metadatas" :key="item.metadataId" :label="item.metadataAlias" :value="item.metadataId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.common.METADATA_INSTANCE')" label-position="left" label-width="120px" prop="metadataId" >
          <MetaInstance :metadata-id="rowData.metadataId" :configure-id="0"  :read-only="false"></MetaInstance>
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
import {createFsuType, deleteFsuType, getAllFsuTypes, updateFsuType} from "@/api/fsu-type";
import {getAllMetadatas} from "@/api/metadata";
import {fsuDataParsingTypeEnum} from "@/utils/agent"
import MetaInstance from "@/components/setting/metadata/MetaInstance.vue";

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

Promise.all([getAllUserAlias(),getAllFsuTypes(),getAllMetadatas()]).then(function (x) {
  users.value = x[0].data.result.items
  tableDataSets.value = x[1].data.result.items
  metadatas.value = x[2].data.result.items
})

let rowData = ref({
  fsuTypeId:null,
  fsuTypeName:"",
  dataAnalysisMode:0,
  newInstanceName:"",
  metadataId:null,
  createdAt:0,
  updatedAt:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  rowData.value = {
    fsuTypeId:null,
    fsuTypeName:"",
    dataAnalysisMode:0,
    newInstanceName:"",
    metadataId:null,
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

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.fsuType.FSU_TYPE')+": "+rowData.value.fsuTypeName, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteFsuType(rowData.value.fsuTypeId).then(x=>{
          if(x.data.isSuccess){
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            getFsuTypes()
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

function handleConfirm() {
  const { createdAt,updatedAt, ...newTable } = rowData.value;
  let fn = createFsuType
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateFsuType
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
        for(let i=0;i<tableDataSets.value.length;i++){
          if(tableDataSets.value[i].fsuTypeId==newTable.fsuTypeId){
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

function getFsuTypes() {
  getAllFsuTypes().then(x=>{
    if(x.data.isSuccess){
      tableDataSets.value = x.data.result.items
    }
  })
}

const getDataAnalysisModeNameByModeId = (modeId) => {
  if(modeId==fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_UNKNOWN)
    return t('language.common.UNKNOWN')
  if(modeId==fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_BOTTOM)
    return t('language.fsuType.ANALYSIS_MODE_BOTTOM')
  if(modeId==fsuDataParsingTypeEnum.FSU_DATA_PARSING_TYPE_CENTER)
    return t('language.fsuType.ANALYSIS_MODE_CENTER')
}

const getMetadataNameById = (metadataId) => {
  if(metadatas.value){
    const l = metadatas.value.filter(x=>x.metadataId == metadataId)
    if(l.length>0)
      return l[0].metadataAlias
  }
  return ""
}

const rules = {
  fsuTypeId: [{ required: true,trigger: 'blur' }],
  fsuTypeName: [{ required: true,trigger: 'blur' }],
  dataAnalysisMode: [{ required: true,trigger: 'blur' }],
  newInstanceName: [{ required: true,trigger: 'blur' }],
}

</script>

<style scoped>

</style>

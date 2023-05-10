<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-metadata'])">{{$t('language.common.CREATE')}}</el-button>
    <div >
      <el-table :data="tableDataSets" :max-height="maxHeight">
        <el-table-column prop="metadataId" :label="$t('language.metadata.METADATA_ID')" width="130" />
        <el-table-column prop="metadataName" :label="$t('language.metadata.METADATA_NAME')" width="230" />
        <el-table-column prop="metadataAlias" :label="$t('language.metadata.METADATA_ALIAS')" width="230" />
        <el-table-column :label="$t('language.metadata.META_INSTANCE')" width="250" >
          <template #default="scope">
            <MetaInstance :metadata-id="scope.row.metadataId" :configure-id="0" :key="keyNum" :read-only="true"></MetaInstance>
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
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-metadata','update-metadata'])">{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-metadata'])">{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="1000px">
      <Metadata :metadata-id="metadataId" :operator-type="operatedType" ref="metadataRef"></Metadata>
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
import {operatedTypeEnum} from "../../utils/common";
import i18n from "@/lang"
import {getAllUserAlias} from "@/api/user";
import {useActionsStore} from "@/stores/action";
import {createMetadata, deleteMetadata, getAllMetadatas, updateMetadata} from "@/api/metadata";
import MetaInstance from "@/components/setting/metadata/MetaInstance.vue";
import Metadata from "@/components/setting/metadata/Metadata.vue";

const t = i18n.global.t
const keyNum = ref(0)
let dialogTitle = ref(t('language.common.CREATE'))
let dialogVisible = ref(false)
let operatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
let maxHeight = ref(window.innerHeight - 165)
const metadataId = ref(0)
const dialogForm = ref()
const users = ref()
const tableDataSets = ref([])
const metadataRef = ref()
const actionsStore = useActionsStore()

Promise.all([getAllUserAlias(),getAllMetadatas()]).then(function (x) {
  users.value = x[0].data.result.items
  x[1].data.result.items.forEach(y=>y.metadataAttribute =  JSON.parse(y.metadataAttributeJson))
  tableDataSets.value = x[1].data.result.items
})

let rowData = ref({
  metadataId:null,
  metadataName:"",
  metadataAlias:"",
  metadataAttribute:[],
  createdAt:0,
  updatedAt:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  dialogVisible.value = true
  nextTick(()=>{
    metadataRef.value.setMetadata({
      metadataId:null,
      metadataName:"",
      metadataAlias:"",
      metadataAttribute:[],
    })
  })
}

function handleView(index,row) {
  metadataId.value = row.metadataId
  dialogTitle.value = t('language.common.VIEW')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_VIEW
  const {...raw} = row
  raw.metadataAttribute = JSON.parse(raw.metadataAttributeJson)
  dialogVisible.value = true
  nextTick(()=>{
    metadataRef.value.setMetadata(raw)
  })
}

function handleEdit(index,row) {
  metadataId.value = row.metadataId
  dialogTitle.value = t('language.common.EDIT')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_EDIT
  const {...raw} = row
  raw.metadataAttribute = JSON.parse(raw.metadataAttributeJson)
  dialogVisible.value = true
  nextTick(()=>{
    metadataRef.value.setMetadata(raw)
  })
}

function handleDelete(index,row) {
  const {...raw} = row
  rowData.value = raw
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_DELETE

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.metadata.METADATA')+": "+rowData.value.metadataAlias, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteMetadata(rowData.value.metadataId).then(x=>{
          if(x.data.isSuccess){
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            getMetadatas()
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

function handleConfirm() {
  const metadata = metadataRef.value.getMetadata()
  metadata.metadataAttributeJson = JSON.stringify(metadata.metadataAttribute)
  delete metadata.metadataAttribute
  delete metadata.createdAt
  delete metadata.updatedAt
  let fn = createMetadata
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateMetadata
  }
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_CREATE){
    delete metadata.metadataId
  }
  fn(JSON.stringify(metadata)).then(x=>{
    if(x.data.isSuccess){
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
        const m = metadataRef.value.getMetadata()
        m.metadataAttributeJson = JSON.stringify(m.metadataAttribute)
        for(let i=0;i<tableDataSets.value.length;i++){
          if(tableDataSets.value[i].metadataId==m.metadataId){
            tableDataSets.value[i] = m
            break
          }
        }
      }else{
        const res = x.data.result
        delete res['@type']
        tableDataSets.value.push(res)
      }
      dialogVisible.value=false
      keyNum.value++
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

function getMetadatas() {
  getAllMetadatas().then(x=>{
    if(x.data.isSuccess){
      x.data.result.items.forEach(x=>x.metadataAttribute =  JSON.parse(x.metadataAttributeJson))
      tableDataSets.value = x.data.result.items
    }
  })
}

const rules = {
  alarmLevelNumber: [{ required: true,trigger: 'blur' }],
  alarmLevelName: [{ required: true,trigger: 'blur' }],
  alarmLevelAlias: [{ required: true,trigger: 'blur' }],
}

</script>

<style scoped>

</style>

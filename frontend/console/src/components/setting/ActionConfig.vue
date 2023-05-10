<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-action'])">{{$t('language.common.CREATE')}}</el-button>
    <div >
      <el-table :data="tableDataSetsWithPage" :max-height="maxHeight">
        <el-table-column prop="actionId" :label="$t('language.action.ACTION_ID')" width="100" />
        <el-table-column prop="actionName" :label="$t('language.action.ACTION_NAME')" width="320" />
        <el-table-column prop="actionAlias" :label="$t('language.action.ACTION_ALIAS')" width="220" />
        <el-table-column :label="$t('language.actionGroup.ACTION_GROUP_PARENT')" width="180" >
          <template #default="scope">
            {{getActionGroupById(scope.row.actionGroupId)}}
          </template>
        </el-table-column>
        <el-table-column prop="url" :label="$t('language.action.URL')" width="340" />
        <el-table-column prop="method" :label="$t('language.action.METHOD')" width="100" />
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
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-action','update-action'])">{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-action'])">{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :page-sizes="[5, 10, 15, 20]" :small="small" :disabled="disabled"
                     :background="background" layout="sizes, prev, pager, next" :total="actionsStore.actions.length"
                     @size-change="handleSizeChange" @current-change="handleCurrentChange" style="margin-top: 3px;"/>
    </div>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="rowData" :rules="rules" ref="dialogForm">
        <el-form-item :label="$t('language.action.ACTION_ID')" label-position="left" label-width="120px" prop="actionId" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_VIEW" required >
          <el-input v-model="rowData.actionId" readonly />
        </el-form-item>
        <el-form-item :label="$t('language.actionGroup.ACTION_GROUP_PARENT')" label-position="left" label-width="120px" prop="actionGroupId" required>
          <el-select v-model="rowData.actionGroupId" class="m-2" :placeholder="$t('language.common.SELECT')" @change="changeParent">
            <el-option v-for="item in actionsStore.actionGroups" :key="item.actionGroupId" :label="item.actionGroupAlias" :value="item.actionGroupId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.action.ACTION_NAME')" label-position="left" label-width="120px" prop="actionName" required>
          <el-input v-model="rowData.actionName"/>
        </el-form-item>
        <el-form-item :label="$t('language.action.ACTION_ALIAS')" label-position="left" label-width="120px" prop="actionAlias" required>
          <el-input v-model="rowData.actionAlias"/>
        </el-form-item>
        <el-form-item :label="$t('language.action.URL')" label-position="left" label-width="120px" prop="url" required>
          <el-input v-model="rowData.url"/>
        </el-form-item>
        <el-form-item :label="$t('language.action.METHOD')" label-position="left" label-width="120px" prop="method" required>
          <el-input v-model="rowData.method"/>
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
import {computed, nextTick, ref} from "vue";
import {ElMessage, ElMessageBox} from "element-plus";
import {operatedTypeEnum} from "../../utils/common";
import i18n from "@/lang"
import {storeToRefs} from "pinia";
import {useActionsStore} from "@/stores/action";
import {getAllActionGroups} from "@/api/action-group";
import {getAllUserAlias} from "@/api/user";
import {createAction, deleteAction, getAllActions, updateAction} from "@/api/action";

const t = i18n.global.t
const currentPage = ref(1)
const pageSize = ref(15)
const small = ref(false)
const background = ref(false)
const disabled = ref(false)
const actionsStore = useActionsStore()
const {actionGroups:tableDataSets} = storeToRefs(actionsStore)
let dialogTitle = ref(t('language.common.CREATE'))
let dialogVisible = ref(false)
let operatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
let maxHeight = ref(window.innerHeight - 165)
const dialogForm = ref()
const users = ref()

Promise.all([getAllUserAlias(),getAllActionGroups(),getAllActions()]).then(function (x) {
  users.value = x[0].data.result.items
  actionsStore.setActionGroups(x[1].data.result.items)
  actionsStore.setActions(x[2].data.result.items)
})

let rowData = ref({
  actionId:null,
  actionName:null,
  actionAlias:"",
  actionGroupId:"",
  url:"",
  method:"",
  createdAt:0,
  updatedAt:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  rowData.value = {
    actionId:null,
    actionName:null,
    actionAlias:"",
    actionGroupId:"",
    url:"",
    method:"",
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
  if(rowData.value.actionGroupId==0)
    rowData.value.actionGroupId = null
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleEdit(index,row) {
  dialogTitle.value = t('language.common.EDIT')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_EDIT
  const {...raw} = row
  rowData.value = raw
  if(rowData.value.actionGroupId==0)
    rowData.value.actionGroupId = null
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleDelete(index,row) {
  const {...raw} = row
  rowData.value = raw
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_DELETE

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.action.ACTION')+": "+rowData.value.actionAlias, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteAction(rowData.value.actionId).then(x=>{
          if(x.data.isSuccess){
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            actionsStore.actions = actionsStore.actions.filter(x=>x.actionId!=rowData.value.actionId)
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

function handleConfirm() {
  const { createdAt,updatedAt, ...newTable } = rowData.value;
  let fn = createAction
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateAction
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
        for(let i=0;i<actionsStore.actions.length;i++){
          if(actionsStore.actions[i].actionId==newTable.actionId){
            actionsStore.actions[i] = newStoreData
            break
          }
        }
      }else{
        const res = x.data.result
        delete res['@type']
        actionsStore.actions.push(res)
      }
      dialogVisible.value=false
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  }).catch(()=>{
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  })
}

const getActionGroupById = computed(()=>(actionGroupId) => {
  let l = actionsStore.actionGroups.filter(x=>x.actionGroupId == actionGroupId)
  if(l.length>0)
    return l[0].actionGroupAlias
  else
    return ""
})

const tableDataSetsWithPage = computed(() => {
  const start = (currentPage.value-1)*pageSize.value
  const end = (currentPage.value-1)*pageSize.value + pageSize.value
  return actionsStore.actions.slice( start,end )
})

const handleSizeChange = (itemCountPerPage) => {
  pageSize.value = itemCountPerPage
}

const handleCurrentChange = (changedCurrentPage) => {
  currentPage.value = changedCurrentPage
}

const changeParent = (currentValue) =>{
  const l = actionsStore.actionGroups.filter(x=>x.actionGroupId == currentValue)
  rowData.value.actionGroupId = l[0].actionGroupId
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
  actionId: [{ required: true,trigger: 'blur' }],
  actionName: [{ required: true,trigger: 'blur' }],
  actionAlias: [{ required: true,trigger: 'blur' }],
  actionGroupId: [{ required: true,trigger: 'blur' }],
  url: [{ required: true,trigger: 'blur' }],
  method: [{ required: true,trigger: 'blur' }]
}

</script>

<style scoped>

</style>

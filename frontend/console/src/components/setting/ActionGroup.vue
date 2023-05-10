<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-action-group'])">{{$t('language.common.CREATE')}}</el-button>
    <div >
      <el-table :data="tableDataSetsWithPage" :max-height="maxHeight">
        <el-table-column prop="actionGroupId" :label="$t('language.actionGroup.ACTION_GROUP_ID')" width="180" />
        <el-table-column prop="actionGroupName" :label="$t('language.actionGroup.ACTION_GROUP_NAME')" width="260" />
        <el-table-column prop="actionGroupAlias" :label="$t('language.actionGroup.ACTION_GROUP_ALIAS')" width="190" />
        <el-table-column :label="$t('language.actionGroup.ACTION_GROUP_PARENT')" width="180" >
          <template #default="scope">
            {{getActionGroupById(scope.row.parentActionGroupId)}}
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
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-action-group','update-action-group'])">{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-action-group'])">{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination v-model:current-page="currentPage" v-model:page-size="pageSize" :page-sizes="[5, 10, 15, 20]" :small="small" :disabled="disabled"
                     :background="background" layout="sizes, prev, pager, next" :total="actionsStore.actionGroups.length"
                     @size-change="handleSizeChange" @current-change="handleCurrentChange" style="margin-top: 3px;"/>
    </div>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="rowData" :rules="rules" ref="dialogForm">
        <el-form-item :label="$t('language.actionGroup.ACTION_GROUP_PARENT')" label-position="left" label-width="120px" prop="parentActionGroupId" required>
          <el-select v-model="rowData.parentActionGroupId" class="m-2" :placeholder="$t('language.common.SELECT')" @change="changeParent">
            <el-option v-for="item in actionsStore.actionGroups" :key="item.actionGroupId" :label="item.actionGroupAlias" :value="item.actionGroupId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.actionGroup.ACTION_GROUP_ID')" label-position="left" label-width="120px" prop="actionGroupId" required >
          <el-input v-model="rowData.actionGroupId"/>
        </el-form-item>
        <el-form-item :label="$t('language.actionGroup.ACTION_GROUP_NAME')" label-position="left" label-width="120px" prop="actionGroupName" required>
          <el-input v-model="rowData.actionGroupName"/>
        </el-form-item>
        <el-form-item :label="$t('language.actionGroup.ACTION_GROUP_ALIAS')" label-position="left" label-width="120px" prop="actionGroupAlias" required>
          <el-input v-model="rowData.actionGroupAlias"/>
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
import {storeToRefs} from "pinia"
import {useActionsStore} from "@/stores/action"
import {createActionGroup, deleteActionGroup, getAllActionGroups, updateActionGroup} from "@/api/action-group"
import {getAllUserAlias} from "@/api/user"

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

Promise.all([getAllUserAlias(),getAllActionGroups()]).then(function (x) {
  users.value = x[0].data.result.items
  actionsStore.setActionGroups(x[1].data.result.items)
})

let rowData = ref({
  actionGroupId:null,
  parentActionGroupId:null,
  parentActionGroupName:"",
  actionGroupName:"",
  actionGroupAlias:"",
  createdAt:0,
  updatedAt:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  rowData.value = {
    actionGroupId:null,
    parentActionGroupId:null,
    parentActionGroupName:"",
    actionGroupName:"",
    actionGroupAlias:"",
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
  if(rowData.value.parentActionGroupId==0)
    rowData.value.parentActionGroupId = null
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleEdit(index,row) {
  dialogTitle.value = t('language.common.EDIT')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_EDIT
  const {...raw} = row
  rowData.value = raw
  if(rowData.value.parentActionGroupId==0)
    rowData.value.parentActionGroupId = null
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleDelete(index,row) {
  const {...raw} = row
  rowData.value = raw
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_DELETE

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.actionGroup.ACTION_GROUP')+": "+rowData.value.actionGroupAlias, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteActionGroup(rowData.value.actionGroupId).then(x=>{
          if(x.data.isSuccess){
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            actionsStore.actionGroups = actionsStore.actionGroups.filter(x=>x.actionGroupId!=rowData.value.actionGroupId)
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

function handleConfirm() {
  const { parentActionGroupName,createdAt,updatedAt, ...newTable } = rowData.value;
  let fn = createActionGroup
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateActionGroup
  }
  fn(newTable).then(x=>{
    if(x.data.isSuccess){
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
        const { parentMenuIdName, ...newStoreData } = rowData.value
        for(let i=0;i<actionsStore.actionGroups.length;i++){
          if(actionsStore.actionGroups[i].actionGroupId==newTable.actionGroupId){
            actionsStore.actionGroups[i] = newStoreData
            break
          }
        }
      }else{
        const res = x.data.result
        delete res['@type']
        actionsStore.actionGroups.push(res)
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
  return actionsStore.actionGroups.slice( start,end )
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
  actionGroupId: [{ required: true,trigger: 'blur' }],
  parentActionGroupId: [{ required: true,trigger: 'blur' }],
  actionGroupName: [{ required: true,trigger: 'blur' }],
  actionGroupAlias: [{ required: true,trigger: 'blur' }]
}

</script>

<style scoped>

</style>

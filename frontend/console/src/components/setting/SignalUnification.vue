<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-signal-unification'])">{{$t('language.common.CREATE')}}</el-button>
    <div >
      <el-table :data="tableDataSets" :max-height="maxHeight">
        <el-table-column prop="signalUnificationId" :label="$t('language.signalUnification.SIGNAL_UNIFICATION_ID')" width="130" />
        <el-table-column prop="signalUnificationName" :label="$t('language.signalUnification.SIGNAL_UNIFICATION_NAME')" width="130" />
        <el-table-column prop="signalUnificationAlias" :label="$t('language.signalUnification.SIGNAL_UNIFICATION_ALIAS')" width="200" />
        <el-table-column prop="signalUnificationRelatedSignalId" :label="$t('language.signalUnification.SIGNAL_UNIFICATION_RELATED_SIGNAL_ID')" width="400" >
          <template #default="scope">
            {{getPathName(scope.row.signalUnificationRelatedSignalId)}}
          </template>
        </el-table-column>
        <el-table-column prop="signalUnificationNote" :label="$t('language.signalUnification.SIGNAL_UNIFICATION_NOTE')" width="250" />
        <el-table-column :label="$t('language.common.TIME')" width="100">
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
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-signal-unification','update-signal-unification'])">{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-signal-unification'])">{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="550px">
      <el-form :model="rowData" :rules="rules" ref="dialogForm">
        <el-form-item :label="$t('language.signalUnification.SIGNAL_UNIFICATION_ID')" label-position="left" label-width="150px" prop="signalUnificationId" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_VIEW" required >
          <el-input v-model="rowData.signalUnificationId" readonly />
        </el-form-item>
        <el-form-item :label="$t('language.signalUnification.SIGNAL_UNIFICATION_NAME')" label-position="left" label-width="150px" prop="signalUnificationName" required>
          <el-input v-model="rowData.signalUnificationName"/>
        </el-form-item>
        <el-form-item :label="$t('language.signalUnification.SIGNAL_UNIFICATION_ALIAS')" label-position="left" label-width="150px" prop="signalUnificationAlias" >
          <el-input v-model="rowData.signalUnificationAlias"/>
        </el-form-item>
        <el-form-item :label="$t('language.signalUnification.SIGNAL_UNIFICATION_RELATED_SIGNAL_ID')" label-position="left" label-width="150px" prop="signalUnificationRelatedSignalId" required>
          <el-input v-model="rowData.signalUnificationRelatedSignalId"/>
        </el-form-item>
        <el-form-item :label="$t('language.signalUnification.SIGNAL_UNIFICATION_NOTE')" label-position="left" label-width="150px" prop="signalUnificationNote" >
          <el-input v-model="rowData.signalUnificationNote"/>
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
import i18n from "@/lang";
import {nextTick, ref} from "vue";
import {getParentObjectPath, operatedTypeEnum} from "@/utils/common";
import {useActionsStore} from "@/stores/action";
import {getAllUserAlias} from "@/api/user";
import {
  createSignalUnification,
  deleteSignalUnification,
  getAllSignalUnifications,
  updateSignalUnification
} from "@/api/signal-unification";
import {ElMessage, ElMessageBox} from "element-plus";
import {useSignalsStore} from "@/stores/signal";
import {getAllLogicObjects} from "@/api/logic-object";
import {getAllSignals} from "@/api/signal";

const t = i18n.global.t
const dialogTitle = ref(t('language.common.CREATE'))
const dialogVisible = ref(false)
const operatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
const maxHeight = ref(window.innerHeight - 165)
const dialogForm = ref()
const users = ref()
const tableDataSets = ref([])
const actionsStore = useActionsStore()
const signalsStore = useSignalsStore()

Promise.all([getAllUserAlias(),getAllSignalUnifications(),getAllLogicObjects(),getAllSignals()]).then(function (x) {
  users.value = x[0].data.result.items
  tableDataSets.value = x[1].data.result.items
  signalsStore.setLogicObjects(x[2].data.result.items)
  signalsStore.setSignals(x[3].data.result.items)
})

let rowData = ref({
  signalUnificationId:null,
  signalUnificationName:"",
  signalUnificationAlias:"",
  signalUnificationRelatedSignalId:null,
  signalUnificationNote:"",
  createdAt:0,
  updatedAt:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  rowData.value = {
    signalUnificationId:null,
    signalUnificationName:"",
    signalUnificationAlias:"",
    signalUnificationRelatedSignalId:null,
    signalUnificationNote:"",
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
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleEdit(index,row) {
  dialogTitle.value = t('language.common.EDIT')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_EDIT
  const {...raw} = row
  rowData.value = raw
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleDelete(index,row) {
  const {...raw} = row
  rowData.value = raw
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_DELETE

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.signalUnification.SIGNAL_UNIFICATION')+": "+rowData.value.signalUnificationAlias, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteSignalUnification(rowData.value.signalUnificationId).then(x=>{
          if(x.data.isSuccess){
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            getSignalUnifications()
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

function handleConfirm() {
  const { createdAt,updatedAt, ...newTable } = rowData.value;
  let fn = createSignalUnification
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateSignalUnification
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
          if(tableDataSets.value[i].signalUnificationId==newTable.signalUnificationId){
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

function getSignalUnifications() {
  getAllSignalUnifications().then(x=>{
    if(x.data.isSuccess){
      tableDataSets.value = x.data.result.items
    }
  })
}

function getPathName(signalUnificationRelatedSignalId) {
  if(signalUnificationRelatedSignalId){
    const signal = signalsStore.findSignalById(signalUnificationRelatedSignalId)
    if(signal){
      return getParentObjectPath(signalUnificationRelatedSignalId,false)+"."+signal.signalName
    }
  }
  return ""
}

const rules = {
  signalUnificationName: [{ required: true,trigger: 'blur' }],
  signalUnificationRelatedSignalId: [{ required: true,trigger: 'blur' }]
}

</script>

<style scoped>

</style>
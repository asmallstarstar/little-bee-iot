<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-department'])">{{$t('language.common.CREATE')}}</el-button>
    <div >
      <el-table :data="tableDataSets" :max-height="maxHeight">
        <el-table-column prop="departmentId" :label="$t('language.department.DEPARTMENT_ID')" width="100" />
        <el-table-column prop="departmentName" :label="$t('language.department.DEPARTMENT_NAME')" width="220" />
        <el-table-column :label="$t('language.department.PARENT_DEPARTMENT')" width="220" >
          <template #default="scope">
            {{getDepartmentNameById(scope.row.parentDepartmentId)}}
          </template>
        </el-table-column>
        <el-table-column prop="positionAmongBrothers" :label="$t('language.department.POSITION_AMONG_BROTHERS')" width="100" />
        <el-table-column prop="leaderName" :label="$t('language.department.LEADER_NAME')" width="200" />
        <el-table-column prop="phone" :label="$t('language.department.PHONE')" width="150" />
        <el-table-column prop="email" :label="$t('language.department.EMAIL')" width="150" />
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
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-department','update-department'])">{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-department'])">{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="rowData" :rules="rules" ref="dialogForm">
        <el-form-item :label="$t('language.department.DEPARTMENT_ID')" label-position="left" label-width="120px" prop="departmentId" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_VIEW" required >
          <el-input v-model="rowData.departmentId" readonly />
        </el-form-item>
        <el-form-item :label="$t('language.department.PARENT_DEPARTMENT')" label-position="left" label-width="120px" prop="parentDepartmentId" >
          <el-select v-model="rowData.parentDepartmentId" class="m-2" :placeholder="$t('language.common.SELECT')" >
            <el-option v-for="item in tableDataSets" :key="item.departmentId" :label="item.departmentName" :value="item.departmentId"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.department.DEPARTMENT_NAME')" label-position="left" label-width="120px" prop="departmentName" required>
          <el-input v-model="rowData.departmentName"/>
        </el-form-item>
        <el-form-item :label="$t('language.department.POSITION_AMONG_BROTHERS')" label-position="left" label-width="120px" prop="positionAmongBrothers" >
          <el-input v-model="rowData.positionAmongBrothers"/>
        </el-form-item>
        <el-form-item :label="$t('language.department.LEADER_NAME')" label-position="left" label-width="120px" prop="leaderName" >
          <el-input v-model="rowData.leaderName"/>
        </el-form-item>
        <el-form-item :label="$t('language.department.PHONE')" label-position="left" label-width="120px" prop="phone" >
          <el-input v-model="rowData.phone"/>
        </el-form-item>
        <el-form-item :label="$t('language.department.EMAIL')" label-position="left" label-width="120px" prop="email" >
          <el-input v-model="rowData.email"/>
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
import {getAllUserAlias} from "@/api/user";
import {createDepartment, deleteDepartment, getAllDepartments, updateDepartment} from "@/api/department";
import {useActionsStore} from "@/stores/action";

const t = i18n.global.t
let dialogTitle = ref(t('language.common.CREATE'))
let dialogVisible = ref(false)
let operatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
let maxHeight = ref(window.innerHeight - 165)
const dialogForm = ref()
const users = ref()
const tableDataSets = ref([])
const actionsStore = useActionsStore()

Promise.all([getAllUserAlias(),getAllDepartments()]).then(function (x) {
  users.value = x[0].data.result.items
  tableDataSets.value = x[1].data.result.items
})

let rowData = ref({
  departmentId:null,
  parentDepartmentId:null,
  departmentName:"",
  positionAmongBrothers:null,
  leaderName:"",
  phone:"",
  email:"",
  createdAt:0,
  updatedAt:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  rowData.value = {
    departmentId:null,
    parentDepartmentId:null,
    departmentName:"",
    positionAmongBrothers:null,
    leaderName:"",
    phone:"",
    email:"",
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
  if(rowData.value.parentDepartmentId==0)
    rowData.value.parentDepartmentId = null
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleEdit(index,row) {
  dialogTitle.value = t('language.common.EDIT')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_EDIT
  const {...raw} = row
  rowData.value = raw
  if(rowData.value.parentDepartmentId==0)
    rowData.value.parentDepartmentId = null
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleDelete(index,row) {
  const {...raw} = row
  rowData.value = raw
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_DELETE

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.department.DEPARTMENT')+": "+rowData.value.departmentName, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteDepartment(rowData.value.departmentId).then(x=>{
          if(x.data.isSuccess){
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            getDepartments()
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

function handleConfirm() {
  const { createdAt,updatedAt, ...newTable } = rowData.value;
  let fn = createDepartment
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateDepartment
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
          if(tableDataSets.value[i].departmentId==newTable.departmentId){
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

const getDepartmentNameById = computed(()=>(departmentId) => {
  let l = tableDataSets.value.filter(x=>x.departmentId == departmentId)
  if(l.length>0)
    return l[0].departmentName
  else
    return ""
})

const getUserNickByUserId =  (userId) => {
  if(users && users.value){
    const l = users.value.filter(x=>x.userId == userId)
    if(l.length>0)
      return l[0].userNick
  }
  return ""
}

function getDepartments() {
  getAllDepartments().then(x=>{
    if(x.data.isSuccess){
      tableDataSets.value = x.data.result.items
    }
  })
}

const rules = {
  departmentId: [{ required: true,trigger: 'blur' }],
  departmentName: [{ required: true,trigger: 'blur' }],
}

</script>

<style scoped>

</style>

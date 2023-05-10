<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-role'])">{{$t('language.common.CREATE')}}</el-button>
    <el-table :data="rolesData" style="width: 100%">
      <el-table-column prop="roleId" :label="$t('language.role.ROLE_ID')" width="180" />
      <el-table-column prop="roleName" :label="$t('language.role.ROLE_NAME')" width="180" />
      <el-table-column prop="roleAlias" :label="$t('language.role.ROLE_ALIAS')" width="180" />
      <el-table-column prop="roleNote" :label="$t('language.role.ROLE_NOTE')" />
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
          <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-role','update-role'])">{{$t('language.common.EDIT')}}</el-button>
          <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-role'])">{{$t('language.common.DELETE')}}</el-button>
          <el-button size="small" @click="handleAuthorization(scope.$index, scope.row)" v-if=" actionsStore.isPermission(['all-action','all-action-group',
          'get-group-grant','grant-group-batch'])">{{$t('language.common.GRANT')}}</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog v-model="roleDialogVisible" :title="roleDialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="role">
        <el-form-item :label="$t('language.role.ROLE_ID')" label-position="left" label-width="80px" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_VIEW">
          <el-input v-model="role.roleId"  readonly/>
        </el-form-item>
        <el-form-item :label="$t('language.role.ROLE_NAME')" label-position="left" label-width="80px">
          <el-input v-model="role.roleName"  />
        </el-form-item>
        <el-form-item :label="$t('language.role.ROLE_ALIAS')" label-position="left" label-width="80px">
          <el-input v-model="role.roleAlias"  />
        </el-form-item>
        <el-form-item :label="$t('language.role.ROLE_NOTE')" label-position="left" label-width="80px">
          <el-input v-model="role.roleNote" type="textarea"/>
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="roleDialogVisible = false">{{$t('language.common.CANCEL')}}</el-button>
        <el-button type="primary" @click="handleConfirm" v-if="operatedType!=operatedTypeEnum.OPERATED_TYPE_VIEW">{{$t('language.common.CONFIRM')}}</el-button>
      </span>
      </template>
    </el-dialog>

    <el-dialog v-model="grantDialogVisible" :title="$t('language.common.GRANT')" :close-on-click-modal="false" align-center draggable width="350px">
      <div class="card">
        <h3 class="title">{{role.roleAlias}}</h3>
        <h3 class="sub-title">{{$t('language.menuAction.ACTION_PERMISS')}}</h3>
        <ActionTree class="action-tree" style="height: calc(65vh);" ref="actionTreeRef"></ActionTree>
        <el-button type="primary" @click="grant">{{$t('language.common.SUBMIT')}}</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script setup>
import {createRole, deleteRole, getAllRoles, getRoleGranted, grantBatchRole, updateRole} from "@/api/role"
import {nextTick, ref} from "vue"
import {operatedTypeEnum} from "@/utils/common"
import {ElMessage, ElMessageBox} from "element-plus"
import i18n from "@/lang"
import ActionTree from '@/components/setting/ActionTree.vue'
import {useActionsStore} from "@/stores/action";
import {getAllUserAlias} from "@/api/user";

const t = i18n.global.t
const actionsStore = useActionsStore()
let rolesData = ref([])
let operatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
let roleDialogVisible = ref(false)
let grantDialogVisible = ref(false)
const users = ref()
let role = ref({
  roleId:0,
  roleName:"",
  roleAlias:"",
  roleNote:"",
})
let roleDialogTitle = ref(t('language.common.CREATE'))
const actionTreeRef = ref()

Promise.all([getAllUserAlias(),getAllRoles()]).then(function (x) {
  users.value = x[0].data.result.items
  rolesData.value = x[1].data.result.items
})

function getRoles() {
  getAllRoles().then(x=>{
    if(x.data.isSuccess){
      rolesData.value = x.data.result.items
    }
  })
}

function handleCreate() {
  roleDialogTitle.value = t('language.common.CREATE')
  role.value = {
    roleId:0,
    roleName:"",
    roleAlias:"",
    roleNote:"",
  }
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  roleDialogVisible.value = true
}

function handleView(index,row) {
  roleDialogTitle.value = t('language.common.VIEW')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_VIEW
  roleDialogVisible.value = true
  const {...raw} = row
  role.value = raw
}

function handleEdit(index,row) {
  roleDialogTitle.value = t('language.common.EDIT')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_EDIT
  roleDialogVisible.value = true
  const {...raw} = row
  role.value = raw
}

function handleDelete(index,row) {
  const {...raw} = row
  role.value = raw
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_DELETE

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.role.ROLE')+":"+role.value.roleName, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteRole(role.value.roleId).then(x=>{
          if(x.data.isSuccess){
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            getRoles()
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

function handleAuthorization(index,row) {
  role.value = row
  grantDialogVisible.value = true
  nextTick(() => {
    actionTreeRef.value.setIsDisabled([], true)
  })
  getRoleGranted(role.value.roleId).then(x=>{
    let keys = []
    x.data.result.items.forEach(k=>{
      keys.push('a-'+k.actionId)
    })
    nextTick(() => {
      actionTreeRef.value.setCheckedNodes([])
      actionTreeRef.value.setCheckedNodes(keys)
    })
  })
}

function grant() {
  let body = {}
  body.items = []
  const actionList = actionTreeRef.value.getCheckedNodes()
  actionList.forEach(x=>{
    const ids = x.id.split("-")
    if(ids[0]==="a"){
      const actions = actionsStore.getActionByAcionId(ids[1])
      let item = {}
      item.groupId = role.value.roleId
      item.actionId = actions[0].actionId
      item.method = actions[0].method
      body.items.push(item)
    }
  })
  grantBatchRole(body).then(x=>{
    if(x.data.isSuccess){
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      grantDialogVisible.value = false
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  }).catch(()=>{
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  })
}

function handleConfirm() {
  roleDialogVisible.value=false
  let fn = createRole
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateRole
  }
  fn(role.value).then(x=>{
    if(x.data.isSuccess){
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      getRoles()
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

</script>

<style scoped>
.title{
  text-align: center;
}

.card > .el-button{
  display:block;
  margin:25px auto 0;
}
</style>

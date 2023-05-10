<template>
  <div>
    <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-user'])">{{$t('language.common.CREATE')}}</el-button>
    <el-table :data="usersData" style="width: 100%">
      <el-table-column prop="userId" :label="$t('language.user.USER_ID')" width="80" />
      <el-table-column prop="userName" :label="$t('language.user.USER_USERNAME')" width="120"/>
      <el-table-column prop="userNick" :label="$t('language.user.USER_NICK')" width="120"/>
      <el-table-column :label="$t('language.user.USER_DEPARTMENT')" width="130" >
        <template #default="scope">
          {{getDepartmentName(scope.row.departmentId)}}
        </template>
      </el-table-column>
      <el-table-column prop="phone" :label="$t('language.user.USER_PHONE')" width="100"/>
      <el-table-column prop="email" :label="$t('language.user.USER_EMAIL')" width="160"/>
      <el-table-column :label="$t('language.user.USER_SEX')" width="80">
        <template #default="scope">
          {{getGenderName(scope.row.sex)}}
        </template>
      </el-table-column>
      <el-table-column prop="status" :label="$t('language.user.USER_STATUS')" width="80">
        <template #default="scope">
          {{getAccountStatus(scope.row.status)}}
        </template>
      </el-table-column>
      <el-table-column prop="remark" :label="$t('language.user.USER_REMARK')" width="150"/>
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
          <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-user','update-user'])">{{$t('language.common.EDIT')}}</el-button>
          <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-user'])">{{$t('language.common.DELETE')}}</el-button>
          <el-button size="small" @click="handleAuthorization(scope.$index, scope.row)" v-if="scope.row.userName!='admin' &&
          actionsStore.isPermission(['all-action','all-action-group','grant-user-batch',
          'get-user-grant','get-join-group','join-group','remove-group'])" >{{$t('language.common.GRANT')}}</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="userDialogVisible" :title="userDialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="user" :rules="rules" ref="userForm">
        <el-form-item :label="$t('language.user.USER_ID')" label-position="left" label-width="135px" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_VIEW" >
          <el-input v-model="user.userId"  readonly/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_DEPARTMENT')" label-position="left" label-width="135px">
          <el-select v-model="user.departmentId" class="m-2" :placeholder="$t('language.common.SELECT')">
            <el-option
                v-for="item in departmentsData"
                :key="item.departmentId"
                :label="item.departmentName"
                :value="item.departmentId"
            />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_USERNAME')" label-position="left" label-width="135px" prop="userName" required>
          <el-input v-model="user.userName"/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_NICK')" label-position="left" label-width="135px" prop="userNick" required>
          <el-input v-model="user.userNick"/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_PASSWORD')" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_CREATE" label-position="left" label-width="135px" prop="password" required >
          <el-input v-model="user.password" type="password"/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_REPEAT_PASSWORD')" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_CREATE" label-position="left" label-width="135px" prop="repeatPassword" required>
          <el-input v-model="user.repeatPassword" type="password"/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_PHONE')" label-position="left" label-width="135px">
          <el-input v-model="user.phone"/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_EMAIL')" label-position="left" label-width="135px">
          <el-input v-model="user.email"/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_SEX')" label-position="left" label-width="135px">
          <el-select v-model="user.sex" class="m-2" :placeholder="$t('language.common.SELECT')">
            <el-option :label="$t('language.user.USER_MALE')" :value="genderTypeEnum.GENDER_TYPE_MALE"/>
            <el-option :label="$t('language.user.USER_FEMALE')" :value="genderTypeEnum.GENDER_TYPE_FEMALE"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_STATUS')" label-position="left" label-width="135px">
          <el-select v-model="user.status" class="m-2" :placeholder="$t('language.common.SELECT')">
            <el-option :label="$t('language.user.USER_NORMAL')" :value="userStatusEnum.USER_STATUS_NORMAL"/>
            <el-option :label="$t('language.user.USER_DEACTIVATED')" :value="userStatusEnum.USER_STATUS_DEACTIVATED"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_REMARK')" label-position="left" label-width="135px">
          <el-input v-model="user.remark" type="textarea"/>
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="userDialogVisible = false">{{$t('language.common.CANCEL')}}</el-button>
        <el-button type="primary" @click="handleConfirm" v-if="operatedType!=operatedTypeEnum.OPERATED_TYPE_VIEW">{{$t('language.common.CONFIRM')}}</el-button>
      </span>
      </template>
    </el-dialog>

    <el-dialog v-model="grantDialogVisible" :title="$t('language.common.GRANT')" :close-on-click-modal="false" align-center draggable width="350px">
      <div class="card">
        <h3 class="title">{{user.userNick}}</h3>
        <h3 class="sub-title">{{$t('language.role.ROLE')}}</h3>
        <div>
          <el-checkbox-group v-model="checkList">
            <el-checkbox v-for="item in rolesData" :key="item.roleId" :label="item.roleAlias"  @change="checked=> handleRoleChange(checked,item)"/>
          </el-checkbox-group>
        </div>
        <h3 class="sub-title">{{$t('language.menuAction.ACTION_PERMISS')}}</h3>
        <ActionTree class="action-tree" style="height: calc(65vh);" ref="actionTreeRef"></ActionTree>
        <el-button type="primary" @click="grant">{{$t('language.common.SUBMIT')}}</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script setup>
import {genderTypeEnum, operatedTypeEnum, userStatusEnum} from "@/utils/common";
import {computed, nextTick, ref} from "vue";
import {
  batchGrantUser,
  createUser, deleteUser,
  getAllUsers,
  getUseGrant,
  getUserGroup,
  joinGroup,
  removeGroup,
  updateUser
} from "@/api/user";
import {getAllDepartments} from "@/api/department";
import i18n from "@/lang";
import {getAllRoles, getRoleGranted} from "@/api/role";
import {ElMessage, ElMessageBox} from "element-plus";
import ActionTree from '@/components/setting/ActionTree.vue'
import {useActionsStore} from "@/stores/action";

const t = i18n.global.t
const actionsStore = useActionsStore()
const usersData = ref([])
const departmentsData = ref([])
const rolesData = ref([])
const grantDialogVisible = ref(false)
const userDialogVisible = ref(false)
const actionTreeRef = ref()
const userForm = ref()
const user = ref({
  userId:0,
  departmentId:null,
  departmentName:"",
  userName:"",
  userNick:"",
  password:"",
  repeatPassword:"",
  phone:"",
  email:"",
  sex:0,
  status:0,
  remark:""
})
const passwordSameCheck = async(rule, value, callback) => {
  if (value.length < 1) {
    return callback(new Error(t('language.user.REPEAT_PASSWORD_NOT_EMPTY')));
  } else if(user.value.password !== user.value.repeatPassword){
    return callback(new Error(t('language.user.REPEAT_PASSWORD_INCONSISTENT')));
  }else{
    callback()
  }
}

const rules = {
  userName: [{ required: true,trigger: 'blur' }],
  userNick: [{ required: true,trigger: 'blur' }],
  password: [{ required: true,trigger: 'blur' }],
  repeatPassword: [{ required: true,validator: passwordSameCheck, trigger: 'blur' }]
}

let userDialogTitle = ref(t('language.common.CREATE'))
let operatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
const checkList = ref([])

Promise.all([getAllDepartments(),getAllRoles(),getAllUsers()])
    .then(function (results) {
      departmentsData.value = results[0].data.result.items
      rolesData.value = results[1].data.result.items
      usersData.value = results[2].data.result.items
    })

function getUsers() {
  getAllUsers().then(x=>{
    if(x.data.isSuccess){
      usersData.value = x.data.result.items
    }
  })
}

function handleCreate() {
  userDialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE

  user.value = {
    userId:0,
    departmentId:null,
    departmentName:"",
    userName:"",
    userNick:"",
    password:"",
    repeatPassword:"",
    phone:"",
    email:"",
    sex:0,
    status:0,
    remark:""
  }
  userDialogVisible.value = true
}

function handleView(index,row) {
  userDialogTitle.value = t('language.common.VIEW')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_VIEW
  userDialogVisible.value = true
  const {...raw} = row
  user.value = raw
}

function handleEdit(index,row) {
  userDialogTitle.value = t('language.common.EDIT')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_EDIT
  userDialogVisible.value = true
  const {...raw} = row
  user.value = raw
}

function handleDelete(index,row) {
  const {...raw} = row
  user.value = raw
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_DELETE

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.user.USER')+":"+user.value.userNick, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteUser(user.value.userId).then(x=>{
          if(x.data.isSuccess){
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            getUsers()
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

function handleConfirm() {
  const { repeatPassword,departmentName, ...newUser } = user.value;
  userDialogVisible.value=false
  let fn = createUser
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateUser
  }
  fn(newUser).then(x=>{
    if(x.data.isSuccess){
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      getUsers()
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  }).catch(()=>{
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  })
}

async function handleAuthorization(index,row) {
  user.value = row
  grantDialogVisible.value = true
  nextTick(() => {
    actionTreeRef.value.setIsDisabled([], true)
  })

  const res = await getUserGroup(user.value.userId)
  checkList.value = rolesData.value.filter(x => {
    let sets = new Set(res.data.result.items.map(y=>y.groupId))
    return sets.has(x.roleId)
  }).map(x=>x.roleAlias)

  let keys = []
  for(let i=0;i<res.data.result.items.length;i++){
    const roleId = res.data.result.items[i].groupId
    const actions = await getRoleGranted(roleId)
    actions.data.result.items.forEach(k=>{
      keys.push('a-'+k.actionId)
    })
  }

  let moreKeys = []
  const moreActions = await getUseGrant(user.value.userId)
  if(moreActions.data.isSuccess){
    moreActions.data.result.items.forEach(x=>{
      moreKeys.push("a-"+x.actionId)
    })
  }

  const allCheckeds = [...keys,...moreKeys]
  nextTick(() => {
    actionTreeRef.value.setCheckedNodes(allCheckeds, true)
    actionTreeRef.value.setIsDisabled(keys,true)
  })

}

async function handleRoleChange(checked, item) {
  const res = await getRoleGranted(item.roleId)
  let keys = []
  res.data.result.items.forEach(k => {
    keys.push('a-' + k.actionId)
  })

  const body = {
    userId:user.value.userId,
    groupId:item.roleId
}
  if (checked) {
    const res1 = await joinGroup(body)
    if(res1.data.isSuccess){
      const checkeds = actionTreeRef.value.getCheckedKeys(true)
      const allCheckeds = [...keys, ...checkeds]
      nextTick(() => {
        actionTreeRef.value.setCheckedNodes(allCheckeds, true)
        actionTreeRef.value.setIsDisabled(keys,true)
      })
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  } else {
    const res1 = await removeGroup(body)
    if(res1.data.isSuccess){
      const checkeds = actionTreeRef.value.getCheckedKeys(true)
      const allCheckeds = checkeds.concat(keys).filter(e => checkeds.includes(e) && !keys.includes(e))
      nextTick(() => {
        actionTreeRef.value.setCheckedNodes(allCheckeds, true)
        actionTreeRef.value.setIsDisabled(keys,false)
      })
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  }
}

async function grant() {
  let keys = []
  const actionList = actionTreeRef.value.getCheckedNodes()
  actionList.forEach(x=>{
    const ids = x.id.split("-")
    if(ids[0]==="a"){
      if(!x.disabled) {
        keys.push(Number(ids[1]))
      }
    }
  })

  const actionsStore = useActionsStore()
  let body = {}
  body.items = []
  keys.forEach(x=>{
    const action = actionsStore.getActionByAcionId(x)
    const userPolicy = {
      userId:user.value.userId,
      actionId:action[0].actionId,
      method:action[0].method
    }
    body.items.push(userPolicy)
  })
  const res = await batchGrantUser(body)
  if(res.data.isSuccess){
    ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
    })
    grantDialogVisible.value = false
  }else{
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  }
}

const getDepartmentName = computed(()=>(departmentId) => {
  let l = departmentsData.value.filter(x=>x.departmentId == departmentId)
  if(l.length>0)
    return l[0].departmentName
  else
    return ""
})

const getUserNickByUserId =  (userId) => {
  if(usersData && usersData.value){
    let l = usersData.value.filter(x=>x.userId == userId)
    if(l.length>0)
      return l[0].userNick
    else
      return ""
  }
}

const getGenderName = computed(()=>(genderIndex) => {
  if(genderTypeEnum.GENDER_TYPE_MALE == genderIndex)
    return t('language.user.USER_MALE')
  if(genderTypeEnum.GENDER_TYPE_FEMALE == genderIndex)
    return t('language.user.USER_FEMALE')
  if(genderTypeEnum.GENDER_TYPE_UNKNOWN == genderIndex)
    return t('language.user.USER_UNKNOWN')
})

const getAccountStatus = computed(()=>(statusIndex) => {
  if(userStatusEnum.USER_STATUS_NORMAL == statusIndex)
    return t('language.user.USER_NORMAL')
  if(userStatusEnum.USER_STATUS_DEACTIVATED == statusIndex)
    return t('language.user.USER_DEACTIVATED')
})

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

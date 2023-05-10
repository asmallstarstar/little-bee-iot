<template>
  <div>
    <el-dialog v-model="userDialogVisible" :title="userDialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="userInfo" :rules="rules" ref="userForm">
        <el-form-item :label="$t('language.user.USER_ID')" label-position="left" label-width="135px">
          <el-input v-model="userInfo.userId"  readonly/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_DEPARTMENT')" label-position="left" label-width="135px">
          <el-select v-model="userInfo.departmentId" class="m-2" :placeholder="$t('language.common.SELECT')" disabled>
            <el-option
                v-for="item in departmentsData"
                :key="item.departmentId"
                :label="item.departmentName"
                :value="item.departmentId"
            />
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_USERNAME')" label-position="left" label-width="135px" prop="userName" required>
          <el-input v-model="userInfo.userName" readonly/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_NICK')" label-position="left" label-width="135px" prop="userNick" required>
          <el-input v-model="userInfo.userNick"/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_PHONE')" label-position="left" label-width="135px">
          <el-input v-model="userInfo.phone"/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_EMAIL')" label-position="left" label-width="135px">
          <el-input v-model="userInfo.email"/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_SEX')" label-position="left" label-width="135px">
          <el-select v-model="userInfo.sex" class="m-2" :placeholder="$t('language.common.SELECT')">
            <el-option :label="$t('language.user.USER_MALE')" :value="genderTypeEnum.GENDER_TYPE_MALE"/>
            <el-option :label="$t('language.user.USER_FEMALE')" :value="genderTypeEnum.GENDER_TYPE_FEMALE"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_STATUS')" label-position="left" label-width="135px">
          <el-select v-model="userInfo.status" class="m-2" :placeholder="$t('language.common.SELECT')" disabled>
            <el-option :label="$t('language.user.USER_NORMAL')" :value="userStatusEnum.USER_STATUS_NORMAL"/>
            <el-option :label="$t('language.user.USER_DEACTIVATED')" :value="userStatusEnum.USER_STATUS_DEACTIVATED"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_REMARK')" label-position="left" label-width="135px">
          <el-input v-model="userInfo.remark" type="textarea"/>
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="userDialogVisible=false">{{$t('language.common.CANCEL')}}</el-button>
        <el-button type="primary" @click="handleUserInfoConfirm">{{$t('language.common.CONFIRM')}}</el-button>
      </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {ref, toRaw, watch} from "vue";
import {useUsersStore} from "@/stores/user";
import i18n from "@/lang";
import {getAllDepartments} from "@/api/department";
import {updateUser} from "@/api/user";
import {ElMessage} from "element-plus";
import {userStatusEnum, genderTypeEnum} from "@/utils/common"

const t = i18n.global.t
const userDialogTitle = ref(t('language.common.EDIT'))
const userForm = ref()
const usersStore = useUsersStore()
const userDialogVisible = ref(false)

const userInfo = ref({
  userId:usersStore.userId,
  departmentId:usersStore.departmentId,
  userName:usersStore.userName,
  userNick:usersStore.userNick,
  phone:usersStore.phone,
  email:usersStore.email,
  sex:usersStore.sex,
  avatar:usersStore.avatar,
  status:usersStore.status,
  remark:usersStore.remark
})
const departmentsData = ref()

getAllDepartments().then(x=>{
  departmentsData.value = x.data.result.items
})

function handleUserInfoConfirm() {
  const {...rawData} = toRaw(userInfo.value)
  updateUser(rawData).then(x=>{
    if(x.data.isSuccess){
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      usersStore.setUserInfo(rawData,null)
      userDialogVisible.value = false
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  }).catch(()=>{
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  })
}

function showUserDialog() {
  userDialogVisible.value = true
}

defineExpose({
  showUserDialog
})

const rules = {
  userNick: [{ required: true,trigger: 'blur' }],
}

</script>

<style scoped>

</style>
<template>
  <div>
    <el-dialog v-model="userDialogVisible" :title="userDialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="userPassword" :rules="rules" ref="userForm">
        <el-form-item :label="$t('language.user.USER_OLD_PASSWORD')" label-position="left" label-width="135px" prop="oldPassword" required>
          <el-input type="password" v-model="userPassword.oldPassword"/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_NEW_PASSWORD')" label-position="left"  prop="newPassword" label-width="135px">
          <el-input type="password" v-model="userPassword.newPassword"/>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_REPEAT_NEW_PASSWORD')" label-position="left" prop="newRepeatPassword" label-width="135px">
          <el-input type="password" v-model="userPassword.newRepeatPassword"/>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="userDialogVisible=false">{{$t('language.common.CANCEL')}}</el-button>
          <el-button type="primary" @click="handleModifyPasswordConfirm">{{$t('language.common.CONFIRM')}}</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {ref, toRaw} from "vue";
import {useUsersStore} from "@/stores/user";
import i18n from "@/lang";
import {modifyPasword} from "@/api/user";
import {ElMessage} from "element-plus";

const t = i18n.global.t
const userDialogTitle = ref(t('language.common.EDIT'))
const userForm = ref()
const usersStore = useUsersStore()
const userDialogVisible = ref(false)

const userPassword = ref({
  userId:usersStore.userId,
  oldPassword:"",
  newPassword:"",
  newRepeatPassword:""
})

function handleModifyPasswordConfirm() {
  const {...rawData} = toRaw(userPassword.value)
  modifyPasword(rawData).then(x=>{
    if(x.data.isSuccess){
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      userDialogVisible.value = false
    }else{
      ElMessage.error(x.data.errorDesc)
    }
  }).catch(()=>{
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  })
}

function showModifyPasswordDialog() {
  userDialogVisible.value = true
}

defineExpose({
  showModifyPasswordDialog
})

const passwordSameCheck = async(rule, value, callback) => {
  if (value.length < 1) {
    return callback(new Error(t('language.user.REPEAT_PASSWORD_NOT_EMPTY')));
  } else if(userPassword.value.newPassword !== userPassword.value.newRepeatPassword){
    return callback(new Error(t('language.user.REPEAT_PASSWORD_INCONSISTENT')));
  }else{
    callback()
  }
}

const rules = {
  oldPassword: [{ required: true,trigger: 'blur' }],
  newPassword: [{ required: true,trigger: 'blur' }],
  newRepeatPassword: [{ required: true,validator: passwordSameCheck, trigger: 'blur' }]
}

</script>

<style scoped>

</style>
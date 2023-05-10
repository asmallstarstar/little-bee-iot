<template>
  <div>
    <el-table :data="realdataLists" style="width: 100%" :cell-style="{ textAlign: 'left' }" :header-cell-style="{ textAlign: 'left' }">
      <el-table-column prop="signalId" :label="$t('language.signal.SIGNAL_ID')" width="150"></el-table-column>
      <el-table-column prop="signalName" :label="$t('language.signal.SIGNAL_NAME')"> </el-table-column>
      <el-table-column prop="signalType" :label="$t('language.signal.SIGNAL_TYPE')" width="150">
        <template #default="scope">
          <span>{{getSignalType(scope.row.signalType)}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="signalValue" :label="$t('language.signal.SIGNAL_VALUE')" width="300">
        <template #default="scope">
          <el-button v-if="scope.row.signalType == signalTypeEnum.SIGNAL_TYPE_AO
          ||scope.row.signalType == signalTypeEnum.SIGNAL_TYPE_DO
          ||scope.row.signalType == signalTypeEnum.SIGNAL_TYPE_SO" type="primary" size="small" @click="controlClick(scope.row)">{{$t('language.signal.SIGNAL_CONTROL')}}</el-button>
          <span :class="signalValueClassName(scope.row.signalRunState,scope.row.signalValueState,new Date(Date.parse(scope.row.valueOccurredTime)))"
                v-else-if="scope.row.signalType==signalTypeEnum.SIGNAL_TYPE_AI || scope.row.signalType==signalTypeEnum.SIGNAL_TYPE_VAI">
            {{scope.row.signalValueState==signalValueStateEnum.SIGNAL_VALUE_STATE_VALID?scope.row.analogValue.toFixed(scope.row.decimalPrecision)+scope.row.valueUnit:"" }}</span>
          <span :class="signalValueClassName(scope.row.signalRunState,scope.row.signalValueState,new Date(Date.parse(scope.row.valueOccurredTime)))"
                v-else-if="scope.row.signalType==signalTypeEnum.SIGNAL_TYPE_DI || scope.row.signalType==signalTypeEnum.SIGNAL_TYPE_VDI">{{scope.row.digitalValueDesc }}</span>
          <span :class="signalValueClassName(scope.row.signalRunState,scope.row.signalValueState,new Date(Date.parse(scope.row.valueOccurredTime)))"
                v-else-if="scope.row.signalType==signalTypeEnum.SIGNAL_TYPE_SI || scope.row.signalType==signalTypeEnum.SIGNAL_TYPE_VSI">
            {{ scope.row.stringValue }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="valueOccurredTime" :label="$t('language.signal.UPDATE_TIME')" width="400">
        <template #default="scope">
          <span v-show="isShowTime(scope.row)">{{new Date(Date.parse(scope.row.valueOccurredTime)).toLocaleString() }}</span>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog v-model="dialogVisible" :title="$t('language.user.VERIFY_USER')" :close-on-click-modal="false" align-center draggable width="350px">
      <el-form :model="verifyForm">
        <el-form-item :label="$t('language.user.USER_USERNAME')" label-width="80px">
          <el-input v-model="verifyForm.userName" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item :label="$t('language.user.USER_PASSWORD')" label-width="80px">
          <el-input type="password" v-model="verifyForm.password" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">{{$t('language.common.CANCEL')}}</el-button>
            <el-button type="primary" @click="verifyClick">{{$t('language.common.CONFIRM')}}</el-button>
          </span>
      </template>
    </el-dialog>

    <el-dialog v-model="controlCommandFormVisible" :title="$t('language.controlCommand.CONTROL_COMMAND')" :close-on-click-modal="false" align-center draggable width="350px">
      <el-form :model="controlCommandForm">
        <el-form-item :label="$t('language.controlCommand.CONTROL_TYPE')" label-width="80px">
          <el-select v-model="controlCommandForm.controlCommandType" :placeholder="$t('language.controlCommand.SELECT_CONTROL_TYPE')" style="width: 230px" @change="changeCommandType">
            <el-option v-for="item in controlCommandTypeText" :key="item.index" :label="item.name" :value="item.index">{{ item.name }}</el-option>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.controlCommand.DELAY_IN_MILLISECONDS')" label-width="80px">
          <el-input v-model="controlCommandForm.delay" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item :label="$t('language.controlCommand.CONTROL_VALUE')" label-width="80px" v-show="isShowControlValueInput">
          <el-input v-model="controlCommandForm.controlValue" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
          <span class="dialog-footer">
            <el-button @click="controlCommandFormVisible = false">{{$t('language.common.CANCEL')}}</el-button>
            <el-button type="primary" @click="controlCommandClick">{{$t('language.common.CONFIRM')}}</el-button>
          </span>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>
import {onDeactivated, onUnmounted, ref, watch} from "vue";
import {useRoute} from "vue-router";
import {
  createConnectionWithoutPredicatesType,
  createFieldValue, createFilterWithoutConnectionType, createOrderBy,
  createPredicate, createQueryCommand,
  fieldValueTypeEnum, orderTypeEnum,
  predicateTypeEnum,
  pushFieldValueToPredicate
} from "@/utils/query";
import {querySignal} from "@/api/signal";
import {controlCommand, getRealdata} from "@/api/realdata";
import {
  controlCommandResultText,
  controlCommandTypeEnum,
  controlCommandTypeText,
  signalRunStatusEnum,
  signalTypeEnum,
  signalTypeText,
  signalValueStateEnum
} from "@/utils/common";
import {verifyUser} from "@/api/user";
import {ElMessage} from "element-plus";
const route = useRoute()
const realdataLists = ref([])
const dialogVisible  = ref(false)
const verifyForm =ref({
  userName: "",
  password: "",
})
const controlCommandFormVisible = ref(false)
const controlCommandForm = ref({
  signalId:0,
  controlCommandType: 0,
  delay: 0,
  controlValue: "",
})
const isShowControlValueInput = ref(false)
let maxHeight = ref(window.innerHeight - 165)
let ids = {
  signalIds:[]
}
let timerId = null
let user = null
let controlCommandSerialNumber = null

watch(() => route.params, (newValue, oldValue) => {
  if(newValue.nodeId){
    onRouteUpdate(newValue.nodeId)
  }
}, {immediate: true, deep: true})

function onRouteUpdate(nodeId) {
  const predicate = createPredicate(predicateTypeEnum.PREDICATE_TYPE_EQ,'parent_logic_id')
  pushFieldValueToPredicate(predicate,createFieldValue(fieldValueTypeEnum.FIELD_VALUE_TYPE_INTEGER,Number(nodeId)))
  let predicates = []
  predicates.push(predicate)
  let connection = createConnectionWithoutPredicatesType(predicates)
  let connects = []
  connects.push(connection)
  let filter = createFilterWithoutConnectionType(connects)
  let orderBy = createOrderBy('signal_id',orderTypeEnum.ORDER_TYPE_ASC)
  let order = []
  order.push(orderBy)
  const query = createQueryCommand(filter,order)
  querySignal(query).then(x=>{
    ids.signalIds = []
    realdataLists.value=[]
    x.data.result.items.forEach(k=>{
      if(k.signalType!=signalTypeEnum.SIGNAL_TYPE_AO && k.signalType!=signalTypeEnum.SIGNAL_TYPE_DO && k.signalType!=signalTypeEnum.SIGNAL_TYPE_SO){
        ids.signalIds.push(k.signalId)
      }
      realdataLists.value.push({
        signalId:k.signalId,
        signalName:k.signalName,
        signalType:k.signalType,
        signalValue:"",
        valueOccurredTime:new Date().toLocaleString(),
        analogValue:0,
      })
    })
    timerId = setInterval(() => {
      requestRealdata()
    }, 5000);
  })
}

function requestRealdata() {
  if(ids.signalIds.length>0){
    getRealdata(ids).then(x=>{
      if( x.data.result){
        x.data.result.items.forEach(y=>{
          for(let i=0;i<realdataLists.value.length;i++){
            if(realdataLists.value[i].signalId == y.signalId){
              realdataLists.value[i] = y
              break
            }
          }
        })
      }
    })
  }
}

function controlClick(row) {
  const {...raw} = row
  //console.log(raw)
  verifyForm.value = {
    userName: "",
    password: "",
  }
  dialogVisible.value = true
  controlCommandForm.value.signalId = row.signalId
}

function verifyClick() {
  verifyUser(verifyForm.value).then(x=>{
    if(x.data.isSuccess){
      dialogVisible.value = false
      controlCommandFormVisible.value = true
      user=x.data.result.user
    }else{
      user = null
      ElMessage({
        showClose: true,
        message: x.data.errorDesc,
        type: 'error',
      })
    }
  })
}

function controlCommandClick() {
  let param = {
    controlCommandSerialNumber:"",
    operatorUserId:user.userId,
    operatorUserName:user.userName,
    signalId:controlCommandForm.value.signalId,
    controlCommandType:controlCommandForm.value.controlCommandType,
    delay:parseInt(controlCommandForm.value.delay,10),
    digitalValue:Math.floor(controlCommandForm.value.controlValue),
    analogValue:Number(controlCommandForm.value.controlValue),
    stringValue:controlCommandForm.value.controlValue
  }
  controlCommand(param).then(x=>{
    //console.log(x)
    controlCommandSerialNumber = x.data.result.controlCommandSerialNumber
    const controlCommandResultType = x.data.result.controlCommandResultType
    controlCommandFormVisible.value = false
    ElMessage({
      showClose: true,
      message: controlCommandResultText[controlCommandResultType].name,
      type: 'success',
    })
  })
}

function changeCommandType() {
  if(controlCommandForm.value.controlCommandType==controlCommandTypeEnum.CONTROL_COMMAND_TYPE_ANALOG
      || controlCommandForm.value.controlCommandType==controlCommandTypeEnum.CONTROL_COMMAND_TYPE_STRING
      || controlCommandForm.value.controlCommandType==controlCommandTypeEnum.CONTROL_COMMAND_TYPE_STRIKE)
  {
    isShowControlValueInput.value = true
  }else{
    isShowControlValueInput.value = false
  }
}

function signalValueClassName(signalRunState,signalValueState,valueTime) {
  const now = new Date()
  if ( (now.getTime()-valueTime.getTime())>1000*60*3 ){//timeout
    return "unknown-state"
  }
  if(signalValueState==signalValueStateEnum.SIGNAL_VALUE_STATE_VALID){
    if (signalRunState == signalRunStatusEnum.SIGNAL_RUN_STATE_UNKNOWN)
      return "unknown-state"
    if (signalRunState == signalRunStatusEnum.SIGNAL_RUN_STATE_ALARM)
      return "alarm-state"
    if (signalRunState == signalRunStatusEnum.SIGNAL_RUN_STATE_NORMAL)
      return "normal-state"
  }
  return "unknown-state"
}

function isShowTime(row) {
  switch (row.signalType) {
    case signalTypeEnum.SIGNAL_TYPE_AO:
    case signalTypeEnum.SIGNAL_TYPE_DO:
    case signalTypeEnum.SIGNAL_TYPE_SO:
      return false;
    default:
      return true;
  }
}

function getSignalType(signalType) {
  for (const i in signalTypeText) {
    if(signalTypeText[i].index==signalType){
      return signalTypeText[i].name
    }
  }
  return signalTypeText[0].name
}

onDeactivated(()=>{
  if(timerId!=null){
    clearInterval(timerId)
  }
})

</script>

<style scoped>

.el-table .unknown-state {
  color: #4c4c4c;
}

.el-table .normal-state {
  color: #00ff00;
}

.el-table .alarm-state {
  color: #ff0000;
}

</style>
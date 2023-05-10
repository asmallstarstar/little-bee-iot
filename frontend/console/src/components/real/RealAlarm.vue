<template>
  <div>
    <el-table :data="tableDataSetsWithPage" :max-height="maxHeight">
<!--      <el-table-column prop="alarmSerialNumber" :label="$t('language.activeAlarm.ALARM_SN')" width="180" />-->
      <el-table-column prop="signalId" :label="$t('language.signal.SIGNAL_ID')" width="70" />
      <el-table-column prop="signalName" :label="$t('language.signal.SIGNAL_NAME')" width="90" />
      <el-table-column prop="signalType" :label="$t('language.signal.SIGNAL_TYPE')" width="60"/>
<!--      <el-table-column prop="signalUnificationId" :label="$t('language.signal.SIGNAL_UNIFICATION')" width="60"/>-->
      <el-table-column prop="alarmBeginTime" :label="$t('language.activeAlarm.BEGIN_TIME')" width="170">
        <template #default="scope">
          <span>{{new Date(Date.parse(scope.row.alarmBeginTime)).toLocaleString()}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="alarmEndTime" :label="$t('language.activeAlarm.END_TIME')" width="170">
        <template #default="scope">
          <span>{{scope.row.alarmEndTime?new Date(Date.parse(scope.row.alarmEndTime)).toLocaleString():""}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="alarmLasting" :label="$t('language.activeAlarm.LASTING_TIME')" width="100" />
      <el-table-column prop="deviceLocation" :label="$t('language.activeAlarm.DEVICE_POSITION')" width="250"/>
      <el-table-column prop="deviceName" :label="$t('language.logicObject.DEVICE_NAME')" width="150"/>
      <el-table-column prop="deviceTypeId" :label="$t('language.deviceType.DEVICE_TYPE')" width="100" >
        <template #default="scope">
          <span>{{getDeviceTypeName(scope.row.deviceTypeId)}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="alarmDesc" :label="$t('language.activeAlarm.ALARM_DESC')" width="100" />
      <el-table-column prop="alarmLevelNumber" :label="$t('language.activeAlarm.ALARM_LEVEL')" width="100">
        <template #default="scope">
          <span  :class="getAlarmLevelStyle(scope.row.alarmLevelNumber)">{{getAlarmLevelName(scope.row.alarmLevelNumber)}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="occurredDigitalValue" :label="$t('language.activeAlarm.DIGITAL_VALUE')" width="50"/>
      <el-table-column prop="occurredAnalogValue" :label="$t('language.activeAlarm.ANALOG_VALUE')" width="50"/>
      <el-table-column prop="occurredStringValue" :label="$t('language.activeAlarm.STRING_VALUE')" width="50"/>
      <el-table-column prop="ackTime" :label="$t('language.activeAlarm.ACK_TIME')" width="170">
        <template #default="scope">
          <span>{{scope.row.ackTime?new Date(Date.parse(scope.row.ackTime)).toLocaleString():""}}</span>
        </template>
      </el-table-column>
      <el-table-column prop="ackUserName" :label="$t('language.activeAlarm.ACK_USERNAME')" width="100"/>
      <el-table-column :label="$t('language.common.OPERATE')" width="100">
        <template #default="scope">
          <el-button size="small" @click="handleAck(scope.$index, scope.row)" v-if="actionsStore.isPermission(['ack-alarm'])">{{$t('language.activeAlarm.ACK')}}</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :page-sizes="[5, 10, 15, 20]"
        :small="small"
        :disabled="disabled"
        :background="background"
        layout="total,sizes, prev, pager, next"
        :total="tableDataSets.length"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        style="margin-top: 3px;"
    />
  </div>
</template>

<script setup>
import {computed, onActivated, onDeactivated, ref} from "vue";
import {
  connectionTypeEnum, createConnection,
  createFieldValue, createFilterWithoutConnectionType, createOrderBy,
  createPredicate, createQueryCommand,
  fieldValueTypeEnum, orderTypeEnum,
  predicateTypeEnum,
  pushFieldValueToPredicate
} from "@/utils/query";
import {ackAlarm, queryAlarm} from "@/api/alarm";
import {useUsersStore} from "@/stores/user";
import {getAllAlarmLevels} from "@/api/alarm-level";
import {getAllDeviceTypes} from "@/api/device-type";
import {useActionsStore} from "@/stores/action";
import {ElMessage} from "element-plus";
import i18n from "@/lang";
let websocket = null
let lastTimerId = null
let deleteAckAlarmTimerId = null
const t = i18n.global.t
const tableDataSets = ref([])
const maxHeight = ref(window.innerHeight - 165)
const currentPage = ref(1)
const pageSize = ref(15)
const small = ref(false)
const background = ref(false)
const disabled = ref(false)
const alarmLevels = ref()
const deviceTypes = ref()
const actionsStore = useActionsStore()

const tableDataSetsWithPage = computed(() => {
  const start = (currentPage.value-1)*pageSize.value
  const end = (currentPage.value-1)*pageSize.value + pageSize.value
  return tableDataSets.value.slice( start,end )
})

let rowData = ref({
  alarmSerialNumber:'',
  signalId:null,
  signalName:"",
  signalType:"",
  signalUnificationId:'',
  alarmBeginTime:null,
  alarmEndTime:null,
  deviceLocation:"",
  deviceName:"",
  deviceTypeId:null,
  alarmDesc:"",
  alarmLevelNumber:null,
  alarmThresholdNumber:null,
  signalValueType:null,
  occurredDigitalValue:null,
  occurredAnalogValue:null,
  valueUnit:"",
  occurredStringValue:"",
  ackState:null,
  ackTime:null,
  ackUserId:null,
  ackUserName:null,
  alarmLasting:"",
})

function handleAck(index,row) {
  const {...raw} = row
  delete raw.alarmLasting
  const usersStore = useUsersStore()
  raw.ackUserName = usersStore.userNick
  ackAlarm(raw).then(x=>{
    if(x.data.isSuccess){
    ElMessage({
      showClose: false,
      message: t('language.common.SUBMIT_SUCCESS'),
      type: 'success',
    })
  }
  })
}

const handleSizeChange = (itemCountPerPage) => {
  pageSize.value = itemCountPerPage
}
const handleCurrentChange = (changedCurrentPage) => {
  currentPage.value = changedCurrentPage
}

onActivated(() => {
  getAllAlarmLevels().then(x=>{
    alarmLevels.value = x.data.result.items
  })

  getAllDeviceTypes().then(x=>{
    deviceTypes.value = x.data.result.items
  })

  getAllActiveAlarm()

  lastTimerId = setInterval(() => {
    tableDataSets.value.map(x=>{
      x.alarmLasting = lastTimeFormat(Date.parse(x.alarmBeginTime),Date.parse(x.alarmEndTime))
    })
  }, 1000);

  deleteAckAlarmTimerId = setInterval(() => {
    for (const i in tableDataSets.value) {
      if(tableDataSets.value[i].alarmEndTime && tableDataSets.value[i].ackState==1){
        tableDataSets.value.splice(Number(i),1)
      }
    }
  }, 5000);

  websocket = new WebSocket(import.meta.env.VITE_WEBSOCKET)
  websocket.onopen=function (){
    console.log("websocket open")
  }
  websocket.onmessage = function(data){
    const msg = JSON.parse(data.data)
    if(msg.messageType==='BEGIN_ALARM'){
      beginAlarm(msg.message)
    }
    if(msg.messageType==='END_ALARM'){
      endAlarm(msg.message)
    }
    if(msg.messageType==='ACK_ALARM'){
      ackAlarmMessage(msg.message)
    }
  }
  websocket.onclose =function (){
    console.log("websocket close")
  }
  websocket.onerror = function (e){
    console.log("websocket error:"+e)
  }
})

onDeactivated(() => {
  console.log("alarm deactivated")
  websocket.close()
  clearInterval(lastTimerId)
  clearInterval(deleteAckAlarmTimerId)
})

function beginAlarm(beginAlarm) {
  delete beginAlarm['@type']
  tableDataSets.value = [beginAlarm, ...tableDataSets.value];
}

function endAlarm(endAlarm) {
  for (const i in tableDataSets.value) {
    if(tableDataSets.value[i].alarmSerialNumber === endAlarm.alarmSerialNumber) {
      tableDataSets.value[i].alarmEndTime = endAlarm.alarmEndTime
      tableDataSets.value[i].alarmDesc = endAlarm.alarmDesc
      break
    }
  }
}

function ackAlarmMessage(ackAlarm) {
  for (const i in tableDataSets.value) {
    if(tableDataSets.value[i].alarmSerialNumber === ackAlarm.alarmSerialNumber && tableDataSets.value[i].ackState==0) {
      tableDataSets.value[i].ackState = ackAlarm.ackState
      tableDataSets.value[i].ackTime = ackAlarm.ackTime
      tableDataSets.value[i].ackUserId = ackAlarm.ackUserId
      tableDataSets.value[i].ackUserName = ackAlarm.ackUserName
      break
    }
  }
}

function getAllActiveAlarm() {
  const predicate = createPredicate(predicateTypeEnum.PREDICATE_TYPE_IS_NULL,'alarm_end_time')

  const predicate1 = createPredicate(predicateTypeEnum.PREDICATE_TYPE_EQ,'ack_state')
  pushFieldValueToPredicate(predicate1,createFieldValue(fieldValueTypeEnum.FIELD_VALUE_TYPE_INTEGER,0))

  let predicates = []
  predicates.push(predicate)
  predicates.push(predicate1)

  let connectionType = []
  connectionType.push(connectionTypeEnum.CONNECT_TYPE_OR)

  let connection = createConnection(connectionType,predicates)
  let connects = []
  connects.push(connection)
  let filter = createFilterWithoutConnectionType(connects)
  let orderBy = createOrderBy('alarm_begin_time',orderTypeEnum.ORDER_TYPE_DESC)
  let order = []
  order.push(orderBy)
  const query = createQueryCommand(filter,order)
  //console.log(JSON.stringify(query))
  queryAlarm(query).then(x=>{
    tableDataSets.value = x.data.result.items
  })
}

function getAlarmLevelStyle(alarmLevelNumber){
  switch (alarmLevelNumber) {
    case 1:
      return "emergency-alarm"
    case 2:
      return "serious-alarm"
    case 3:
      return "normal-alarm"
    case 4:
      return "running-event"
  }
}

function getAlarmLevelName(alarmLevelNumber) {
  for (const i in alarmLevels.value) {
    if(alarmLevels.value[i].alarmLevelNumber == alarmLevelNumber){
      return alarmLevels.value[i].alarmLevelAlias
    }
  }
  return ""
}

function getDeviceTypeName(deviceTypeId) {
  for (const i in deviceTypes.value) {
    if(deviceTypes.value[i].deviceTypeId == deviceTypeId){
      return deviceTypes.value[i].deviceTypeName
    }
  }
  return ""
}

function lastTimeFormat(alarmBegin,alarmEnd){
  let current = new Date();
  if(alarmEnd){
    current  = alarmEnd;
  }
  let difftime = (current - alarmBegin)/1000;
  let days = parseInt(difftime/86400);
  let hours = parseInt(difftime/3600)-24*days;
  let minutes = parseInt(difftime%3600/60);
  let seconds = parseInt(difftime%60);
  let ds = ''+days+'d';
  let hs = ''+hours+'h';
  let ms = ''+minutes+'m';
  let ss = ''+seconds+'s';
  let str = '';
  if(days!=0)
    str+=ds;
  if(hours!=0)
    str+=hs;
  if(minutes!=0)
    str+=ms
  if(seconds!=0)
    str+=ss;
  return str;
}

</script>

<style scoped>

.el-table .running-event {
  color: #b1e80c;
}

.el-table .normal-alarm {
  color: #fda501;
}

.el-table .serious-alarm {
  color: #ff6a00;
}

.el-table .emergency-alarm {
  color: #ff0000;
}

</style>
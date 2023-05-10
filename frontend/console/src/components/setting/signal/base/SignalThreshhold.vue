<template>
  <div>
    <div >
      <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-signal-threshold'])">{{$t('language.common.CREATE')}}</el-button>
      <el-table :data="tableDataSets" >
        <el-table-column prop="thresholdId" :label="$t('language.threshold.THRESHOLD_ID')" width="80" />
        <el-table-column prop="thresholdNumber" :label="$t('language.threshold.THRESHOLD_NUMBER')" width="80" />
        <el-table-column prop="thresholdValueAi" :label="$t('language.threshold.THRESHOLD_AI_VALUE')" width="110"
                         v-if="signalObjectType===signalTypeEnum.SIGNAL_TYPE_AI || signalObjectType===signalTypeEnum.SIGNAL_TYPE_VAI"/>
        <el-table-column prop="thresholdValueDi" :label="$t('language.threshold.THRESHOLD_DI_VALUE')" width="110"
                         v-if="signalObjectType===signalTypeEnum.SIGNAL_TYPE_DI || signalObjectType===signalTypeEnum.SIGNAL_TYPE_VDI"/>
        <el-table-column prop="thresholdDirection" :label="$t('language.threshold.THRESHOLD_DIRECTION')" width="80" >
          <template #default="scope">
            {{thresholdDirectionText[scope.row.thresholdDirection].name}}
          </template>
        </el-table-column>
        <el-table-column prop="thresholdDeadArea" :label="$t('language.threshold.THRESHOLD_DEAD_AREA')" width="80"
                         v-if="signalObjectType===signalTypeEnum.SIGNAL_TYPE_AI || signalObjectType===signalTypeEnum.SIGNAL_TYPE_VAI"/>
        <el-table-column prop="isNotify" :label="$t('language.threshold.THRESHOLD_IS_NOTIFY')" width="80" >
          <template #default="scope">
            {{scope.row.isNotify?$t('language.common.YES'):$t('language.common.NO')}}
          </template>
        </el-table-column>
        <el-table-column prop="alarmLevelNumber" :label="$t('language.threshold.THRESHOLD_ALARM_LEVEL_NUMBER')" width="80" >
          <template #default="scope">
            {{getAlarmLevelName(scope.row.alarmLevelNumber)}}
          </template>
        </el-table-column>
        <el-table-column prop="upNotifyNote" :label="$t('language.threshold.THRESHOLD_UP_NOTIFY_NOTE')" width="80" />
        <el-table-column prop="downNotifyNote" :label="$t('language.threshold.THRESHOLD_DOWN_NOTIFY_NOTE')" width="80" />
        <el-table-column prop="alarmDelay" :label="$t('language.threshold.THRESHOLD_ALARM_DELAY')" width="80" />
        <el-table-column prop="recoverAlarmDelay" :label="$t('language.threshold.THRESHOLD_RECOVER_ALARM_DELAY')" width="110" />
        <el-table-column :label="$t('language.common.OPERATE')">
          <template #default="scope">
            <el-button size="small" @click="handleView(scope.$index, scope.row)">{{$t('language.common.VIEW')}}</el-button>
            <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-signal-threshold','update-signal-threshold'])">{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-signal-threshold'])">{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
      <el-form :model="rowData" :rules="rules" ref="dialogForm">
        <el-form-item :label="$t('language.threshold.THRESHOLD_ID')" label-position="left" label-width="120px" prop="thresholdId" v-if="thresholdOperatedType==operatedTypeEnum.OPERATED_TYPE_VIEW" required >
          <el-input v-model="rowData.thresholdId" readonly />
        </el-form-item>
        <el-form-item :label="$t('language.threshold.THRESHOLD_NUMBER')" label-position="left" label-width="120px" prop="thresholdNumber" required>
          <el-input v-model="rowData.thresholdNumber"/>
        </el-form-item>
        <el-form-item :label="$t('language.threshold.THRESHOLD_AI_VALUE')" label-position="left" label-width="120px" prop="thresholdValueAi"
                      v-if="signalObjectType===signalTypeEnum.SIGNAL_TYPE_AI || signalObjectType===signalTypeEnum.SIGNAL_TYPE_VAI">
          <el-input v-model="rowData.thresholdValueAi"/>
        </el-form-item>
        <el-form-item :label="$t('language.threshold.THRESHOLD_DI_VALUE')" label-position="left" label-width="120px" prop="thresholdValueDi"
                      v-if="signalObjectType===signalTypeEnum.SIGNAL_TYPE_DI || signalObjectType===signalTypeEnum.SIGNAL_TYPE_VDI">
          <el-input v-model="rowData.thresholdValueDi"/>
        </el-form-item>
        <el-form-item :label="$t('language.threshold.THRESHOLD_DIRECTION')" label-position="left" label-width="120px" prop="thresholdDirection" >
          <el-select v-model="rowData.thresholdDirection" class="m-2" :placeholder="$t('language.common.SELECT')">
            <el-option v-for="item in thresholdDirectionText" :key="item.index" :label="item.name" :value="item.index"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.threshold.THRESHOLD_DEAD_AREA')" label-position="left" label-width="120px" prop="thresholdDeadArea"
                      v-if="signalObjectType===signalTypeEnum.SIGNAL_TYPE_AI || signalObjectType===signalTypeEnum.SIGNAL_TYPE_VAI">
          <el-input v-model="rowData.thresholdDeadArea"/>
        </el-form-item>
        <el-form-item :label="$t('language.threshold.THRESHOLD_NOTIFY')" label-position="left" label-width="120px" prop="isNotify" >
          <el-checkbox v-model="rowData.isNotify" :label="$t('language.threshold.THRESHOLD_IS_NOTIFY')" />
        </el-form-item>
        <el-form-item :label="$t('language.threshold.THRESHOLD_ALARM_LEVEL_NUMBER')" label-position="left" label-width="120px" prop="alarmLevelNumber" required>
          <el-select v-model="rowData.alarmLevelNumber" class="m-2" :placeholder="$t('language.common.SELECT')">
            <el-option v-for="item in alarmLevels" :key="item.alarmLevelNumber" :label="item.alarmLevelAlias" :value="item.alarmLevelNumber"/>
          </el-select>
        </el-form-item>
        <el-form-item :label="$t('language.threshold.THRESHOLD_UP_NOTIFY_NOTE')" label-position="left" label-width="120px" prop="upNotifyNote" >
          <el-input v-model="rowData.upNotifyNote"/>
        </el-form-item>
        <el-form-item :label="$t('language.threshold.THRESHOLD_DOWN_NOTIFY_NOTE')" label-position="left" label-width="120px" prop="downNotifyNote" >
          <el-input v-model="rowData.downNotifyNote"/>
        </el-form-item>
        <el-form-item :label="$t('language.threshold.THRESHOLD_ALARM_DELAY')" label-position="left" label-width="120px" prop="alarmDelay" >
          <el-input v-model="rowData.alarmDelay"/>
        </el-form-item>
        <el-form-item :label="$t('language.threshold.THRESHOLD_RECOVER_ALARM_DELAY')" label-position="left" label-width="120px" prop="recoverAlarmDelay" >
          <el-input v-model="rowData.recoverAlarmDelay"/>
        </el-form-item>
      </el-form>
      <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">{{$t('language.common.CANCEL')}}</el-button>
            <el-button type="primary" @click="handleConfirm" v-if="thresholdOperatedType!=operatedTypeEnum.OPERATED_TYPE_VIEW">{{$t('language.common.CONFIRM')}}</el-button>
          </span>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>
import {nextTick, ref, toRaw, watch} from "vue";
import {operatedTypeEnum, signalTypeEnum, thresholdDirectionText} from "@/utils/common";
import i18n from "@/lang";
import {useActionsStore} from "@/stores/action";
import {getAllAlarmLevels} from "@/api/alarm-level";
import {ElMessage, ElMessageBox} from "element-plus";
import {getAllSignalThreshholds, querySignalThreshhold} from "@/api/signal-threshhold";
import {retrieveSignalAcquisition} from "@/api/signal-acquisition";
import {retrieveFsu} from "@/api/fsu";
import {
  createConnectionWithoutPredicatesType, createEqualQueryCommand,
  createFieldValue, createFilterWithoutConnectionType, createOrderBy,
  createPredicate, createQueryCommand,
  fieldValueTypeEnum, orderTypeEnum,
  predicateTypeEnum,
  pushFieldValueToPredicate
} from "@/utils/query";

const props = defineProps(['signalId','operatedType','signalObjectType'])
const t = i18n.global.t
const dialogTitle = ref(t('language.common.CREATE'))
const dialogVisible = ref(false)
const thresholdOperatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
const dialogForm = ref()
const users = ref()
const alarmLevels = ref()
const tableDataSets = ref([])
const actionsStore = useActionsStore()

getAllAlarmLevels().then(x=>{
  alarmLevels.value = x.data.result.items
})

watch(props,(newX)=>{
  if(newX.operatedType!==operatedTypeEnum.OPERATED_TYPE_CREATE){
    const query = createEqualQueryCommand(fieldValueTypeEnum.FIELD_VALUE_TYPE_INTEGER,props.signalId,"signal_id")
    querySignalThreshhold(query).then(x=>{
      tableDataSets.value = x.data.result.items.filter(k => k.signalId == props.signalId)
    })
  }else{
    tableDataSets.value = []
  }
},{immediate: true, deep: true})

const rowData = ref({
  thresholdId:null,
  signalId:props.signalId,
  signalType:props.signalObjectType,
  thresholdNumber:1,
  thresholdValueAi:0.0,
  thresholdValueDi:1,
  thresholdDirection:(props.signalObjectType===signalTypeEnum.SIGNAL_TYPE_DI || props.signalObjectType===signalTypeEnum.SIGNAL_TYPE_VDI)?3:1,
  thresholdDeadArea:0.0,
  isNotify:true,
  alarmLevelNumber:1,
  upNotifyNote:t('language.threshold.THRESHOLD_UP_NOTIFY_NOTE'),
  downNotifyNote:t('language.threshold.THRESHOLD_DOWN_NOTIFY_NOTE'),
  alarmDelay:0,
  recoverAlarmDelay:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  thresholdOperatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  rowData.value = {
    thresholdId:null,
    signalId:props.signalId,
    signalType:props.signalObjectType,
    thresholdNumber:1,
    thresholdValueAi:0.0,
    thresholdValueDi:1,
    thresholdDirection:(props.signalObjectType===signalTypeEnum.SIGNAL_TYPE_DI || props.signalObjectType===signalTypeEnum.SIGNAL_TYPE_VDI)?3:1,
    thresholdDeadArea:0.0,
    isNotify:true,
    alarmLevelNumber:1,
    upNotifyNote:t('language.threshold.THRESHOLD_UP_NOTIFY_NOTE'),
    downNotifyNote:t('language.threshold.THRESHOLD_DOWN_NOTIFY_NOTE'),
    alarmDelay:0,
    recoverAlarmDelay:0,
  }
  dialogVisible.value = true
}

function handleView(index,row) {
  dialogTitle.value = t('language.common.VIEW')
  thresholdOperatedType.value = operatedTypeEnum.OPERATED_TYPE_VIEW
  const {...raw} = row
  rowData.value = raw
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleEdit(index,row) {
  dialogTitle.value = t('language.common.EDIT')
  thresholdOperatedType.value = operatedTypeEnum.OPERATED_TYPE_EDIT
  const {...raw} = row
  rowData.value = raw
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleDelete(index,row) {
  ElMessageBox.confirm(t('language.threshold.THRESHOLD_DELETE'), t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        tableDataSets.value = tableDataSets.value.filter(x=>x.thresholdNumber!==row.thresholdNumber)
      })
}

function handleConfirm() {
  if(thresholdOperatedType.value == operatedTypeEnum.OPERATED_TYPE_CREATE){
    const l = tableDataSets.value.filter(x=>x.thresholdNumber===rowData.value.thresholdNumber)
    if(l.length>0){
      ElMessage.error(t('language.threshold.THRESHOLD_EXIST_SAME_NUMBER'))
    }else{
      tableDataSets.value.push(rowData.value)
      dialogVisible.value = false
    }
  }else{
    for(let i in tableDataSets.value){
      if( tableDataSets.value[i].thresholdNumber === rowData.value.thresholdNumber ){
        tableDataSets.value[i] = rowData.value
      }
    }
    dialogVisible.value = false
  }
}

function getAlarmLevelName(alarmLevelNumber) {
  if(alarmLevels && alarmLevels.value){
    const l = alarmLevels.value.filter(x=>x.alarmLevelNumber==alarmLevelNumber)
    if(l.length>0)
      return l[0].alarmLevelAlias
  }
  return ""
}

const rules = {
  thresholdNumber: [{ required: true,trigger: 'blur' }],
  alarmLevelName: [{ required: true,trigger: 'blur' }],
  alarmLevelAlias: [{ required: true,trigger: 'blur' }],
}

const getSignalThreshold = () => {
  const {...rawData} = toRaw(tableDataSets.value)
  let l = []
  for(let i in rawData){
    l.push(rawData[i])
  }
  return l
}

defineExpose({
  getSignalThreshold
})

</script>

<style scoped>

</style>
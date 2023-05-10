<template>
  <el-scrollbar :max-height="maxHeight">
    <el-collapse v-model="activeNames" >
      <el-collapse-item v-if="operatedType!=operatedTypeEnum.OPERATED_TYPE_CREATE" :title="$t('language.common.CREATE_OBJECT')" name="createObject">
        <div>
          <NewInstance :parent-object-type="logicObject.logicObjectTypeId"
                       :parent-logic-object-id="logicObject.logicObjectId">
          </NewInstance>
        </div>
      </el-collapse-item>
      <el-collapse-item :title="$t('language.logicObject.LOGIC_OBJECT_PROPERTY')" name="logicObjectProperty">
        <el-form :model="logicObject" :rules="rules" ref="formRef">
          <el-form-item :label="$t('language.logicObject.LOGIC_OBJECT_ID')" label-position="left" label-width="120px" prop="logicObjectId" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_EDIT" required >
            <el-input v-model="logicObject.logicObjectId" readonly />
          </el-form-item>
          <el-form-item :label="$t('language.logicObject.LOGIC_OBJECT_NAME')" label-position="left" label-width="120px" prop="logicObjectName" required>
            <el-input v-model="logicObject.logicObjectName" />
          </el-form-item>
          <el-form-item :label="$t('language.common.ALIAS')" label-position="left" label-width="120px" prop="logicObjectAlias">
            <el-input v-model="logicObject.logicObjectAlias" />
          </el-form-item>
          <el-form-item :label="$t('language.common.INDUSTRIAL_INTERNET_ID')" label-position="left" label-width="120px" prop="industrialInternetId">
            <el-input v-model="logicObject.industrialInternetId" />
          </el-form-item>
          <el-form-item :label="$t('language.common.PARENT_OBJECT')" label-position="left" label-width="120px" prop="parentLogicObjectId" required>
            <el-input v-model="logicObject.parentLogicObjectPath" readonly/>
          </el-form-item>
          <el-form-item :label="$t('language.logicObject.LOGIC_OBJECT_TYPE')" label-position="left" label-width="120px" prop="logicObjectTypeId" required>
            <el-select v-model="logicObject.logicObjectTypeId" class="m-2" :placeholder="$t('language.common.SELECT')">
              <el-option v-for="item in signalTreeObjectTypeText" :key="item.index" :label="item.name" :value="item.index"/>
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('language.common.POSITION_AMONG_BROTHERS')" label-position="left" label-width="120px" prop="positionAmongBrothers">
            <el-input v-model="logicObject.positionAmongBrothers"/>
          </el-form-item>
          <el-form-item :label="$t('language.common.METADATA')" label-position="left" label-width="120px" prop="metadataId" >
            <el-select v-model="logicObject.metadataId" class="m-2" :placeholder="$t('language.common.SELECT')" >
              <el-option v-for="item in metadatas" :key="item.metadataId" :label="item.metadataAlias" :value="item.metadataId"/>
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('language.common.METADATA_INSTANCE')" label-position="left" label-width="120px" prop="configureId" >
            <MetaInstance :metadata-id="logicObject.metadataId" :configure-id="0" :read-only="false" v-if="operatedType==operatedTypeEnum.OPERATED_TYPE_CREATE" ref="metaInstanceRef"></MetaInstance>
            <MetaInstance :metadata-id="logicObject.metadataId" :configure-id="logicObject.configureId"  :read-only="false" v-else ref="metaInstanceRef"></MetaInstance>
          </el-form-item>
          <el-form-item v-if="operatedType!=operatedTypeEnum.OPERATED_TYPE_CREATE" :label="$t('language.common.TIME')" label-position="left" label-width="120px">
            <div v-if="logicObject.createdAt">{{$t('language.common.CREATE_TIME')}}: {{ new Date(logicObject.createdAt).toLocaleString()}}</div>
            <div v-if="logicObject.createdBy">{{$t('language.common.CREATOR')}}:{{getUserNickByUserId(logicObject.createdBy) }}</div>
            <div v-if="logicObject.updatedAt">{{$t('language.common.UPDATE_TIME')}}: {{ new Date(logicObject.updatedAt).toLocaleString() }}</div>
            <div v-if="logicObject.updatedBy">{{$t('language.common.MODIFIER')}}:{{ getUserNickByUserId(logicObject.updatedBy) }}</div>
          </el-form-item>
        </el-form>
      </el-collapse-item>
      <el-collapse-item v-if="logicObject.logicObjectTypeId==signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_AREA" :title="$t('language.logicObject.AREA_PROPERTY')" name="AreaProperty">
        <div>
          <Area :logic-object-id="logicObject.logicObjectId" :operated-type="operatedType" ref="areaRef"></Area>
        </div>
      </el-collapse-item>
      <el-collapse-item v-if="logicObject.logicObjectTypeId==signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_DEVICE" :title="$t('language.logicObject.DEVICE_PROPERTY')" name="DeviceProperty">
        <div>
          <Device :logic-object-id="logicObject.logicObjectId" :operated-type="operatedType" ref="deviceRef" ></Device>
        </div>
      </el-collapse-item>
    </el-collapse>
    <div class="operate-group" v-if="logicObject.logicObjectId!=0">
      <el-button type="danger" size="small" @click="handleDelete" v-if="operatedType!=operatedTypeEnum.OPERATED_TYPE_CREATE && actionsStore.isPermission(['delete-logic-area','delete-logic-device'])">{{$t('language.common.DELETE')}}</el-button>
      <el-button type="primary" size="small" @click="handleConfirm" v-if="actionsStore.isPermission(['create-logic-area','update-logic-area','create-logic-device','update-logic-device'])">{{$t('language.common.CONFIRM')}}</el-button>
    </div>
  </el-scrollbar>
</template>

<script setup>
import NewInstance from "@/components/setting/logic/NewInstance.vue";
import Area from "@/components/setting/logic/Area.vue"
import Device from "@/components/setting/logic/Device.vue"
import MetaInstance from "@/components/setting/metadata/MetaInstance.vue"
import {
  getParentObjectPath,
  operatedTypeEnum,
  signalTreeObjectTypeEnum, signalTreeObjectTypeText
} from "@/utils/common";
import {ref, watch} from "vue";
import {createLogicObject, deleteLogicObject, retrieveLogicObject, updateLogicObject} from "@/api/logic-object";
import i18n from "@/lang";
import {getAllUserAlias} from "@/api/user";
import {getAllMetadatas} from "@/api/metadata";
import { getAllConfigures} from "@/api/configure";
import router from "@/router";
import {useRoute} from "vue-router";
import {ElMessage, ElMessageBox} from "element-plus";
import {useSignalsStore} from "@/stores/signal";
import {useTabsStore} from "@/stores/tab";
import {useActionsStore} from "@/stores/action";

let maxHeight = ref(window.innerHeight-120)
const actionsStore = useActionsStore()
const signalsStore = useSignalsStore()
const t = i18n.global.t
const route = useRoute()
const formRef = ref()
const activeNames = ref("logicObjectProperty")
const users = ref()
const metadatas = ref()
const configures = ref()
const operatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
const areaRef = ref()
const deviceRef = ref()
const metaInstanceRef = ref()
const  logicObject = ref({
  logicObjectId:-1,
  logicObjectName:"",
  logicObjectAlias:"",
  industrialInternetId:"",
  parentLogicObjectId:0,
  parentLogicObjectPath:"",
  logicObjectTypeId:0,
  positionAmongBrothers:0,
  metadataId:null,
  configureId:0,
  createdAt:0,
  createdBy:0,
  updatedAt:0,
  updatedBy:0
})

Promise.all([getAllUserAlias(),getAllMetadatas(),getAllConfigures()]).then(function (x) {
  users.value = x[0].data.result.items
  metadatas.value = x[1].data.result.items
  configures.value = x[2].data.result.items
})

const resetLogicObject = () =>{
  logicObject.value = {
    logicObjectId:-1,
    logicObjectName:"",
    parentLogicObjectId:0,
    parentLogicObjectPath:"",
    logicObjectTypeId:0,
    positionAmongBrothers:0,
    metadataId:null,
    configureId:0,
    createdAt:0,
    createdBy:0,
    updatedAt:0,
    updatedBy:0
  }
}

const setLogicObjectByZeroId = () =>{
  logicObject.value = {
    logicObjectId:0,
    logicObjectName:t('language.common.SURVEILLANCE_CENTER_SIGNAL'),
    parentLogicObjectId:0,
    parentLogicObjectPath:t('language.common.SURVEILLANCE_CENTER_SIGNAL'),
    logicObjectTypeId:signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_CENTER,
    positionAmongBrothers:0,
    metadataId:null,
    configureId:0,
    createdAt:0,
    createdBy:0,
    updatedAt:0,
    updatedBy:0
  }
}

const setLogicObjectById = (logicObjectId) =>{
  retrieveLogicObject(logicObjectId).then(async x => {
    delete x.data.result["@type"]
    logicObject.value = x.data.result
    if (logicObject.value.metadataId == 0) {
      logicObject.value.metadataId = null
    }
    logicObject.value.parentLogicObjectPath =  getParentLogicObjectPath(logicObject.value.parentLogicObjectId)
  })
}

async function onRouteUpdate(params) {
  const paramLogicObjectId = Number(params.logicObjectId)
  const paramsOperatedType = Number(params.operatedType)
  const paramsObjectType = Number(params.objectType)

  operatedType.value = paramsOperatedType
  resetLogicObject()
  if (paramsOperatedType == operatedTypeEnum.OPERATED_TYPE_CREATE) {
    logicObject.value.parentLogicObjectId = paramLogicObjectId
    logicObject.value.logicObjectTypeId = paramsObjectType
    logicObject.value.parentLogicObjectPath = getParentLogicObjectPath(logicObject.value.parentLogicObjectId)
  } else {
    logicObject.value.logicObjectId = paramLogicObjectId
    if (paramLogicObjectId == 0) {
      setLogicObjectByZeroId()
    } else {
      setLogicObjectById(paramLogicObjectId)
    }
  }
}

watch(() => route.params, (newValue, oldValue) => {
  if(newValue.logicObjectId){
    onRouteUpdate(newValue)
  }
}, {immediate: true, deep: true})

function handleDelete() {
  const childLogicObjects = signalsStore.findChildLogicObjectByLogicObjectId(logicObject.value.logicObjectId)
  const childSignals = signalsStore.findChildSignalByLogicObjectId(logicObject.value.logicObjectId)
  if(childLogicObjects.length>0 || childSignals.length>0){
    ElMessage.warning(t('language.common.EXIST_CHILD'))
  }else{
    ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.logicObject.LOGIC_OBJECT')+": "+logicObject.value.logicObjectName, t('language.common.DELETE_WARNING'), {
      confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
        .then(() => {
          deleteLogicObject(logicObject.value.logicObjectId).then(x=>{
            if(x.data.isSuccess){
              ElMessage({
                showClose: false,
                message: t('language.common.DELETE_SUCCESS'),
                type: 'success',
              })
              signalsStore.logicObjects = signalsStore.logicObjects.filter(x=>x.logicObjectId!=logicObject.value.logicObjectId)
              const tabsStore = useTabsStore()
              tabsStore.removeTab(route.path)
            }else{
              ElMessage.error(t('language.common.DELETE_FAI'))
            }
          })
        })
  }
}

function handleConfirm() {
  const { parentLogicObjectPath,createdAt,updatedAt,createdBy,updatedBy, ...newData } = logicObject.value;
  let obj = {}
  obj.logicObject = newData
  if(logicObject.value.metadataId){
    const metadataAttributeJson = JSON.stringify(metaInstanceRef.value.getMetadataAttribute())
    obj.logicConfigure = {
      configureId:0,
      metadataId:logicObject.value.metadataId,
      configureJson:metadataAttributeJson
    }
  }
  if(logicObject.value.logicObjectTypeId==signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_AREA){
    obj.logicArea = areaRef.value.getArea()
  }
  if(logicObject.value.logicObjectTypeId==signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_DEVICE){
    obj.logicDevice = deviceRef.value.getDevice()
  }
  let fn = createLogicObject
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateLogicObject
  }
  fn(obj).then(x=>{
    if(x.data.isSuccess){
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
        const {parentLogicObjectPath, ...data } = logicObject.value
        for(let i in signalsStore.logicObjects){
          if(signalsStore.logicObjects[i].logicObjectId==data.logicObjectId){
            signalsStore.logicObjects[i] = data
            break
          }
        }
      }else{
        const res = x.data.result
        delete res['@type']
        signalsStore.logicObjects.push(res.logicObject)
      }
      if(operatedType.value==operatedTypeEnum.OPERATED_TYPE_CREATE){
        const tabsStore = useTabsStore()
        tabsStore.removeTab(route.path)
        router.replace({ name: 'logicObject',
          params: {
            logicObjectId: x.data.result.logicObject.logicObjectId,
            operatedType:operatedTypeEnum.OPERATED_TYPE_EDIT,
            objectType:x.data.result.logicObject.logicObjectTypeId
          }
        })
      }
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
  }).catch(()=>{
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  })
}

const getParentLogicObjectPath = (parentLogicObjectId) => {
  if(parentLogicObjectId==0){
    return t('language.common.SURVEILLANCE_CENTER_SIGNAL')
  }else{
    return getParentObjectPath(parentLogicObjectId,true)
  }
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
  logicObjectName: [{ required: true,trigger: 'blur' }],
  parentLogicObjectId: [{ required: true,trigger: 'blur' }],
  logicObjectTypeId: [{ required: true,trigger: 'blur' }],
}

</script>

<style scoped>
  .operate-group{
    margin: 10px 10px;
  }
  .el-input ,.el-select {
    width: 500px;
  }
</style>
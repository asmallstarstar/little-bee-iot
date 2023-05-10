import i18n from '@/lang/index'
import {getLogicObjectPath, getSignalPath} from "@/api/logic-object";
import {useSignalsStore} from "@/stores/signal";
const t = i18n.global.t

export const SURVEILLANCE_CENTER_ID = 0

export const operatedTypeEnum = {
    OPERATED_TYPE_UNKNOWN: 0,
    OPERATED_TYPE_CREATE: 1,
    OPERATED_TYPE_EDIT: 2,
    OPERATED_TYPE_VIEW: 3,
    OPERATED_TYPE_DELETE: 4,
}

export const genderTypeEnum = {
    GENDER_TYPE_MALE: 0,
    GENDER_TYPE_FEMALE: 1,
    GENDER_TYPE_UNKNOWN: 2,
}

export const userStatusEnum = {
    USER_STATUS_NORMAL: 0,
    USER_STATUS_DEACTIVATED: 1,
}

export const signalTreeObjectTypeEnum = {
    SIGNAL_TREE_OBJECT_TYPE_UNKNOWN : 0,
    SIGNAL_TREE_OBJECT_TYPE_CENTER :  1,
    SIGNAL_TREE_OBJECT_TYPE_AREA:     2,
    SIGNAL_TREE_OBJECT_TYPE_DEVICE:   3,
    SIGNAL_TREE_OBJECT_TYPE_SIGNAL:   4
}

export const signalTreeObjectTypeText = [
    {index:0,name:t('language.signalTreeObjectType.UNKNOWN')},
    {index:1,name:t('language.signalTreeObjectType.CENTER')},
    {index:2,name:t('language.signalTreeObjectType.AREA')},
    {index:3,name:t('language.signalTreeObjectType.DEVICE')},
    {index:4,name:t('language.signalTreeObjectType.SIGNAL')},
]

export const signalTypeEnum = {
    SIGNAL_TYPE_UNKNOWN : 0,
    SIGNAL_TYPE_AI      : 1,
    SIGNAL_TYPE_DI      : 2,
    SIGNAL_TYPE_SI      : 3,
    SIGNAL_TYPE_AO      : 4,
    SIGNAL_TYPE_DO      : 5,
    SIGNAL_TYPE_SO      : 6,
    SIGNAL_TYPE_VAI     : 7,
    SIGNAL_TYPE_VDI     : 8,
    SIGNAL_TYPE_VSI     : 9,
    SIGNAL_TYPE_VIDEO   : 10,
}

export const signalTypeText = [
    {index:0,name:t('language.signalType.UNKNOWN')},
    {index:1,name:t('language.signalType.AI')},
    {index:2,name:t('language.signalType.DI')},
    {index:3,name:t('language.signalType.SI')},
    {index:4,name:t('language.signalType.AO')},
    {index:5,name:t('language.signalType.DO')},
    {index:6,name:t('language.signalType.SO')},
    {index:7,name:t('language.signalType.VAI')},
    {index:8,name:t('language.signalType.VDI')},
    {index:9,name:t('language.signalType.VSI')},
    {index:10,name:t('language.signalType.VIDEO')},
]

export const signalRunStatusEnum = {
    SIGNAL_RUN_STATE_UNKNOWN     : 0,
    SIGNAL_RUN_STATE_NORMAL      : 1,
    SIGNAL_RUN_STATE_ALARM       : 2,
}

export const signalValueStateEnum = {
    SIGNAL_VALUE_STATE_UNKNOWN      : 0,
    SIGNAL_VALUE_STATE_INVALID      : 1,
    SIGNAL_VALUE_STATE_VALID        : 2,
}

export const controlCommandTypeEnum = {
    CONTROL_COMMAND_TYPE_UNKNOWN     : 0,
    CONTROL_COMMAND_TYPE_TURN_ON     : 1,
    CONTROL_COMMAND_TYPE_TURN_OFF    : 2,
    CONTROL_COMMAND_TYPE_STRIKE      : 3,
    CONTROL_COMMAND_TYPE_ANALOG      : 4,
    CONTROL_COMMAND_TYPE_STRING      : 5,
}

export const controlCommandTypeText = [
    {index:0,name:t('language.controlCommand.CONTROL_COMMAND_TYPE_UNKNOWN')},
    {index:1,name:t('language.controlCommand.CONTROL_COMMAND_TYPE_TURN_ON')},
    {index:2,name:t('language.controlCommand.CONTROL_COMMAND_TYPE_TURN_OFF')},
    {index:3,name:t('language.controlCommand.CONTROL_COMMAND_TYPE_STRIKE')},
    {index:4,name:t('language.controlCommand.CONTROL_COMMAND_TYPE_ANALOG')},
    {index:5,name:t('language.controlCommand.CONTROL_COMMAND_TYPE_STRING')},
]



export const controlCommandResultEnum = {
    CONTROL_COMMAND_RESULT_UNKNOWN     : 0,
    CONTROL_COMMAND_RESULT_RECEIVED    : 1,
    CONTROL_COMMAND_RESULT_DELIVERING  : 2,
    CONTROL_COMMAND_RESULT_RUNNING     : 3,
    CONTROL_COMMAND_RESULT_SUCCESS     : 4,
    CONTROL_COMMAND_RESULT_FAIL        : 5,
    CONTROL_COMMAND_RESULT_TIMEOUT     : 6,
}

export const controlCommandResultText = [
    {index:0,name:t('language.controlCommand.CONTROL_COMMAND_RESULT_UNKNOWN')},
    {index:1,name:t('language.controlCommand.CONTROL_COMMAND_RESULT_RECEIVED')},
    {index:2,name:t('language.controlCommand.CONTROL_COMMAND_RESULT_DELIVERING')},
    {index:3,name:t('language.controlCommand.CONTROL_COMMAND_RESULT_RUNNING')},
    {index:4,name:t('language.controlCommand.CONTROL_COMMAND_RESULT_SUCCESS')},
    {index:5,name:t('language.controlCommand.CONTROL_COMMAND_RESULT_FAIL')},
    {index:6,name:t('language.controlCommand.CONTROL_COMMAND_RESULT_TIMEOUT')},
]

export function getParentObjectPath(objectId,isLogicObject) {
    const signalsStore = useSignalsStore()
    let path = ""
    if(isLogicObject){
        path =  signalsStore.getLogicObjectParentPath(objectId)
    }else{
        path =   signalsStore.getSignalParentPath(objectId)
    }
    if(path===""){
        return  ""
    }
    let res = ""
    const arrPath = path.split(".")
    for(let i in arrPath){
        const id = Number(arrPath[i])
        if(id==0){
            res = signalsStore.surveillanceCenter.label
        }else {
            const logicObject = signalsStore.findLogicObjectById(id)
            const logicObjectName = logicObject == null ? "" : logicObject.logicObjectName
            res = res + "." + logicObjectName
        }
    }
    return res
}

export const fsuDataParsingTypeEnum = {
    FSU_DATA_PARSING_TYPE_UNKNOWN     : 0,
    FSU_DATA_PARSING_TYPE_BOTTOM      : 1,
    FSU_DATA_PARSING_TYPE_CENTER      : 2,
}

export const thresholdDirectionEnum = {
    THRESHOLD_DIRECTION_UNKNOWN  : 0,
    THRESHOLD_DIRECTION_LESS     : 1,
    THRESHOLD_DIRECTION_GREAT    : 2,
    THRESHOLD_DIRECTION_EQUAL    : 3,
}

export const thresholdDirectionText = [
    {index:0,name:t('language.common.UNKNOWN')},
    {index:1,name:t('language.threshold.THRESHOLD_DIRECTION_LESS')},
    {index:2,name:t('language.threshold.THRESHOLD_DIRECTION_GREAT')},
    {index:3,name:t('language.threshold.THRESHOLD_DIRECTION_EQUAL')},
]

export const videoPlayTypeEnum = {
    VIDEO_PAY_TYPE_UNKNOWN : 0,
    VIDEO_PAY_TYPE_HTTP_M3U8 : 1,
    VIDEO_PAY_TYPE_RTSP:     2,
}

export const videoPlayTypeText = [
    {index:0,name:t('language.signalType.UNKNOWN')},
    {index:1,name:"Http m3u8"},
    {index:2,name:"RTSP"}
]
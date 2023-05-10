import { defineStore } from 'pinia'
import {operatedTypeEnum, signalTreeObjectTypeEnum, signalTypeEnum, SURVEILLANCE_CENTER_ID} from "@/utils/common";
import i18n from '@/lang/index'

export const useSignalsStore = defineStore('signals', {
    state: () => {
        return {
            surveillanceCenter:{
                id:SURVEILLANCE_CENTER_ID,
                label:i18n.global.t('language.common.SURVEILLANCE_CENTER_SIGNAL'),
                type:signalTreeObjectTypeEnum.SIGNAL_TREE_OBJECT_TYPE_CENTER,
                signalType:null,
                children:[],
                data:{},
            },
            logicObjects:[],
            signals:[]
        }
    },
    getters:{

    },
    actions: {
        setLogicObjects(logicObjects){
            this.logicObjects = logicObjects;
        },
        setSignals(signals){
            this.signals = signals
        },
        resetSurveillanceCenter(){
            this.surveillanceCenter.children = []
        },
        findLogicObjectById(logicObjectId){
            const list = this.logicObjects.filter(x=>x.logicObjectId==logicObjectId)
            return list.length>0?list[0]:null
        },
        findSignalById(signalId){
            const list = this.signals.filter(x=>x.signalId==signalId)
            return list.length>0?list[0]:null
        },
        findParentLogicObjectBySignalId(signalId) {
            const list = this.signals.filter(x=>x.signalId==signalId)
            if(list.length==0){
                return null
            }
            return this.findLogicObjectById(list[0].parentLogicId)
        },
        findChildLogicObjectByLogicObjectId(logicObjectId){
            return this.logicObjects.filter(x=>x.parentLogicObjectId==logicObjectId)
        },
        findChildSignalByLogicObjectId(logicObjectId){
            return this.signals.filter(x=>x.parentLogicId == logicObjectId)
        },
        getSignalParentPath(signalId){
            const list = this.signals.filter(x=>x.signalId==signalId)
            if(list.length==0){
                return ""
            }else{
                let p = this.getLogicObjectParentPath(list[0].parentLogicId)
                return p
            }
        },
        getLogicObjectParentPath(logicObjectId){
            let strId = logicObjectId+""
            let id = logicObjectId
            do{
                const l1 = this.logicObjects.filter(x=>x.logicObjectId==id)
                if(l1.length==0){
                    break
                }else{
                    id = l1[0].parentLogicObjectId
                    strId = id+"."+strId
                }

            }while (id!=0)
            return strId
        }
    },
    persist: {
        enabled: true,
        strategies: [
            { storage: sessionStorage, paths: ['logicObjects','signals'] }
        ],
    }
})

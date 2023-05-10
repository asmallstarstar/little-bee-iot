import { defineStore } from 'pinia'
import {SURVEILLANCE_CENTER_ID} from "@/utils/common";
import {agentObjectTypeEnum} from "@/utils/agent";
import i18n from '@/lang/index'

export const useAgentsStore = defineStore('agents', {
    state: () => {
        return {
            surveillanceCenter:{
                id:SURVEILLANCE_CENTER_ID,
                label:i18n.global.t('language.common.SURVEILLANCE_CENTER_AGENT'),
                type:agentObjectTypeEnum.AGENT_OBJECT_TYPE_SURVEILLANCE_CENTER,
                children:[],
                data:{},
            },
            agents:[],
            fsus:[],
            drivers:[],
        }
    },
    getters:{

    },
    actions: {
        setAgents(agents){
            this.agents = agents;
        },
        setFsus(fsus){
            this.fsus = fsus
        },
        setDrivers(drivers){
            this.drivers = drivers
        },
        resetSurveillanceCenter(){
            this.surveillanceCenter.children = []
        },
        findAgentById(agentId){
            const n = this.agents.filter(x=>x.agentId==agentId)
            return n.length>0?n[0]:null
        },
        findFsuById(fsuId){
            const n = this.fsus.filter(x=>x.fsuId==fsuId)
            return n.length>0?n[0]:null
        }
    }
})

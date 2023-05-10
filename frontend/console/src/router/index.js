import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '@/views/LoginView.vue'
import MainView from '@/views/MainView.vue'
import Home from "@/components/Home.vue"
import FsuTree from "@/components/layout/side/FsuTree.vue";
import Agent from "@/components/setting/agent/Agent.vue";
import Driver from "@/components/setting/agent/Driver.vue";
import Fsu from "@/components/setting/agent/Fsu.vue";
import LogicObject from "@/components/setting/logic/LogicObject.vue"
import SignalTree from "@/components/layout/side/SignalTree.vue"
import Framework from "@/components/setting/signal/Framework.vue"
import RealData from "@/components/real/RealData.vue";
import ConfigurationSetting from "@/components/system/ConfigurationSetting.vue";
import  i18n from '@/lang/index'

const router = createRouter({
  history: createWebHistory(import.meta.env.VITE_BASE_PATH),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/main',
      name: 'main',
      component: MainView,
      children:[
        {
          path:'/home',
          name:'home',
          components:{default:Home},
          meta: { title: 'Home' }
        },
        {
          path:'/center',
          name:'center',
          components:{
            default: Agent ,
            sidebar: FsuTree
          },
          meta:{title:i18n.global.t('language.common.SURVEILLANCE_CENTER_AGENT')}
        },
        {
          path:'/agent/:agentId',
          name:'agent',
          components:{
            default: Fsu ,
            sidebar: FsuTree
          },
          meta:{title:'FSU'}
        },
        {
          path:'/fsu/:fsuId',
          name:'fsu',
          components:{
            default: Driver ,
            sidebar: FsuTree
          },
          meta:{title:'Driver'}
        },
        {
          path:'/logic/:logicObjectId/:operatedType/:objectType',
          name:'logicObject',
          components:{
            default: LogicObject ,
            sidebar: SignalTree
          },
          meta:{title:'LogicObject'}
        },
        {
          path:'/signal/:signalId/:operatedType/:objectType/:parentLogicObjectId',
          name:'signal',
          components:{
            default: Framework ,
            sidebar: SignalTree
          },
          meta:{title:'Signal'}
        },
        {
          path:'/real-data/:nodeType/:nodeId',
          name:'realdata',
          components:{
            default: RealData ,
            sidebar: SignalTree
          },
          meta:{title:'real data'}
        },
        {
          path:'/configuration-setting/:nodeType/:nodeId',
          name:'configuration',
          components:{
            default: ConfigurationSetting ,
            sidebar: SignalTree
          },
          meta:{title:'configuration setting'}
        }
      ]
    }
  ]
})

export default router

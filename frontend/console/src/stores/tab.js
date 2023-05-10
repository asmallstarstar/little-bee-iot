import { defineStore } from 'pinia'
import router from "@/router";
import i18n from '@/lang/index'
import {useMenusStore} from "@/stores/menu";

export const useTabsStore = defineStore('tabs', {
    state: () => {
        return {
            tabs:[], // name,title,path,
            home:{
                name:'home',
                title:i18n.global.t('language.common.HOME'),
                path:'/home'
            },
            signalTree:['real-data','monitor-object','configuration-setting']
        }
    },

    getters:{
        tabsNameList:(state) => {
            return () => state.tabs.map(x=>x.name)
        }
    },
    actions: {
        resetTabs(){
          this.tabs = []
        },
        addTab(name,title,path){
            const isExist = this.tabs.filter(x=>x.path===path)
            if(isExist.length===0){
                this.tabs.push({
                    name:name,
                    title:title,
                    path:path
                })
            }
        },
        removeTab(tabPath){
            this.tabs = this.tabs.filter(x=>x.path!=tabPath)
            if(this.tabs.length>0){
                router.push(this.tabs[this.tabs.length-1].path)
            }
        },
        removeAll(){
            this.tabs = this.tabs.filter(x=>x.name==this.home.name)
            if(this.tabs.length>0){
                router.push(this.tabs[this.tabs.length-1].path)
            }
        },
        removeOthers(path){
            this.tabs = this.tabs.filter(x=>(x.name==this.home.name || x.path==path))
            if(this.tabs.length>0){
                router.push(this.tabs[this.tabs.length-1].path)
            }
        },
        findTabByPath(path){
            const res = this.tabs.filter(x=>x.path==path)
            if(res.length>0){
                return res[0]
            }else{
                return null
            }
        },
        getSignalTreeMenu(){
            const menusStore = useMenusStore()
            const allMenus = menusStore.menus
            return allMenus.filter(x=>this.signalTree.includes(x.menuName))
        }
    },
    persist: {
        enabled: true,
        strategies: [
            { storage: sessionStorage, paths: ['tabs'] }
        ],
    }
})

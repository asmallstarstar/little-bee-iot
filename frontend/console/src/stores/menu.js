import { defineStore } from 'pinia'
import {getAllMenus} from "@/api/menu"
import router from "@/router";

export const useMenusStore = defineStore('menus', {
    state: () => {
        return {
            menus:null
        }
    },

    getters:{
        menuTree(state) {
            let menuTreeData = []
            for(let i = 0;i<this.menus.length;i++){
                if(this.menus[i].parentMenuId==0){
                    let node = this.createNode(this.menus[i])
                    this.getChildrenNode(this.menus[i],node)
                    menuTreeData.push(node)
                }
            }
            return menuTreeData
        },
        getMenuByMenuId: (state) => {
            return (menuId) => state.menus.filter(x=>x.menuId==menuId)
        },
        getMenusByLevel:(state) => {
            return (menuLevel) => state.menus.filter(x=>x.menuLevel==menuLevel)
        },
        getMenusByParentId:(state) => {
            return (parentId) => state.menus.filter(x=>x.parentMenuId==parentId)
        },
        haveChildMenu:(state) => {
            return (parentId) => (state.menus.filter(x=>x.parentMenuId==parentId)).length>0
        },
    },
    actions: {
        resetMenus(){
          this.menus = null
        },
        setMenus(menus){
            this.menus = menus;
        },
        async load(){
            if(this.menus==null){
                this.menus = []
                const res = await getAllMenus()
                if(res.data.isSuccess){
                    this.menus = res.data.result.items
                    this.menus.forEach(x=>x.show = false)
                }
            }
        },
        geMenusByParent(parentId){
            return this.menus.filter(x=>x.parentMenuId==parentId)
        },
        getChildrenNode(menu, parentNode){
            const childMenus = this.geMenusByParent(menu.menuId)
            if(childMenus.length>0){
                for(let i = 0;i<childMenus.length;i++){
                    let node = this.createNode(childMenus[i])
                    this.getChildrenNode(childMenus[i],node)
                    parentNode.children.push(node)
                }
            }
        },
        createNode(menu){
            let node = {}
            node.id = menu.menuId
            node.label = menu.menuAlias
            node.children =  []
            return node
        },

        //Deprecated: generated menus by the server
        assembledMenus(userActions,menuActions){
            const userActionNames = userActions.map(x=>x.actionName)
            for(let i=3;i>=1;i--){
                this.getMenusByLevel(i).forEach(x=>{
                    const currentMenuActions = menuActions.filter(y=>y.menuName == x.menuName).map(y=>y.actionName)
                    const permission = currentMenuActions.filter(z=>userActionNames.includes(z)).length>0
                    const isHaveChildrens = this.getMenusByParentId(x.menuId).filter(m=>m.show).length>0
                    if(permission || isHaveChildrens){
                        x.show = true
                    }
                })
            }
        },

        assembledRoutes() {
            const mains = import.meta.glob(`../components/**/*.vue`)
            const sidebars = import.meta.glob(`../components/layout/side/*.vue`)
            //console.log(mains)
            //console.log(sidebars)
            this.menus.filter(x=>x.isShow && x.component!=='' ).forEach(x=>{
                let r = null
                if(x.sidebar==''){
                    r = {
                        path:x.path,
                        name:x.menuName,
                        components:{
                            default: mains[`../components/${x.component}`]
                        },
                        meta:{title:x.title}
                    }
                }else{
                    r = {
                        path:x.path,
                        name:x.menuName,
                        components:{
                            default: mains[`../components/${x.component}`] ,
                            sidebar: sidebars[`../components/${x.sidebar}`]
                        },
                        meta:{title:x.title}
                    }
                    //console.log(r)
                }
                router.addRoute('main', r)
            })
        }
    }
})

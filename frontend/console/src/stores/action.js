import { defineStore } from 'pinia'

export const useActionsStore = defineStore('actions', {
    state: () => {
        return {
            actions:[],
            actionGroups:[],
            userActions:[], //All permissions owned by the user
        }
    },
    getters:{
        actionTree(state) {
            let actionTreeData = []
            for(let i = 0;i<this.actionGroups.length;i++){
                if(this.actionGroups[i].parentActionGroupId==0){
                    let node = this.createNode(this.actionGroups[i])
                    this.getChildrenNode(this.actionGroups[i],node)
                    actionTreeData.push(node)
                }
            }
            return actionTreeData
        },
        getActionByAcionId: (state) => {
            return (actionId) => state.actions.filter(x=>x.actionId==actionId)
        },
        getActionByAcionName: (state) => {
            return (actionName) => state.actions.filter(x=>x.actionName===actionName)
        },
    },
    actions: {
        setActions(actions){
          this.actions = actions;
        },
        setActionGroups(actionGroups){
            this.actionGroups = actionGroups
        },
        setUserActions(userActions){
          this.userActions = userActions
        },
        getActionGroupsByParent(parentId){
            return this.actionGroups.filter(x=>x.parentActionGroupId==parentId)
        },
        getActionsByParent(parentId){
            return this.actions.filter(x=>x.actionGroupId==parentId)
        },
        getChildrenNode(actionGroup, parentNode){
            const childActionGroups = this.getActionGroupsByParent(actionGroup.actionGroupId)
            if(childActionGroups.length>0){
                for(let i = 0;i<childActionGroups.length;i++){
                    let node = this.createNode(childActionGroups[i])
                    this.getChildrenNode(childActionGroups[i],node)
                    parentNode.children.push(node)
                }
            }
        },
        createNode(actionGroup){
            let node = {}
            node.id = "g-"+actionGroup.actionGroupId //It's group of action
            node.disabled = false
            node.label = actionGroup.actionGroupAlias
            node.isLeaf = false
            node.actionGroup = actionGroup
            node.actions = this.getActionsByParent(actionGroup.actionGroupId)
            node.children =  []
            node.actions.forEach(x=>{
                let n = {}
                n.id = "a-"+x.actionId //It's action
                n.disabled = false
                n.label = x.actionAlias
                n.isLeaf = true
                n.children = []
                node.children.push(n)
            })
            return node
        },
        isPermission(buttonActions) {
            const subs = buttonActions.filter(x=>this.userActions.includes(x))
            if(subs.length==buttonActions.length){
                return true
            }else{
                return false
            }
        }
    }
})

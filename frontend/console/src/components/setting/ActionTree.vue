<template>
  <el-scrollbar>
    <el-tree
        ref="treeRef"
        :data="actionTree"
        show-checkbox
        node-key="id"
        default-expand-all
        :expand-on-click-node="false"
        :props="{ class: customNodeClass,defaultProps }"
    />
  </el-scrollbar>
</template>

<script setup>
  import {useActionsStore} from "@/stores/action";
  import {computed, ref} from "vue";
  import {storeToRefs} from "pinia";
  import {getAllActions} from "@/api/action";
  import {getAllActionGroups} from "@/api/action-group";
  import {getAllUserAlias} from "@/api/user";
  const treeRef = ref()
  const actionsStore =  useActionsStore()
  //const {actionTree} = storeToRefs(actionsStore)

  const actionTree = computed(() => {
    let actionTreeData = []
    for(let i = 0;i<actionsStore.actionGroups.length;i++){
      if(actionsStore.actionGroups[i].parentActionGroupId==0){
        let node = actionsStore.createNode(actionsStore.actionGroups[i])
        actionsStore.getChildrenNode(actionsStore.actionGroups[i],node)
        actionTreeData.push(node)
      }
    }
    return actionTreeData
  })

  Promise.all([getAllActions(),getAllActionGroups()]).then(function (x) {
    actionsStore.setActions(x[0].data.result.items)
    actionsStore.setActionGroups(x[1].data.result.items)
  })

  const defaultProps = {
    children: 'children',
    label: 'label',
    disabled: 'disabled',
  }

  const customNodeClass = (data, node) => {
    if (data.isLeaf) {
      return 'is-leaf'
    }
    return null
  }

  const setIsDisabled = (ids,isDisabled) =>{
    const setIds = new Set(ids)
    for(let i= 0;i<actionTree.value.length;i++){
      setTreeNodeDisable(setIds,actionTree.value[i],isDisabled)
    }
  }

  function setTreeNodeDisable(setIds,node,isDisabled) {
    if(setIds.has(node.id)){
      node.disabled = isDisabled
    }else{
      node.disabled = ! isDisabled
      for(let i=0;i<node.children.length;i++){
        setTreeNodeDisable(setIds,node.children[i],isDisabled)
      }
    }
  }

  const getCheckedNodes = () => {
    return treeRef.value.getCheckedNodes(true, false)
  }

  const getCheckedKeys = (leafOnly) => {
    return treeRef.value.getCheckedKeys(leafOnly)
  }

  const setCheckedNodes = (keys) => {
    return treeRef.value.setCheckedKeys(keys, false)
  }

  const setIsChecked = (key,isCheck) =>{
    treeRef.value.setChecked(key,isCheck ,false)
  }

  defineExpose({
    setIsDisabled,getCheckedNodes,getCheckedKeys,setCheckedNodes,setIsChecked
  })

</script>

<style>
.is-leaf > .el-tree-node__content {
  color: #409EFF;
}

.el-tree-node.is-expanded.is-leaf > .el-tree-node__children {
  display: flex;
  flex-direction: row;
}
.is-leaf > .el-tree-node__children > div {
  width: 25%;
}
</style>

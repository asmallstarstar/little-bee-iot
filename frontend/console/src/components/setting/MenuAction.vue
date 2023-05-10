<template>
  <div class="card">
    <h3 class="title">{{$t('language.menuAction.SETTING_MENU_ACTION_PERMISS')}}</h3>
    <h4 class="sub-title">{{$t('language.menuAction.MENU')}}</h4>
    <MenuTree class="menu-tree" @select-menu="handleMenu" ref="menuTreeRef"></MenuTree>
    <h4 class="sub-title">{{$t('language.menuAction.ACTION_PERMISS')}}</h4>
    <ActionTree class="action-tree" :max-height="maxHeight" ref="actionTreeRef"></ActionTree>
    <el-button type="primary" @click="submit">{{$t('language.common.SUBMIT')}}</el-button>
  </div>
</template>

<script setup>
import {batchCreateMenuAction,queryMenuAction} from "@/api/menu";
import ActionTree from '@/components/setting/ActionTree.vue'
import MenuTree from '@/components/setting/MenuTree.vue'
import {useMenusStore} from "@/stores/menu";
import {useActionsStore} from "@/stores/action";
import {ref} from "vue";
import {ElMessage} from "element-plus";
import i18n from "@/lang";
import {
  createConnectionWithoutPredicatesType,
  createFieldValue, createFilterWithoutConnectionType,
  createOrderBy,
  createPredicate, createQueryCommand,
  fieldValueTypeEnum, orderTypeEnum,
  predicateTypeEnum,
  pushFieldValueToPredicate
} from "@/utils/query"
const menusStore =  useMenusStore()
const actionsStore = useActionsStore()
const t = i18n.global.t
let maxHeight = ref(window.innerHeight - 400)

function handleMenu(menuId) {
  const menus = menusStore.getMenuByMenuId(menuId)
  if(menus.length>0){
    const predicate = createPredicate(predicateTypeEnum.PREDICATE_TYPE_EQ,'menu_name')
    pushFieldValueToPredicate(predicate,createFieldValue(fieldValueTypeEnum.FIELD_VALUE_TYPE_STRING,menus[0].menuName))
    let predicates = []
    predicates.push(predicate)
    let connection = createConnectionWithoutPredicatesType(predicates)
    let connects = []
    connects.push(connection)
    let filter = createFilterWithoutConnectionType(connects)
    let orderBy = createOrderBy('menu_action_id',orderTypeEnum.ORDER_TYPE_ASC)
    let order = []
    order.push(orderBy)
    const query = createQueryCommand(filter,order)
    queryMenuAction(query).then(x=>{
      let keys = []
      x.data.result.items.forEach(k=>{
        const action = actionsStore.getActionByAcionName(k.actionName)
        if(action.length>0){
          keys.push('a-'+action[0].actionId)
        }
      })
      actionTreeRef.value.setCheckedNodes(keys)
    })
  }
}

const menuTreeRef = ref(null)
const actionTreeRef = ref()

function submit(){
  const currentMenuId = menuTreeRef.value.getCurrentMenuId()
  const menu = getMenuByMenuId(currentMenuId)
  if(menu!=null){
    const actionList = actionTreeRef.value.getCheckedNodes()
    let body = {}
    body.menuName = menu.menuName
    body.actionNameItems = []
    actionList.forEach(x=>{
      const ids = x.id.split("-")
      if(ids[0]==="a"){
        const actions = actionsStore.getActionByAcionId(ids[1])
        body.actionNameItems.push(actions[0].actionName)
      }
    })

    batchCreateMenuAction(body).then(x=>{
      if(x.data.isSuccess){
        ElMessage({
          showClose: false,
          message: t('language.common.SUBMIT_SUCCESS'),
          type: 'success',
        })
      }else{
        ElMessage.error(t('language.common.SUBMIT_FAIL'))
      }
    }).catch(()=>{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    })
  }
}

function getMenuByMenuId(menuId) {
  const menus = menusStore.getMenuByMenuId(menuId)
  if(menus.length>0){
    return menus[0]
  }
  return null
}

</script>

<style scoped>
  .title{
    margin: 10px 0;
    text-align: center;
  }
  .card{
    width: 300px;
    padding: 0 5px;
  }
  .menu-tree{
    margin: 5px 0;
  }
  .action-tree{
    margin: 5px 0;
  }
  .sub-title{
    margin: 5px 0;
  }
  .el-button{
    display:block;
    margin:25px auto 0;
  }

</style>

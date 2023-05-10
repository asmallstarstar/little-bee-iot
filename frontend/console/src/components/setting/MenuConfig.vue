<template>
    <div>
      <el-button size="small" type="primary" @click="handleCreate" v-if="actionsStore.isPermission(['create-menu'])">{{$t('language.common.CREATE')}}</el-button>
      <el-button size="small" type="primary" @click="handlePermiss"
                 v-if="actionsStore.isPermission(['all-action','all-action-group','query-menu-action'])">{{$t('language.menu.MENU_PERMISSION')}}</el-button>
      <div>
        <el-table :data="tableDataSetsWithPage" :max-height="maxHeight">
          <el-table-column prop="menuId" :label="$t('language.menu.MENU_ID')" width="80" />
          <el-table-column prop="menuName" :label="$t('language.menu.MENU_NAME')" width="160" />
          <el-table-column prop="menuAlias" :label="$t('language.menu.MENU_ALIAS')" width="90" />
          <el-table-column :label="$t('language.menu.MENU_PARENT')" width="80" >
            <template #default="scope">
              {{getMenuNameById(scope.row.parentMenuId)}}
            </template>
          </el-table-column>
          <el-table-column prop="menuLevel" :label="$t('language.menu.MENU_LEVEL')" width="60"/>
          <el-table-column prop="path" :label="$t('language.menu.MENU_PATH')" width="220"/>
          <el-table-column prop="component" :label="$t('language.menu.MENU_COMPONENT')" width="250"/>
          <el-table-column prop="title" :label="$t('language.menu.MENU_TITLE')" width="160"/>
          <el-table-column prop="sidebar" :label="$t('language.menu.MENU_SIDEBAR')" width="250"/>
          <el-table-column prop="icon" :label="$t('language.menu.MENU_ICON')" width="150"/>
          <el-table-column :label="$t('language.common.TIME')" width="60">
            <template #default="scope">
              <el-popover effect="light" trigger="hover" placement="top" width="auto">
                <template #default>
                  <div>{{$t('language.common.CREATE_TIME')}}:
                    {{ scope.row.createdAt?(new Date(scope.row.createdAt)).toLocaleString():""}}</div>
                  <div>{{$t('language.common.UPDATE_TIME')}}:
                    {{ scope.row.updatedAt?(new Date(scope.row.updatedAt)).toLocaleString():"" }}</div>
                  <div>{{$t('language.common.CREATOR')}}:
                    {{ scope.row.createdBy?getUserNickByUserId(scope.row.createdBy):"" }}</div>
                  <div>{{$t('language.common.MODIFIER')}}:
                    {{ scope.row.updatedBy?getUserNickByUserId(scope.row.updatedBy):"" }}</div>
                </template>
                <template #reference>
                  <el-tag>...</el-tag>
                </template>
              </el-popover>
            </template>
          </el-table-column>
          <el-table-column :label="$t('language.common.OPERATE')" width="200">
            <template #default="scope">
              <el-button size="small" @click="handleView(scope.$index, scope.row)" >{{$t('language.common.VIEW')}}</el-button>
              <el-button size="small" @click="handleEdit(scope.$index, scope.row)" v-if="actionsStore.isPermission(['create-menu','update-menu'])">{{$t('language.common.EDIT')}}</el-button>
              <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)" v-if="actionsStore.isPermission(['delete-menu'])">{{$t('language.common.DELETE')}}</el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-pagination
            v-model:current-page="currentPage"
            v-model:page-size="pageSize"
            :page-sizes="[5, 10, 15, 20]"
            :small="small"
            :disabled="disabled"
            :background="background"
            layout="sizes, prev, pager, next"
            :total="menusStore.menus.length"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            style="margin-top: 3px;"
        />
      </div>

      <el-dialog v-model="dialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="500px">
        <el-form :model="rowData" :rules="rules" ref="dialogForm">
          <el-form-item :label="$t('language.menu.MENU_PARENT')" label-position="left" label-width="100px" prop="parentMenuId" required>
            <el-select v-model="rowData.parentMenuId" class="m-2" :placeholder="$t('language.common.SELECT')" @change="changeParent" >
              <el-option
                  v-for="item in menusStore.menus"
                  :key="item.menuId"
                  :label="item.menuAlias"
                  :value="item.menuId"
              />
            </el-select>
          </el-form-item>
          <el-form-item :label="$t('language.menu.MENU_ID')" label-position="left" label-width="100px" prop="menuId" required >
            <el-input v-model="rowData.menuId" :readonly="operatedType===operatedTypeEnum.OPERATED_TYPE_CREATE?false:true"/>
          </el-form-item>
          <el-form-item :label="$t('language.menu.MENU_NAME')" label-position="left" label-width="100px" prop="menuName" required>
            <el-input v-model="rowData.menuName"/>
          </el-form-item>
          <el-form-item :label="$t('language.menu.MENU_ALIAS')" label-position="left" label-width="100px" prop="menuAlias" required>
            <el-input v-model="rowData.menuAlias"/>
          </el-form-item>
          <el-form-item :label="$t('language.menu.MENU_LEVEL')" label-position="left" label-width="100px" prop="menuLevel" required>
            <el-input v-model="rowData.menuLevel"/>
          </el-form-item>
          <el-form-item :label="$t('language.menu.MENU_PATH')" label-position="left" label-width="100px" prop="path" required>
            <el-input v-model="rowData.path"/>
          </el-form-item>
          <el-form-item :label="$t('language.menu.MENU_COMPONENT')" label-position="left" label-width="100px" prop="component" required>
            <el-input v-model="rowData.component"/>
          </el-form-item>
          <el-form-item :label="$t('language.menu.MENU_TITLE')" label-position="left" label-width="100px" prop="title" required>
            <el-input v-model="rowData.title"/>
          </el-form-item>
          <el-form-item :label="$t('language.menu.MENU_SIDEBAR')" label-position="left" label-width="100px" prop="sidebar" >
            <el-input v-model="rowData.sidebar"/>
          </el-form-item>
          <el-form-item :label="$t('language.menu.MENU_ICON')" label-position="left" label-width="100px" prop="icon" required>
            <el-input v-model="rowData.icon"/>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="dialogVisible = false">{{$t('language.common.CANCEL')}}</el-button>
            <el-button type="primary" @click="handleConfirm" v-if="operatedType!=operatedTypeEnum.OPERATED_TYPE_VIEW">{{$t('language.common.CONFIRM')}}</el-button>
          </span>
        </template>
      </el-dialog>

      <el-dialog v-model="grantDialogVisible" :title="t('language.menu.MENU_PERMISSION')" :close-on-click-modal="false"
                 align-center draggable width="320px" >
        <MenuAction></MenuAction>
      </el-dialog>

    </div>
  </template>

<script setup>
import {computed, nextTick, ref, toRaw} from "vue";
import {useMenusStore} from "../../stores/menu";
import {ElMessage, ElMessageBox} from "element-plus";
import {operatedTypeEnum} from "../../utils/common";
import {createMenu, deleteMenu, updateMenu} from "../../api/menu";
import i18n from "@/lang"
import MenuAction from "./MenuAction.vue";
import {getAllUserAlias} from "@/api/user";
import {useActionsStore} from "@/stores/action";

const t = i18n.global.t
const currentPage = ref(1)
const pageSize = ref(15)
const small = ref(false)
const background = ref(false)
const disabled = ref(false)
const grantDialogVisible = ref(false)
const menusStore = useMenusStore()
const actionsStore = useActionsStore()
//const {menus:tableDataSets} = storeToRefs(menusStore)
let dialogTitle = ref(t('language.common.CREATE'))
let dialogVisible = ref(false)
let operatedType = ref(operatedTypeEnum.OPERATED_TYPE_UNKNOWN)
let maxHeight = ref(window.innerHeight - 165)
const dialogForm = ref()
const users = ref()
getAllAlias()
let rowData = ref({
  menuId:null,
  parentMenuId:null,
  parentMenuIdName:"",
  menuName:"",
  menuAlias:"",
  menuLevel:"",
  path:"",
  component:"",
  title:"",
  icon:"",
  sidebar:"",
  createdAt:0,
  updatedAt:0,
})

function handleCreate() {
  dialogTitle.value = t('language.common.CREATE')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_CREATE
  rowData.value = {
    menuId:null,
    parentMenuId:null,
    parentMenuIdName:"",
    menuName:"",
    menuAlias:"",
    menuLevel:"",
    path:"",
    component:"",
    title:"",
    icon:"",
    sidebar:null,
    createdAt:0,
    updatedAt:""
  }
  dialogVisible.value = true
}

function handleView(index,row) {
  dialogTitle.value = t('language.common.VIEW')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_VIEW
  const {...raw} = row
  rowData.value = raw
  if(rowData.value.parentMenuId==0)
    rowData.value.parentMenuId = null
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleEdit(index,row) {
  dialogTitle.value = t('language.common.EDIT')
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_EDIT
  const {...raw} = row
  rowData.value = raw
  if(rowData.value.parentMenuId==0)
    rowData.value.parentMenuId = null
  nextTick(()=>{
    dialogVisible.value = true
  })
}

function handleDelete(index,row) {
  const {...raw} = row
  rowData.value = raw
  operatedType.value = operatedTypeEnum.OPERATED_TYPE_DELETE

  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.menu.MENU')+": "+rowData.value.menuAlias, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        deleteMenu(rowData.value.menuId).then(x=>{
          if(x.data.isSuccess){
            ElMessage({
              showClose: false,
              message: t('language.common.DELETE_SUCCESS'),
              type: 'success',
            })
            menusStore.menus = menusStore.menus.filter(x=>x.menuId!=rowData.value.menuId)
          }else{
            ElMessage.error(t('language.common.DELETE_FAI'))
          }
        })
      })
}

function handleConfirm() {
  const { parentMenuIdName,createdAt,updatedAt, ...newTable } = rowData.value;
  let fn = createMenu
  if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
    fn = updateMenu
  }
  fn(newTable).then(x=>{
    if(x.data.isSuccess){
      ElMessage({
        showClose: false,
        message: t('language.common.SUBMIT_SUCCESS'),
        type: 'success',
      })
      if(operatedType.value == operatedTypeEnum.OPERATED_TYPE_EDIT){
        const { parentMenuIdName, ...newStoreData } = rowData.value
        for(let i=0;i<menusStore.menus.length;i++){
          if(menusStore.menus[i].menuId==newTable.menuId){
            menusStore.menus[i] = newStoreData
            break
          }
        }
      }else{
        const res = x.data.result
        delete res['@type']
        menusStore.menus.push(res)
      }
    }else{
      ElMessage.error(t('language.common.SUBMIT_FAIL'))
    }
    dialogVisible.value=false
  }).catch(()=>{
    ElMessage.error(t('language.common.SUBMIT_FAIL'))
  })
}

function handlePermiss() {
  grantDialogVisible.value = true
}

const changeParent = (currentValue) =>{
  const l = menusStore.menus.filter(x=>x.menuId == currentValue)
  rowData.value.menuLevel = l[0].menuLevel+1
  rowData.value.menuId = l[0].menuId
}

const getMenuNameById = computed(()=>(menuId) => {
  let l = menusStore.menus.filter(x=>x.menuId == menuId)
  if(l.length>0)
    return l[0].menuAlias
  else
    return ""
})

const tableDataSetsWithPage = computed(() => {
  const start = (currentPage.value-1)*pageSize.value
  const end = (currentPage.value-1)*pageSize.value + pageSize.value
  return menusStore.menus.slice( start,end )
})

const handleSizeChange = (itemCountPerPage) => {
  pageSize.value = itemCountPerPage
}
const handleCurrentChange = (changedCurrentPage) => {
  currentPage.value = changedCurrentPage
}

const getUserNickByUserId =  (userId) => {
  if(users && users.value){
    const l = users.value.filter(x=>x.userId == userId)
    if(l.length>0)
      return l[0].userNick
  }
  return ""
}

function getAllAlias(){
  getAllUserAlias().then(x=>{
    users.value = x.data.result.items
  })
}

const rules = {
  menuId: [{ required: true,trigger: 'blur' }],
  menuName: [{ required: true,trigger: 'blur' }],
  menuAlias: [{ required: true,trigger: 'blur' }],
  parentMenuId: [{ required: true,trigger: 'blur' }],
  menuLevel: [{ required: true,trigger: 'blur' }],
  path: [{ required: true,trigger: 'blur' }],
  component: [{ required: true,trigger: 'blur' }],
  title: [{ required: true,trigger: 'blur' }],
  sidebar: [{ required: false,trigger: 'blur' }],
  icon: [{ required: true,trigger: 'blur' }],
}

</script>

<style scoped>

</style>

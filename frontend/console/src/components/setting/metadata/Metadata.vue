<template>
  <el-form :model="metadata" :rules="metadataRules" ref="metadataDialogForm">
    <el-form-item :label="$t('language.metadata.METADATA_ID')" label-position="left" label-width="120px" prop="metadataId"  required
    v-if="operatorType==operatedTypeEnum.OPERATED_TYPE_VIEW">
      <el-input v-model="metadata.metadataId"  readonly/>
    </el-form-item>
    <el-form-item :label="$t('language.metadata.METADATA_NAME')" label-position="left" label-width="120px" prop="metadataName"  required >
      <el-input v-model="metadata.metadataName"  />
    </el-form-item>
    <el-form-item :label="$t('language.metadata.METADATA_ALIAS')" label-position="left" label-width="120px" prop="metadataAlias"  required >
      <el-input v-model="metadata.metadataAlias"  />
    </el-form-item>
    <el-form-item :label="$t('language.common.OPERATE')" label-position="left" label-width="120px">
      <el-button size="small" type="primary" @click="handleAddAttribute">{{$t('language.metadata.ADD_ATTRIBUTE')}}</el-button>
    </el-form-item>
    <el-form-item :label="$t('language.metadata.ATTRIBUTE')" label-position="left" label-width="120px">
      <el-table :data="metadata.metadataAttribute" >
        <el-table-column prop="metadataAttributeName" :label="$t('language.metadata.ATTRIBUTE_NAME')" width="120" />
        <el-table-column prop="metadataAttributeAlias" :label="$t('language.metadata.ATTRIBUTE_ALIAS')" width="120" />
        <el-table-column :label="$t('language.metadata.ATTRIBUTE_TYPE')" width="120" >
          <template #default="scope">
            {{getAttribute(scope.row.metadataAttributeTypeEnum)}}
          </template>
        </el-table-column>
        <el-table-column :label="$t('language.metadata.ATTRIBUTE_VALUE_TYPE')" width="120" >
          <template #default="scope">
            {{getValueType(scope.row.metadataValueTypeEnum)}}
          </template>
        </el-table-column>
        <el-table-column :label="$t('language.metadata.ATTRIBUTE_ENUM_ITEM')" width="180" >
          <template #default="scope">
            <el-tag
                v-for="tag in scope.row.enumValueItem"
                :key="tag.enumValueItemName"
                class="mx-1"
                type="success"
            >
              {{ tag.enumValueItemAlias+":" +tag.enumValueItemValue}}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column :label="$t('language.common.OPERATE')">
          <template #default="scope">
            <el-button size="small" @click="handleEditAttribute(scope.$index, scope.row)" >{{$t('language.common.EDIT')}}</el-button>
            <el-button size="small" type="danger" @click="handleDeleteAttribute(scope.$index, scope.row)" >{{$t('language.common.DELETE')}}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-form-item>
  </el-form>

  <el-dialog v-model="attributeDialogVisible" :title="dialogTitle" :close-on-click-modal="false" align-center draggable width="350px">
    <el-form :model="attribute" :rules="rules" ref="dialogForm">
      <el-form-item :label="$t('language.metadata.ATTRIBUTE_NAME')" label-position="left" label-width="120px" prop="metadataAttributeName"  required >
        <el-input v-model="attribute.metadataAttributeName"  />
      </el-form-item>
      <el-form-item :label="$t('language.metadata.ATTRIBUTE_ALIAS')" label-position="left" label-width="120px" prop="metadataAttributeAlias"  required >
        <el-input v-model="attribute.metadataAttributeAlias"  />
      </el-form-item>
      <el-form-item :label="$t('language.metadata.ATTRIBUTE_TYPE')" label-position="left" label-width="120px" prop="metadataAttributeTypeEnum" required>
        <el-select v-model="attribute.metadataAttributeTypeEnum" class="m-2" :placeholder="$t('language.common.SELECT')" >
          <el-option v-for="item in metadataAttributeTypeText" :key="item.index" :label="item.name" :value="item.index"/>
        </el-select>
      </el-form-item>
      <el-form-item :label="$t('language.metadata.ATTRIBUTE_VALUE_TYPE')" label-position="left" label-width="120px" prop="metadataValueTypeEnum" required>
        <el-select v-model="attribute.metadataValueTypeEnum" class="m-2" :placeholder="$t('language.common.SELECT')" >
          <el-option v-for="item in metadataValueTypeText" :key="item.index" :label="item.name" :value="item.index"/>
        </el-select>
      </el-form-item>
      <template v-if="attribute.metadataAttributeTypeEnum==metadataAttributeTypeEnum.METADATA_ATTRIBUTE_TYPE_ENUM">
        <el-form-item :label="$t('language.common.OPERATE')" label-position="left" label-width="120px">
          <el-button size="small" type="primary" @click="addEnumItem">{{$t('language.metadata.ADD_ATTRIBUTE_ENUM_ITEM')}}</el-button>
        </el-form-item>
        <el-form-item :label="$t('language.metadata.ATTRIBUTE_ENUM_ITEM')" label-position="left" label-width="120px" >
          <el-tag
              v-for="tag in attribute.enumValueItem"
              :key="tag.enumValueItemName"
              class="mx-1"
              closable
              type="success"
              @close="handleDeleteTag(tag,tag.enumValueItemName)"
          >
            {{ tag.enumValueItemAlias+":" +tag.enumValueItemValue}}
          </el-tag>
        </el-form-item>
      </template>
    </el-form>
    <template #footer>
          <span class="dialog-footer">
            <el-button @click="attributeDialogVisible = false">{{$t('language.common.CANCEL')}}</el-button>
            <el-button type="primary" @click="handleConfirmAttribute">{{$t('language.common.CONFIRM')}}</el-button>
          </span>
    </template>
  </el-dialog>

  <el-dialog v-model="enumValueDialogVisible" :title="$t('language.common.CREATE')" :close-on-click-modal="false" align-center draggable width="350px">
    <el-form :model="enumValue" :rules="enumValueRules" ref="enumValueDialogForm">
      <el-form-item :label="$t('language.metadata.ENUM_ITEM_NAME')" label-position="left" label-width="120px" prop="enumValueItemName"  required >
        <el-input v-model="enumValue.enumValueItemName"  />
      </el-form-item>
      <el-form-item :label="$t('language.metadata.ENUM_ITEM_ALIAS')" label-position="left" label-width="120px" prop="enumValueItemAlias"  required >
        <el-input v-model="enumValue.enumValueItemAlias"  />
      </el-form-item>
      <el-form-item :label="$t('language.metadata.ENUM_ITEM_VALUE')" label-position="left" label-width="120px" prop="enumValueItemValue"  required >
        <el-input v-model="enumValue.enumValueItemValue"  />
      </el-form-item>
    </el-form>
    <template #footer>
          <span class="dialog-footer">
            <el-button @click="enumValueDialogVisible = false">{{$t('language.common.CANCEL')}}</el-button>
            <el-button type="primary" @click="handleEnumItemConfirm">{{$t('language.common.CONFIRM')}}</el-button>
          </span>
    </template>
  </el-dialog>

</template>

<script setup>
import {nextTick, ref, toRaw, watch} from "vue";
import {metadataAttributeTypeText,metadataValueTypeText,metadataAttributeTypeEnum} from "@/utils/metadata"
import {operatedTypeEnum} from "@/utils/common";
import i18n from "@/lang";
import {retrieveMetadata} from "@/api/metadata";
import {ElMessage, ElMessageBox} from "element-plus";
import {deleteAction} from "@/api/action";
import {deleteActionGroup} from "@/api/action-group";

const t = i18n.global.t
let dialogTitle = ref(t('language.common.CREATE'))
const props = defineProps(['metadataId','operatorType'])
const dialogForm = ref()
const enumValueDialogForm = ref()
const metadataDialogForm = ref()
const attributeDialogVisible = ref(false)
const enumValueDialogVisible = ref(false)
let attributeOperatorType = operatedTypeEnum.OPERATED_TYPE_UNKNOWN
let currentAttributeIndex = 0
const metadata = ref({
  metadataId:null,
  metadataName:"",
  metadataAlias:"",
  metadataAttribute:[]
})
const attribute = ref({
  metadataAttributeName:"",
  metadataAttributeAlias:"",
  metadataAttributeTypeEnum:0,
  metadataValueTypeEnum:0,
  enumValueItem:[]
})

const enumValue = ref({
  enumValueItemName:"",
  enumValueItemAlias:"",
  enumValueItemValue:""
})

function handleAddAttribute() {
  attributeOperatorType = operatedTypeEnum.OPERATED_TYPE_CREATE
  attribute.value = {
    metadataAttributeName:"",
    metadataAttributeAlias:"",
    metadataAttributeTypeEnum:0,
    metadataValueTypeEnum:0,
    enumValueItem:[]
  }
  attributeDialogVisible.value = true
}

function handleEditAttribute(index, row) {
  currentAttributeIndex = index
  attributeOperatorType = operatedTypeEnum.OPERATED_TYPE_EDIT
  const {...raw} = row
  attribute.value = raw
  attributeDialogVisible.value = true
}

function handleDeleteAttribute(index, row) {
  ElMessageBox.confirm(t('language.common.DELETE_TITLE')+t('language.metadata.ATTRIBUTE')+": "+row.metadataAttributeAlias, t('language.common.DELETE_WARNING'), {
    confirmButtonText: t('language.common.CONFIRM_BUTTON_TEXT'), cancelButtonText: t('language.common.CANCEL_BUTTON_TEXT'), type: 'warning',})
      .then(() => {
        metadata.value.metadataAttribute = metadata.value.metadataAttribute.filter(x=>x.metadataAttributeName!=row.metadataAttributeName)
      })
}

function handleConfirmAttribute() {
  const {...raw} = toRaw(attribute.value)
  if(attributeOperatorType==operatedTypeEnum.OPERATED_TYPE_CREATE){
    metadata.value.metadataAttribute.push(raw)
  }else if(attributeOperatorType==operatedTypeEnum.OPERATED_TYPE_EDIT){
    metadata.value.metadataAttribute[currentAttributeIndex] = raw
  }
  attributeDialogVisible.value = false
}

function addEnumItem() {
  enumValue.value = {
    enumValueItemName:"",
    enumValueItemAlias:"",
    enumValueItemValue:""
  }
  enumValueDialogVisible.value = true
}

function handleDeleteTag(tag,tagName) {
  attribute.value.enumValueItem = attribute.value.enumValueItem.filter(x=>x.enumValueItemName!=tagName)
}

function handleEnumItemConfirm() {
  const {...raw} = toRaw(enumValue.value)
  attribute.value.enumValueItem.push(raw)
  enumValueDialogVisible.value = false
}

function getAttribute(index) {
  for(let i=0;i<metadataAttributeTypeText.length;i++){
    if(i==index) {
      return metadataAttributeTypeText[i].name
    }
  }
  return ""
}

function getValueType(index) {
  for(let i=0;i<metadataValueTypeText.length;i++){
    if(i==index) {
      return metadataValueTypeText[i].name
    }
  }
  return ""
}

function setMetadata(m) {
  metadata.value = m
}

const getMetadata = () => {
  const {...raw} = toRaw(metadata.value)
  return raw
}

defineExpose({
  getMetadata,setMetadata
})

const metadataRules = {
  metadataName: [{ required: true,trigger: 'blur' }],
  metadataAlias: [{ required: true,trigger: 'blur' }],
}

const rules = {
  metadataAttributeName: [{ required: true,trigger: 'blur' }],
  metadataAttributeAlias: [{ required: true,trigger: 'blur' }],
  metadataAttributeTypeEnum: [{ required: true,trigger: 'blur' }],
  metadataValueTypeEnum: [{ required: true,trigger: 'blur' }]
}

const enumValueRules = {
  enumValueItemName: [{ required: true,trigger: 'blur' }],
  enumValueItemAlias: [{ required: true,trigger: 'blur' }],
  enumValueItemValue: [{ required: true,trigger: 'blur' }]
}
</script>

<style scoped>

</style>
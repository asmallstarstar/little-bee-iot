<template>
  <el-form :model="area">
    <el-form-item :label="$t('language.logicObject.AREA_NAME')" label-position="left" label-width="120px" prop="areaAlias">
      <el-input v-model="area.areaAlias" />
    </el-form-item>
    <el-form-item :label="$t('language.logicObject.AREA_TYPE_NAME')" label-position="left" label-width="120px" prop="areaType">
      <el-select v-model="area.areaType" class="m-2" :placeholder="$t('language.common.SELECT')">
        <el-option v-for="item in areaTypes" :key="item.areaTypeId" :label="item.areaTypeName" :value="item.areaTypeId"/>
      </el-select>
    </el-form-item>
  </el-form>
</template>

<script setup>
import {ref, toRaw, watch} from "vue";
import {getAllAreaTypes} from "@/api/area-type";
import {operatedTypeEnum} from "@/utils/common";
import {retrieveLogicArea} from "@/api/logic-area";

const areaTypes = ref()
const props = defineProps(['logicObjectId','operatedType'])
const  area = ref({
  logicObjectId:0,
  areaAlias:"",
  areaType:null
})

watch(props,(newX)=>{
  if(newX.operatedType!==operatedTypeEnum.OPERATED_TYPE_CREATE){
    retrieveLogicArea(newX.logicObjectId).then(x=>{
      delete x.data.result['@type']
      area.value = x.data.result
    })
  }else{
    area.value = {
      logicObjectId:0,
      areaAlias:"",
      areaType:null
    }
  }
},{immediate: true, deep: true})

getAllAreaTypes().then(function (x) {
  areaTypes.value = x.data.result.items
})

const getArea = () => {
  const {...rawData} = toRaw(area.value)
  return rawData
}

defineExpose({
  getArea
})

</script>

<style scoped>
.el-input ,.el-select {
  width: 500px;
}
</style>
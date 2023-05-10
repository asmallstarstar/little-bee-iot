<template>
  <div>
    <el-form size="small">
      <el-form-item v-for="item in newMetadataAttribute" :label="item.metadataAttributeAlias" label-position="left" >
        <el-input v-if="item.metadataAttributeTypeEnum==metadataAttributeTypeEnum.METADATA_ATTRIBUTE_TYPE_SIMPLE" v-model="item.metadataAttributeValue" :disabled="readOnly"/>

        <template v-else-if="item.metadataAttributeTypeEnum==metadataAttributeTypeEnum.METADATA_ATTRIBUTE_TYPE_ARRAY">
          <el-input
              v-model="item.metadataAttributeInputValue"
              :placeholder="$t('language.metadata.ARRAY_VALUE')"
              class="input-with-select"
              :disabled="readOnly"
              v-if="!readOnly"
          >
            <template #append>
              <el-button @click="item.metadataAttributeArrayValue.push(item.metadataAttributeInputValue)" :disabled="readOnly"> {{$t('language.common.ADD')}}</el-button>
            </template>
          </el-input>

          <el-tag
              v-for="tag in item.metadataAttributeArrayValue"
              class="mx-1"
              :closable="!readOnly"
              type="info"
              @close="item.metadataAttributeArrayValue=item.metadataAttributeArrayValue.filter(x=>x!=tag)"
          >
            {{ tag }}
          </el-tag>
        </template>

        <el-select v-else-if="item.metadataAttributeTypeEnum==metadataAttributeTypeEnum.METADATA_ATTRIBUTE_TYPE_ENUM"
                   :placeholder="$t('language.common.SELECT')"
                   v-model="item.metadataAttributeValue" :disabled="readOnly">
          <el-option v-for="r in item.enumValueItem" :label="r.enumValueItemAlias" :key="r.enumValueItemName" :value="r.enumValueItemValue"></el-option>
        </el-select>
      </el-form-item>
    </el-form>

  </div>
</template>

<script setup>
import {initMetadataAttributeValue, metadataAttributeTypeEnum} from "@/utils/metadata"
import {ref, toRaw, watch} from "vue";
import {retrieveMetadata} from "@/api/metadata";
import {retrieveConfigure} from "@/api/configure";
const props = defineProps(['metadataId','configureId','readOnly'])
const newMetadataAttribute = ref()
transformMetadataAttribute(props.metadataId,props.configureId)

watch(props,(newX)=>{
  transformMetadataAttribute(newX.metadataId,newX.configureId)
})

function transformMetadataAttribute(metadataId, configureId) {
  if (metadataId && configureId) {
    Promise.all([retrieveMetadata(metadataId),retrieveConfigure(configureId)]).then(function (x) {
      if (metadataId != x[1].data.result.metadataId) {
        newMetadataAttribute.value = initMetadataAttributeValue(JSON.parse(x[0].data.result.metadataAttributeJson))
      } else {
        newMetadataAttribute.value = JSON.parse(x[1].data.result.configureJson)
      }
    })
  }else if (metadataId && !configureId) {
    retrieveMetadata(metadataId).then(x=>{
      newMetadataAttribute.value = initMetadataAttributeValue(JSON.parse(x.data.result.metadataAttributeJson))
    })
  }else if (!metadataId && !configureId) {
    newMetadataAttribute.value = []
  }
}

const getMetadataAttribute = () => {
  const {...rawData} = toRaw(newMetadataAttribute.value)
  return rawData
}

defineExpose({
  getMetadataAttribute
})

</script>

<style scoped>

</style>
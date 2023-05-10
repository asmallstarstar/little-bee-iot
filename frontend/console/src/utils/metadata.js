import i18n from "@/lang";

export const metadataAttributeTypeEnum = {
    METADATA_ATTRIBUTE_TYPE_UNKNOWN: 0,
    METADATA_ATTRIBUTE_TYPE_SIMPLE:  1,
    METADATA_ATTRIBUTE_TYPE_ARRAY:   2,
    METADATA_ATTRIBUTE_TYPE_ENUM:    3
}
const t = i18n.global.t

export const metadataAttributeTypeText = [
    {index:0,name:t('language.metadata.UNKNOWN')},
    {index:1,name:t('language.metadata.SIMPLE')},
    {index:2,name:t('language.metadata.ARRAY')},
    {index:3,name:t('language.metadata.ENUM')}
]

export const metadataValueTypeEnum = {
    METADATA_VALUE_TYPE_UNKNOWN : 0,
    METADATA_VALUE_TYPE_INT     : 1,
    METADATA_VALUE_TYPE_LONG    : 2,
    METADATA_VALUE_TYPE_FLOAT   : 3,
    METADATA_VALUE_TYPE_DOUBLE  : 4,
    METADATA_VALUE_TYPE_STRING  : 5,
    METADATA_VALUE_TYPE_BOOLEAN : 6,
}

export const metadataValueTypeText = [
    {index:0,name:t('language.metadata.UNKNOWN')},
    {index:1,name:t('language.metadata.INT')},
    {index:2,name:t('language.metadata.LONG')},
    {index:3,name:t('language.metadata.FLOAT')},
    {index:4,name:t('language.metadata.DOUBLE')},
    {index:5,name:t('language.metadata.STRING')},
    {index:6,name:t('language.metadata.BOOLEAN')}
]

export function initMetadataAttributeValue(metadataAttribute){
    metadataAttribute.forEach(x=>{
        x.metadataAttributeValue = ""
        if(x.metadataAttributeTypeEnum==metadataAttributeTypeEnum.METADATA_ATTRIBUTE_TYPE_SIMPLE ||
            x.metadataAttributeTypeEnum==metadataAttributeTypeEnum.METADATA_ATTRIBUTE_TYPE_ENUM){
           switch (x.metadataValueTypeEnum) {
               case metadataValueTypeEnum.METADATA_VALUE_TYPE_INT:
                   x.metadataAttributeValue = 0
                   break
               case metadataValueTypeEnum.METADATA_VALUE_TYPE_LONG:
                   x.metadataAttributeValue = 0
                   break
               case metadataValueTypeEnum.METADATA_VALUE_TYPE_FLOAT:
                   x.metadataAttributeValue = 0.0
                   break
               case metadataValueTypeEnum.METADATA_VALUE_TYPE_DOUBLE:
                   x.metadataAttributeValue = 0.0
                   break
               case metadataValueTypeEnum.METADATA_VALUE_TYPE_STRING:
                   x.metadataAttributeValue = ""
                   break
               case metadataValueTypeEnum.METADATA_VALUE_TYPE_BOOLEAN:
                   x.metadataAttributeValue = false
                   break
               default:
                   x.metadataAttributeValue = ""
                   break
           }
        }else if(x.metadataAttributeTypeEnum==metadataAttributeTypeEnum.METADATA_ATTRIBUTE_TYPE_ARRAY){
            x.metadataAttributeValue = ""
            x.metadataAttributeArrayValue = []
            x.metadataAttributeInputValue = ""
        }
    })
    return metadataAttribute
}
export const predicateTypeEnum = {
    PREDICATE_TYPE_UNKNOWN: "PREDICATE_TYPE_UNKNOWN",
    PREDICATE_TYPE_EQ: "PREDICATE_TYPE_EQ",
    PREDICATE_TYPE_LT: "PREDICATE_TYPE_LT",
    PREDICATE_TYPE_GT: "PREDICATE_TYPE_GT",
    PREDICATE_TYPE_LE: "PREDICATE_TYPE_LE",
    PREDICATE_TYPE_GE: "PREDICATE_TYPE_GE",
    PREDICATE_TYPE_BETWEEN: "PREDICATE_TYPE_BETWEEN",
    PREDICATE_TYPE_LIKE: "PREDICATE_TYPE_LIKE",
    PREDICATE_TYPE_NE: "PREDICATE_TYPE_NE",
    PREDICATE_TYPE_IS_NULL: "PREDICATE_TYPE_IS_NULL",
    PREDICATE_TYPE_IS_NOT_NULL: "PREDICATE_TYPE_IS_NOT_NULL",
    PREDICATE_TYPE_IN: "PREDICATE_TYPE_IN",
    PREDICATE_TYPE_NOT_IN: "PREDICATE_TYPE_NOT_IN",
}

export const connectionTypeEnum = {
    CONNECT_TYPE_UNKNOWN: "CONNECT_TYPE_UNKNOWN",
    CONNECT_TYPE_AND: "CONNECT_TYPE_AND",
    CONNECT_TYPE_OR: "CONNECT_TYPE_OR",
}

export const fieldValueTypeEnum = {
    FIELD_VALUE_TYPE_UNKNOWN: "FIELD_VALUE_TYPE_UNKNOWN",
    FIELD_VALUE_TYPE_INTEGER:"FIELD_VALUE_TYPE_INTEGER",
    FIELD_VALUE_TYPE_FLOAT: "FIELD_VALUE_TYPE_FLOAT",
    FIELD_VALUE_TYPE_STRING: "FIELD_VALUE_TYPE_STRING",
    FIELD_VALUE_TYPE_TIME: "FIELD_VALUE_TYPE_TIME",
}

export const orderTypeEnum = {
    ORDER_TYPE_UNKNOWN: "ORDER_TYPE_UNKNOWN",
    ORDER_TYPE_ASC: "ORDER_TYPE_ASC",
    ORDER_TYPE_DESC: "ORDER_TYPE_DESC",
}

export function createFieldValue(fieldValueType,fieldValue) {
    let fieldVale = {}
    fieldVale.fieldType = fieldValueType
    fieldVale.fieldValueInteger = 0
    fieldVale.fieldValueFloat = 0.0
    fieldVale.fieldValueString = ""
    fieldVale.fieldValueTime = (new Date()).toISOString()
    if(fieldValueType==fieldValueTypeEnum.FIELD_VALUE_TYPE_INTEGER){
        fieldVale.fieldValueInteger = fieldValue
    }
    if(fieldValueType==fieldValueTypeEnum.FIELD_VALUE_TYPE_FLOAT){
        fieldVale.fieldValueFloat = fieldValue
    }
    if(fieldValueType==fieldValueTypeEnum.FIELD_VALUE_TYPE_STRING){
        fieldVale.fieldValueString = fieldValue
    }
    if(fieldValueType==fieldValueTypeEnum.FIELD_VALUE_TYPE_TIME){
        fieldVale.fieldValueTime = fieldValue
    }
    return fieldVale
}

export function createPredicate(predicateType,fieldName) {
    let predicate = {}
    predicate.predicateType = predicateType
    predicate.fieldName = fieldName
    predicate.fieldValues = []
    return predicate
}

export function pushFieldValueToPredicate(predicate,fieldVale) {
    predicate.fieldValues.push(fieldVale)
}

export function createConnectionWithoutPredicatesType(predicates){
    let connection = {}
    connection.predicates = predicates
    return connection
}

export function createConnection(predicatesConnectionTypes,predicates){
    let connection = {}
    connection.predicatesConnectionType = predicatesConnectionTypes
    connection.predicates = predicates
    return connection
}

export function createFilterWithoutConnectionType(connections){
    let filter = {}
    filter.connection = connections
    return filter
}

export function createFilter(connectionTypes,connections){
    let filter = {}
    filter.connectionType = connectionTypes
    filter.connection = connections
    return filter
}

export function createOrderBy(fieldName,order){
    let orderBy = {}
    orderBy.fieldName = fieldName
    orderBy.order = order
    return orderBy
}

export function createPaginationQueryCommand(filter,orderBy,pageNumber,rowsPerPage){
    let c = {}
    c.filter = filter
    c.orderBy = orderBy
    c.pageNumber = pageNumber
    c.rowsPerPage = rowsPerPage
    return c
}

export function createQueryCommand(filter,orderBy){
    let c = {}
    c.filter = filter
    c.orderBy = orderBy
    return c
}

export function createEqualQueryCommand(valueType,value,name){
    const fieldValue = createFieldValue(valueType,value)
    const predicate = createPredicate(predicateTypeEnum.PREDICATE_TYPE_EQ,name)
    pushFieldValueToPredicate(predicate,fieldValue)
    let predicates = []
    predicates.push(predicate)
    let connection = createConnectionWithoutPredicatesType(predicates)
    let connects = []
    connects.push(connection)
    let filter = createFilterWithoutConnectionType(connects)
    let order = []
    const query = createQueryCommand(filter,order)
    return query
}

/*
//Sample:

{
    "filter": {
    "connection": [{
        "predicatesConnectionType": [
            "CONNECT_TYPE_AND"
        ],
        "predicates": [{
            "predicateType": "PREDICATE_TYPE_GT",
            "fieldName": "position_among_brothers",
            "fieldValues": [{
                "fieldType": "FIELD_VALUE_TYPE_INTEGER",
                "fieldValueInteger": 100
            }]
        },
            {
                "predicateType": "PREDICATE_TYPE_GT",
                "fieldName": "created_at",
                "fieldValues": [{
                    "fieldType": "FIELD_VALUE_TYPE_TIME",
                    "fieldValueTime": "2022-07-18T00:30:17Z"
                }]
            }
        ]
    },
        {
            "predicatesConnectionType": [
                "CONNECT_TYPE_AND"
            ],
            "predicates": [{
                "predicateType": "PREDICATE_TYPE_BETWEEN",
                "fieldName": "metadata_id",
                "fieldValues": [{
                    "fieldType": "FIELD_VALUE_TYPE_INTEGER",
                    "fieldValueInteger": 152
                }, {
                    "fieldType": "FIELD_VALUE_TYPE_INTEGER",
                    "fieldValueInteger": 154
                }]
            },
                {
                    "predicateType": "PREDICATE_TYPE_LIKE",
                    "fieldName": "logic_object_name",
                    "fieldValues": [{
                        "fieldType": "FIELD_VALUE_TYPE_STRING",
                        "fieldValueString": "%ab%"
                    }]
                }
            ]
        }
    ],
        "connectionType": ["CONNECT_TYPE_OR"]
},
    "orderBy": [{
    "fieldName": "created_at",
    "order": "ORDER_TYPE_DESC"
}]
}

 */

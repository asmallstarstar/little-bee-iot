package queryclause

import (
	query "message/querycommand"
)

func QueryCommandString(q *query.QueryCommand) (string, []interface{}) {
	v := make([]interface{}, 0, 5)
	var s string = ""
	if q.Filter != nil {
		w := WhereClauseString(q, &v)
		if w != "" {
			s = " WHERE " + w
		}
	}
	order := OrderClauseString(q)
	if order != "" {
		s = s + " ORDER BY " + order
	}
	p := Pagination(q, &v)
	if p != "" {
		s = s + p
	}

	return s, v
}

func WhereClauseString(q *query.QueryCommand, v *[]interface{}) string {
	return FilterStatementString(q.Filter, v)
}

func FilterStatementString(f *query.Filter, v *[]interface{}) string {
	var s string = ""
	if len(f.Connection) > 0 && len(f.ConnectionType)+1 == len(f.Connection) {
		for index, value := range f.Connection {
			if index == 0 {
				s = " (" + ConnectionStatementString(value, v) + ") "
			} else {
				if f.ConnectionType[index-1] == query.ConnectionTypeEnum_CONNECT_TYPE_AND {
					s = s + " AND (" + ConnectionStatementString(value, v) + ") "
				}
				if f.ConnectionType[index-1] == query.ConnectionTypeEnum_CONNECT_TYPE_OR {
					s = s + " OR (" + ConnectionStatementString(value, v) + ") "
				}
			}
		}
	}
	return s
}

func ConnectionStatementString(c *query.Connection, v *[]interface{}) string {
	var s string = ""
	if len(c.Predicates) > 0 && len(c.PredicatesConnectionType)+1 == len(c.Predicates) {
		for index, value := range c.Predicates {
			if index == 0 {
				s = PredicateString(value, v)
			} else {
				if c.PredicatesConnectionType[index-1] == query.ConnectionTypeEnum_CONNECT_TYPE_AND {
					s = s + " AND " + PredicateString(value, v)
				}
				if c.PredicatesConnectionType[index-1] == query.ConnectionTypeEnum_CONNECT_TYPE_OR {
					s = s + " OR " + PredicateString(value, v)
				}
			}
		}
	}
	if s != "" {
		return " " + s + " "
	}
	return s
}

func PredicateString(p *query.Predicate, v *[]interface{}) string {
	var w string = ""
	switch p.PredicateType {
	case query.PredicateTypeEnum_PREDICATE_TYPE_UNKNOWN:
	case query.PredicateTypeEnum_PREDICATE_TYPE_EQ:
		if len(p.FieldValues) == 1 {
			w = p.FieldName + " = " + "?"
			ExtractFieldValue(p.FieldValues, v)
		}
	case query.PredicateTypeEnum_PREDICATE_TYPE_LT:
		if len(p.FieldValues) == 1 {
			w = p.FieldName + " < " + "?"
			ExtractFieldValue(p.FieldValues, v)
		}
	case query.PredicateTypeEnum_PREDICATE_TYPE_GT:
		if len(p.FieldValues) == 1 {
			w = p.FieldName + " > " + "?"
			ExtractFieldValue(p.FieldValues, v)
		}
	case query.PredicateTypeEnum_PREDICATE_TYPE_LE:
		if len(p.FieldValues) == 1 {
			w = p.FieldName + " <= " + "?"
			ExtractFieldValue(p.FieldValues, v)
		}
	case query.PredicateTypeEnum_PREDICATE_TYPE_GE:
		if len(p.FieldValues) == 1 {
			w = p.FieldName + " >= " + "?"
			ExtractFieldValue(p.FieldValues, v)
		}
	case query.PredicateTypeEnum_PREDICATE_TYPE_BETWEEN:
		if len(p.FieldValues) == 2 {
			w = "(" + p.FieldName + " BETWEEN " + "?" + " AND " + "?)"
			ExtractFieldValue(p.FieldValues, v)
		}
	case query.PredicateTypeEnum_PREDICATE_TYPE_LIKE:
		if len(p.FieldValues) == 1 {
			w = p.FieldName + ` LIKE? `
			ExtractFieldValue(p.FieldValues, v)
		}
	case query.PredicateTypeEnum_PREDICATE_TYPE_NE:
		if len(p.FieldValues) == 1 {
			w = p.FieldName + " <> " + "?"
			ExtractFieldValue(p.FieldValues, v)
		}
	case query.PredicateTypeEnum_PREDICATE_TYPE_IS_NULL:
		if len(p.FieldValues) == 0 {
			w = p.FieldName + " IS NULL "
		}
	case query.PredicateTypeEnum_PREDICATE_TYPE_IS_NOT_NULL:
		if len(p.FieldValues) == 0 {
			w = p.FieldName + " IS NOT NULL "
		}
	case query.PredicateTypeEnum_PREDICATE_TYPE_IN:
		var s string = ""
		if len(p.FieldValues) > 0 {
			for index := range p.FieldValues {
				if index == 0 {
					s = "?"
				} else {
					s = s + ", ?"
				}
			}
			w = p.FieldName + " IN(" + s + ")"
			ExtractFieldValue(p.FieldValues, v)
		}
	case query.PredicateTypeEnum_PREDICATE_TYPE_NOT_IN:
		var s string = ""
		if len(p.FieldValues) > 0 {
			for index := range p.FieldValues {
				if index == 0 {
					s = "?"
				} else {
					s = s + ", ?"
				}
			}
			w = p.FieldName + " NOT IN(" + s + ")"
			ExtractFieldValue(p.FieldValues, v)
		}
	}
	if w != "" {
		return " " + w
	}
	return w
}

func ExtractFieldValue(fvs []*query.FieldValue, v *[]interface{}) {
	for i := 0; i < len(fvs); i++ {
		fv := fvs[i]
		switch fv.FieldType {
		case query.FieldValueTypeEnum_FIELD_VALUE_TYPE_UNKNOWN:
		case query.FieldValueTypeEnum_FIELD_VALUE_TYPE_INTEGER:
			*v = append(*v, fv.FieldValueInteger)
		case query.FieldValueTypeEnum_FIELD_VALUE_TYPE_FLOAT:
			*v = append(*v, fv.FieldValueFloat)
		case query.FieldValueTypeEnum_FIELD_VALUE_TYPE_STRING:
			*v = append(*v, fv.FieldValueString)
		case query.FieldValueTypeEnum_FIELD_VALUE_TYPE_TIME:
			s := fv.FieldValueTime.AsTime().Format("2006-01-02 03:04:05")
			*v = append(*v, s)
		}
	}
}

func OrderClauseString(q *query.QueryCommand) string {
	var o string = ""
	for index, value := range q.OrderBy {

		if value.Order == query.OrderTypeEnum_ORDER_TYPE_ASC {
			if index != 0 {
				o = o + ", " + value.FieldName + " ASC" + " "
			} else {
				o = value.FieldName + " ASC" + " "
			}
		}

		if value.Order == query.OrderTypeEnum_ORDER_TYPE_DESC {
			if index != 0 {
				o = o + ", " + value.FieldName + " DESC" + " "
			} else {
				o = value.FieldName + " DESC" + " "
			}
		}
	}
	return o
}

func Pagination(q *query.QueryCommand, v *[]interface{}) string {
	if q.PageNumber != 0 && q.RowsPerPage != 0 {
		*v = append(*v, (q.PageNumber-1)*q.RowsPerPage)
		*v = append(*v, q.RowsPerPage)
		return " LIMIT ?,? "
	} else {
		return ""
	}
}

package error

type MessageCodeValueEnum int32

const (
	SUCCESS                    MessageCodeValueEnum = 0
	ERROR_URL_QUERY_PARAMETER  MessageCodeValueEnum = 1
	ERROR_URL_PATH_PARAMETER   MessageCodeValueEnum = 2
	ERROR_BODY_PARSING         MessageCodeValueEnum = 3
	ERROR_WRITE_DATABASE       MessageCodeValueEnum = 4
	ERROR_READ_DATABASE        MessageCodeValueEnum = 5
	ERROR_OPEN_DATABAE         MessageCodeValueEnum = 6
	ERROR_TOKEN_NOT_FOUND      MessageCodeValueEnum = 7
	ERROR_TOKEN_HAS_EXPIRED    MessageCodeValueEnum = 8
	ERROR_TOKEN_PARSED         MessageCodeValueEnum = 9
	ERROR_TOKEN_GENERATE       MessageCodeValueEnum = 10
	ERROR_TOKEN_EXCEPTION      MessageCodeValueEnum = 11
	ERROR_USERNAME_OR_PASSWORD MessageCodeValueEnum = 12
	ERROR_CASBIN_ADAPTER       MessageCodeValueEnum = 13
	ERROR_CASBIN_NEW_ENFORCER  MessageCodeValueEnum = 14
	ERROR_CASBIN_LOAD_POLICY   MessageCodeValueEnum = 15
	ERROR_CASBIN_GRANT         MessageCodeValueEnum = 16
	SUCCESS_GRANT              MessageCodeValueEnum = 17
	FAIL_GRANT                 MessageCodeValueEnum = 18
	SUCCESS_JOIN_GROUP         MessageCodeValueEnum = 19
	FAIL_JOIN_GROUP            MessageCodeValueEnum = 20
	FAIL_READ_ACTION_TABLE     MessageCodeValueEnum = 21
	FAIL_JUDGMENT_PERMISSION   MessageCodeValueEnum = 22
	FAIL_PERMISSION_DENIED     MessageCodeValueEnum = 23
	FAIL_NOTFOUND_ACTION       MessageCodeValueEnum = 24
	FAIL_MODIFY_PASSWORD       MessageCodeValueEnum = 25
)

const ERROR_DESC_STRING string = "error desc"

// Enum value maps for ErrorCodeValueEnum.
var (
	MessageCodeValueEnum_desc_id = map[int32]string{
		0:  "id_success",
		1:  "id_error_url_query_parameter",
		2:  "id_error_url_path_parameter",
		3:  "id_error_body_parsing",
		4:  "id_error_write_database",
		5:  "id_error_read_database",
		6:  "id_error_open_database",
		7:  "id_error_token_not_found",
		8:  "id_error_token_has_expired",
		9:  "id_error_token_parsed",
		10: "id_error_token_generate",
		11: "id_error_token_exception",
		12: "id_error_user_name_or_password",
		13: "id_error_casbin_adapter",
		14: "id_error_casbin_new_enforcer",
		15: "id_error_casbin_load_policy",
		16: "id_error_casbin_grant",
		17: "id_success_grant",
		18: "id_fail_grant",
		19: "id_success_join_group",
		20: "id_fail_join_group",
		21: "id_fail_read_action_table",
		22: "id_fail_judgment_permission",
		23: "id_fail_permission_denied",
		24: "id_fail_notfound_action",
		25: "id_fail_modify_password",
	}
	MessageCodeValueEnum_value = map[string]int32{
		"SUCCESS":                    0,
		"ERROR_URL_QUERY_PARAMETER":  1,
		"ERROR_URL_PATH_PARAMETER":   2,
		"ERROR_BODY_PARSING":         3,
		"ERROR_WRITE_DATABASE":       4,
		"ERROR_READ_DATABASE":        5,
		"ERROR_OPEN_DATABAE":         6,
		"ERROR_TOKEN_NOT_FOUND":      7,
		"ERROR_TOKEN_HAS_EXPIRED":    8,
		"ERROR_TOKEN_PARSED":         9,
		"ERROR_TOKEN_GENERATE":       10,
		"ERROR_TOKEN_EXCEPTION":      11,
		"ERROR_USERNAME_OR_PASSWORD": 12,
		"ERROR_CASBIN_ADAPTER":       13,
		"ERROR_CASBIN_NEW_ENFORCER":  14,
		"ERROR_CASBIN_LOAD_POLICY":   15,
		"ERROR_CASBIN_GRANT":         16,
		"SUCCESS_GRANT":              17,
		"FAIL_GRANT":                 18,
		"SUCCESS_JOIN_GROUP":         19,
		"FAIL_JOIN_GROUP":            20,
		"FAIL_READ_ACTION_TABLE":     21,
		"FAIL_JUDGMENT_PERMISSION":   22,
		"FAIL_PERMISSION_DENIED":     23,
		"FAIL_NOTFOUND_ACTION":       24,
		"FAIL_MODIFY_PASSWORD":       25,
	}
)

package test

import (
	"controller/testengine"
	"message/littlebee"
	"net/http"
	"testing"
)

func TestCreateLogicObject(t *testing.T) {
	e, token := testengine.GetTestEngine(t)

	t.Log("create logic object")
	e.POST("/little-bee/config/logic-object").
		WithHeader("token", token).
		WithJSON(&littlebee.LogicObject{
			LogicObjectName:       "logic object name",
			ParentLogicObjectId:   1,
			LogicObjectTypeId:     2,
			MetadataId:            3,
			ConfigureId:           4,
			PositionAmongBrothers: 5,
		}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("isSuccess", true)
}

func TestRetrieveLogicObject(t *testing.T) {
	e, token := testengine.GetTestEngine(t)

	t.Log("retrieve logic object")
	e.GET("/little-bee/config/logic-object/1").
		WithHeader("token", token).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("isSuccess", true)
}

func TestUpdateLogicObject(t *testing.T) {
	e, token := testengine.GetTestEngine(t)

	t.Log("update logic object")
	e.PUT("/little-bee/config/logic-object").
		WithHeader("token", token).
		WithJSON(&littlebee.LogicObject{
			LogicObjectId:         1,
			LogicObjectName:       "logic object name",
			ParentLogicObjectId:   1,
			LogicObjectTypeId:     2,
			MetadataId:            3,
			ConfigureId:           4,
			PositionAmongBrothers: 5,
		}).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("isSuccess", true)
}

func TestDeleteLogicObject(t *testing.T) {
	e, token := testengine.GetTestEngine(t)

	t.Log("delete logic object")
	e.DELETE("/little-bee/config/logic-object/2").
		WithHeader("token", token).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("isSuccess", true)
}

func TestDeleteLogicObjectWithFlag(t *testing.T) {
	e, token := testengine.GetTestEngine(t)

	t.Log("set logic object deleted flag")
	e.PUT("/little-bee/config/logic-object/3").
		WithHeader("token", token).
		Expect().
		Status(http.StatusOK).
		JSON().Object().
		ValueEqual("isSuccess", true)
}

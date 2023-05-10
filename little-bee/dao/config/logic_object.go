package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
	"gorm.io/gorm"
)

func CreateLogicObject(l *littlebee.LogicObjectExtend) (*littlebee.LogicObjectExtend, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if l.LogicConfigure != nil {
			if err := tx.Create(l.LogicConfigure).Error; err != nil {
				return err
			}
			l.LogicObject.ConfigureId = l.LogicConfigure.ConfigureId
		} else {
			l.LogicObject.ConfigureId = 0
		}

		if err := tx.Omit("LogicObjectId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l.LogicObject).Error; err != nil {
			return err
		}
		if l.LogicArea != nil {
			l.LogicArea.LogicObjectId = l.LogicObject.LogicObjectId
			if err := tx.Create(l.LogicArea).Error; err != nil {
				return err
			}
		}
		if l.LogicDevice != nil {
			l.LogicDevice.LogicObjectId = l.LogicObject.LogicObjectId
			if err := tx.Create(l.LogicDevice).Error; err != nil {
				return err
			}
		}

		return nil
	})
	return l, err
}

func RetrieveLogicObject(logicObjectId int32) (*littlebee.LogicObject, error) {
	r := &littlebee.LogicObject{}
	result := database.DB.First(&r, logicObjectId)
	return r, result.Error
}

func UpdateLogicObject(l *littlebee.LogicObjectExtend) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if l.LogicConfigure != nil {
			if l.LogicConfigure.ConfigureId == 0 {
				if err := tx.Create(l.LogicConfigure).Error; err != nil {
					return err
				}
				l.LogicObject.ConfigureId = l.LogicConfigure.ConfigureId
			} else {
				m := structs.Map(l.LogicConfigure)
				if err := tx.Model(l.LogicConfigure).Updates(m).Error; err != nil {
					return err
				}
			}
		}

		m := structs.Map(l.LogicObject)
		m["created_at"] = l.LogicObject.CreatedAt.AsTime().UTC()
		m["updated_at"] = l.LogicObject.UpdatedAt.AsTime().UTC()
		m["deleted_at"] = l.LogicObject.DeletedAt.AsTime().UTC()
		if err := tx.Model(l.LogicObject).Omit("LogicObjectId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(m).Error; err != nil {
			return err
		}
		if l.LogicArea != nil {
			m = structs.Map(l.LogicArea)
			if err := tx.Model(l.LogicArea).Updates(m).Error; err != nil {
				return err
			}
		}
		if l.LogicDevice != nil {
			m = structs.Map(l.LogicDevice)
			if err := tx.Model(l.LogicDevice).Updates(m).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func DeleteLogicObjectWithFlag(l *littlebee.LogicObject) error {
	result := database.DB.Model(l).Updates(littlebee.LogicObject{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteLogicObject(l *littlebee.LogicObject) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		r := &littlebee.LogicObject{}
		if err := tx.First(&r, l.LogicObjectId).Error; err != nil {
			return err
		}
		if r.ConfigureId != 0 {
			c := &littlebee.Configure{
				ConfigureId: r.ConfigureId,
			}
			if err := tx.Delete(c).Error; err != nil {
				return err
			}
		}
		a := &littlebee.LogicArea{
			LogicObjectId: r.LogicObjectId,
		}
		if err := tx.Delete(a).Error; err != nil {
			return err
		}
		d := &littlebee.LogicDevice{
			LogicObjectId: r.LogicObjectId,
		}
		if err := tx.Delete(d).Error; err != nil {
			return err
		}
		if err := tx.Delete(l).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func QueryLogicObject(q *querycommand.QueryCommand) (*littlebee.LogicObjectList, error) {
	var logicObjects []*littlebee.LogicObject
	s := "SELECT * FROM logic_objects "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&logicObjects)
	return &littlebee.LogicObjectList{Items: logicObjects}, result.Error
}

func GetAllLogicObjects() (*littlebee.LogicObjectList, error) {
	var logicObjects []*littlebee.LogicObject
	result := database.DB.Find(&logicObjects)
	return &littlebee.LogicObjectList{Items: logicObjects}, result.Error
}

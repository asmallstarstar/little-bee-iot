package config

import (
	"dao/database"
	"message/littlebee"

	"github.com/fatih/structs"
	"gorm.io/gorm"
)

func CreateVsi(l *littlebee.VSi) (*littlebee.VSi, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("SignalId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l.Signal).Error; err != nil {
			return err
		}

		l.SignalSi.SignalId = l.Signal.SignalId
		l.SignalVirtual.SignalId = l.Signal.SignalId
		if err := tx.Create(l.SignalVirtual).Error; err != nil {
			return err
		}
		if err := tx.Create(l.SignalSi).Error; err != nil {
			return err
		}
		return nil
	})
	return l, err
}

func RetrieveVsi(vsiId int32) (*littlebee.VSi, error) {
	r := &littlebee.VSi{}
	if err := database.DB.First(r.Signal, vsiId).Error; err != nil {
		return nil, err
	}
	if err := database.DB.First(r.SignalVirtual, vsiId).Error; err != nil {
		return nil, err
	}
	if err := database.DB.First(r.SignalSi, vsiId).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func UpdateVsi(l *littlebee.VSi) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		m := structs.Map(l.Signal)
		m["created_at"] = l.Signal.CreatedAt.AsTime().UTC()
		m["updated_at"] = l.Signal.UpdatedAt.AsTime().UTC()
		m["deleted_at"] = l.Signal.DeletedAt.AsTime().UTC()
		if err := tx.Model(l.Signal).Omit("SignalId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(m).Error; err != nil {
			return err
		}
		m = structs.Map(l.SignalVirtual)
		if err := tx.Model(l.SignalVirtual).Updates(m).Error; err != nil {
			return err
		}
		m = structs.Map(l.SignalSi)
		if err := tx.Model(l.SignalSi).Updates(m).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func DeleteVsiWithFlag(l *littlebee.VSi) error {
	result := database.DB.Model(l).Updates(littlebee.Signal{
		IsDeleted: l.Signal.IsDeleted,
		DeletedAt: l.Signal.DeletedAt,
		DeletedBy: l.Signal.DeletedBy,
	})
	return result.Error
}

func DeleteVsi(l *littlebee.VSi) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(l.Signal).Error; err != nil {
			return err
		}
		l.SignalSi.SignalId = l.Signal.SignalId
		l.SignalVirtual.SignalId = l.Signal.SignalId
		if err := tx.Delete(l.SignalVirtual).Error; err != nil {
			return err
		}
		if err := tx.Delete(l.SignalSi).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

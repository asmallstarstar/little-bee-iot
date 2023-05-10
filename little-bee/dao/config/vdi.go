package config

import (
	"dao/database"
	"message/littlebee"

	"github.com/fatih/structs"
	"gorm.io/gorm"
)

func CreateVdi(l *littlebee.VDi) (*littlebee.VDi, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("SignalId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l.Signal).Error; err != nil {
			return err
		}

		l.SignalDi.SignalId = l.Signal.SignalId
		l.SignalVirtual.SignalId = l.Signal.SignalId
		if err := tx.Create(l.SignalVirtual).Error; err != nil {
			return err
		}
		if err := tx.Create(l.SignalDi).Error; err != nil {
			return err
		}
		st := &littlebee.SignalThreshold{}
		if err := tx.Where("signal_id = ?", l.Signal.SignalId).Delete(st).Error; err != nil {
			return err
		}
		for _, v := range l.SignalThreshold {
			v.SignalId = l.Signal.SignalId
			if err := tx.Create(v).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return l, err
}

func RetrieveVdi(vdiId int32) (*littlebee.VDi, error) {
	r := &littlebee.VDi{}
	if err := database.DB.First(r.Signal, vdiId).Error; err != nil {
		return nil, err
	}
	if err := database.DB.First(r.SignalVirtual, vdiId).Error; err != nil {
		return nil, err
	}
	if err := database.DB.First(r.SignalDi, vdiId).Error; err != nil {
		return nil, err
	}
	if err := database.DB.Where("signal_id = ?", vdiId).Find(r.SignalThreshold).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func UpdateVdi(l *littlebee.VDi) error {
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
		m = structs.Map(l.SignalDi)
		if err := tx.Model(l.SignalDi).Updates(m).Error; err != nil {
			return err
		}
		st := &littlebee.SignalThreshold{}
		if err := tx.Where("signal_id = ?", l.Signal.SignalId).Delete(st).Error; err != nil {
			return err
		}
		for _, v := range l.SignalThreshold {
			v.SignalId = l.Signal.SignalId
			if err := tx.Create(v).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func DeleteVdiWithFlag(l *littlebee.VDi) error {
	result := database.DB.Model(l).Updates(littlebee.Signal{
		IsDeleted: l.Signal.IsDeleted,
		DeletedAt: l.Signal.DeletedAt,
		DeletedBy: l.Signal.DeletedBy,
	})
	return result.Error
}

func DeleteVdi(l *littlebee.VDi) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(l.Signal).Error; err != nil {
			return err
		}
		l.SignalDi.SignalId = l.Signal.SignalId
		l.SignalVirtual.SignalId = l.Signal.SignalId
		if err := tx.Delete(l.SignalVirtual).Error; err != nil {
			return err
		}
		if err := tx.Delete(l.SignalDi).Error; err != nil {
			return err
		}
		st := &littlebee.SignalThreshold{}
		if err := tx.Where("signal_id = ?", l.Signal.SignalId).Delete(st).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

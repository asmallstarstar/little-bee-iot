package config

import (
	"dao/database"
	"message/littlebee"

	"github.com/fatih/structs"
	"gorm.io/gorm"
)

func CreateSi(l *littlebee.Si) (*littlebee.Si, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if l.SignalBindConfigure != nil {
			if err := tx.Create(l.SignalBindConfigure).Error; err != nil {
				return err
			}
			l.SignalAcquisition.SignalBindConfigureId = l.SignalBindConfigure.ConfigureId
		} else {
			l.SignalAcquisition.SignalBindConfigureId = 0
		}

		if err := tx.Omit("SignalId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l.Signal).Error; err != nil {
			return err
		}

		l.SignalSi.SignalId = l.Signal.SignalId
		l.SignalAcquisition.SignalId = l.Signal.SignalId
		if err := tx.Create(l.SignalAcquisition).Error; err != nil {
			return err
		}
		if err := tx.Create(l.SignalSi).Error; err != nil {
			return err
		}
		return nil
	})
	return l, err
}

func RetrieveSi(siId int32) (*littlebee.Si, error) {
	r := &littlebee.Si{}
	if err := database.DB.First(r.Signal, siId).Error; err != nil {
		return nil, err
	}
	if err := database.DB.First(r.SignalAcquisition, siId).Error; err != nil {
		return nil, err
	}
	if err := database.DB.First(r.SignalSi, siId).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func UpdateSi(l *littlebee.Si) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if l.SignalBindConfigure != nil {
			if l.SignalBindConfigure.ConfigureId == 0 {
				if err := tx.Create(l.SignalBindConfigure).Error; err != nil {
					return err
				}
				l.SignalAcquisition.SignalBindConfigureId = l.SignalBindConfigure.ConfigureId
			} else {
				m := structs.Map(l.SignalBindConfigure)
				m["created_at"] = l.Signal.CreatedAt.AsTime().UTC()
				m["updated_at"] = l.Signal.UpdatedAt.AsTime().UTC()
				m["deleted_at"] = l.Signal.DeletedAt.AsTime().UTC()
				if err := tx.Model(l.SignalBindConfigure).Updates(m).Error; err != nil {
					return err
				}
			}
		}

		m := structs.Map(l.Signal)
		m["created_at"] = l.Signal.CreatedAt.AsTime().UTC()
		m["updated_at"] = l.Signal.UpdatedAt.AsTime().UTC()
		m["deleted_at"] = l.Signal.DeletedAt.AsTime().UTC()
		if err := tx.Model(l.Signal).Omit("SignalId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(m).Error; err != nil {
			return err
		}
		m = structs.Map(l.SignalAcquisition)
		if err := tx.Model(l.SignalAcquisition).Updates(m).Error; err != nil {
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

func DeleteSiWithFlag(l *littlebee.Si) error {
	result := database.DB.Model(l).Updates(littlebee.Signal{
		IsDeleted: l.Signal.IsDeleted,
		DeletedAt: l.Signal.DeletedAt,
		DeletedBy: l.Signal.DeletedBy,
	})
	return result.Error
}

func DeleteSi(l *littlebee.Si) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {

		r := &littlebee.SignalAcquisition{
			SignalId: l.Signal.SignalId,
		}
		if err := tx.First(&r, r.SignalId).Error; err != nil {
			return err
		}
		if r.SignalBindConfigureId != 0 {
			c := &littlebee.Configure{
				ConfigureId: r.SignalBindConfigureId,
			}
			if err := tx.Delete(c).Error; err != nil {
				return err
			}
		}

		if err := tx.Delete(l.Signal).Error; err != nil {
			return err
		}
		l.SignalSi.SignalId = l.Signal.SignalId
		l.SignalAcquisition.SignalId = l.Signal.SignalId
		if err := tx.Delete(l.SignalAcquisition).Error; err != nil {
			return err
		}
		if err := tx.Delete(l.SignalSi).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

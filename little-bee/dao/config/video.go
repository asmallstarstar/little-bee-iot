package config

import (
	"dao/database"
	"message/littlebee"

	"github.com/fatih/structs"
	"gorm.io/gorm"
)

func CreateVideo(l *littlebee.Video) (*littlebee.Video, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if l.VideoBindConfigure != nil {
			if err := tx.Create(l.VideoBindConfigure).Error; err != nil {
				return err
			}
			l.SignalVideo.VideoBindConfigureId = l.VideoBindConfigure.ConfigureId
		} else {
			l.SignalVideo.VideoBindConfigureId = 0
		}

		if err := tx.Omit("SignalId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l.Signal).Error; err != nil {
			return err
		}

		l.SignalVideo.SignalId = l.Signal.SignalId
		if err := tx.Create(l.SignalVideo).Error; err != nil {
			return err
		}
		return nil
	})
	return l, err
}

func RetrieveVideo(videoId int32) (*littlebee.Video, error) {
	r := &littlebee.Video{}
	if err := database.DB.First(r.Signal, videoId).Error; err != nil {
		return nil, err
	}
	if err := database.DB.First(r.SignalVideo, videoId).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func UpdateVideo(l *littlebee.Video) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if l.VideoBindConfigure != nil {
			if l.VideoBindConfigure.ConfigureId == 0 {
				if err := tx.Create(l.VideoBindConfigure).Error; err != nil {
					return err
				}
				l.SignalVideo.VideoBindConfigureId = l.VideoBindConfigure.ConfigureId
			} else {
				m := structs.Map(l.VideoBindConfigure)
				m["created_at"] = l.Signal.CreatedAt.AsTime().UTC()
				m["updated_at"] = l.Signal.UpdatedAt.AsTime().UTC()
				m["deleted_at"] = l.Signal.DeletedAt.AsTime().UTC()
				if err := tx.Model(l.VideoBindConfigure).Updates(m).Error; err != nil {
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
		m = structs.Map(l.SignalVideo)
		if err := tx.Model(l.SignalVideo).Updates(m).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func DeleteVideoWithFlag(l *littlebee.Video) error {
	result := database.DB.Model(l).Updates(littlebee.Signal{
		IsDeleted: l.Signal.IsDeleted,
		DeletedAt: l.Signal.DeletedAt,
		DeletedBy: l.Signal.DeletedBy,
	})
	return result.Error
}

func DeleteVideo(l *littlebee.Video) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {

		r := &littlebee.SignalVideo{
			SignalId: l.Signal.SignalId,
		}
		if err := tx.First(&r, r.SignalId).Error; err != nil {
			return err
		}
		if r.VideoBindConfigureId != 0 {
			c := &littlebee.Configure{
				ConfigureId: r.VideoBindConfigureId,
			}
			if err := tx.Delete(c).Error; err != nil {
				return err
			}
		}

		if err := tx.Delete(l.Signal).Error; err != nil {
			return err
		}
		l.SignalVideo.SignalId = l.Signal.SignalId
		if err := tx.Delete(l.SignalVideo).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

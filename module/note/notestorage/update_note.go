package notestorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

func (store *sqlStore) Update(ctx context.Context, id int, data *notemodel.NoteUpdate) error {
	db := store.db

	if err := db.Table(data.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

package notestorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

func (store *sqlStore) CreateNote(context context.Context, data *notemodel.NoteCreate) error {
	db := store.db

	if err := db.Table(data.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

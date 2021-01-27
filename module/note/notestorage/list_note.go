package notestorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/module/note/notemodel"
)

func (store *sqlStore) ListNote(
	context context.Context,
	filter *notemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]notemodel.Note, error) {
	db := store.db
	var notes []notemodel.Note

	db = db.Table(notemodel.Note{}.TableName()).Where("status not in (0)")

	if v := filter; v != nil {
		if v.CategoryId > 0 {
			db = db.Where("category_id = ?", v.CategoryId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db = db.Limit(paging.Limit)

	for _, k := range moreKeys {
		db = db.Preload(k)
	}

	if paging.Cursor > 0 {
		db = db.Where("id < ?", paging.Cursor)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Order("id desc").Find(&notes).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return notes, nil
}

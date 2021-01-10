package notestorage

import "fooddlv/module/note/notemodel"

func (store *sqlStore) ListNote() ([]notemodel.Note, error) {
	db := store.db
	var notes []notemodel.Note

	if err := db.Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

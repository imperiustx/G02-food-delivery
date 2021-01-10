package notestorage

import "fooddlv/module/note/notemodel"

func (store *sqlStore) FindNote(id int) (*notemodel.Note, error) {
	db := store.db
	var note notemodel.Note

	if err := db.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, err
	}

	return &note, nil
}

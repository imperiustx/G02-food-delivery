package notestorage

func (store *sqlStore) DeleteNote(id int) error {
	db := store.db

	if err := db.Table("notes").Where("id = ?", id).Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}

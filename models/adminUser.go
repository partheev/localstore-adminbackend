package models

func (db *DBModel) GetAdminUser(id int) AdminUser {
	stmt := `SELECT id,first_name,last_name,email,phone_number FROM adminusers WHERE id = $1`
	row := db.client.QueryRow(stmt, id)
	var aU AdminUser
	row.Scan(
		&aU.Id,
		&aU.First_name,
		&aU.Last_name,
		&aU.Email,
		&aU.Phone_number,
	)
	return aU
}

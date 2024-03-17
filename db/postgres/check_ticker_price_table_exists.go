package postgres

import "log"

func (db *Database) CheckTickerPriceTableExists(table string) (bool, error) {

	rows, err := db.conn.Raw(
		"SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'public';",
	).Rows()

	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		var existingTable string
		err := rows.Scan(&existingTable)
		if err != nil {
			return false, err
		}
		if existingTable == table {
			log.Printf("%s price table already exists...fetching from database.", table)
			return true, nil
		}
	}
	log.Printf("%s does not exist in db...fetching from external API.", table)
	return false, nil
}

package data

import "database/sql"

type URLData struct {
	OriginalUrl string
	ShortUrl    string
	Clicks      int
}

type URLDataModel struct {
	DB *sql.DB
}

func (um *URLDataModel) GetAll() ([]*URLData, error) {
	stmt := `SELECT original_url, shortened_url, clicks from urls`

	rows, err := um.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	urls := []*URLData{}

	for rows.Next() {
		url := &URLData{}
		err := rows.Scan(&url.OriginalUrl, &url.ShortUrl, &url.Clicks)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

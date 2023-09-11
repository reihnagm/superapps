package models

import (
	"errors"
	helper "superapps/helpers"
)

type All struct {
	Uid			string `json:"id"`
	Title 		string `json:"title"`
	Description string `json:"description"`
}


func (n *News) GetNews() (map[string]interface{}, error) {

	var news All
	var assign All

	var data []All

	query := `SELECT uid, title, description FROM news`

	// var data []All
	// var rest All

	rows, err := db.Debug().Raw(query).Rows()

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, errors.New(err.Error())
	}

	for rows.Next() {
		db.ScanRows(rows, &news)

		assign.Uid = news.Uid
		assign.Title = news.Title
		assign.Description = news.Description

		data = append(data, assign)
	}

	return map[string]interface{}{"news": &data}, nil
}
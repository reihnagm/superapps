package services

import (
	"os"
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
	"strconv"
	"errors"
	"math"
	models "superapps/models"
	helper "superapps/helpers"
)

func GetMembernear(originLat, originLng, appName string) (map[string]interface{}, error) {

	var allMembernear []models.Membernear

	var membernearAssign models.MembernearAssign
	var appendMembernearAssign = make([]models.MembernearAssign, 0) 

	var applications []models.Application
	errCheckApp := db.Debug().Raw(`SELECT uid, username FROM applications WHERE username = '`+appName+`'`).Scan(&applications).Error
	
	if errCheckApp != nil {
		helper.Logger("error", "In Server: "+errCheckApp.Error())
		return nil, errors.New(errCheckApp.Error())
	}

	isAppExist := len(applications)

	if isAppExist == 0 {
		return nil, errors.New("App not found")
	} 

	ApplicationId := applications[0].Uid

	errAllMembernearQuery := db.Debug().Raw(`SELECT lat, lng FROM fcms WHERE app_id = '`+ApplicationId+`' `).Scan(&allMembernear).Error

	if errAllMembernearQuery != nil {
		helper.Logger("error", "In Server: "+errAllMembernearQuery.Error())
		return nil, errors.New(errAllMembernearQuery.Error())
	}

	for i, _ := range allMembernear {

		var destLat = fmt.Sprint(allMembernear[i].Lat) // Format as string
		var destLng = fmt.Sprint(allMembernear[i].Lng) // Format as string

		var url = os.Getenv("GMAP_URL") + "?origin=" + originLat + "," + originLng + "&destination=" + destLat + "," + destLng + "&key=" + os.Getenv("GMAP_KEY")

		resp, errRes := http.NewRequest(http.MethodGet, url, nil)  

		if errRes != nil {
			helper.Logger("error", "In Server: "+errRes.Error())
		}

		var GmapsResp models.GmapsDirection

		if err := json.NewDecoder(resp.Body).Decode(&GmapsResp); err != nil {
			helper.Logger("error", "In Server: "+err.Error())
		 }

		for z, _ := range GmapsResp.Routes {
			var txt = GmapsResp.Routes[z].Legs[0].Distance.Text

			parseDistance, _ := strconv.ParseFloat(strings.Split(txt, " ")[0], 64)

			distance := math.Ceil(parseDistance)

			if distance <= 5 {	
				membernearAssign.Lat = allMembernear[i].Lat
				membernearAssign.Lng = allMembernear[i].Lng
				membernearAssign.Distance = txt	
				
				appendMembernearAssign = append(appendMembernearAssign, membernearAssign)
			} 
		}
	}

	return map[string]any {
		"data": appendMembernearAssign,
	}, nil
}
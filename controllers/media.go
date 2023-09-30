package controllers

import (
	"net/http"
	"io"
	"os"
	"path/filepath"
	helper "superapps/helpers"
)

type MediaResponseEntity struct {
	Path 	 string `json:"path"`
	Filename string `json:"filename"`
	Size 	 int64 	`json:"size"`
	Mime     string `json:"mime"`
}

func Upload(w http.ResponseWriter, r *http.Request) {

	url := os.Getenv("API_URL")

	err := r.ParseMultipartForm(10 << 20) // Max 10 MB
	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]interface{}{})
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]interface{}{})
		return
	}
	defer file.Close()

	parentFolder := "public"
	subFolder := r.FormValue("folder")

 	os.Mkdir(parentFolder, os.ModePerm)

	subFolderPath := filepath.Join(parentFolder, subFolder)

	os.Mkdir(subFolderPath, os.ModePerm)

	filePath := filepath.Join(subFolderPath, handler.Filename)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]interface{}{})
		return
	}
	defer f.Close()

	io.Copy(f, file)

	var mediaAssign MediaResponseEntity

	mediaAssign.Path = url + "/" + subFolder + "/" + handler.Filename
	mediaAssign.Filename = handler.Filename
	mediaAssign.Size = handler.Size 
	mediaAssign.Mime = handler.Header.Get("Content-Type")

	helper.Logger("info", "Upload success")
	helper.Response(w, http.StatusOK, false, "Successfully", mediaAssign)
}
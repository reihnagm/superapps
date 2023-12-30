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

 	errParentMkdir := os.Mkdir(parentFolder, os.ModePerm)

	if errParentMkdir != nil {
		helper.Response(w, 400, true, errParentMkdir.Error(), map[string]interface{}{})
		return
	}

	subFolderPath := filepath.Join(parentFolder, subFolder)

	errSubFolder := os.Mkdir(subFolderPath, os.ModePerm)
	
	if errSubFolder != nil {
		helper.Response(w, 400, true, errSubFolder.Error(), map[string]interface{}{})
		return
	}

	filePath := filepath.Join(subFolderPath, handler.Filename)
	f, err := os.OpenFile(filepath.Clean(filePath), os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		helper.Response(w, 400, true, err.Error(), map[string]interface{}{})
		return
	}

	defer f.Close()

	_, errCopy := io.Copy(f, file)

	if errCopy != nil {
		helper.Response(w, 400, true, errCopy.Error(), map[string]interface{}{})
		return
	}

	var mediaAssign MediaResponseEntity

	mediaAssign.Path = url + "/" + subFolder + "/" + handler.Filename
	mediaAssign.Filename = handler.Filename
	mediaAssign.Size = handler.Size 
	mediaAssign.Mime = handler.Header.Get("Content-Type")

	helper.Logger("info", "Upload success")
	helper.Response(w, http.StatusOK, false, "Successfully", mediaAssign)
}
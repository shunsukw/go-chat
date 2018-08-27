package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/shunsukw/go-chat/common/utility"
	"github.com/shunsukw/go-chat/tasks"
)

type UploadImageForm struct {
	PageTitle  string
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		DisplayUploadImageForm(w, r, &u)
	case "POST":
		ValidateUploadImageForm(w, r, &u)
	default:
		DisplayUploadImageForm(w, r, &u)
	}
}

// GET--------------------
// DisplayUploadImageForm ...
func DisplayUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {
	RenderGatedTemplate(w, "./templates/uploadimageform.html", u)
}

// POST
// ValidateUploadImageForm ...
func ValidateUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {
	ProcessUploadImage(w, r, u)
}

func ProcessUploadImage(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {
	shouldProcessThumbnailAsynchronously := false

	file, fileheader, err := r.FormFile("imagefile")
	if err != nil {
		log.Println("Encountered error when attempting to read uploaded file: ", err)
	}

	randomFileName := utility.GenerateUUID()

	if fileheader != nil {
		extension := filepath.Ext(fileheader.Filename)
		// TODO: ここで何をやっているのかわからない。
		r.ParseMultipartForm(32 << 20)

		defer file.Close()

		imageFilePathWithoutExtensions := "./static/uploads/images/" + randomFileName
		f, err := os.OpenFile(imageFilePathWithoutExtension+extension, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)

		thumbnailResizeTask := tasks.NewImageResizeTask(imageFilePathWithoutExtension, extension)

		// Thumbnail画像を非同期でリサイズ作成をする処理をここにかく
		if shouldProcessThumbnailAsynchronously == true {
			// async 処理

		} else {
			thumbnailResizeTask.Perform()
		}
		// ------------------------------------------------

		m := make(map[string]string)
		m["thumbnailPath"] = strings.TrimPrefix(imageFilePathWithoutExtension, ".") + "_thumb.png"
		m["imagePath"] = strings.TrimPrefix(imageFilePathWithoutExtension, ".") + ".png"
		m["PageTitle"] = "Image Preview"

		RenderGatedTemplate(w, "./templates/imagepreview.html", m)
	} else {
		w.Write([]byte("Failed to process uploaded file!"))
	}
}

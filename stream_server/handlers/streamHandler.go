package handlers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/just-lym/video_server/stream_server/customlog"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	VideoDir      = "./videoFile/"
	MaxUploadSize = 10
)

func StreamHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vid := ps.ByName("vid-id")
	vl := VideoDir + vid
	video, err := os.Open(vl)
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func UploadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 设置request body缓冲区最大容量，如果超过最大容量，将会关闭reader
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "file is too large")
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		customlog.Errorln("read file error:%s", err)
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	videoName := ps.ByName("vid-id")
	err = os.WriteFile(VideoDir+videoName, data, 0666)
	if err != nil {
		customlog.Errorln("write file error%s", err.Error())
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Uploaded Successfully")
}

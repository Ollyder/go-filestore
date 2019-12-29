package handler

import (
	"encoding/json"
	"fmt"
	"go-filestore/meta"
	"go-filestore/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// UploadHandler is used to handler upload
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// return input html
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// receive input file stream
		fromFile, head, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Get upload file error, because: %s\n", err.Error())
		}
		// 保证文件一定被关闭
		defer fromFile.Close()
		fileLocation := "./tmp/" + head.Filename
		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: fileLocation,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"), // go 的格式化也太奇葩了
		}

		toFile, err := os.Create(fileLocation)
		if err != nil {
			fmt.Printf("create tmp file err : %s\n", err.Error())
		}
		defer toFile.Close()
		fileMeta.FileSize, err = io.Copy(toFile, fromFile)
		if err != nil {
			fmt.Printf("save file err : %s\n", err.Error())
		}
		toFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(toFile)
		meta.UpdateFileMeta(fileMeta)
		UploadSucHandler(w, r)

	}

}

// UploadSucHandler : return success info
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished")
}

func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileHash := r.Form["filehash"][0] // r.Form["xxx"] 对应的是一个数组
	fMeta := meta.GetFileMeta(fileHash)
	data, err := json.Marshal(fMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)

	f, err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octect-stream")
	// 指定此文件下载名
	w.Header().Set("content-disposition", "attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}

func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")

	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	curFileMeta := meta.GetFileMeta(fileSha1)
	curFileMEta.FileName = newFileName
	meta.UpdateFileMeta(curFileMeta)

	data, err := json.Marshal(curFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// TODO fileMeta删除
}

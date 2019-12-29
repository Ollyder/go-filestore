package db

import (
	"fmt"
	mydb "go-filestore/db/mysql"
)

func OnFileUploadFinished(filehash string, filename string,
	filesize int64, fileaddr string) bool {
	stme, err := mydb.DBConn().Prepare(
		"INSERT IGNORE INTO tbl_file(`file_sha1`, `file_name`, `file_size`" +
			"`file_addr`, `status`) VALUES(?,?,?,?,1)")
	if err != nil {
		fmt.Printf("Prepare staement err : %s \n", err.Error())
		return false
	}

	defer stme.Close()

	ret, err := stme.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(errError())
		return false
	}

	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Printf("File with hash has been upload before: %s\n", filehash)
		}
		return true
	}
	return false
}

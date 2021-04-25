package db

import "os"

/*
DB_File 数据持久化
*/
type DB_File struct {
	FileName string
	File     *os.File
}

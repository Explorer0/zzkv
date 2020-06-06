package bitcask

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

const DefaultFileMode os.FileMode = 0666


type Position struct {
	FileName 	string		`json:"file_name"`
	Pos 		uint64		`json:"position"`
}

type IndexItem struct {
	Key 		string		`json:"key"`
	CreateTime 	int64		`json:"create_time"`
	PosItem		Position	`json:"pos_item"`
}

func (index *IndexItem) pack() string {
	packedData, packErr := json.Marshal(index)
	if packErr != nil {
		panic(fmt.Sprintf("failed to pack index item. error:[%s]", packErr))
	}

	return string(packedData)
}

func (index *IndexItem) unpack(data string) {
	if index == nil {
		index = new(IndexItem)
	}

	unpackErr := json.Unmarshal([]byte(data), index)
	if unpackErr != nil {
		panic(fmt.Sprintf("failed to unpack index item. error:[%s]", unpackErr))
	}

	return
}


func (pos *Position) getFileName() string {
	return pos.FileName
}

func (pos *Position) getPosition() uint64 {
	return pos.Pos
}

func (pos *Position) pack() string {
	packedData, packErr := json.Marshal(pos)
	if packErr != nil {
		panic(fmt.Sprintf("failed to pack position. error:[%s]", packErr))
	}

	return string(packedData)
}

func (pos *Position) unpack(data string) {
	if pos == nil {
		pos = new(Position)
	}

	unpackErr := json.Unmarshal([]byte(data), pos)
	if unpackErr != nil {
		panic(fmt.Sprintf("failed to unpack position. error:[%s]", unpackErr))
	}

	return
}


type IndexManager map[string]IndexItem

type Storager struct {}

func (s *Storager) Read(pos Position) []byte {
	if pos.getFileName() == "" || pos.getPosition() <= 0 {
		return nil
	}

	fileHandle, openErr := os.OpenFile(pos.getFileName(), os.O_RDONLY, DefaultFileMode)
	if openErr != nil {
		panic(fmt.Sprintf("Occur fatal error while opening data file. error[%s]", openErr))
	}
	defer fileHandle.Close()

	//move pointer to pos relative to the origin of the file
	_, seekErr := fileHandle.Seek(int64(pos.getPosition()), os.SEEK_SET)
	if seekErr != nil {
		return nil
	}

	result, readErr := ioutil.ReadAll(fileHandle)
	if readErr != nil {
		panic(fmt.Sprintf("Occur fatal error while read file. error[%s]", readErr))
	}

	return result
}

func (s *Storager) Write(data []byte, syncFlag bool) (*Position, error) {
	curDataFileName := s.generateCurrentDataFileName()
	fileHandle, openErr := os.OpenFile(curDataFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, DefaultFileMode)
	if openErr != nil {
		return nil, openErr
	}
	defer fileHandle.Close()

	curPos, seekErr := fileHandle.Seek(0, os.SEEK_CUR)
	if seekErr != nil {
		return nil, seekErr
	}

	_, writeErr := io.Copy(io.Writer(fileHandle), bytes.NewReader(data))
	if writeErr != nil {
		return nil, writeErr
	}

	if syncFlag {
		syncErr := fileHandle.Sync()
		if syncErr != nil {
			return nil, syncErr
		}
	}

	posItem := &Position{FileName:curDataFileName, Pos:uint64(curPos)}
	return posItem, nil
}

func (s *Storager) generateCurrentDataFileName() string {
	now := time.Now()
	return fmt.Sprintf("bitcask_%d-%2d-%2d.zzkv", now.Year(), now.Month(), now.Day())
}





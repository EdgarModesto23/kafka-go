package main

import (
	"bytes"
	"encoding/binary"
)

type Headers struct {
	api_key     int16
	api_version int16
	corr_id     int32
	client_id   string
}

type Size int32

type Body []byte

func byteSliceToInt[T any](data T, bslice []byte, start int, end int) error {
  err := binary.Read(bytes.NewReader(bslice[start:end]), binary.BigEndian, data)
  if err != nil {
    return err
  }
  return nil
}

func checkVersion(v int16) bool {
  if v < 0 || v > 4 {
    return false
  }
  return true
}

func ParseRequest(r []byte) (Size, *Headers, Body, int) {

  if (len(r) < 12) {
    return 0, nil, nil, 5
  }

  var rsize Size
  //get size of request
  err := byteSliceToInt(&rsize, r, 0, 4)
  if err != nil {
    return 0, nil, nil, 500
  }

  rheaders := Headers{}

  // parse request's api_key
  err = byteSliceToInt(&rheaders.api_key, r, 4, 6)
  if err != nil {
    return 0, nil, nil, 500
  }
  // parse api_version
  err = byteSliceToInt(&rheaders.api_version, r, 6, 8)
  if err != nil {
    return 0, nil, nil, 500
  }
  
  // parse correlation id
  err = byteSliceToInt(&rheaders.corr_id, r, 8, 12)
  if err != nil {
    return 0, nil, nil, 500
  }


  if !checkVersion(rheaders.api_version) {
    return 0, &rheaders, Body{}, 35
  }

  return 0, &rheaders, Body{}, 0
} 

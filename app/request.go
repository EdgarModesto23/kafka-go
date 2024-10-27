package main

import (
	"bytes"
	"encoding/binary"
	"errors"
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

func ParseRequest(r []byte) (Size, *Headers, Body, error) {

  if (len(r) < 12) {
    return 0, nil, nil, errors.New("Wrong request size")
  }

  var rsize Size
  //get size of request
  err := byteSliceToInt(&rsize, r, 0, 4)
  if err != nil {
    return 0, nil, nil, err
  }

  rheaders := Headers{}

  // parse request's api_key
  err = byteSliceToInt(&rheaders.api_key, r, 4, 6)
  if err != nil {
    return 0, nil, nil, err
  }
  // parse api_version
  err = byteSliceToInt(&rheaders.api_version, r, 6, 8)
  if err != nil {
    return 0, nil, nil, err
  }
  // parse correlation id
  err = byteSliceToInt(&rheaders.corr_id, r, 8, 12)
  if err != nil {
    return 0, nil, nil, err
  }

  return 0, &rheaders, Body{}, nil
} 

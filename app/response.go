package main

import (
	"encoding/binary"
)

type Response struct {
	length   int32
	corr_id  int32
	err_code int16
	body     []byte
}

func ResponseToByte(r Response) []byte {
	s := make([]byte, 4)
	cid := make([]byte, 4)
  error := make([]byte, 2)

	// response size
	binary.BigEndian.PutUint32(s, uint32(r.length))

	// correlation id
	binary.BigEndian.PutUint32(cid, uint32(r.corr_id))

	// error
	binary.BigEndian.PutUint16(error, uint16(r.err_code))

	headers := append(s, cid...)

	herror := append(headers, error...)

  res := append(herror, r.body...)

	return res
}

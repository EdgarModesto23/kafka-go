package main

import (
	"encoding/binary"
)


func Response(size, corr_id int32, error_code int32) []byte {
  
  s := make([]byte, 4)
  cid := make([]byte, 4)
  b := make([]byte, 4)

  // response size
  binary.BigEndian.PutUint32(s, uint32(size))

  // correlation id 
  binary.BigEndian.PutUint32(cid, uint32(corr_id))

  //body 
  binary.BigEndian.PutUint16(b, uint16(error_code))

  headers := append(s, cid...)

  res := append(headers, b...)

  return res
}



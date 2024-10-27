package main

import (
	"encoding/binary"
)




func Response(size, corr_id int32) []byte {
  
  s := make([]byte, 4)
  cid := make([]byte, 4)

  // response size
  binary.BigEndian.PutUint32(s, uint32(size))

  // correlation id 
  binary.BigEndian.PutUint32(cid, uint32(corr_id))


  return append(s, cid...)

}



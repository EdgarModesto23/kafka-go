package main

import (
	"bytes"
	"encoding/binary"
)

type APIVersions struct {
	request Request
}

func (h *APIVersions) Execute() Response {
	body := []int16{18, 3, 4}

	buf := new(bytes.Buffer)

  buf.WriteByte(byte(2))

	for _, v := range body {
		if err := binary.Write(buf, binary.BigEndian, v); err != nil {
      return Response{
				corr_id:  h.request.headers.corr_id,
				err_code: -1,
				body:     []byte{},
				length:   6,
			}
		}
	}

  buf.WriteByte(byte(0))
  buf.Write([]byte{0,0,0,0})
  buf.WriteByte(byte(0))

	size := int32(len(buf.Bytes()) + 6)

	return Response{
		corr_id:  h.request.headers.corr_id,
		err_code: 0,
		body:     buf.Bytes(),
		length:   size,
	}
}

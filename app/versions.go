package main

import (
	"bytes"
	"encoding/binary"
)

type APIVersions struct {
	request Request
}

func writeAPIKeysSupportedNum(buf *bytes.Buffer, n int){
  buf.WriteByte(byte(n))
}

func supportApiKey18(buf *bytes.Buffer) error {
  body := []int16{18, 3, 4}
  for _, v := range body {
    if err := binary.Write(buf, binary.BigEndian, v); err != nil {
      return err
    }
  }

  buf.WriteByte(byte(0))
  
  return nil
}

func supportApiKey75(buf *bytes.Buffer) error{
  body := []int16{75,0,4}

  for _, v := range body {
    if err := binary.Write(buf, binary.BigEndian, v); err != nil {
      return err
    }
  }

  buf.WriteByte(byte(0))

  return nil
}

func writeSupportedAPIKeys(buf *bytes.Buffer) (int32, error) {
  writeAPIKeysSupportedNum(buf,3)

  supportApiKey18(buf)
  supportApiKey75(buf)

  buf.Write([]byte{0,0,0,0})
  buf.WriteByte(byte(0))

  size := int32(len(buf.Bytes()) + 6)

  return size, nil

}

func (h *APIVersions) Execute() Response {

	buf := new(bytes.Buffer)

  size, err := writeSupportedAPIKeys(buf)
  if err != nil {
    return Response{
				corr_id:  h.request.headers.corr_id,
				err_code: -1,
				body:     []byte{},
				length:   6,
			}
  }


	return Response{
		corr_id:  h.request.headers.corr_id,
		err_code: 0,
		body:     buf.Bytes(),
		length:   size,
	}
}

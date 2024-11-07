package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ListPartitions struct {
	request Request
}

func getTopicData(buf *bytes.Buffer, topic string) {
	buf.Write(binary.BigEndian.AppendUint16([]byte{}, uint16(3)))
}

func getTopicsFromRequest(b Body) []string {
	// read first 4 bytes for n of topics
  res := []string{}

  n_topics := binary.BigEndian.Uint32(b[:5])

  topics := b[5:n_topics + 1]

  offset := 0

  for i := 0; i < int(n_topics); i++ {
    slen, _, _ := ParseUnsignedVarint(topics[offset:])
    topic := string(topics[offset:slen + 1])
    fmt.Println(topic)
  }

	return res
}

func addThrottle(b *bytes.Buffer){
	b.Write([]byte{0, 0, 0, 0})
}

func (h *ListPartitions) Execute() Response {
	buf := new(bytes.Buffer)

	getTopicsFromRequest(h.request.body)
  
  addThrottle(buf)

	getTopicData(buf, "")

	// cursor
	buf.WriteByte(byte(0))

	// tag buffer
	buf.WriteByte(byte(0))

	size := int32(len(buf.Bytes()) + 6)

	return Response{
		corr_id:  h.request.headers.corr_id,
		err_code: 0,
		body:     buf.Bytes(),
		length:   size,
	}
}

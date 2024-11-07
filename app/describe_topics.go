package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"golang.org/x/text/encoding/charmap"
	_ "golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	_ "golang.org/x/text/transform"
)

type ListPartitions struct {
	request Request
}

func getTopicData(buf *bytes.Buffer, topic string) {
	buf.Write(binary.BigEndian.AppendUint16([]byte{}, uint16(3)))
	buf.WriteByte(byte(len(topic) + 1))
	buf.Write([]byte(topic))
	buf.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	buf.WriteByte(byte(0))
	buf.WriteByte(byte(1))
	buf.Write(binary.BigEndian.AppendUint32([]byte{}, uint32(0x00000df8)))
	buf.WriteByte(byte(0))

	// cursor
	buf.WriteByte(byte(0xff))

	// tag buffer
	buf.WriteByte(byte(0))
}

func getTopicsFromRequest(b Body) []string {
	res := []string{}
	decoder := charmap.ISO8859_1.NewDecoder()

	// Transform the byte slice and decode it into UTF-8

	// Return the result as a string

	// read n of topics
	n_topics := int32(b[1]) - 1

	offset := 1

	for i := 0; i < int(n_topics); i++ {
		slen := int(b[offset+1])
		fmt.Println(slen)
		utf8Data, _, _ := transform.Bytes(decoder, b[offset+2:slen+2])
		res = append(res, string(utf8Data))
	}

	return res
}

func addThrottle(b *bytes.Buffer) {
	b.Write([]byte{0, 0, 0})
}

func (h *ListPartitions) Execute() Response {
	buf := new(bytes.Buffer)

	topics := getTopicsFromRequest(h.request.body)

	addThrottle(buf)

	// arr len
	buf.WriteByte(byte(len(topics) + 1))

	for _, v := range topics {
		getTopicData(buf, v)
	}

	size := int32(len(buf.Bytes()) + 6)

	return Response{
		corr_id:  h.request.headers.corr_id,
		err_code: 0,
		body:     buf.Bytes(),
		length:   size,
	}
}

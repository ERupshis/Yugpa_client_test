package messages

import (
	"encoding/binary"
	"fmt"
	"net"
)

// Request structure
type Request struct {
	Path string
}

func (r *Request) Serialize() ([]byte, error) {
	pathLength := int32(len(r.Path))
	buffer := make([]byte, 4+len(r.Path))

	binary.LittleEndian.PutUint32(buffer[0:4], uint32(pathLength))
	copy(buffer[4:], r.Path)

	return buffer, nil
}

func (r *Request) Deserialize(conn net.Conn) error {
	var pathLen int32
	if err := binary.Read(conn, binary.LittleEndian, &pathLen); err != nil {
		return fmt.Errorf("read Path length: %w", err)
	}

	entryNameBytes := make([]byte, pathLen)
	if _, err := conn.Read(entryNameBytes); err != nil {
		return fmt.Errorf("read Path: %w", err)
	}
	r.Path = string(entryNameBytes)

	return nil
}

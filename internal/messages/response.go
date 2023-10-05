package messages

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

// EntryType enumeration
type EntryType int32

const (
	Undefined EntryType = iota
	DIRECTORY
	FILE
)

// Response structure
type Response struct {
	Status           int32
	DirectoryEntries map[string]EntryType
}

func (r *Response) Serialize() ([]byte, error) {
	var buf bytes.Buffer

	// Status handling.
	if err := binary.Write(&buf, binary.LittleEndian, r.Status); err != nil {
		return nil, err
	}

	// Entries handling.
	numEntries := int32(len(r.DirectoryEntries))
	if err := binary.Write(&buf, binary.LittleEndian, numEntries); err != nil {
		return nil, fmt.Errorf("write entries count: %w", err)
	}

	for entryName, entryType := range r.DirectoryEntries {
		// Entry's entryName.
		entryNameLen := int32(len(entryName))
		if err := binary.Write(&buf, binary.LittleEndian, entryNameLen); err != nil {
			return nil, fmt.Errorf("write entry '%s' name length: %w", entryName, err)
		}

		if _, err := buf.WriteString(entryName); err != nil {
			return nil, fmt.Errorf("write entry '%s' name: %w", entryName, err)
		}

		// Entry's type.
		if err := binary.Write(&buf, binary.LittleEndian, int32(entryType)); err != nil {
			return nil, fmt.Errorf("write entry '%s' type: %w", entryName, err)
		}
	}

	return buf.Bytes(), nil
}

func (r *Response) Deserialize(buf io.Reader) error {
	r.DirectoryEntries = map[string]EntryType{}

	if err := binary.Read(buf, binary.LittleEndian, &r.Status); err != nil {
		return fmt.Errorf("read response status: %w", err)
	}

	var numEntries int32
	if err := binary.Read(buf, binary.LittleEndian, &numEntries); err != nil {
		return fmt.Errorf("read entries count: %w", err)
	}

	for i := 0; i < int(numEntries); i++ {
		var entryNameLen int32
		if err := binary.Read(buf, binary.LittleEndian, &entryNameLen); err != nil {
			return fmt.Errorf("read entry[%d] name length: %w", i, err)
		}

		entryNameBytes := make([]byte, entryNameLen)
		if _, err := buf.Read(entryNameBytes); err != nil {
			return fmt.Errorf("read entry[%d] name: %w", i, err)
		}
		entryName := string(entryNameBytes)

		var entryType EntryType
		if err := binary.Read(buf, binary.LittleEndian, &entryType); err != nil {
			return fmt.Errorf("read entry[%d] type: %w", i, err)
		}

		r.DirectoryEntries[entryName] = entryType
	}

	return nil
}

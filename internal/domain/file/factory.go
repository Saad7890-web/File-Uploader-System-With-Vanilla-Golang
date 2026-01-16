package file

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewFile(
	ownerID string,
	name string,
	size int64,
	mimeType string,
	path string,
	checksum string,
) (*File, error) {

	if ownerID == "" {
		return nil, ErrInvalidOwner
	}

	if strings.TrimSpace(name) == "" {
		return nil, ErrInvalidName
	}

	if size <= 0 {
		return nil, ErrInvalidSize
	}

	return &File{
		ID:        uuid.NewString(),
		OwnerID:  ownerID,
		Name:      name,
		Size:      size,
		MimeType:  mimeType,
		Path:      path,
		Checksum:  checksum,
		CreatedAt: time.Now().UTC(),
	}, nil
}

package file

import "time"


type File struct {
	ID        string
	OwnerID  string
	Name      string
	Size      int64
	MimeType  string
	Path      string
	Checksum  string
	CreatedAt time.Time
}

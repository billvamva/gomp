package components

import "bytes"

type Component interface {
	Render(buf *bytes.Buffer) error
}

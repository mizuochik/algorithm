package gapbuffer

type GapBuffer struct {
	Buf      []rune
	GapLeft  int
	GapRight int
}

func New(bufSize int) *GapBuffer {
	return &GapBuffer{
		Buf:      make([]rune, bufSize),
		GapLeft:  0,
		GapRight: bufSize,
	}
}

func (b *GapBuffer) String() string {
	return string(b.Buf[:b.GapLeft]) + string(b.Buf[b.GapRight:])
}

func (b *GapBuffer) Insert(s string) {
	r := []rune(s)
	if len(r) > b.GapSize() {
		// TBD
	}
	for i := 0; i < len(r); i++ {
		b.Buf[b.GapLeft] = r[i]
		b.GapLeft++
	}
}

func (b *GapBuffer) GapSize() int {
	return b.GapRight - b.GapLeft
}

func (b *GapBuffer) Cursor() int {
	return b.GapLeft
}

func (b *GapBuffer) SetCursor(pos int) {
	if pos > b.GapLeft {
		l := pos - b.GapLeft
		for i := 0; i < l; i++ {
			b.Buf[b.GapLeft] = b.Buf[b.GapRight]
			b.GapLeft++
			b.GapRight++
		}
	} else if pos < b.GapLeft {
		l := b.GapLeft - pos
		for i := 0; i < l; i++ {
			b.Buf[b.GapRight-1] = b.Buf[b.GapLeft-1]
			b.GapLeft--
			b.GapRight--
		}
	}
}

func (b *GapBuffer) Grow(newSize int) {
	oldSize := len(b.Buf)
	newGapRight := b.GapRight + (newSize - oldSize)
	newBuf := make([]rune, newSize)
	for i := 0; i < b.GapLeft; i++ {
		newBuf[i] = b.Buf[i]
	}
	for i := 0; i < len(b.Buf[b.GapRight:]); i++ {
		newBuf[newGapRight+i] = b.Buf[b.GapRight+1]
	}
	b.Buf = newBuf
	b.GapRight = newGapRight
}

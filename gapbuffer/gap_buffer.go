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

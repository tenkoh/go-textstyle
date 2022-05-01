package textstyle

import (
	"reflect"
	"testing"

	"golang.org/x/text/transform"
)

type testReplacer struct {
}

func (rep *testReplacer) LowerFunc(r rune) rune {
	return r + 1
}
func (rep *testReplacer) UpperFunc(r rune) rune {
	return r + 2
}
func (rep *testReplacer) DigitFunc(r rune) rune {
	return r + 3
}

func TestReplacer_Replace(t *testing.T) {
	type args struct {
		r   Replacer
		src []byte
	}
	r1 := &testReplacer{}
	tests := []struct {
		arg  args
		want []byte
	}{
		{args{r1, []byte("aA1あ")}, []byte("bC4あ")},
		{args{r1, []byte{250, 100, 70, 50, 250}}, []byte{250, 101, 72, 53, 250}},
	}
	for _, tt := range tests {
		if got := Replace(tt.arg.r, tt.arg.src); !reflect.DeepEqual(tt.want, got) {
			t.Errorf("want %v, got %v\n", tt.want, got)
		}
	}
}

func TestTransformer_Transform(t *testing.T) {
	type args struct {
		src   []byte
		dst   []byte
		atEOF bool
	}
	type wants struct {
		wrote        []byte
		nSrc         int
		nDst         int
		stockToWrite []byte
		err          error
	}
	r1 := &testReplacer{}
	tr1 := NewTransformer(r1)
	tests := []struct {
		tr    *Transformer
		args  []args
		wants []wants
	}{
		{
			tr1,
			[]args{{[]byte("aA1あ"), make([]byte, 10), false}},
			[]wants{{fillBytes([]byte("bC4あ"), 10), 6, 6, nil, nil}},
		},
		{
			tr1,
			[]args{{[]byte("aA1あ"), make([]byte, 5), false}},
			[]wants{{[]byte{98, 67, 52, 227, 129}, 6, 5, []byte{130}, transform.ErrShortDst}},
		},
		{
			tr1,
			[]args{
				{[]byte("aA1あ"), make([]byte, 5), false},
				{[]byte(""), make([]byte, 5), false},
			},
			[]wants{
				{[]byte{98, 67, 52, 227, 129}, 6, 5, []byte{130}, transform.ErrShortDst},
				{fillBytes([]byte{130}, 5), 0, 1, nil, nil},
			},
		},
	}

	for _, tt := range tests {
		if len(tt.args) != len(tt.wants) {
			t.Fatal("invalid test condition: len(args) must be same as len(wants)")
		}
		for i, arg := range tt.args {
			nDst, nSrc, err := tt.tr.Transform(arg.dst, arg.src, arg.atEOF)
			want := tt.wants[i]
			if !reflect.DeepEqual(arg.dst, want.wrote) {
				t.Errorf("wrote bytes mismatched; want %v, got %v\n", want.wrote, arg.dst)
			}
			if nDst != want.nDst {
				t.Errorf("nDst mismatched; want %d, got %d\n", want.nDst, nDst)
			}
			if nSrc != want.nSrc {
				t.Errorf("nSrc mismatched; want %d, got %d\n", want.nSrc, nSrc)
			}
			if err != want.err {
				t.Errorf("error mismatched; want %s, got %s\n", want.err, err)
			}
			if got := tt.tr.stockToWrite; !reflect.DeepEqual(got, want.stockToWrite) {
				t.Errorf("stockToWrite mismatched; want %v, got %v\n", got, want.stockToWrite)
			}
		}
		tt.tr.Reset()
	}
}

func fillBytes(src []byte, nDst int) []byte {
	var n int
	dst := make([]byte, nDst)
	if nDst < len(src) {
		n = nDst
	} else {
		n = len(src)
	}
	for i := 0; i < n; i++ {
		dst[i] = src[i]
	}
	return dst
}

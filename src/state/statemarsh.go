package state

import (
	"github.com/acharapko/bytepack"
	"io"
	"log"
)

var pack = bytepack.NewBytePack(5)

func (t *Command) Marshal(w io.Writer) {
	var b [8]byte
	bs := b[:1]
	b[0] = byte(t.Op)
	w.Write(bs)
	bs = b[:8]

	packedBytes, err := pack.Pack(t.K)
	if err != nil {
		panic(err)
	}
	w.Write(packedBytes)

	//binary.LittleEndian.PutUint64(bs, uint64(t.K))
	//w.Write(bs)
	//binary.LittleEndian.PutUint64(bs, uint64(t.V))
	//w.Write(bs)

	packedBytes, err = pack.Pack(t.V)
	if err != nil {
		panic(err)
	}
	w.Write(packedBytes)
}

func (t *Command) Unmarshal(r io.Reader) error {
	var b [1]byte
	bs := b[:1]
	if _, err := io.ReadFull(r, bs); err != nil {
		return err
	}
	t.Op = Operation(b[0])

	/*if _, err := io.ReadFull(r, bs); err != nil {
		return err
	}
	t.K = Key(binary.LittleEndian.Uint64(bs))
	if _, err := io.ReadFull(r, bs); err != nil {
		return err
	}
	t.V = Value(binary.LittleEndian.Uint64(bs))*/
	//var k Key

	err := pack.UnpackFromIOReader(r, &t.K)
	if err != nil {
		return err
	}

	err = pack.UnpackFromIOReader(r, &t.V)
	if err != nil {
		return err
	}

	return nil
}

func (t *Key) Marshal(w io.Writer) {
	/*var b [8]byte
	bs := b[:8]
	binary.LittleEndian.PutUint64(bs, uint64(*t))*/
	//w.Write(bs)

	packedBytes, err := pack.Pack(*t)
	if err != nil {
		panic(err)
	}
	w.Write(packedBytes)
}

func (t *Value) Marshal(w io.Writer) {
	/*var b [8]byte
	bs := b[:8]
	binary.LittleEndian.PutUint64(bs, uint64(*t))
	w.Write(bs)*/

	packedBytes, err := pack.Pack(*t)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(packedBytes)
	if err != nil {
		log.Panicln(err)
	}
}

func (t *Key) Unmarshal(r io.Reader) error {
	/*var b [8]byte
	bs := b[:8]
	if _, err := io.ReadFull(r, bs); err != nil {
		return err
	}
	*t = Key(binary.LittleEndian.Uint64(bs))
	return nil*/

	err := pack.UnpackFromIOReader(r, t)
	if err != nil {
		return err
	}
	return nil

}

func (t *Value) Unmarshal(r io.Reader) error {
	/*var b [8]byte
	bs := b[:8]
	if _, err := io.ReadFull(r, bs); err != nil {
		return err
	}
	*t = Value(binary.LittleEndian.Uint64(bs))
	return nil*/

	err := pack.UnpackFromIOReader(r, t)
	if err != nil {
		return err
	}
	return nil

}

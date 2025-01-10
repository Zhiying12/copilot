package state

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"testing"
)

func Test_Command(t *testing.T) {
	rnd := rand.New(rand.NewSource(42))
	k := make([]byte, 8)
	val := make([]byte, 8)
	binary.LittleEndian.PutUint64(k, rnd.Uint64())
	binary.LittleEndian.PutUint64(val, rnd.Uint64())

	cmd := Command{
		Op: PUT,
		K:  k,
		V:  val,
	}

	b := make([]byte, 0, 1024)
	buf := bytes.NewBuffer(b)

	cmd.Marshal(buf)

	if buf.Len() == 0 {
		t.Fail()
	}

	cmd2 := new(Command)

	err := cmd2.Unmarshal(buf)
	if err != nil {
		t.Errorf("Unmarshal error: %v", err)
		t.Fail()
	}

	if cmd.Op != cmd2.Op {
		t.Errorf("Commands do not match")
	}

	if !bytes.Equal(cmd.K, cmd2.K) {
		t.Errorf("Keys do not match")
	}

	if !bytes.Equal(cmd.V, cmd2.V) {
		t.Errorf("Values do not match")
	}
}

func Test_Key(t *testing.T) {
	rnd := rand.New(rand.NewSource(42))
	k := make([]byte, 8)

	binary.LittleEndian.PutUint64(k, rnd.Uint64())

	key := Key(k)

	b := make([]byte, 0, 1024)
	buf := bytes.NewBuffer(b)

	key.Marshal(buf)

	if buf.Len() == 0 {
		t.Fail()
	}

	key2 := new(Key)

	err := key2.Unmarshal(buf)
	if err != nil {
		t.Errorf("Unmarshal error: %v", err)
		t.Fail()
	}

	if !bytes.Equal(key, *key2) {
		t.Errorf("Keys do not match")
	}
}

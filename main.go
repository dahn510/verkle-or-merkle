package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	v "github.com/gballet/go-verkle"
	m "github.com/wealdtech/go-merkletree/v2"
	"github.com/wealdtech/go-merkletree/v2/sha3"
)

var seed = rand.NewSource(1234)

const chunkSize = 32

func genData(size int64) ([][]byte, error) {
	r := rand.New(seed)

	remaining := size

	var sliceSize = size / chunkSize
	if size%chunkSize > 0 {
		sliceSize++
	}

	data := make([][]byte, 0, sliceSize)

	for i := 0; i < int(sliceSize); i++ {
		bufSize := chunkSize
		if remaining < chunkSize {
			bufSize = int(remaining)
		}
		b := make([]byte, bufSize)

		read, err := r.Read(b)
		if err != nil {
			return nil, err
		}

		if read != bufSize {
			return nil, fmt.Errorf("unequivalent amount read to buf: %d/%d", read, bufSize)
		}

		data = append(data, b)

	}

	return data, nil
}

// genKeys generates keys for data insert to verkle node
func genKeys(data [][]byte) [][]byte {
	keys := make([][]byte, 0, len(data))

	for i := range data {
		key := make([]byte, 32) //it uses specific length of key for some reason
		hex.Encode(key, []byte(fmt.Sprintf("%d", i)))
		keys = append(keys, key)
	}

	return keys
}

func benchCreateMerkleTree(size int64, b *testing.B) error {
	b.StopTimer()
	data, err := genData(size)
	if err != nil {
		return err
	}
	b.StartTimer()

	_, err = m.NewTree(
		m.WithData(data),
		m.WithHashType(sha3.New512()),
		m.WithSalt(false),
	)

	return err
}

func benchCreateVerkleTree(size int64, b *testing.B) error {
	b.StopTimer()
	data, err := genData(size)
	if err != nil {
		return err
	}
	keys := genKeys(data)
	b.StartTimer()

	node := v.New()

	for i, d := range data {
		err = node.Insert(keys[i], d, nil)
		if err != nil {
			return err
		}
	}

	_ = node.Commit()

	return nil
}

func benchJustCreateMerkleProof(size int64, b *testing.B) {
	b.StopTimer()
	data, err := genData(size)
	if err != nil {
		b.Fatal(err)
	}

	root, err := m.NewTree(
		m.WithData(data),
		m.WithHashType(sha3.New512()),
		m.WithSalt(false),
	)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if _, err := root.GenerateProof(data[0], 0); err != nil {
			b.Fatal(err)
		}
	}
}

func benchJustCreateVerkleProof(size int64, b *testing.B) {
	b.StopTimer()
	data, err := genData(size)
	if err != nil {
		b.Fatal(err)
	}
	keys := genKeys(data)

	node := v.New()

	for i, d := range data {
		err = node.Insert(keys[i], d, nil)
		if err != nil {
			b.Fatal(err)
		}
	}

	_ = node.Commit()

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		if _, _, _, _, err := v.MakeVerkleMultiProof(node, nil, [][]byte{keys[0]}, nil); err != nil {
			b.Fatal(err)
		}
	}
}

func benchCreateAndSerializeMerkleProof(size int64, b *testing.B) error {
	b.StopTimer()
	data, err := genData(size)
	if err != nil {
		return err
	}

	root, err := m.NewTree(
		m.WithData(data),
		m.WithHashType(sha3.New512()),
		m.WithSalt(false),
	)
	b.StartTimer()

	proof, err := root.GenerateProof(data[0], 0)
	if err != nil {
		return err
	}

	proofMarshal, err := json.Marshal(*proof)
	if err != nil {
		return err
	}
	b.Logf("proof size: %d", len(proofMarshal))

	return nil
}

func benchCreateAndSerializeVerkleProof(size int64, b *testing.B) error {
	b.StopTimer()
	data, err := genData(size)
	if err != nil {
		return err
	}
	keys := genKeys(data)

	node := v.New()

	for i, d := range data {
		err = node.Insert(keys[i], d, nil)
		if err != nil {
			return err
		}
	}

	_ = node.Commit()

	b.StartTimer()

	proof, _, _, _, err := v.MakeVerkleMultiProof(node, nil, [][]byte{keys[0]}, nil)
	if err != nil {
		return err
	}

	verkleProof, stateDiff, err := v.SerializeProof(proof)
	if err != nil {
		return err
	}

	vproof, err := verkleProof.MarshalJSON()
	if err != nil {
		return err
	}

	stateMarshal, err := json.Marshal(stateDiff)
	if err != nil {
		return err
	}

	b.Logf("verkle proof marshalled size: %d", len(vproof))
	b.Logf("stateDiff marshalled size: %d", len(stateMarshal))

	return err
}

func main() {
}

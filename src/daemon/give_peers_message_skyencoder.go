// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package daemon

import (
	"errors"
	"math"

	"github.com/SkycoinProject/cx-chains/src/cipher/encoder"
)

// encodeSizeGivePeersMessage computes the size of an encoded object of type GivePeersMessage
func encodeSizeGivePeersMessage(obj *GivePeersMessage) uint64 {
	i0 := uint64(0)

	// obj.Peers
	i0 += 4
	{
		i1 := uint64(0)

		// x.IP
		i1 += 4

		// x.Port
		i1 += 2

		i0 += uint64(len(obj.Peers)) * i1
	}

	return i0
}

// encodeGivePeersMessage encodes an object of type GivePeersMessage to a buffer allocated to the exact size
// required to encode the object.
func encodeGivePeersMessage(obj *GivePeersMessage) ([]byte, error) {
	n := encodeSizeGivePeersMessage(obj)
	buf := make([]byte, n)

	if err := encodeGivePeersMessageToBuffer(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}

// encodeGivePeersMessageToBuffer encodes an object of type GivePeersMessage to a []byte buffer.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeGivePeersMessageToBuffer(buf []byte, obj *GivePeersMessage) error {
	if uint64(len(buf)) < encodeSizeGivePeersMessage(obj) {
		return encoder.ErrBufferUnderflow
	}

	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Peers maxlen check
	if len(obj.Peers) > 512 {
		return encoder.ErrMaxLenExceeded
	}

	// obj.Peers length check
	if uint64(len(obj.Peers)) > math.MaxUint32 {
		return errors.New("obj.Peers length exceeds math.MaxUint32")
	}

	// obj.Peers length
	e.Uint32(uint32(len(obj.Peers)))

	// obj.Peers
	for _, x := range obj.Peers {

		// x.IP
		e.Uint32(x.IP)

		// x.Port
		e.Uint16(x.Port)

	}

	return nil
}

// decodeGivePeersMessage decodes an object of type GivePeersMessage from a buffer.
// Returns the number of bytes used from the buffer to decode the object.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
func decodeGivePeersMessage(buf []byte, obj *GivePeersMessage) (uint64, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Peers

		ul, err := d.Uint32()
		if err != nil {
			return 0, err
		}

		length := int(ul)
		if length < 0 || length > len(d.Buffer) {
			return 0, encoder.ErrBufferUnderflow
		}

		if length > 512 {
			return 0, encoder.ErrMaxLenExceeded
		}

		if length != 0 {
			obj.Peers = make([]IPAddr, length)

			for z1 := range obj.Peers {
				{
					// obj.Peers[z1].IP
					i, err := d.Uint32()
					if err != nil {
						return 0, err
					}
					obj.Peers[z1].IP = i
				}

				{
					// obj.Peers[z1].Port
					i, err := d.Uint16()
					if err != nil {
						return 0, err
					}
					obj.Peers[z1].Port = i
				}

			}
		}
	}

	return uint64(len(buf) - len(d.Buffer)), nil
}

// decodeGivePeersMessageExact decodes an object of type GivePeersMessage from a buffer.
// If the buffer not long enough to decode the object, returns encoder.ErrBufferUnderflow.
// If the buffer is longer than required to decode the object, returns encoder.ErrRemainingBytes.
func decodeGivePeersMessageExact(buf []byte, obj *GivePeersMessage) error {
	if n, err := decodeGivePeersMessage(buf, obj); err != nil {
		return err
	} else if n != uint64(len(buf)) {
		return encoder.ErrRemainingBytes
	}

	return nil
}

package cbordemo

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math/big"
)

type Decoder struct {
	bf     *bufio.Reader
	result []any
}

func NewDecoder(data io.Reader) *Decoder {
	return &Decoder{
		bf: bufio.NewReader(data),
	}
}

func (dec *Decoder) Decode() ([]any, error) {

	dec.result = make([]any, 0)

	result, err := dec.decode()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (dec *Decoder) decode() ([]any, error) {

	for {
		initialByte, err := dec.bf.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		majorType := initialByte >> 5
		additionalInfo := initialByte & 0x1f

		var v interface{}

		switch majorType {

		case MajorTypeUnsignedInt:
			v, err = dec.decodeUnsignedInteger(dec.bf, int(additionalInfo))
			if err != nil {
				return nil, err
			}
		case MajorTypeNegativeInt:
			v, err = dec.decodeNegativeInteger(dec.bf, int(additionalInfo))
			if err != nil {
				return nil, err
			}
		case MajorTypeByteString:
			v, err = dec.decodeByteString(dec.bf, int(additionalInfo))
			if err != nil {
				return nil, err
			}
		case MajorTypeTextString:
			v, err = dec.decodeTextString(dec.bf, int(additionalInfo))
			if err != nil {
				return nil, err
			}
		case MajorTypeArray:
			v, err = dec.decodeArray(dec.bf, int(additionalInfo))
			if err != nil {
				return nil, err
			}
		case MajorTypeMap:
			// handle map
		case MajorTypeTag:
			// handle tag
		case MajorTypeOther:
			// handle other
		}

		dec.result = append(dec.result, v)
	}

	return dec.result, nil
}

func (dec *Decoder) decodeItem() (interface{}, error) {
	initialByte, err := dec.bf.ReadByte()
	if err != nil {
		return nil, err
	}

	majorType := initialByte >> 5
	additionalInfo := initialByte & 0x1f

	var v interface{}

	switch majorType {
	case MajorTypeUnsignedInt:
		v, err = dec.decodeUnsignedInteger(dec.bf, int(additionalInfo))
	case MajorTypeNegativeInt:
		v, err = dec.decodeNegativeInteger(dec.bf, int(additionalInfo))
	case MajorTypeByteString:
		v, err = dec.decodeByteString(dec.bf, int(additionalInfo))
	case MajorTypeTextString:
		v, err = dec.decodeTextString(dec.bf, int(additionalInfo))
	case MajorTypeArray:
		v, err = dec.decodeArray(dec.bf, int(additionalInfo))
		// Add other cases as needed
	}

	if err != nil {
		return nil, err
	}

	return v, nil
}

func (dec Decoder) decodeUnsignedInteger(bf *bufio.Reader, additionalInfo int) (uint64, error) {
	var value uint64
	var err error

	switch additionalInfo {
	case 24:
		// handle 1 byte
		b, err := bf.ReadByte()
		if err != nil {
			return 0, err
		}
		value = uint64(b)
	case 25:
		// handle 2 bytes
		var b [2]byte
		_, err = io.ReadFull(bf, b[:])
		if err != nil {
			return 0, err
		}
		value = uint64(binary.BigEndian.Uint16(b[:]))
	case 26:
		// handle 4 bytes
		var b [4]byte
		_, err = io.ReadFull(bf, b[:])
		if err != nil {
			return 0, err
		}
		value = uint64(binary.BigEndian.Uint32(b[:]))
	case 27:
		// handle 8 bytes
		var b [8]byte
		_, err = io.ReadFull(bf, b[:])
		if err != nil {
			return 0, err
		}
		value = uint64(binary.BigEndian.Uint64(b[:]))
	case 28:
		// handle 16 bytes
		// CBOR specification does not define integers of more than 8 bytes
		return 0, fmt.Errorf("additional info of 28 is not supported")
	default:
		// handle 0-23
		value = uint64(additionalInfo)
	}

	return value, nil
}

func (dec Decoder) decodeNegativeInteger(bf *bufio.Reader, additionalInfo int) (*big.Int, error) {
	v, err := dec.decodeUnsignedInteger(bf, additionalInfo)
	if err != nil {
		return nil, err
	}
	negV := big.NewInt(0).Neg(new(big.Int).SetUint64(v)) // Negate the value
	negV.Sub(negV, big.NewInt(1))                        // Subtract 1
	return negV, nil
}

func (dec Decoder) decodeByteString(bf *bufio.Reader, additionalInfo int) ([]byte, error) {
	length, err := dec.decodeUnsignedInteger(bf, additionalInfo)
	if err != nil {
		return nil, err
	}

	data := make([]byte, length)
	_, err = io.ReadFull(bf, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (dec Decoder) decodeTextString(bf *bufio.Reader, additionalInfo int) (string, error) {
	length, err := dec.decodeUnsignedInteger(bf, additionalInfo)
	if err != nil {
		return "", err
	}

	data := make([]byte, length)
	_, err = io.ReadFull(bf, data)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (dec *Decoder) decodeArray(bf *bufio.Reader, additionalInfo int) ([]interface{}, error) {
	length, err := dec.decodeUnsignedInteger(bf, additionalInfo)
	if err != nil {
		return nil, err
	}

	array := make([]interface{}, length)
	for i := uint64(0); i < length; i++ {
		item, err := dec.decodeItem()
		if err != nil {
			return nil, err
		}
		array[i] = item
	}

	return array, nil
}

package crc16itu

import (
	"testing"
)

type testpair struct {
	values []byte
	result uint16
}

var tests = []testpair{
	{[]byte{0x05, 0x13, 0x00, 0x05}, 0xafd5},
	{[]byte{ 0x0D, 0x01, 0x03, 0x53, 0x41 ,0x35, 0x32 ,0x15 ,0x03 ,0x62 ,0x00 ,0x02 }, 0x2d06},
}

func TestCRC16ITU(t *testing.T) {
	for _, pair := range tests {
		v := CRC16ITU(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package rtp

import "testing"

func FuzzUnmarshalPacket(f *testing.F) {
	f.Add([]byte{
		0x90, 0xe0, 0x69, 0x8f, 0xd9, 0xc2, 0x93, 0xda, 0x1c, 0x64,
		0x27, 0x82, 0x00, 0x01, 0x00, 0x01, 0xFF, 0xFF, 0xFF, 0xFF, 0x98, 0x36, 0xbe, 0x88, 0x9e,
	})

	f.Fuzz(func(t *testing.T, data []byte) {
		packet := &Packet{}
		err := packet.Unmarshal(data)
		if err != nil {
			return
		}

		_, err = packet.Marshal()
		if err != nil {
			return
		}
	})
}

func FuzzUnmarshalOneByteHeaderExtension(f *testing.F) {
	f.Add([]byte{
		0xBE, 0xDE, 0x00, 0x01, 0x10, 0xAA, 0x20, 0xBB,
	})

	f.Fuzz(func(t *testing.T, data []byte) {
		e := &OneByteHeaderExtension{}
		_, err := e.Unmarshal(data)
		if err != nil {
			return
		}

		_, err = e.Marshal()
		if err != nil {
			return
		}
	})
}

func FuzzUnmarshalTwoByteHeaderExtension(f *testing.F) {
	f.Add([]byte{
		0xBE, 0xDE, 0x00, 0x01, 0x50, 0xAA, 0x00, 0x00,
		0x98, 0x36, 0xbe, 0x88, 0x9e,
	})

	f.Fuzz(func(t *testing.T, data []byte) {
		e := &TwoByteHeaderExtension{}
		_, err := e.Unmarshal(data)
		if err != nil {
			return
		}

		_, err = e.Marshal()
		if err != nil {
			return
		}
	})
}

func FuzzUnmarshalPlayoutDelayExtension(f *testing.F) {
	f.Add([]byte{
		0x01, 0x01, 0x00,
	})

	f.Fuzz(func(t *testing.T, data []byte) {
		e := &PlayoutDelayExtension{}
		err := e.Unmarshal(data)
		if err != nil {
			return
		}

		_, err = e.Marshal()
		if err != nil {
			return
		}
	})
}

func FuzzUnmarshalTransportCCExtension(f *testing.F) {
	f.Add([]byte{
		0x00, 0x02,
	})

	f.Fuzz(func(t *testing.T, data []byte) {
		e := &TransportCCExtension{}
		err := e.Unmarshal(data)
		if err != nil {
			return
		}

		_, err = e.Marshal()
		if err != nil {
			return
		}
	})
}

// The MIT License (MIT)
//
// Copyright (c) 2016 winlin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// @remark we must use another packet for utest, because the cgo will dup symbols.
package fdkaac_test

import (
	"github.com/onedss/go-fdkaac/fdkaac"
	"testing"
)

func TestRawInit(t *testing.T) {
	d := fdkaac.NewAacDecoder()

	asc := []byte{0x12, 0x10}
	if err := d.InitRaw(asc); err != nil {
		t.Error("init failed, err is", err)
		return
	}
	defer d.Close()

	if d.AacSampleRate() != 0 {
		t.Error("AacSampleRate", d.AacSampleRate(), "is not 0")
	}
	if d.Profile() != 0 {
		t.Error("Profile", d.Profile(), "is not 0")
	}
	if d.AudioObjectType() != 0 {
		t.Error("AudioObjectType", d.AudioObjectType(), "is not 0")
	}
	if d.ChannelConfig() != 0 {
		t.Error("ChannelConfig", d.ChannelConfig(), "is not 0")
	}
	if d.AacSamplesPerFrame() != 0 {
		t.Error("AacSamplesPerFrame", d.AacSamplesPerFrame(), "is not 0")
	}
}

func TestAdtsInit(t *testing.T) {
	d := fdkaac.NewAacDecoder()

	if err := d.InitAdts(); err != nil {
		t.Error("init failed, err is", err)
		return
	}
	defer d.Close()

	if d.AacSampleRate() != 0 {
		t.Error("AacSampleRate", d.AacSampleRate(), "is not 0")
	}
	if d.Profile() != 0 {
		t.Error("Profile", d.Profile(), "is not 0")
	}
	if d.AudioObjectType() != 0 {
		t.Error("AudioObjectType", d.AudioObjectType(), "is not 0")
	}
	if d.ChannelConfig() != 0 {
		t.Error("ChannelConfig", d.ChannelConfig(), "is not 0")
	}
	if d.AacSamplesPerFrame() != 0 {
		t.Error("AacSamplesPerFrame", d.AacSamplesPerFrame(), "is not 0")
	}
}

func TestRawDecode(t *testing.T) {
	d := fdkaac.NewAacDecoder()

	asc := []byte{0x12, 0x10}
	if err := d.InitRaw(asc); err != nil {
		t.Error("init failed, err is", err)
		return
	}
	defer d.Close()

	if b, err := d.Decode([]byte{
		0x21, 0x17, 0x55, 0x35, 0xa1, 0x0c, 0x2f, 0x00, 0x00, 0x50, 0x23, 0xa6, 0x81, 0xbf, 0x9c, 0xbf,
		0x13, 0x73, 0xa9, 0xb0, 0x41, 0xed, 0x60, 0x23, 0x48, 0xf7, 0x34, 0x07, 0x12, 0x53, 0xd8, 0xeb,
		0x49, 0xf4, 0x1e, 0x73, 0xc9, 0x01, 0xfd, 0x16, 0x9f, 0x8e, 0xb5, 0xd5, 0x9b, 0xb6, 0x49, 0xdb,
		0x35, 0x61, 0x3b, 0x54, 0xad, 0x5f, 0x9d, 0x34, 0x94, 0x88, 0x58, 0x89, 0x33, 0x54, 0x89, 0xc4,
		0x09, 0x80, 0xa2, 0xa1, 0x28, 0x81, 0x42, 0x10, 0x48, 0x94, 0x05, 0xfb, 0x03, 0xc7, 0x64, 0xe1,
		0x54, 0x17, 0xf6, 0x65, 0x15, 0x00, 0x48, 0xa9, 0x80, 0x00, 0x38}); err != nil {
		t.Error("decode failed, err is", err)
		return
	} else if len(b) != 4096 {
		t.Error("pcm size invalid, expect 4096, actual is", len(b))
		return
	}
}

func TestRawDecode_MultipleFrames(t *testing.T) {
	d := fdkaac.NewAacDecoder()

	asc := []byte{0x12, 0x10}
	if err := d.InitRaw(asc); err != nil {
		t.Error("init failed, err is", err)
		return
	}
	defer d.Close()

	if b, err := d.Decode([]byte{
		0x21, 0x17, 0x55, 0x55, 0x19, 0x1a, 0x2a, 0x2d, 0x54, 0xce, 0x00, 0x58, 0x1a, 0x1e, 0x42, 0x0e,
		0x1f, 0xd2, 0xd4, 0x9c, 0x15, 0x77, 0xf4, 0x07, 0x38, 0x3d, 0xc5, 0x04, 0x19, 0x64, 0x39, 0x98,
		0x01, 0xae, 0x2e, 0xb1, 0xd0, 0x87, 0xca, 0x33, 0x17, 0xfb, 0x05, 0x00, 0x7a, 0x60, 0x47, 0x79,
		0x6b, 0x9b, 0xdf, 0x2d, 0xfd, 0x32, 0xc6, 0x9f, 0x1f, 0x21, 0x4b, 0x04, 0x9b, 0xe2, 0x4d, 0x62,
		0xc8, 0x01, 0xe0, 0x98, 0x0a, 0x37, 0x48, 0x44, 0x42, 0x02, 0x00, 0xd0, 0x7d, 0xae, 0xb4, 0x32,
		0xf1, 0xcc, 0x76, 0x5f, 0x18, 0xac, 0xae, 0x0e}); err != nil {
		t.Error("fill failed, err is", err)
		return
	} else if len(b) != 4096 {
		t.Error("pcm size invalid, expect 4096, actual is", len(b))
		return
	}

	if b, err := d.Decode([]byte{
		0x21, 0x17, 0x55, 0x5c, 0x21, 0x12, 0xc2, 0x15, 0x04, 0x17, 0x94, 0x50, 0xb0, 0xaf, 0x3a, 0x34,
		0x12, 0x7f, 0xee, 0x54, 0xac, 0xe2, 0x57, 0x57, 0xf7, 0x19, 0x18, 0xc5, 0x08, 0xc9, 0xaa, 0x21,
		0x75, 0x2c, 0xc9, 0x4f, 0x6f, 0xc7, 0xe2, 0xfb, 0x44, 0x72, 0x47, 0x71, 0x4a, 0x88, 0x9b, 0xfe,
		0x0c, 0x83, 0x02, 0x1a, 0xc9, 0x59, 0x7a, 0x48, 0x98, 0xac, 0x02, 0xab, 0x64, 0x22, 0x32, 0xcd,
		0x50, 0x3d, 0x80, 0x16, 0x22, 0x70, 0xb0, 0x7b, 0x00, 0x53, 0xef, 0x7c, 0xbc}); err != nil {
		t.Error("fill failed, err is", err)
		return
	} else if len(b) != 4096 {
		t.Error("pcm size invalid, expect 4096, actual is", len(b))
		return
	}

	if b, err := d.Decode([]byte{
		0x21, 0x17, 0x55, 0x4d, 0x1d, 0x4d, 0x01, 0x42, 0x8a, 0x80, 0x19, 0x01, 0x8b, 0x0b, 0xe0, 0x02,
		0x24, 0x7d, 0x8e, 0x08, 0xf2, 0x65, 0x64, 0xef, 0x02, 0x80, 0xf2, 0x72, 0xe4, 0xea, 0x19, 0x9c,
		0xd6, 0x90, 0xb8, 0x6f, 0xd4, 0x28, 0x74, 0xb9, 0xdd, 0x80, 0x6a, 0xfe, 0x09, 0x0e, 0xa4, 0xb7,
		0x83, 0x7f, 0xf8, 0x80, 0xa4, 0xa1, 0xd6, 0xb3, 0x6d, 0xbd, 0xe5, 0xe3, 0xc7, 0x00, 0xa0, 0x50,
		0x17, 0x49, 0x96, 0x8b, 0x9a, 0x17, 0x40, 0x02, 0xa8, 0x50, 0x15, 0x03, 0x7a, 0x1c, 0x01, 0x5c,
		0x9c}); err != nil {
		t.Error("fill failed, err is", err)
		return
	} else if len(b) != 4096 {
		t.Error("pcm size invalid, expect 4096, actual is", len(b))
		return
	}
}

func TestAdtsDecode_Partial(t *testing.T) {
	d := fdkaac.NewAacDecoder()

	if err := d.InitAdts(); err != nil {
		t.Error("init failed, err is", err)
		return
	}
	defer d.Close()

	if b, err := d.Decode([]byte{0xff, 0xf1, 0x50, 0x80, 0x0c, 0x40, 0xfc,
		0x21, 0x17, 0x55, 0x35, 0xa1, 0x0c, 0x2f, 0x00, 0x00, 0x50, 0x23, 0xa6, 0x81, 0xbf, 0x9c, 0xbf,
		0x13, 0x73, 0xa9, 0xb0, 0x41, 0xed, 0x60, 0x23, 0x48, 0xf7, 0x34, 0x07, 0x12, 0x53, 0xd8, 0xeb,
		0x49, 0xf4, 0x1e, 0x73, 0xc9, 0x01, 0xfd, 0x16, 0x9f, 0x8e, 0xb5, 0xd5, 0x9b, 0xb6, 0x49, 0xdb,
		0x35, 0x61, 0x3b, 0x54, 0xad, 0x5f, 0x9d, 0x34, 0x94, 0x88, 0x58, 0x89, 0x33, 0x54, 0x89, 0xc4,
		0x09, 0x80, 0xa2, 0xa1, 0x28, 0x81, 0x42, 0x10, 0x48, 0x94, 0x05, 0xfb, 0x03, 0xc7, 0x64, 0xe1,
		0x54}); err != nil {
		t.Error("fill failed, err is", err)
		return
	} else if len(b) != 0 {
		t.Error("pcm size invalid, expect 0, actual is", len(b))
		return
	}

	if b, err := d.Decode([]byte{0x17, 0xf6, 0x65, 0x15, 0x00, 0x48, 0xa9, 0x80, 0x00, 0x38,
		0xff, 0xf1, 0x50, 0x80, 0x0c, 0x40, 0xfc,
		0x21, 0x17, 0x55, 0x35, 0xa1, 0x0c, 0x2f, 0x00, 0x00, 0x50, 0x23, 0xa6, 0x81, 0xbf, 0x9c, 0xbf}); err != nil {
		t.Error("fill failed, err is", err)
		return
	} else if len(b) != 4096 {
		t.Error("pcm size invalid, expect 4096, actual is", len(b))
		return
	}

	if b, err := d.Decode([]byte{
		0x13, 0x73, 0xa9, 0xb0, 0x41, 0xed, 0x60, 0x23, 0x48, 0xf7, 0x34, 0x07, 0x12, 0x53, 0xd8, 0xeb,
		0x49, 0xf4, 0x1e, 0x73, 0xc9, 0x01, 0xfd, 0x16, 0x9f, 0x8e, 0xb5, 0xd5, 0x9b, 0xb6, 0x49, 0xdb,
		0x35, 0x61, 0x3b, 0x54, 0xad, 0x5f, 0x9d, 0x34, 0x94, 0x88, 0x58, 0x89, 0x33, 0x54, 0x89, 0xc4,
		0x09, 0x80, 0xa2, 0xa1, 0x28, 0x81, 0x42, 0x10, 0x48, 0x94, 0x05, 0xfb, 0x03, 0xc7, 0x64, 0xe1}); err != nil {
		t.Error("fill failed, err is", err)
		return
	} else if len(b) != 0 {
		t.Error("pcm size invalid, expect 0, actual is", len(b))
		return
	}

	if b, err := d.Decode([]byte{0x54, 0x17, 0xf6, 0x65, 0x15, 0x00, 0x48, 0xa9, 0x80, 0x00, 0x38}); err != nil {
		t.Error("fill failed, err is", err)
		return
	} else if len(b) != 4096 {
		t.Error("pcm size invalid, expect 4096, actual is", len(b))
		return
	}
}

func TestAdtsDecode_Partial2(t *testing.T) {
	d := fdkaac.NewAacDecoder()

	if err := d.InitAdts(); err != nil {
		t.Error("init failed, err is", err)
		return
	}
	defer d.Close()

	if b, err := d.Decode([]byte{
		// frame#0
		0xff, 0xf1, 0x50, 0x80, 0x0e, 0x60, 0xfc,
		0x21, 0x17, 0x55, 0x45, 0x0d, 0x88, 0x90, 0x13, 0x04, 0x2c, 0xa4, 0x01, 0x01, 0xd0, 0x20, 0x3e,
		0x27, 0x6d, 0x38, 0x35, 0x4a, 0x0b, 0x59, 0xb5, 0xde, 0x8d, 0xad, 0x72, 0x7b, 0xa6, 0xe4, 0xd7,
		0xbe, 0x0c, 0xfa, 0xe8, 0x0e, 0x1d, 0xaa, 0xc7, 0x0a, 0x44, 0xd2, 0x33, 0x81, 0xd8, 0x24, 0x81,
		0xd4, 0xc1, 0x76, 0x9b, 0x5b, 0x88, 0x58, 0x9c, 0x23, 0x82, 0xf5, 0x2c, 0x26, 0x04, 0x94, 0x80,
		0xab, 0x7b, 0x28, 0x0a, 0x66, 0x30, 0x90, 0x0a, 0x6a, 0x02, 0x16, 0xb0, 0x50, 0x06, 0x83, 0x6e,
		0xfa, 0xea, 0xe1, 0xd7, 0x30, 0xf0, 0x9b, 0x18, 0x25, 0xfc, 0x6b, 0x42, 0x5a, 0x3c, 0x5e, 0x3c,
		0x18, 0xe7, 0xad, 0xda, 0xc2, 0xcc, 0x09, 0x04, 0xa6, 0x90, 0x91, 0xc0,
		// frame#1
		0xff, 0xf1, 0x50, 0x80,
	}); err != nil {
		t.Error("fill failed, err is", err)
		return
	} else if len(b) != 4096 {
		t.Error("pcm size invalid, expect 4096, actual is", len(b))
		return
	}

	// this will drop frame#1 and cause sync error.
	// but we fix it, so it will return an empty pcm
	// and the frame#1 is keep in internal buffer.
	if _, err := d.Decode(nil); err != nil {
		t.Error("fill failed, err is", err)
		return
	}

	if b, err := d.Decode([]byte{
		// frame#1 continue
		0x0d, 0x40, 0xfc, 0x21, 0x17, 0x55, 0x45, 0x95, 0x18, 0x2c, 0x05, 0x44, 0x10, 0x00, 0xd6, 0x97,
		0x40, 0x7b, 0xe4, 0xb1, 0xcb, 0xcb, 0xd1, 0xa8, 0xc6, 0x40, 0x7d, 0x7c, 0xb3, 0x64, 0xd1, 0x4a,
		0xdc, 0x48, 0x53, 0xfc, 0x32, 0xaa, 0x0a, 0xe9, 0x25, 0xca, 0x7b, 0x4e, 0x5a, 0xa7, 0x4b, 0x52,
		0x96, 0xce, 0xee, 0x5a, 0xe8, 0xde, 0xe6, 0x0c, 0x7f, 0xc0, 0x70, 0x10, 0x6d, 0x54, 0x10, 0x12,
		0x0c, 0xc0, 0x52, 0x4b, 0x04, 0xe0, 0x31, 0x18, 0x08, 0x44, 0x00, 0x00, 0x3a, 0x5e, 0x82, 0xff,
		0xc8, 0xe9, 0x6b, 0x77, 0x18, 0xdd, 0x64, 0xe6, 0x00, 0x15, 0x88, 0x02, 0x62, 0x74, 0x3d, 0xd3,
		0x90, 0x02, 0x65, 0x98, 0x80, 0x1c,
		// frame#2, decoded
		0xff, 0xf1, 0x50, 0x80, 0x0e, 0x60, 0xfc,
		0x21, 0x17, 0x55, 0x45, 0x0d, 0x88, 0x90, 0x13, 0x04, 0x2c, 0xa4, 0x01, 0x01, 0xd0, 0x20, 0x3e,
		0x27, 0x6d, 0x38, 0x35, 0x4a, 0x0b, 0x59, 0xb5, 0xde, 0x8d, 0xad, 0x72, 0x7b, 0xa6, 0xe4, 0xd7,
		0xbe, 0x0c, 0xfa, 0xe8, 0x0e, 0x1d, 0xaa, 0xc7, 0x0a, 0x44, 0xd2, 0x33, 0x81, 0xd8, 0x24, 0x81,
		0xd4, 0xc1, 0x76, 0x9b, 0x5b, 0x88, 0x58, 0x9c, 0x23, 0x82, 0xf5, 0x2c, 0x26, 0x04, 0x94, 0x80,
		0xab, 0x7b, 0x28, 0x0a, 0x66, 0x30, 0x90, 0x0a, 0x6a, 0x02, 0x16, 0xb0, 0x50, 0x06, 0x83, 0x6e,
		0xfa, 0xea, 0xe1, 0xd7, 0x30, 0xf0, 0x9b, 0x18, 0x25, 0xfc, 0x6b, 0x42, 0x5a, 0x3c, 0x5e, 0x3c,
		0x18, 0xe7, 0xad, 0xda, 0xc2, 0xcc, 0x09, 0x04, 0xa6, 0x90, 0x91, 0xc0,
	}); err != nil {
		t.Error("fill failed, err is", err)
		return
	} else if len(b) != 4096 {
		t.Error("pcm size invalid, expect 4096, actual is", len(b))
		return
	}

	// for we fix the bug, we can got the frame#2.
	if b, err := d.Decode(nil); err != nil {
		t.Error("fill failed, err is", err)
		return
	} else if len(b) != 4096 {
		t.Error("pcm size invalid, expect 4096, actual is", len(b))
		return
	}
}

func TestRawDecode_Partial(t *testing.T) {
	d := fdkaac.NewAacDecoder()

	asc := []byte{0x12, 0x10}
	if err := d.InitRaw(asc); err != nil {
		t.Error("init failed, err is", err)
		return
	}
	defer d.Close()

	if _, err := d.Decode([]byte{0xff, 0xf1, 0x50, 0x80, 0x0c, 0x40, 0xfc,
		0x21, 0x17, 0x55, 0x35, 0xa1, 0x0c, 0x2f, 0x00, 0x00, 0x50, 0x23, 0xa6, 0x81, 0xbf, 0x9c, 0xbf,
		0x13, 0x73, 0xa9, 0xb0, 0x41, 0xed, 0x60, 0x23, 0x48, 0xf7, 0x34, 0x07, 0x12, 0x53, 0xd8, 0xeb,
		0x49, 0xf4, 0x1e, 0x73, 0xc9, 0x01, 0xfd, 0x16, 0x9f, 0x8e, 0xb5, 0xd5, 0x9b, 0xb6, 0x49, 0xdb,
		0x35, 0x61, 0x3b, 0x54, 0xad, 0x5f, 0x9d, 0x34, 0x94, 0x88, 0x58, 0x89, 0x33, 0x54, 0x89, 0xc4,
		0x09, 0x80, 0xa2, 0xa1, 0x28, 0x81, 0x42, 0x10, 0x48, 0x94, 0x05, 0xfb, 0x03, 0xc7, 0x64, 0xe1,
		0x54}); err == nil {
		t.Error("fill failed, err is", err)
		return
	}

	if _, err := d.Decode([]byte{0x17, 0xf6, 0x65, 0x15, 0x00, 0x48, 0xa9, 0x80, 0x00, 0x38}); err == nil {
		t.Error("fill failed, err is", err)
		return
	}
}

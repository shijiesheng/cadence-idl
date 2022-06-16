// The MIT License (MIT)
// 
// Copyright (c) 2021 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: uber/cadence/api/v1/domain.proto

package apiv1

var yarpcFileDescriptorClosure824795d6ae7d8e2f = [][]byte{
	// uber/cadence/api/v1/domain.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xd1, 0x6e, 0xdb, 0x36,
		0x14, 0x9d, 0xec, 0x24, 0x73, 0xae, 0x1c, 0xd7, 0x65, 0x9a, 0x46, 0xf1, 0x86, 0x45, 0x4d, 0x51,
		0xc0, 0x2b, 0x30, 0x79, 0xf1, 0x86, 0xad, 0xdd, 0xb0, 0x07, 0xc7, 0x52, 0x3b, 0x0f, 0x59, 0x16,
		0xc8, 0x6e, 0x1e, 0xd6, 0x07, 0x81, 0x96, 0x68, 0x9b, 0xa8, 0x2c, 0x0a, 0x14, 0xed, 0x34, 0x6f,
		0xc3, 0x3e, 0x62, 0x1f, 0xb3, 0xc7, 0x7d, 0xd9, 0x20, 0x8a, 0x52, 0x1c, 0x5b, 0x48, 0xfa, 0x46,
		0xde, 0x7b, 0xcf, 0xe1, 0xd1, 0xd1, 0xbd, 0x94, 0xc0, 0x5c, 0x8c, 0x09, 0xef, 0xf8, 0x38, 0x20,
		0x91, 0x4f, 0x3a, 0x38, 0xa6, 0x9d, 0xe5, 0x69, 0x27, 0x60, 0x73, 0x4c, 0x23, 0x2b, 0xe6, 0x4c,
		0x30, 0xb4, 0x9f, 0x56, 0x58, 0xaa, 0xc2, 0xc2, 0x31, 0xb5, 0x96, 0xa7, 0xad, 0xaf, 0xa6, 0x8c,
		0x4d, 0x43, 0xd2, 0x91, 0x25, 0xe3, 0xc5, 0xa4, 0x13, 0x2c, 0x38, 0x16, 0x94, 0x29, 0x50, 0xeb,
		0x78, 0x3d, 0x2f, 0xe8, 0x9c, 0x24, 0x02, 0xcf, 0xe3, 0xac, 0xe0, 0xe4, 0x9f, 0x1a, 0xec, 0xd8,
		0xf2, 0x18, 0xd4, 0x80, 0x0a, 0x0d, 0x0c, 0xcd, 0xd4, 0xda, 0xbb, 0x6e, 0x85, 0x06, 0x08, 0xc1,
		0x56, 0x84, 0xe7, 0xc4, 0xa8, 0xc8, 0x88, 0x5c, 0xa3, 0xd7, 0xb0, 0x93, 0x08, 0x2c, 0x16, 0x89,
		0x51, 0x35, 0xb5, 0x76, 0xa3, 0xfb, 0xcc, 0x2a, 0x51, 0x65, 0x65, 0x84, 0x43, 0x59, 0xe8, 0x2a,
		0x00, 0x32, 0x41, 0x0f, 0x48, 0xe2, 0x73, 0x1a, 0xa7, 0xfa, 0x8c, 0x2d, 0xc9, 0xba, 0x1a, 0x42,
		0xc7, 0xa0, 0xb3, 0xeb, 0x88, 0x70, 0x8f, 0xcc, 0x31, 0x0d, 0x8d, 0x6d, 0x59, 0x01, 0x32, 0xe4,
		0xa4, 0x11, 0xf4, 0x1a, 0xb6, 0x02, 0x2c, 0xb0, 0xb1, 0x63, 0x56, 0xdb, 0x7a, 0xf7, 0xc5, 0x3d,
		0x67, 0x5b, 0x36, 0x16, 0xd8, 0x89, 0x04, 0xbf, 0x71, 0x25, 0x04, 0xcd, 0xe0, 0xf9, 0x35, 0xe3,
		0x1f, 0x26, 0x21, 0xbb, 0xf6, 0xc8, 0x47, 0xe2, 0x2f, 0xd2, 0x13, 0x3d, 0x4e, 0x04, 0x89, 0xe4,
		0x2a, 0x26, 0x9c, 0xb2, 0xc0, 0xf8, 0xdc, 0xd4, 0xda, 0x7a, 0xf7, 0xc8, 0xca, 0x6c, 0xb3, 0x72,
		0xdb, 0x2c, 0x5b, 0xd9, 0xea, 0x9a, 0x39, 0x8b, 0x93, 0x93, 0xb8, 0x39, 0xc7, 0xa5, 0xa4, 0x40,
		0x7d, 0xa8, 0x8f, 0x71, 0xe0, 0x8d, 0x69, 0x84, 0x39, 0x25, 0x89, 0x51, 0x93, 0x94, 0x66, 0xa9,
		0xd8, 0x33, 0x1c, 0x9c, 0xa9, 0x3a, 0x57, 0x1f, 0xdf, 0x6e, 0xd0, 0x7b, 0x38, 0x9c, 0xd1, 0x44,
		0x30, 0x7e, 0xe3, 0x61, 0xee, 0xcf, 0xe8, 0x12, 0x87, 0x9e, 0x32, 0x7e, 0x57, 0x1a, 0xff, 0xbc,
		0x94, 0xaf, 0xa7, 0x6a, 0x95, 0xf5, 0x07, 0x8a, 0xe3, 0x6e, 0x18, 0x7d, 0x0b, 0x4f, 0x36, 0xc8,
		0x17, 0x9c, 0x1a, 0x20, 0x0d, 0x47, 0x6b, 0xa0, 0x77, 0x9c, 0x22, 0x0c, 0xad, 0x25, 0x4d, 0xe8,
		0x98, 0x86, 0x54, 0x6c, 0x2a, 0xd2, 0x3f, 0x5d, 0x91, 0x71, 0x4b, 0xb3, 0x26, 0xea, 0x07, 0x38,
		0x2c, 0x3b, 0x22, 0xd5, 0x55, 0x97, 0xba, 0x0e, 0x36, 0xa1, 0xa9, 0x34, 0x0b, 0xf6, 0xb1, 0x2f,
		0xe8, 0x92, 0x78, 0x7e, 0xb8, 0x48, 0x04, 0xe1, 0x9e, 0x6c, 0xda, 0x3d, 0x89, 0x79, 0x9c, 0xa5,
		0xfa, 0x59, 0xe6, 0x22, 0xed, 0xe0, 0x4b, 0xa8, 0xa9, 0xc2, 0xc4, 0x68, 0xc8, 0x3e, 0xfa, 0xbe,
		0x54, 0xb8, 0xc2, 0xb8, 0x24, 0x0e, 0xa9, 0x2f, 0xdf, 0x7d, 0x9f, 0x45, 0x13, 0x3a, 0xcd, 0x1b,
		0xa1, 0x60, 0x41, 0x5f, 0x43, 0x73, 0x82, 0x69, 0xc8, 0x96, 0x84, 0x7b, 0x4b, 0xc2, 0x93, 0xb4,
		0xbb, 0x1f, 0x99, 0x5a, 0xbb, 0xea, 0x3e, 0xca, 0xe3, 0x57, 0x59, 0x18, 0xb5, 0xa1, 0x49, 0x13,
		0x6f, 0x1a, 0xb2, 0x31, 0x0e, 0xbd, 0x6c, 0xba, 0x8d, 0xa6, 0xa9, 0xb5, 0x6b, 0x6e, 0x83, 0x26,
		0x6f, 0x65, 0x58, 0x0d, 0xe3, 0x1b, 0xd8, 0x2b, 0x48, 0x69, 0x34, 0x61, 0xc6, 0x63, 0xd9, 0x46,
		0xe5, 0xf3, 0xf6, 0x46, 0x55, 0x0e, 0xa2, 0x09, 0x73, 0xeb, 0x93, 0x95, 0x5d, 0xeb, 0x47, 0xd8,
		0x2d, 0x46, 0x01, 0x35, 0xa1, 0xfa, 0x81, 0xdc, 0xa8, 0x11, 0x4f, 0x97, 0xe8, 0x09, 0x6c, 0x2f,
		0x71, 0xb8, 0xc8, 0x87, 0x3c, 0xdb, 0xfc, 0x54, 0x79, 0xa5, 0x9d, 0xd8, 0x70, 0xfc, 0x80, 0x05,
		0xe8, 0x19, 0xd4, 0xef, 0x78, 0x9e, 0xf1, 0xea, 0xfe, 0xad, 0xdb, 0x27, 0xff, 0x6a, 0xa0, 0xaf,
		0x34, 0x39, 0xfa, 0x0d, 0x6a, 0xc5, 0x60, 0x68, 0xd2, 0x7d, 0xeb, 0xa1, 0xc1, 0xb0, 0xf2, 0x45,
		0x36, 0xce, 0x05, 0xbe, 0xe5, 0xc1, 0xde, 0x9d, 0x54, 0xc9, 0xe3, 0xbd, 0x5a, 0x7d, 0x3c, 0xbd,
		0x7b, 0x72, 0xef, 0x59, 0x37, 0xd2, 0xbe, 0x15, 0x0b, 0xfe, 0xd6, 0x60, 0xef, 0x4e, 0x12, 0x3d,
		0x85, 0x1d, 0x4e, 0x70, 0xc2, 0x22, 0x75, 0x88, 0xda, 0xa1, 0x16, 0xd4, 0x58, 0x4c, 0x38, 0x16,
		0x8c, 0x2b, 0x27, 0x8b, 0x3d, 0xfa, 0x05, 0xea, 0x3e, 0x27, 0x58, 0x90, 0xc0, 0x4b, 0x2f, 0x5f,
		0x79, 0x71, 0xea, 0xdd, 0xd6, 0xc6, 0x15, 0x33, 0xca, 0x6f, 0x66, 0x57, 0x57, 0xf5, 0x69, 0xe4,
		0xe4, 0xbf, 0x0a, 0xd4, 0x57, 0xdf, 0x6f, 0x69, 0xbb, 0x69, 0xe5, 0xed, 0x36, 0x02, 0xa3, 0x28,
		0x4d, 0x04, 0xe6, 0xc2, 0x2b, 0xae, 0x7f, 0xe5, 0xc8, 0x7d, 0x32, 0x9e, 0xe6, 0xd8, 0x61, 0x0a,
		0x2d, 0xe2, 0xe8, 0x0a, 0x8e, 0x0a, 0x56, 0xf2, 0x31, 0xa6, 0x9c, 0xac, 0xd0, 0x3e, 0xfc, 0x74,
		0x87, 0x39, 0xd8, 0x91, 0xd8, 0x5b, 0xde, 0x2e, 0x1c, 0xf8, 0x6c, 0x1e, 0x87, 0x24, 0xb5, 0x2a,
		0x99, 0x61, 0x1e, 0x78, 0x3e, 0x5b, 0x44, 0x42, 0x7e, 0x2a, 0xb6, 0xdd, 0xfd, 0x22, 0x39, 0x4c,
		0x73, 0xfd, 0x34, 0x85, 0x5e, 0x40, 0x23, 0x26, 0x51, 0x40, 0xa3, 0x69, 0x86, 0x48, 0x8c, 0x6d,
		0xb3, 0xda, 0xde, 0x76, 0xf7, 0x54, 0x54, 0x96, 0x26, 0x2f, 0xff, 0xd2, 0xa0, 0xbe, 0xfa, 0x51,
		0x42, 0x47, 0x70, 0x60, 0xff, 0xf1, 0x7b, 0x6f, 0x70, 0xe1, 0x0d, 0x47, 0xbd, 0xd1, 0xbb, 0xa1,
		0x37, 0xb8, 0xb8, 0xea, 0x9d, 0x0f, 0xec, 0xe6, 0x67, 0xe8, 0x4b, 0x30, 0xee, 0xa6, 0x5c, 0xe7,
		0xed, 0x60, 0x38, 0x72, 0x5c, 0xc7, 0x6e, 0x6a, 0x9b, 0x59, 0xdb, 0xb9, 0x74, 0x9d, 0x7e, 0x6f,
		0xe4, 0xd8, 0xcd, 0xca, 0x26, 0xad, 0xed, 0x9c, 0x3b, 0x69, 0xaa, 0xfa, 0x72, 0x06, 0x8d, 0xb5,
		0x1b, 0xef, 0x0b, 0x38, 0xec, 0xb9, 0xfd, 0x5f, 0x07, 0x57, 0xbd, 0xf3, 0x52, 0x15, 0xeb, 0x49,
		0x7b, 0x30, 0xec, 0x9d, 0x9d, 0x4b, 0x15, 0x25, 0x50, 0xe7, 0x22, 0x4b, 0x56, 0xce, 0xde, 0xc3,
		0xa1, 0xcf, 0xe6, 0x65, 0xad, 0x7e, 0xa6, 0x67, 0x26, 0x5c, 0xa6, 0x6f, 0xe5, 0x52, 0xfb, 0xf3,
		0x74, 0x4a, 0xc5, 0x6c, 0x31, 0xb6, 0x7c, 0x36, 0xef, 0xac, 0xfe, 0x7f, 0x7c, 0x43, 0x83, 0xb0,
		0x33, 0x65, 0xd9, 0x5f, 0x83, 0xfa, 0x19, 0xf9, 0x19, 0xc7, 0x74, 0x79, 0x3a, 0xde, 0x91, 0xb1,
		0xef, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x01, 0x9f, 0x23, 0xb0, 0x08, 0x00, 0x00,
	},
	// google/protobuf/duration.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a,
		0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0x56,
		0x5c, 0x1c, 0x2e, 0x50, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0xc5,
		0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5e, 0x62, 0x5e,
		0x7e, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0xe3, 0xd4, 0xcc, 0xc8, 0x25, 0x9c,
		0x9c, 0x9f, 0xab, 0x87, 0x66, 0xa6, 0x13, 0x2f, 0xcc, 0xc4, 0x00, 0x90, 0x48, 0x00, 0x63, 0x94,
		0x21, 0x54, 0x45, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0x3a, 0xc2, 0x81, 0x25,
		0x95, 0x05, 0xa9, 0xc5, 0xfa, 0xd9, 0x79, 0xf9, 0xe5, 0x79, 0x70, 0xc7, 0x16, 0x24, 0xfd, 0x60,
		0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce, 0x1d, 0xa2, 0x39, 0x00, 0xaa,
		0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4, 0x35, 0x89, 0x0d, 0x6c, 0x94,
		0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8a, 0xb4, 0xc3, 0xfb, 0x00, 0x00, 0x00,
	},
	// google/protobuf/timestamp.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xc9, 0xcc, 0x4d,
		0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x03, 0x0b, 0x09, 0xf1, 0x43, 0x14, 0xe8, 0xc1, 0x14, 0x28,
		0x59, 0x73, 0x71, 0x86, 0xc0, 0xd4, 0x08, 0x49, 0x70, 0xb1, 0x17, 0xa7, 0x26, 0xe7, 0xe7, 0xa5,
		0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x79, 0x89,
		0x79, 0xf9, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x53, 0x2b, 0x23, 0x97,
		0x70, 0x72, 0x7e, 0xae, 0x1e, 0x9a, 0xa1, 0x4e, 0x7c, 0x70, 0x23, 0x03, 0x40, 0x42, 0x01, 0x8c,
		0x51, 0x46, 0x50, 0x25, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0x48, 0x6e,
		0xac, 0x2c, 0x48, 0x2d, 0xd6, 0xcf, 0xce, 0xcb, 0x2f, 0xcf, 0x43, 0xb8, 0xb7, 0x20, 0xe9, 0x07,
		0x23, 0xe3, 0x22, 0x26, 0x66, 0xf7, 0x00, 0xa7, 0x55, 0x4c, 0x72, 0xee, 0x10, 0xdd, 0x01, 0x50,
		0x2d, 0x7a, 0xe1, 0xa9, 0x39, 0x39, 0xde, 0x20, 0x0d, 0x21, 0x20, 0xbd, 0x49, 0x6c, 0x60, 0xb3,
		0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0x65, 0xce, 0x7d, 0xff, 0x00, 0x00, 0x00,
	},
}

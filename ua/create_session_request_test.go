// Copyright 2018-2020 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ua

import (
	"testing"
	"time"
)

func TestCreateSessionRequest(t *testing.T) {
	cases := []CodecTestCase{
		{
			Name: "normal",
			Struct: &CreateSessionRequest{
				RequestHeader: &RequestHeader{
					AuthenticationToken: NewByteStringNodeID(0x00, []byte{
						0x08, 0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11,
						0xa6, 0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
					}),
					Timestamp:        time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
					RequestHandle:    1,
					AdditionalHeader: NewExtensionObject(nil),
				},
				ClientDescription: &ApplicationDescription{
					ApplicationURI:      "app-uri",
					ProductURI:          "prod-uri",
					ApplicationName:     NewLocalizedText("app-name"),
					ApplicationType:     ApplicationTypeClient,
					GatewayServerURI:    "gw-uri",
					DiscoveryProfileURI: "profile-uri",
					DiscoveryURLs:       []string{"1", "2"},
				},
				ServerURI:               "server-uri",
				EndpointURL:             "endpoint-url",
				SessionName:             "session-name",
				RequestedSessionTimeout: 6000000,
				MaxResponseMessageSize:  65534,
			},
			Bytes: []byte{
				// AuthenticationToken
				0x05, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x08,
				0x22, 0x87, 0x62, 0xba, 0x81, 0xe1, 0x11, 0xa6,
				0x43, 0xf8, 0x77, 0x7b, 0xc6, 0x2f, 0xc8,
				// Timestamp
				0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
				// RequestHandle
				0x01, 0x00, 0x00, 0x00,
				// ReturnDiagnostics
				0x00, 0x00, 0x00, 0x00,
				// AuditEntryID
				0xff, 0xff, 0xff, 0xff,
				// TimeoutHint
				0x00, 0x00, 0x00, 0x00,
				// AdditionalHeader
				0x00, 0x00, 0x00,
				// ClientDescription: ApplicationDescription
				// ApplicationURI
				0x07, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x75, 0x72, 0x69,
				// ProductURI
				0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x64, 0x2d, 0x75, 0x72, 0x69,
				// ApplicationName
				0x02, 0x08, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x6e, 0x61, 0x6d, 0x65,
				// ApplicationType
				0x01, 0x00, 0x00, 0x00,
				// GatewayServerURI
				0x06, 0x00, 0x00, 0x00, 0x67, 0x77, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryProfileURI
				0x0b, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryURLs
				0x02, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x31, 0x01, 0x00, 0x00, 0x00, 0x32,
				// ServerURI
				0x0a, 0x00, 0x00, 0x00, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x75, 0x72, 0x69,
				// EndpointURL
				0x0c, 0x00, 0x00, 0x00, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2d, 0x75, 0x72, 0x6c,
				// SessionName
				0x0c, 0x00, 0x00, 0x00, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x6e, 0x61, 0x6d, 0x65,
				// ClientNonce
				0xff, 0xff, 0xff, 0xff,
				// ClientCertificate
				0xff, 0xff, 0xff, 0xff,
				// RequestedTimeout
				0x00, 0x00, 0x00, 0x00, 0x60, 0xe3, 0x56, 0x41,
				// MaxResponseMessageSize
				0xfe, 0xff, 0x00, 0x00,
			},
		},
	}
	RunCodecTest(t, cases)
}

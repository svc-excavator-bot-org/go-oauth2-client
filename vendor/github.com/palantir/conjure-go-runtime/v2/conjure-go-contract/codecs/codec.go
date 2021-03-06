// Copyright (c) 2018 Palantir Technologies. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codecs

import (
	"io"
)

// Codec embeds both Decoder and Encoder.
type Codec interface {
	Decoder
	Encoder
}

// A Decoder decodes a serialized message.
type Decoder interface {
	Accept() string
	Decode(r io.Reader, v interface{}) error
	Unmarshal(data []byte, v interface{}) error
}

// An Encoder serializes a message.
type Encoder interface {
	ContentType() string
	Encode(w io.Writer, v interface{}) error
	Marshal(v interface{}) ([]byte, error)
}

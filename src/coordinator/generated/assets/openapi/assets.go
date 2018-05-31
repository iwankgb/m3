// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by "esc -prefix openapi/ -pkg openapi -ignore .go -o openapi/assets.go ."; DO NOT EDIT.

package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/asset-gen.sh": {
		local:   "asset-gen.sh",
		size:    221,
		modtime: 1527706041,
		compressed: `
H4sIAAAAAAAC/0zKQUrGMBDF8X1O8TrNqpAG18WFiBew7kSktpN0EGdKEkEQ7y5W0G82A7//67v4Ihrr
7tzNPN89PM/3t9f+yiUrEIiC/J9THCZs5gBAEjo8ImzwgqcJbWc9w8+tpk30nU9I4s7P626gzMplaaIZ
qdgbvNBvrSvCUTjJB8h/St8P8SsSwvGa/4EQJKsVxpgNwS6mS63c6piNQCO5zZTddwAAAP//XfBIdt0A
AAA=
`,
	},

	"/coordinator.yml": {
		local:   "openapi/coordinator.yml",
		size:    9182,
		modtime: 1527699890,
		compressed: `
H4sIAAAAAAAC/+xZT3ObSBa/61P0kj1sDhaO7c1WcZNjRUuV4rhsV6omqTk09AN1BvqR7kcSJzXffapB
Qkgw/IkSO+MZDragf+//j/caMJ94HIP2mHMyPXYmUkXoTRgjSQl4zHl1enHuTBgTYEItM5KoPObMmJCG
tAxyAsFIpsAMaAmGCU484AZYbqSK2avT25u3LEqQ0/MzFmKaaTBGopqyXzBnIVcskkowzImlqIHxwP60
Vhkn9m5FlHmum56KYBpLWuXBVGJx6v76n47Fpww1Q8XeLST9Pw9KrPFcd40LMS2B9s/TqY3wI2hTRvds
emxTwViIinhINh+MKZ6WCTm/YAvEOAG20JhnTrGa68RjTmXFLphpXMAKYxHqPHWf/Kv8b61auUSGoAzs
GJhlPFwBW5ZL7KR0pWGhJQ43SDBwU24ItLv0X8wvb+bOZIWGrCAaKiz87+T4mTOxNbritPKY4/JMuh+f
ORPisfEmRxtH7D+T8RCa9X+BKpJxrssSX5yzCmucrYIs4SGkoGiAghrWhCtIofCkCNeZZJxWxibJreyU
KYthXRzGSufZ+jjac98eJk9Tru885iyAdjwu1zEDza1/vqhHvwDaIEJUJi9cq6zwLEtkWIi57w2qDTTT
KPJwEFSDyVAZqLl/cny8PdnPnVNbKXLF61jG/q0h8pjzxBUQSSWtlHEva+Fcrw1uFf33u9tbgAItw7nW
qEsFmeVhS7G6SzUTgnHWAPxJrWZC/NhaZVzzFAh0Dbxme4DibpsqqRqXmrnrrtRMiGv4kIOhn4gpZ/fA
lPth5LaXuF+rn/7F76UqAQkQjOfrRSE3grKlwIOxthb5Hnlt091e0vAhlxqEx0jnUF2mu8xqsXsBFd8r
Tcu8FXNEp0XI90/Ss+Ozx3IzVPO3c7Ae7Q315lilFbAGZJf61fLjmKxXtXB+usnaUa1isiomlSGuQmCE
o4r3GEbtVS2Yhxi13dR5PKO2Z5x2kHQ9TscQsxT54fR88CF3n0PBtRjvgGbjWxs8kV/G1dKKPZ4uY6P5
p83cC1+/bsbasA19fweqT8pIY/oNLenBiLzNxV9ri//34G1txcq2PP6XKtd1wOA9hGvCZdoSjuS2DEXF
u0uHWWVryFuA1yXcqbv2uq5ikF8BIhnSPJsrHiQgGj4GiAnwis1RkpvVQOwnLQnMLb7ANJW0xLhPILQn
+VBXNGRc6sFgAmWT83pQlq/34FX7UTwzK6SBVqUS8HmYRb8GteLXrQ4PqmkV6xVoieKSKzQNT6UiiKF2
O0VodzLlyvOzzfUgwfC3G/kFDtOSRxHolznl+nsouuKGDo/qghOff86kvusr4x58FhHoS6RZGIIxBybZ
b1BkUI1hGAEPL1/bW8dRXIyloXqKu7va9Rq/Y/p6R8ngflt+RWgEXRe0Bxei8IInVw01o9pw86FghMPl
NqizoG1jd4SFvTdJ/dN9k6LNB7hR7Dk92XF5hJ+bXdEPqpy/Vl8bI3bf9pKHhLovRpWnNyuuRe+tJE2B
679Dw5zwI+hb2bI/GNjNpHkl7Zal31jKPxdu3QD5osvcJkljyiZ69jfSYFLcFcXX2R7wF1R9+6VPIOMV
9SUNlMhQKupRZtqryrXm9WdNgrSfYUWKdxT35tse1Yfgbk8z1J1BF9YPq9ueY4Y4QU/TKGllgZUU5joE
v48Va/4fNM6tjij6ZhVb33fSVvcTVJ6Wi0fM8S/9W3+29N/6lwtnc3H2ZuYvZ+fLeXVlOZ+9WSNa3md+
l4b4TfSsN8C2VyA/h2ejum3b6Kn1dqt9WH9vU1R/KhyzQ9viW0j1RwAAAP//kfi6qd4jAAA=
`,
	},

	"/doc.html": {
		local:   "openapi/doc.html",
		size:    560,
		modtime: 1527699890,
		compressed: `
H4sIAAAAAAAC/0ySQY8TMQyF7/wKk8te2smgRQKVpCAoxxVoxYVjmrgdL5l4FLtdVcB/R520S0+xv5c8
P1lxrzffvvz4+f0rDDrm9SvXDgA3YEjnAsApacb1w/3mMzzihqOzjTR1RA0Qh1AF1ZuD7pbvjb3VShjR
myPh88RVDUQuikW9eaakg094pIjLuVkAFVIKeSkxZPRvzMUoU/kFQ8WdN4PqJCtrd1xUuj3zPmOYSLrI
o40iH3dhpHzyD2cdaw26uu/7xdu+X7zr+z+PvGXlW2SgYvZG9JRRBkS9Dp1JqwG2nE7w+9IAjKHuqayg
//CCppASlf0N+9t87IuRs9e1urPfZU7FxBFkwrg81OzvEkexokEp2shcE5WgXLvTmO/Wzs7XrxFjpUlB
avy/mJhK9yQJMx1rV1Btmcb26FMOiqJ2eygpozTYiYaSQuaC3ZOY9Tnv7NoCt5zOto/xLwAA//9xV16p
MAIAAA==
`,
	},

	"/": {
		isDir: true,
		local: "openapi",
	},
}
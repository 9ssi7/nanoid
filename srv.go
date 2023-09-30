/*
Package nanoid provides a simple API to generate short, unique IDs.

Installation:

	go get github.com/9ssi7/nanoid

Usage:

	import "github.com/9ssi7/nanoid"

	id, err := nanoid.New()
	if err != nil {
		log.Fatal(err)
	}

Custom Length ID:

	id, err := nanoid.NewWithLength(nanoid.LengthMedium)
	if err != nil {
		log.Fatal(err)
	}

Validating ID:

	if !nanoid.Validate(id) {
		log.Fatal("Invalid ID")
	}

Benchmark Results:

	goos: linux
	goarch: amd64
	pkg: github.com/9ssi7/nanoid
	cpu: 12th Gen Intel(R) Core(TM) i7-12700H
	BenchmarkNanoID
	BenchmarkNanoID-20        	 1240796	       933.8 ns/op
	BenchmarkGoogleUUID
	BenchmarkGoogleUUID-20    	 1312452	       903.1 ns/op

Security:

NanoID uses crypto/rand for generating secure random IDs, reducing the likelihood of collision. However, the users should consider the trade-off between ID length and collision probability.

License:

Apache-2.0 License
*/
package nanoid

import (
	"crypto/rand"
	"encoding/base64"
	"regexp"
)

type Length int

const (
	LengthDefault Length = 21
	LengthShort   Length = 5
	LengthMedium  Length = 10
	LengthLong    Length = 15
)

var instanceID string

func newInstanceId() (string, error) {
	bytes := make([]byte, 2)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

func new(length Length) (string, error) {
	if instanceID == "" {
		var err error
		instanceID, err = newInstanceId()
		if err != nil {
			return "", err
		}
	}
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return instanceID + base64.RawURLEncoding.EncodeToString(bytes), nil
}

func New() (string, error) {
	return new(LengthShort)
}

func NewWithLength(length Length) (string, error) {
	return new(length)
}

func MustNew() string {
	id, err := New()
	if err != nil {
		panic(err)
	}
	return id
}

func Validate(id string) bool {
	if len(id) < 8 {
		return false
	}
	matched, err := regexp.MatchString(`^[A-Za-z0-9_-]+$`, id)
	if err != nil || !matched {
		return false
	}
	_, err = base64.RawURLEncoding.DecodeString(id)
	return err == nil
}

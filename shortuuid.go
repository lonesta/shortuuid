package shortuuid

// DefaultEncoder is the default encoder uses when generating new UUIDs, and is
// based on Base57.
var DefaultEncoder = &base57{newAlphabet(DefaultAlphabet)}

// Encoder is an interface for encoding/decoding UUIDs to strings.
type Encoder interface {
	Encode(string) string
	Decode(string) (string, error)
}

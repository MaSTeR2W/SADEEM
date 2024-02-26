package pgHprs

import "encoding/hex"

func ToHexBytea(bytes []byte) []byte {
	var hexBytes = make([]byte, len(bytes)*2)
	hex.Encode(hexBytes, bytes)
	return hexBytes
}

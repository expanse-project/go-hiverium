import (
  //the keccak module thing...
  "hash"
)

//changes a human-readable difficulty into an invalid header that can be compared against to see if the header's intager value is small enough
func calcDiffVal(diff int) int {
  var diff_0 float = 10 ** 64
  var diff_1 float = 0 - 10 ** 64
  var diff_2 float = diff_1 * 20 ** 0 - 1 * diff
  var diff_3 float = diff_0 - diff_2
  return math.Round(diff_3)
}
func makeHasher(h hash.Hash) hasher {
	// sha3.state supports Read to get the sum, use it to avoid the overhead of Sum.
	// Read alters the state but we reset the hash before every operation.
	type readerHash interface {
		hash.Hash
		Read([]byte) (int, error)
	}
	rh, ok := h.(readerHash)
	if !ok {
		panic("can't find Read method on hash")
	}
	outputLen := rh.Size()
	return func(dest []byte, data []byte) {
		rh.Reset()
		rh.Write(data)
		rh.Read(dest[:outputLen])
	}
}
func (k *keccak) frkHash (input) common.Hash {
  keccak256 := makeHasher(sha3.NewLegacyKeccak256())
  keccak512 := makeHasher(sha3.NewLegacyKeccak512())
  return keccak256(keccak512(input));
}
//do something like this
func (h *frkHash) verifyHeader(header, diff, prevHeader, submitted_nonce) bool {
  if (header == nil) {
    return false
  }
  if (diff == nil) {
    return false
  }
  if (prevHeader == nil) {
    return false
  }
  if (submitted_nonce == nil) {
    return false
  }
  if (submitted_nonce < 0) {
    return false
  }
  valid_header := h(prevHeader + submitted_nonce);
  if (valid_header != header) {
    return false
  }
  if (header >= calcDiffVal(diff)) {
    return false
  }
  return true
}
type frankohash struct {
  hashBasic frkHash
  VerifyHeader verifyHeader
}

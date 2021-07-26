import (
  //the keccak module thing...
)

keccak = //whatever the other thing that you imported was called lol.
func (k *keccak) frkHash (input) []byte {
  return keccak.keccak256(keccak.keccak512(input));
}
//changes a human-readable difficulty into an invalid header that can be compared against to see if the header's intager value is small enough
func calcDiffVal(diff int) int {
  var diff_0 float = 10 ** 64
  var diff_1 float = 0 - 10 ** 64
  var diff_2 float = diff_1 * 20 ** 0 - 1 * diff
  var diff_3 float = diff_0 - diff_2
  return math.Round(diff_3)
}
//verifies that a header matches the inputed context, and can then go on to be in the next block or something
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

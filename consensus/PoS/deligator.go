package pos

import (
	"github.com/expanse-org/go-expanse/common"
)

func deligate(deligatee validator, amount uint64, Deligator deligator) {
  if(amount <= SetUint64(GetBalance(Deligator.address)) - deligated) {
    deligatee.votes += amount
    Deligator.deligated += amount
    if (deligatee.votes >= topValidator) {
      Authorize(deligatee.address, func(signer accounts.Account, mimeType string, message []byte) ([]byte, error) {
        sig = Sign(deligatee.key.PrivateKey, message)
        if (sig != nil) {
          return sig
        }
        return error.new("uh... hate to be 'that guy' but something went wrong...")
      })
      //fire something to the nodes to start minting on the proxy.
    }
  }
}

type deligator struct {
  address common.Address
  deligated uint64
}

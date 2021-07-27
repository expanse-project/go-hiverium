package pos

import (
	"github.com/expanse-org/go-expanse/common"
)

func deligate(deligatee validator, amount uint64, Deligator deligator) {
  if(amount <= SetUint64(GetBalance(Deligator.address)) - deligated) {
    deligatee.votes += amount
    Deligator.deligated += amount
    if (deligatee.votes >= topValidator) {
      Authorize(deligatee.address, SignerFn)
      //       find out what this is ^^^^^
    }
  }
}

type deligator struct {
  address common.Address
  deligated uint64
}

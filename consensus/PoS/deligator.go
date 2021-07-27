package pos

import (
  //add stuff
  "core\state"
  "deligator"
)

func deligate(deligatee validator, amount uint64, Deligator deligator) {
  if(amount <= SetUint64(GetBalance(Deligator.address)) - deligated) {
    deligatee.votes += amount
    Deligator.deligated += amount
  }
}

type deligator struct {
  address common.Address
  deligated uint64
}

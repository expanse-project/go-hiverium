package pos

import (
  "pos"
)

func Author(header *types.Header) (common.Address, error) {
  return header.Coinbase, nil
}
func VerifyHeader(chain ChainHeaderReader, header *types.Header, seal bool) error

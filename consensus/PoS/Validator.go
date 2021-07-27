package pos
import (

	"github.com/expanse-org/go-expanse/common"
)

type validator struct {
  votes uint64
  address common.Address
  nonce uint64
}
var TopValidator validator

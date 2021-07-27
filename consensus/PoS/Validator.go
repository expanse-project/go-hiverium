package pos
import (

	"github.com/expanse-org/go-expanse/common"

)

type validator struct {
  votes uint64
  address common.Address
  key keything
  nonce uint64
}
func init() {
  validator.key, _ = crypto.GenerateKey()
}

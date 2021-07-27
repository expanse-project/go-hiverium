package pos
import (

	"github.com/expanse-org/go-expanse/common"
  "github.com/expanse-org/go-expanse/crypto"
)

type validator struct {
  votes uint64
  address common.Address
  key crypto.PublicKey
  prvkey crypto.PrivateKey
  nonce uint64
}
func init() {
  validator.key, validator.prvkey = crypto.GenerateKey()
}

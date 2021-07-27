package pos
import (
	"bytes"
	"errors"
	"io"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/expanse-org/go-expanse/accounts"
	"github.com/expanse-org/go-expanse/common"
	"github.com/expanse-org/go-expanse/common/hexutil"
	"github.com/expanse-org/go-expanse/consensus"
	"github.com/expanse-org/go-expanse/consensus/misc"
	"github.com/expanse-org/go-expanse/core/state"
	"github.com/expanse-org/go-expanse/core/types"
	"github.com/expanse-org/go-expanse/crypto"
	"github.com/expanse-org/go-expanse/ethdb"
	"github.com/expanse-org/go-expanse/log"
	"github.com/expanse-org/go-expanse/params"
	"github.com/expanse-org/go-expanse/rlp"
	"github.com/expanse-org/go-expanse/rpc"
	"github.com/expanse-org/go-expanse/trie"
	lru "github.com/hashicorp/golang-lru"
	"golang.org/x/crypto/sha3"
)

type validator struct {
  votes uint64
  common.Address
}

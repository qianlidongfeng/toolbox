package toolbox

import (
	"crypto/md5"
	"fmt"
)

func MD5(b []byte) string{
	return fmt.Sprintf("%x", md5.Sum(b))
}

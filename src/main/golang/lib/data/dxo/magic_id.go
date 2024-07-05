package dxo

import "github.com/starter-go/base/lang"

// ShortLinkMagic ... 是一个 base64 形式的 Short-URL-ID
type ShortLinkMagic lang.Base64

func (mid ShortLinkMagic) String() string {
	return string(mid)
}

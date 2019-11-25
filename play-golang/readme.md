# collection of the play.golang.org

collections of the small golang tests


1. how to send []byte as string and rollback to []byte
- link: https://play.golang.org/p/AFSiSOICUDY
  1) []byte to string with hex.EncodeToString()
  2) back to []byte with hex.DecodeString()


2. need Reset(), if you want to reuse hash.Hash
- reference link: https://golang.org/pkg/crypto/hmac/ 
- example link: https://play.golang.org/p/apc25S7du9t
- after create hmac.New(), then `Write()` -> you need to `Reset()` for later use
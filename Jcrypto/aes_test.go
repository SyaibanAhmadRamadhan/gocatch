package Jcrypto

import (
	"testing"
)

var key string
var strEnc string

func init() {
	ip := "127.0.0.1"
	device := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"

	key = GenerateKey(ip + device)

}
func TestAesEncryptCFB(t *testing.T) {
	userID := "01HG3MMXSSFEB3VG3NANYRSY74"
	sessionID := "01HG3MNEW08K3Y2ARZB4RF3J9F"

	webUniq, err := AesEncryptCFB(sessionID+"|"+userID, key)
	strEnc = webUniq
	t.Log(err)
	t.Log(webUniq)
}

func TestAesDecryptCFB(t *testing.T) {
	decrypt, err := AesDecryptCFB(strEnc, key)
	t.Log(err)
	t.Log(decrypt)
}

func TestGenerateKey(t *testing.T) {
	str := "rama"
	key := GenerateKey(str)
	t.Log(key)
}

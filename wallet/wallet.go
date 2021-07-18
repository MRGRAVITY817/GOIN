package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/MRGRAVITY817/goin/utils"
)

// 1) hash the message
// F(message) -> "hashed_message"
// 2) generate the key pair
// (privKey, pubKey) (save priv to a file)
// 3) sign the hash
// ("hashed_message", privKey) -> "signature"
// 4) verify the signature
// ("hashed_message", "signature", pubKey) -> true/false

const (
	privateKey string = "307702010104201c11b8ebc0e123ba66194f16dd1a8f67f1d9828d358a5253917cf4dee913de82a00a06082a8648ce3d030107a144034200041550289d1bc03cd51e1204719fcac79aede87e5a6244e84bce422c548e2257150bb3dcf663c86313044c56496b29f1e2853e8a8128357667b1cff20653b2d98a"
	hashedMsg  string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
	signature  string = "13449dd6911ea621961294fa73513c1da1b7d2669c14aacc0b0483eb51a2df04d17ca513ebb63635b2feaa1f66b273e7fb77b8a835e7d74c5ca6ff0d0f8ee0b6"
)

func Start() {
	// Confirming privatekey is hexadecimal and correct
	privByte, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	_, err = x509.ParseECPrivateKey([]byte(privByte))
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	// sigBytes = [rBytes, sBytes]
	rBytes := sigBytes[:len(sigBytes)/2] // getting the first half of the sig
	sBytes := sigBytes[len(sigBytes)/2:] // getting the last half of the sig
	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
}

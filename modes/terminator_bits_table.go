package modes

import (
	"github.com/harogaston/go-mosaic/bitseq"
	"github.com/harogaston/go-mosaic/version"
)

func GetTerminatorBits(qrversion version.QRVersion, mode QRMode) bitseq.BitSeq {
	// TODO: Implement different terminator bits for Micro QR Codes
	return bitseq.FromInt(0, 4)
}

package main

import (
	"testing"
)

// TestTotalCodewordsIntegrity checks that for every entry, the sum of total codewords
// across all blocks matches the defined total codewords for that version.
func TestTotalCodewordsIntegrity(t *testing.T) {
	ecNames := map[errcorr]string{
		ERR_CORR_L: "L",
		ERR_CORR_M: "M",
		ERR_CORR_Q: "Q",
		ERR_CORR_H: "H",
	}

	for version, data := range capacityData {
		expectedTotal := data.totalCodewords

		for ecLevel, info := range data.ecInfo {
			calculatedTotal := 0
			for _, bg := range info.BlockGroups {
				// Calculate total codewords: Number of Blocks * Total Codewords per Block
				calculatedTotal += bg.NumBlocks * bg.TotalCodewords
			}

			if calculatedTotal != expectedTotal {
				t.Errorf("Mismatch in Version %d [%s]: Expected Total Codewords %d, but calculated sum of blocks is %d",
					version, ecNames[ecLevel], expectedTotal, calculatedTotal)
			}
		}
	}
}

// TestECCodewordsIntegrity checks that for every entry, the sum of calculated error correction codewords
// (Total Codewords - Data Codewords) across all blocks matches the defined TotalECCodewords.
func TestECCodewordsIntegrity(t *testing.T) {
	ecNames := map[errcorr]string{
		ERR_CORR_L: "L",
		ERR_CORR_M: "M",
		ERR_CORR_Q: "Q",
		ERR_CORR_H: "H",
	}

	for version, data := range capacityData {
		for ecLevel, info := range data.ecInfo {
			expectedECCodewords := info.TotalECCodewords
			calculatedECCodewords := 0

			for _, bg := range info.BlockGroups {
				// Calculate EC codewords for this group: NumBlocks * (Total - Data)
				ecPerBlock := bg.TotalCodewords - bg.DataCodewords
				calculatedECCodewords += bg.NumBlocks * ecPerBlock
			}

			if calculatedECCodewords != expectedECCodewords {
				t.Errorf("Mismatch in Version %d [%s]: Expected Total EC Codewords %d, but calculated sum is %d",
					version, ecNames[ecLevel], expectedECCodewords, calculatedECCodewords)
			}
		}
	}
}

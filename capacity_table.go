package main

import (
	"fmt"

	"github.com/harogaston/go-mosaic/version"
)

// BlockGroup represents a group of blocks with the same characteristics
type BlockGroup struct {
	NumBlocks      int
	TotalCodewords int
	DataCodewords  int
}

// ECInfo stores error correction information for a specific level
type ECInfo struct {
	TotalECCodewords int
	BlockGroups      []BlockGroup
}

var microCapacityData = map[int]struct {
	totalCodewords int
	ecInfo         map[errcorr]ECInfo
}{
	1: {
		totalCodewords: 5,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 2,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 5, DataCodewords: 3},
				},
			},
		},
	},
	2: {
		totalCodewords: 10,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 5,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 10, DataCodewords: 5},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 6,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 10, DataCodewords: 4},
				},
			},
		},
	},
	3: {
		totalCodewords: 17,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 6,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 17, DataCodewords: 11},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 8,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 17, DataCodewords: 9},
				},
			},
		},
	},
	4: {
		totalCodewords: 24,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 8,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 24, DataCodewords: 16},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 10,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 24, DataCodewords: 14},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 14,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 14, DataCodewords: 10},
				},
			},
		},
	},
}

// capacityData maps QR Code Version (1-40) to its capacity and error correction info
var capacityData = map[int]struct {
	totalCodewords int
	ecInfo         map[errcorr]ECInfo
}{
	1: {
		totalCodewords: 26,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 7,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 26, DataCodewords: 19},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 10,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 26, DataCodewords: 16},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 13,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 26, DataCodewords: 13},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 17,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 26, DataCodewords: 9},
				},
			},
		},
	},
	2: {
		totalCodewords: 44,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 10,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 44, DataCodewords: 34},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 16,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 44, DataCodewords: 28},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 22,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 44, DataCodewords: 22},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 28,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 44, DataCodewords: 16},
				},
			},
		},
	},
	3: {
		totalCodewords: 70,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 15,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 70, DataCodewords: 55},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 26,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 70, DataCodewords: 44},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 36,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 35, DataCodewords: 17},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 44,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 35, DataCodewords: 13},
				},
			},
		},
	},
	4: {
		totalCodewords: 100,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 20,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 100, DataCodewords: 80},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 36,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 50, DataCodewords: 32},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 52,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 50, DataCodewords: 24},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 64,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 25, DataCodewords: 9},
				},
			},
		},
	},
	5: {
		totalCodewords: 134,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 26,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 134, DataCodewords: 108},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 48,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 67, DataCodewords: 43},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 72,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 33, DataCodewords: 15},
					{NumBlocks: 2, TotalCodewords: 34, DataCodewords: 16},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 88,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 33, DataCodewords: 11},
					{NumBlocks: 2, TotalCodewords: 34, DataCodewords: 12},
				},
			},
		},
	},
	6: {
		totalCodewords: 172,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 36,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 86, DataCodewords: 68},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 64,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 43, DataCodewords: 27},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 96,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 43, DataCodewords: 19},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 112,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 43, DataCodewords: 15},
				},
			},
		},
	},
	7: {
		totalCodewords: 196,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 40,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 98, DataCodewords: 78},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 72,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 49, DataCodewords: 31},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 108,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 32, DataCodewords: 14},
					{NumBlocks: 4, TotalCodewords: 33, DataCodewords: 15},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 130,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 39, DataCodewords: 13},
					{NumBlocks: 1, TotalCodewords: 40, DataCodewords: 14},
				},
			},
		},
	},
	8: {
		totalCodewords: 242,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 48,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 121, DataCodewords: 97},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 88,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 60, DataCodewords: 38},
					{NumBlocks: 2, TotalCodewords: 61, DataCodewords: 39},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 132,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 40, DataCodewords: 18},
					{NumBlocks: 2, TotalCodewords: 41, DataCodewords: 19},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 156,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 40, DataCodewords: 14},
					{NumBlocks: 2, TotalCodewords: 41, DataCodewords: 15},
				},
			},
		},
	},
	9: {
		totalCodewords: 292,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 60,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 146, DataCodewords: 116},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 110,
				BlockGroups: []BlockGroup{
					{NumBlocks: 3, TotalCodewords: 58, DataCodewords: 36},
					{NumBlocks: 2, TotalCodewords: 59, DataCodewords: 37},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 160,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 36, DataCodewords: 16},
					{NumBlocks: 4, TotalCodewords: 37, DataCodewords: 17},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 192,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 36, DataCodewords: 12},
					{NumBlocks: 4, TotalCodewords: 37, DataCodewords: 13},
				},
			},
		},
	},
	10: {
		totalCodewords: 346,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 72,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 86, DataCodewords: 68},
					{NumBlocks: 2, TotalCodewords: 87, DataCodewords: 69},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 130,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 69, DataCodewords: 43},
					{NumBlocks: 1, TotalCodewords: 70, DataCodewords: 44},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 192,
				BlockGroups: []BlockGroup{
					{NumBlocks: 6, TotalCodewords: 43, DataCodewords: 19},
					{NumBlocks: 2, TotalCodewords: 44, DataCodewords: 20},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 224,
				BlockGroups: []BlockGroup{
					{NumBlocks: 6, TotalCodewords: 43, DataCodewords: 15},
					{NumBlocks: 2, TotalCodewords: 44, DataCodewords: 16},
				},
			},
		},
	},
	11: {
		totalCodewords: 404,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 80,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 101, DataCodewords: 81},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 150,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 80, DataCodewords: 50},
					{NumBlocks: 4, TotalCodewords: 81, DataCodewords: 51},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 224,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 50, DataCodewords: 22},
					{NumBlocks: 4, TotalCodewords: 51, DataCodewords: 23},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 264,
				BlockGroups: []BlockGroup{
					{NumBlocks: 3, TotalCodewords: 36, DataCodewords: 12},
					{NumBlocks: 8, TotalCodewords: 37, DataCodewords: 13},
				},
			},
		},
	},
	12: {
		totalCodewords: 466,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 96,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 116, DataCodewords: 92},
					{NumBlocks: 2, TotalCodewords: 117, DataCodewords: 93},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 176,
				BlockGroups: []BlockGroup{
					{NumBlocks: 6, TotalCodewords: 58, DataCodewords: 36},
					{NumBlocks: 2, TotalCodewords: 59, DataCodewords: 37},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 260,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 46, DataCodewords: 20},
					{NumBlocks: 6, TotalCodewords: 47, DataCodewords: 21},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 308,
				BlockGroups: []BlockGroup{
					{NumBlocks: 7, TotalCodewords: 42, DataCodewords: 14},
					{NumBlocks: 4, TotalCodewords: 43, DataCodewords: 15},
				},
			},
		},
	},
	13: {
		totalCodewords: 532,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 104,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 133, DataCodewords: 107},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 198,
				BlockGroups: []BlockGroup{
					{NumBlocks: 8, TotalCodewords: 59, DataCodewords: 37},
					{NumBlocks: 1, TotalCodewords: 60, DataCodewords: 38},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 288,
				BlockGroups: []BlockGroup{
					{NumBlocks: 8, TotalCodewords: 44, DataCodewords: 20},
					{NumBlocks: 4, TotalCodewords: 45, DataCodewords: 21},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 352,
				BlockGroups: []BlockGroup{
					{NumBlocks: 12, TotalCodewords: 33, DataCodewords: 11},
					{NumBlocks: 4, TotalCodewords: 34, DataCodewords: 12},
				},
			},
		},
	},
	14: {
		totalCodewords: 581,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 120,
				BlockGroups: []BlockGroup{
					{NumBlocks: 3, TotalCodewords: 145, DataCodewords: 115},
					{NumBlocks: 1, TotalCodewords: 146, DataCodewords: 116},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 216,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 64, DataCodewords: 40},
					{NumBlocks: 5, TotalCodewords: 65, DataCodewords: 41},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 320,
				BlockGroups: []BlockGroup{
					{NumBlocks: 11, TotalCodewords: 36, DataCodewords: 16},
					{NumBlocks: 5, TotalCodewords: 37, DataCodewords: 17},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 384,
				BlockGroups: []BlockGroup{
					{NumBlocks: 11, TotalCodewords: 36, DataCodewords: 12},
					{NumBlocks: 5, TotalCodewords: 37, DataCodewords: 13},
				},
			},
		},
	},
	15: {
		totalCodewords: 655,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 132,
				BlockGroups: []BlockGroup{
					{NumBlocks: 5, TotalCodewords: 109, DataCodewords: 87},
					{NumBlocks: 1, TotalCodewords: 110, DataCodewords: 88},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 240,
				BlockGroups: []BlockGroup{
					{NumBlocks: 5, TotalCodewords: 65, DataCodewords: 41},
					{NumBlocks: 5, TotalCodewords: 66, DataCodewords: 42},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 360,
				BlockGroups: []BlockGroup{
					{NumBlocks: 5, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 7, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 432,
				BlockGroups: []BlockGroup{
					{NumBlocks: 11, TotalCodewords: 36, DataCodewords: 12},
					{NumBlocks: 7, TotalCodewords: 37, DataCodewords: 13},
				},
			},
		},
	},
	16: {
		totalCodewords: 733,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 144,
				BlockGroups: []BlockGroup{
					{NumBlocks: 5, TotalCodewords: 122, DataCodewords: 98},
					{NumBlocks: 1, TotalCodewords: 123, DataCodewords: 99},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 280,
				BlockGroups: []BlockGroup{
					{NumBlocks: 7, TotalCodewords: 73, DataCodewords: 45},
					{NumBlocks: 3, TotalCodewords: 74, DataCodewords: 46},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 408,
				BlockGroups: []BlockGroup{
					{NumBlocks: 15, TotalCodewords: 43, DataCodewords: 19},
					{NumBlocks: 2, TotalCodewords: 44, DataCodewords: 20},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 480,
				BlockGroups: []BlockGroup{
					{NumBlocks: 3, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 13, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	17: {
		totalCodewords: 815,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 168,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 135, DataCodewords: 107},
					{NumBlocks: 5, TotalCodewords: 136, DataCodewords: 108},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 308,
				BlockGroups: []BlockGroup{
					{NumBlocks: 10, TotalCodewords: 74, DataCodewords: 46},
					{NumBlocks: 1, TotalCodewords: 75, DataCodewords: 47},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 448,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 50, DataCodewords: 22},
					{NumBlocks: 15, TotalCodewords: 51, DataCodewords: 23},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 532,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 42, DataCodewords: 14},
					{NumBlocks: 17, TotalCodewords: 43, DataCodewords: 15},
				},
			},
		},
	},
	18: {
		totalCodewords: 901,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 180,
				BlockGroups: []BlockGroup{
					{NumBlocks: 5, TotalCodewords: 150, DataCodewords: 120},
					{NumBlocks: 1, TotalCodewords: 151, DataCodewords: 121},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 338,
				BlockGroups: []BlockGroup{
					{NumBlocks: 9, TotalCodewords: 69, DataCodewords: 43},
					{NumBlocks: 4, TotalCodewords: 70, DataCodewords: 44},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 504,
				BlockGroups: []BlockGroup{
					{NumBlocks: 17, TotalCodewords: 50, DataCodewords: 22},
					{NumBlocks: 1, TotalCodewords: 51, DataCodewords: 23},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 588,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 42, DataCodewords: 14},
					{NumBlocks: 19, TotalCodewords: 43, DataCodewords: 15},
				},
			},
		},
	},
	19: {
		totalCodewords: 991,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 196,
				BlockGroups: []BlockGroup{
					{NumBlocks: 3, TotalCodewords: 141, DataCodewords: 113},
					{NumBlocks: 4, TotalCodewords: 142, DataCodewords: 114},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 364,
				BlockGroups: []BlockGroup{
					{NumBlocks: 3, TotalCodewords: 70, DataCodewords: 44},
					{NumBlocks: 11, TotalCodewords: 71, DataCodewords: 45},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 546,
				BlockGroups: []BlockGroup{
					{NumBlocks: 17, TotalCodewords: 47, DataCodewords: 21},
					{NumBlocks: 4, TotalCodewords: 48, DataCodewords: 22},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 650,
				BlockGroups: []BlockGroup{
					{NumBlocks: 9, TotalCodewords: 39, DataCodewords: 13},
					{NumBlocks: 16, TotalCodewords: 40, DataCodewords: 14},
				},
			},
		},
	},
	20: {
		totalCodewords: 1085,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 224,
				BlockGroups: []BlockGroup{
					{NumBlocks: 3, TotalCodewords: 135, DataCodewords: 107},
					{NumBlocks: 5, TotalCodewords: 136, DataCodewords: 108},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 416,
				BlockGroups: []BlockGroup{
					{NumBlocks: 3, TotalCodewords: 67, DataCodewords: 41},
					{NumBlocks: 13, TotalCodewords: 68, DataCodewords: 42},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 600,
				BlockGroups: []BlockGroup{
					{NumBlocks: 15, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 5, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 700,
				BlockGroups: []BlockGroup{
					{NumBlocks: 15, TotalCodewords: 43, DataCodewords: 15},
					{NumBlocks: 10, TotalCodewords: 44, DataCodewords: 16},
				},
			},
		},
	},
	21: {
		totalCodewords: 1156,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 224,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 144, DataCodewords: 116},
					{NumBlocks: 4, TotalCodewords: 145, DataCodewords: 117},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 442,
				BlockGroups: []BlockGroup{
					{NumBlocks: 17, TotalCodewords: 68, DataCodewords: 42},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 644,
				BlockGroups: []BlockGroup{
					{NumBlocks: 17, TotalCodewords: 50, DataCodewords: 22},
					{NumBlocks: 6, TotalCodewords: 51, DataCodewords: 23},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 750,
				BlockGroups: []BlockGroup{
					{NumBlocks: 19, TotalCodewords: 46, DataCodewords: 16},
					{NumBlocks: 6, TotalCodewords: 47, DataCodewords: 17},
				},
			},
		},
	},
	22: {
		totalCodewords: 1258,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 252,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 139, DataCodewords: 111},
					{NumBlocks: 7, TotalCodewords: 140, DataCodewords: 112},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 476,
				BlockGroups: []BlockGroup{
					{NumBlocks: 17, TotalCodewords: 74, DataCodewords: 46},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 690,
				BlockGroups: []BlockGroup{
					{NumBlocks: 7, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 16, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 816,
				BlockGroups: []BlockGroup{
					{NumBlocks: 34, TotalCodewords: 37, DataCodewords: 13},
				},
			},
		},
	},
	23: {
		totalCodewords: 1364,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 270,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 151, DataCodewords: 121},
					{NumBlocks: 5, TotalCodewords: 152, DataCodewords: 122},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 504,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 75, DataCodewords: 47},
					{NumBlocks: 14, TotalCodewords: 76, DataCodewords: 48},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 750,
				BlockGroups: []BlockGroup{
					{NumBlocks: 11, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 14, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 900,
				BlockGroups: []BlockGroup{
					{NumBlocks: 16, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 14, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	24: {
		totalCodewords: 1474,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 300,
				BlockGroups: []BlockGroup{
					{NumBlocks: 6, TotalCodewords: 147, DataCodewords: 117},
					{NumBlocks: 4, TotalCodewords: 148, DataCodewords: 118},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 560,
				BlockGroups: []BlockGroup{
					{NumBlocks: 6, TotalCodewords: 73, DataCodewords: 45},
					{NumBlocks: 14, TotalCodewords: 74, DataCodewords: 46},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 810,
				BlockGroups: []BlockGroup{
					{NumBlocks: 11, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 16, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 960,
				BlockGroups: []BlockGroup{
					{NumBlocks: 30, TotalCodewords: 46, DataCodewords: 16},
					{NumBlocks: 2, TotalCodewords: 47, DataCodewords: 17},
				},
			},
		},
	},
	25: {
		totalCodewords: 1588,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 312,
				BlockGroups: []BlockGroup{
					{NumBlocks: 8, TotalCodewords: 132, DataCodewords: 106},
					{NumBlocks: 4, TotalCodewords: 133, DataCodewords: 107},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 588,
				BlockGroups: []BlockGroup{
					{NumBlocks: 8, TotalCodewords: 75, DataCodewords: 47},
					{NumBlocks: 13, TotalCodewords: 76, DataCodewords: 48},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 870,
				BlockGroups: []BlockGroup{
					{NumBlocks: 7, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 22, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1050,
				BlockGroups: []BlockGroup{
					{NumBlocks: 22, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 13, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	26: {
		totalCodewords: 1706,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 336,
				BlockGroups: []BlockGroup{
					{NumBlocks: 10, TotalCodewords: 142, DataCodewords: 114},
					{NumBlocks: 2, TotalCodewords: 143, DataCodewords: 115},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 644,
				BlockGroups: []BlockGroup{
					{NumBlocks: 19, TotalCodewords: 74, DataCodewords: 46},
					{NumBlocks: 4, TotalCodewords: 75, DataCodewords: 47},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 952,
				BlockGroups: []BlockGroup{
					{NumBlocks: 28, TotalCodewords: 50, DataCodewords: 22},
					{NumBlocks: 6, TotalCodewords: 51, DataCodewords: 23},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1110,
				BlockGroups: []BlockGroup{
					{NumBlocks: 33, TotalCodewords: 46, DataCodewords: 16},
					{NumBlocks: 4, TotalCodewords: 47, DataCodewords: 17},
				},
			},
		},
	},
	27: {
		totalCodewords: 1828,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 360,
				BlockGroups: []BlockGroup{
					{NumBlocks: 8, TotalCodewords: 152, DataCodewords: 122},
					{NumBlocks: 4, TotalCodewords: 153, DataCodewords: 123},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 700,
				BlockGroups: []BlockGroup{
					{NumBlocks: 22, TotalCodewords: 73, DataCodewords: 45},
					{NumBlocks: 3, TotalCodewords: 74, DataCodewords: 46},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1020,
				BlockGroups: []BlockGroup{
					{NumBlocks: 8, TotalCodewords: 53, DataCodewords: 23},
					{NumBlocks: 26, TotalCodewords: 54, DataCodewords: 24},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1200,
				BlockGroups: []BlockGroup{
					{NumBlocks: 12, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 28, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	28: {
		totalCodewords: 1921,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 390,
				BlockGroups: []BlockGroup{
					{NumBlocks: 3, TotalCodewords: 147, DataCodewords: 117},
					{NumBlocks: 10, TotalCodewords: 148, DataCodewords: 118},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 728,
				BlockGroups: []BlockGroup{
					{NumBlocks: 3, TotalCodewords: 73, DataCodewords: 45},
					{NumBlocks: 23, TotalCodewords: 74, DataCodewords: 46},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1050,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 31, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1260,
				BlockGroups: []BlockGroup{
					{NumBlocks: 11, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 31, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	29: {
		totalCodewords: 2051,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 420,
				BlockGroups: []BlockGroup{
					{NumBlocks: 7, TotalCodewords: 146, DataCodewords: 116},
					{NumBlocks: 7, TotalCodewords: 147, DataCodewords: 117},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 784,
				BlockGroups: []BlockGroup{
					{NumBlocks: 21, TotalCodewords: 73, DataCodewords: 45},
					{NumBlocks: 7, TotalCodewords: 74, DataCodewords: 46},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1140,
				BlockGroups: []BlockGroup{
					{NumBlocks: 1, TotalCodewords: 53, DataCodewords: 23},
					{NumBlocks: 37, TotalCodewords: 54, DataCodewords: 24},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1350,
				BlockGroups: []BlockGroup{
					{NumBlocks: 19, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 26, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	30: {
		totalCodewords: 2185,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 450,
				BlockGroups: []BlockGroup{
					{NumBlocks: 5, TotalCodewords: 145, DataCodewords: 115},
					{NumBlocks: 10, TotalCodewords: 146, DataCodewords: 116},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 812,
				BlockGroups: []BlockGroup{
					{NumBlocks: 19, TotalCodewords: 75, DataCodewords: 47},
					{NumBlocks: 10, TotalCodewords: 76, DataCodewords: 48},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1200,
				BlockGroups: []BlockGroup{
					{NumBlocks: 15, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 25, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1440,
				BlockGroups: []BlockGroup{
					{NumBlocks: 23, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 25, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	31: {
		totalCodewords: 2323,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 480,
				BlockGroups: []BlockGroup{
					{NumBlocks: 13, TotalCodewords: 145, DataCodewords: 115},
					{NumBlocks: 3, TotalCodewords: 146, DataCodewords: 116},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 868,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 74, DataCodewords: 46},
					{NumBlocks: 29, TotalCodewords: 75, DataCodewords: 47},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1290,
				BlockGroups: []BlockGroup{
					{NumBlocks: 42, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 1, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1530,
				BlockGroups: []BlockGroup{
					{NumBlocks: 23, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 28, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	32: {
		totalCodewords: 2465,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 510,
				BlockGroups: []BlockGroup{
					{NumBlocks: 17, TotalCodewords: 145, DataCodewords: 115},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 924,
				BlockGroups: []BlockGroup{
					{NumBlocks: 10, TotalCodewords: 74, DataCodewords: 46},
					{NumBlocks: 23, TotalCodewords: 75, DataCodewords: 47},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1350,
				BlockGroups: []BlockGroup{
					{NumBlocks: 10, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 35, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1620,
				BlockGroups: []BlockGroup{
					{NumBlocks: 19, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 35, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	33: {
		totalCodewords: 2611,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 540,
				BlockGroups: []BlockGroup{
					{NumBlocks: 17, TotalCodewords: 145, DataCodewords: 115},
					{NumBlocks: 1, TotalCodewords: 146, DataCodewords: 116},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 980,
				BlockGroups: []BlockGroup{
					{NumBlocks: 14, TotalCodewords: 74, DataCodewords: 46},
					{NumBlocks: 21, TotalCodewords: 75, DataCodewords: 47},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1440,
				BlockGroups: []BlockGroup{
					{NumBlocks: 29, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 19, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1710,
				BlockGroups: []BlockGroup{
					{NumBlocks: 11, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 46, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	34: {
		totalCodewords: 2761,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 570,
				BlockGroups: []BlockGroup{
					{NumBlocks: 13, TotalCodewords: 145, DataCodewords: 115},
					{NumBlocks: 6, TotalCodewords: 146, DataCodewords: 116},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 1036,
				BlockGroups: []BlockGroup{
					{NumBlocks: 14, TotalCodewords: 74, DataCodewords: 46},
					{NumBlocks: 23, TotalCodewords: 75, DataCodewords: 47},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1530,
				BlockGroups: []BlockGroup{
					{NumBlocks: 44, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 7, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1800,
				BlockGroups: []BlockGroup{
					{NumBlocks: 59, TotalCodewords: 46, DataCodewords: 16},
					{NumBlocks: 1, TotalCodewords: 47, DataCodewords: 17},
				},
			},
		},
	},
	35: {
		totalCodewords: 2876,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 570,
				BlockGroups: []BlockGroup{
					{NumBlocks: 12, TotalCodewords: 151, DataCodewords: 121},
					{NumBlocks: 7, TotalCodewords: 152, DataCodewords: 122},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 1064,
				BlockGroups: []BlockGroup{
					{NumBlocks: 12, TotalCodewords: 75, DataCodewords: 47},
					{NumBlocks: 26, TotalCodewords: 76, DataCodewords: 48},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1590,
				BlockGroups: []BlockGroup{
					{NumBlocks: 39, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 14, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1890,
				BlockGroups: []BlockGroup{
					{NumBlocks: 22, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 41, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	36: {
		totalCodewords: 3034,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 600,
				BlockGroups: []BlockGroup{
					{NumBlocks: 6, TotalCodewords: 151, DataCodewords: 121},
					{NumBlocks: 14, TotalCodewords: 152, DataCodewords: 122},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 1120,
				BlockGroups: []BlockGroup{
					{NumBlocks: 6, TotalCodewords: 75, DataCodewords: 47},
					{NumBlocks: 34, TotalCodewords: 76, DataCodewords: 48},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1680,
				BlockGroups: []BlockGroup{
					{NumBlocks: 46, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 10, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 1980,
				BlockGroups: []BlockGroup{
					{NumBlocks: 2, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 64, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	37: {
		totalCodewords: 3196,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 630,
				BlockGroups: []BlockGroup{
					{NumBlocks: 17, TotalCodewords: 152, DataCodewords: 122},
					{NumBlocks: 4, TotalCodewords: 153, DataCodewords: 123},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 1204,
				BlockGroups: []BlockGroup{
					{NumBlocks: 29, TotalCodewords: 74, DataCodewords: 46},
					{NumBlocks: 14, TotalCodewords: 75, DataCodewords: 47},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1770,
				BlockGroups: []BlockGroup{
					{NumBlocks: 49, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 10, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 2100,
				BlockGroups: []BlockGroup{
					{NumBlocks: 24, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 46, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	38: {
		totalCodewords: 3362,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 660,
				BlockGroups: []BlockGroup{
					{NumBlocks: 4, TotalCodewords: 152, DataCodewords: 122},
					{NumBlocks: 18, TotalCodewords: 153, DataCodewords: 123},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 1260,
				BlockGroups: []BlockGroup{
					{NumBlocks: 13, TotalCodewords: 74, DataCodewords: 46},
					{NumBlocks: 32, TotalCodewords: 75, DataCodewords: 47},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1860,
				BlockGroups: []BlockGroup{
					{NumBlocks: 48, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 14, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 2220,
				BlockGroups: []BlockGroup{
					{NumBlocks: 42, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 32, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	39: {
		totalCodewords: 3532,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 720,
				BlockGroups: []BlockGroup{
					{NumBlocks: 20, TotalCodewords: 147, DataCodewords: 117},
					{NumBlocks: 4, TotalCodewords: 148, DataCodewords: 118},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 1316,
				BlockGroups: []BlockGroup{
					{NumBlocks: 40, TotalCodewords: 75, DataCodewords: 47},
					{NumBlocks: 7, TotalCodewords: 76, DataCodewords: 48},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 1950,
				BlockGroups: []BlockGroup{
					{NumBlocks: 43, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 22, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 2310,
				BlockGroups: []BlockGroup{
					{NumBlocks: 10, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 67, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
	40: {
		totalCodewords: 3706,
		ecInfo: map[errcorr]ECInfo{
			ERR_CORR_L: {
				TotalECCodewords: 750,
				BlockGroups: []BlockGroup{
					{NumBlocks: 19, TotalCodewords: 148, DataCodewords: 118},
					{NumBlocks: 6, TotalCodewords: 149, DataCodewords: 119},
				},
			},
			ERR_CORR_M: {
				TotalECCodewords: 1372,
				BlockGroups: []BlockGroup{
					{NumBlocks: 18, TotalCodewords: 75, DataCodewords: 47},
					{NumBlocks: 31, TotalCodewords: 76, DataCodewords: 48},
				},
			},
			ERR_CORR_Q: {
				TotalECCodewords: 2040,
				BlockGroups: []BlockGroup{
					{NumBlocks: 34, TotalCodewords: 54, DataCodewords: 24},
					{NumBlocks: 34, TotalCodewords: 55, DataCodewords: 25},
				},
			},
			ERR_CORR_H: {
				TotalECCodewords: 2430,
				BlockGroups: []BlockGroup{
					{NumBlocks: 20, TotalCodewords: 45, DataCodewords: 15},
					{NumBlocks: 61, TotalCodewords: 46, DataCodewords: 16},
				},
			},
		},
	},
}

func getTotalCodewords(v version.QRVersion) int {
	switch v.Format {
	case version.FORMAT_MICRO_QR:
		data := microCapacityData[v.Number]
		return data.totalCodewords
	case version.FORMAT_QR:
		fallthrough
	case version.FORMAT_QR_MODEL_2:
		data := capacityData[v.Number]
		return data.totalCodewords
	}
	return 0
}

func getTotalECCodewords(v version.QRVersion, ecLevel errcorr) int {
	var ecInfo ECInfo
	switch v.Format {
	case version.FORMAT_MICRO_QR:
		capacityData := microCapacityData[v.Number]
		ecInfo = capacityData.ecInfo[ecLevel]
	case version.FORMAT_QR:
		fallthrough
	case version.FORMAT_QR_MODEL_2:
		capacityData := capacityData[v.Number]
		ecInfo = capacityData.ecInfo[ecLevel]
	}

	return ecInfo.TotalECCodewords
}

func getTotalDataCodewords(v version.QRVersion, ecLevel errcorr) int {
	return getTotalCodewords(v) - getTotalECCodewords(v, ecLevel)
}

func ValidateCapacityData() {
	fmt.Println("Starting validation of QR Code capacity data...")
	hasError := false

	// Map to print string representation of EC levels
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
			calculatedTotalDataCodewords := 0
			for _, bg := range info.BlockGroups {
				// Calculate total codewords: Number of Blocks * Total Codewords per Block
				calculatedTotal += bg.NumBlocks * bg.TotalCodewords
				calculatedTotalDataCodewords += bg.NumBlocks * (bg.TotalCodewords - bg.DataCodewords)
			}

			if calculatedTotal != expectedTotal {
				fmt.Printf("Mismatch in Version %d [%s]: Expected %d, but calculated sum of blocks is %d\n",
					version, ecNames[ecLevel], expectedTotal, calculatedTotal)
				hasError = true
			}

			if calculatedTotalDataCodewords != info.TotalECCodewords {
				fmt.Printf("Mismatch in Version %d [%s]: Expected EC Codewords %d, but calculated sum is %d\n",
					version, ecNames[ecLevel], info.TotalECCodewords, calculatedTotalDataCodewords)
				hasError = true
			}
		}
	}

	if !hasError {
		fmt.Println("Success: All Version and Error Correction Level block calculations match the expected total codewords.")
	} else {
		fmt.Println("Validation finished with errors.")
	}
}

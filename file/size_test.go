package file

import (
	"testing"
)

type sizeResult struct {
	inStr     string
	outResult Size
}

func TestParseSize(t *testing.T) {

	result := []sizeResult{
		{"1k", KByte},
		{"1K", KByte},
		{"1m", MByte},
		{"1M", MByte},
		{"1g", GByte},
		{"1G", GByte},
		{"1kb", KByte},
		{"1KB", KByte},
		{"1mb", MByte},
		{"1MB", MByte},
		{"1g", GByte},
		{"1Gb", GByte},
	}

	for _, v := range result {
		o, err := ParseSize(v.inStr)
		if o != v.outResult {
			t.Fatalf("input arg(%s), parseSize (%d), The expected value is (%d)\n",
				v.inStr, int(o), v.outResult)
		}
		if err != nil {
			t.Logf("input arg(%s), parseSize (%d), The expected value is (%d):err(%s)\n",
				v.inStr, int(o), v.outResult, err)
		}
	}
}
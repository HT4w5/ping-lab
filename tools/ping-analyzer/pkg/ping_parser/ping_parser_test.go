package ping_parser_test

import (
	"math"
	"testing"

	"github.com/HT4w5/ping-lab/tools/ping-analyzer/pkg/ping_parser"
)

const eps = 1e-9

func TestDeliveryRate(t *testing.T) {
	p1 := ping_parser.New("testdata/rate5in7.txt")
	if math.Abs(5./7.-p1.DeliveryRate()) > eps {
		t.Error("Incorrect delivery rate")
	}

	p2 := ping_parser.New("testdata/rate3in7.txt")
	if math.Abs(3./7.-p2.DeliveryRate()) > eps {
		t.Error("Incorrect delivery rate")
	}

	p3 := ping_parser.New("testdata/rate4in6.txt")
	if math.Abs(4./6.-p3.DeliveryRate()) > eps {
		t.Error("Incorrect delivery rate")
	}
}

func TestLongestConsecutive(t *testing.T) {
	p1 := ping_parser.New("testdata/consec5.txt")
	if p1.LongestConsecutive() != 5 {
		t.Error("Incorrect longest consecutive")
	}

	p2 := ping_parser.New("testdata/consec4.txt")
	if p2.LongestConsecutive() != 4 {
		t.Error("Incorrect longest consecutive")
	}

	p3 := ping_parser.New("testdata/consec4-2.txt")
	if p3.LongestConsecutive() != 4 {
		t.Error("Incorrect longest consecutive")
	}

	p4 := ping_parser.New("testdata/consec3.txt")
	if p4.LongestConsecutive() != 3 {
		t.Error("Incorrect longest consecutive")
	}

	p5 := ping_parser.New("testdata/consec2.txt")
	if p5.LongestConsecutive() != 2 {
		t.Error("Incorrect longest consecutive")
	}
}

func TestLongestLostBurst(t *testing.T) {
	p1 := ping_parser.New("testdata/lost0.txt")
	if p1.LongestLostBurst() != 0 {
		t.Error("Incorrect longest consecutive")
	}

	p2 := ping_parser.New("testdata/lost1.txt")
	if p2.LongestLostBurst() != 1 {
		t.Error("Incorrect longest consecutive")
	}

	p3 := ping_parser.New("testdata/lost2.txt")
	if p3.LongestLostBurst() != 2 {
		t.Error("Incorrect longest consecutive")
	}

	p4 := ping_parser.New("testdata/lost3.txt")
	if p4.LongestLostBurst() != 3 {
		t.Error("Incorrect longest consecutive")
	}

	p5 := ping_parser.New("testdata/lost3-2.txt")
	if p5.LongestLostBurst() != 3 {
		t.Error("Incorrect longest consecutive")
	}

	p6 := ping_parser.New("testdata/lost4.txt")
	if p6.LongestLostBurst() != 4 {
		t.Error("Incorrect longest consecutive")
	}
}

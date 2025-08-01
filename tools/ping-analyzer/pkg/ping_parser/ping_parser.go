package ping_parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type PingParser struct {
	File        string
	PingEntries []PingEntry
}

// Example: "[1754011851.617662] 64 bytes from 127.0.0.1: icmp_seq=169 ttl=64 time=0.044 ms"
// Only parse fields that we are concerned with.
type PingEntry struct {
	Lost      bool
	Timestamp time.Time
	IcmpSeq   int
	RTT       float64
}

func New(file string) *PingParser {
	p := &PingParser{
		File: file,
	}
	p.init()

	return p
}

func (p *PingParser) init() {
	file, err := os.Open(p.File)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // Skip first line.
	lastSeq := 0
	for scanner.Scan() {
		line := scanner.Text()
		entry, err := parseLine(line)
		if err != nil {
			fmt.Printf("Error parsing line '%s': %v\n", line, err)
			continue
		}
		// Fill in lost entries.
		if entry.IcmpSeq != lastSeq {
			for i := lastSeq + 1; i < entry.IcmpSeq; i++ {
				p.PingEntries = append(p.PingEntries, PingEntry{
					Lost:    true,
					IcmpSeq: i,
				})
			}
		}
		p.PingEntries = append(p.PingEntries, entry)
		lastSeq = entry.IcmpSeq
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}

func parseLine(line string) (PingEntry, error) {
	// Regular expression to capture all the fields
	re := regexp.MustCompile(`^\[(\d+\.\d+)\] (\d+) bytes from ([\d\.]+): icmp_seq=(\d+) ttl=(\d+) time=([\d.]+) (.*)$`)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 8 {
		return PingEntry{}, fmt.Errorf("line format mismatch")
	}

	// Parse timestamp
	timestampFloat, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return PingEntry{}, fmt.Errorf("invalid timestamp: %w", err)
	}
	sec := int64(timestampFloat)
	nsec := int64((timestampFloat - float64(sec)) * 1e9)
	timestamp := time.Unix(sec, nsec)

	// Parse icmp_seq
	icmpSeq, err := strconv.Atoi(matches[4])
	if err != nil {
		return PingEntry{}, fmt.Errorf("invalid icmp_seq: %w", err)
	}

	// Parse time
	pingTime, err := strconv.ParseFloat(matches[6], 64)
	if err != nil {
		return PingEntry{}, fmt.Errorf("invalid ping time: %w", err)
	}

	return PingEntry{
		Lost:      false,
		Timestamp: timestamp,
		IcmpSeq:   icmpSeq,
		RTT:       pingTime,
	}, nil
}

func (p *PingParser) DeliveryRate() float64 {
	lost := 0
	for _, e := range p.PingEntries {
		if e.Lost {
			lost++
		}
	}

	return float64(len(p.PingEntries)-lost) / float64(len(p.PingEntries))
}

func (p *PingParser) LongestConsecutive() int {
	longest := 0
	current := 0

	for _, e := range p.PingEntries {
		// Reset counter on lost.
		if e.Lost {
			longest = max(longest, current)
			current = 0
		} else {
			current++
		}
	}

	return longest
}

func (p *PingParser) LongestLostBurst() int {
	longest := 0
	current := 0

	for _, e := range p.PingEntries {
		//Reset counter on not lost.
		if !e.Lost {
			longest = max(longest, current)
			current = 0
		} else {
			current++
		}
	}

	return longest
}

func (p *PingParser) MinRTT() float64 {
	minRTT := 999.
	for _, e := range p.PingEntries {
		if !e.Lost {
			minRTT = min(minRTT, e.RTT)
		}

	}

	return minRTT
}

func (p *PingParser) MaxRTT() float64 {
	maxRTT := 0.
	for _, e := range p.PingEntries {
		if !e.Lost {
			maxRTT = max(maxRTT, e.RTT)
		}
	}

	return maxRTT
}

func (p *PingParser) AutocorrelationReplied(k int) float64 {
	if len(p.PingEntries) < k {
		return -1.
	}

	// totalCount: total number of replied requests #N.
	// repliedCount: number of replied to #(N+k) requests for replied request #N.
	totalCount := 0
	repliedCount := 0

	if k > 0 {
		for i := 0; i < len(p.PingEntries)-k; i++ {
			if p.PingEntries[i].Lost {
				continue
			}
			totalCount++
			if !p.PingEntries[i+k].Lost {
				repliedCount++
			}
		}
	} else if k < 0 {
		for i := -k; i < len(p.PingEntries); i++ {
			if p.PingEntries[i].Lost {
				continue
			}
			totalCount++
			if !p.PingEntries[i+k].Lost {
				repliedCount++
			}
		}
	} else {
		return 1.
	}

	return float64(repliedCount) / float64(totalCount)
}

func (p *PingParser) AutocorrelationLost(k int) float64 {
	if len(p.PingEntries) < k {
		return -1.
	}

	// totalCount: total number of lost requests #N.
	// repliedCount: number of lost to #(N+k) requests for lost request #N.
	totalCount := 0
	lostCount := 0

	if k > 0 {
		for i := 0; i < len(p.PingEntries)-k; i++ {
			if !p.PingEntries[i].Lost {
				continue
			}
			totalCount++
			if p.PingEntries[i+k].Lost {
				lostCount++
			}
		}
	} else if k < 0 {
		for i := -k; i < len(p.PingEntries); i++ {
			if !p.PingEntries[i].Lost {
				continue
			}
			totalCount++
			if p.PingEntries[i+k].Lost {
				lostCount++
			}
		}
	} else {
		return 1.
	}

	return float64(lostCount) / float64(totalCount)
}

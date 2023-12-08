package d05

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	//go:embed example.txt
	example []byte

	//go:embed input.txt
	input []byte
)

func ReadExample() []string {
	s := bufio.NewScanner(bytes.NewBuffer(example))
	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}
	return out
}

func ReadInput() []string {
	s := bufio.NewScanner(bytes.NewBuffer(input))
	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}
	return out
}

type Category string

const (
	CategorySEED        Category = "seed"
	CategorySOIL        Category = "soil"
	CategoryFERTILIZER  Category = "fertilizer"
	CategoryWATER       Category = "water"
	CategoryLIGHT       Category = "light"
	CategoryTEMPERATURE Category = "temperature"
	CategoryHUMIDITY    Category = "humidity"
	CategoryLOCATION    Category = "location"
)

type Almanac struct {
	SeedRanges []*SeedRange

	// Head map is the seed-to-soil map
	Head *Map

	// Tail map is the humidity-to-location map
	Tail *Map
}

func (a *Almanac) LookupDstID(src, dst Category, srcID int) int {
	m := a.Head
	n := srcID

	var started bool
	for {
		if m.Src == src {
			started = true
		}
		if started {
			n = m.DstID(n)
		}
		if m.Dst == dst {
			return n
		}
		m = m.Next
	}
}

func (a *Almanac) LookupSrcID(dst, src Category, dstID int) int {
	m := a.Tail
	n := dstID

	var started bool
	for {
		if m.Dst == dst {
			started = true
		}
		if started {
			n = m.SrcID(n)
		}
		if m.Src == src {
			return n
		}
		m = m.Prev
	}
}

func (a *Almanac) SeedAtLowestLocation() (seed int, location int) {
	for ; location < math.MaxInt; location++ {
		seed = a.LookupSrcID(CategoryLOCATION, CategorySEED, location)
		if a.ContainsSeed(seed) {
			return seed, location
		}
	}
	return -1, -1
}

func (a *Almanac) ContainsSeed(seed int) bool {
	for _, seedRange := range a.SeedRanges {
		if seedRange.ContainsSeed(seed) {
			return true
		}
	}
	return false
}

type SeedRange struct {
	// Num is the start number of the seed range
	Num int
	// Len is the length of the seed range
	Len int
}

func (r *SeedRange) ContainsSeed(seed int) bool {
	return r.Num <= seed && seed < r.Num+r.Len
}

type Map struct {
	Src, Dst Category
	Ranges   []*Range

	Prev *Map
	Next *Map
}

// DstID returns the destination ID for the given source ID.
func (m *Map) DstID(srcID int) int {
	for _, r := range m.Ranges {
		if id, ok := r.DstID(srcID); ok {
			return id
		}
	}
	return srcID
}

// SrcID returns the source ID for the given destination ID.
func (m *Map) SrcID(dstID int) int {
	for _, r := range m.Ranges {
		if id, ok := r.SrcID(dstID); ok {
			return id
		}
	}
	return dstID
}

func (m *Map) String() string {
	var out []string
	for _, r := range m.Ranges {
		out = append(out, fmt.Sprintf("{D:%d, S:%d, L:%d}", r.DstNum, r.SrcNum, r.Len))
	}
	return fmt.Sprintf("%s-to-%s: %v", m.Src, m.Dst, out)
}

type Range struct {
	// DstNum is the destination start number.
	DstNum int
	// SrcNum is the source start number.
	SrcNum int
	// Len is the range length.
	Len int
}

func (r *Range) DstID(srcID int) (int, bool) {
	if r.SrcNum <= srcID && srcID < r.SrcNum+r.Len {
		return r.DstNum + srcID - r.SrcNum, true
	}
	return srcID, false
}

func (r *Range) SrcID(dstID int) (int, bool) {
	if r.DstNum <= dstID && dstID < r.DstNum+r.Len {
		return r.SrcNum + dstID - r.DstNum, true
	}
	return dstID, false
}

func ParseAlmanac(data []string) (*Almanac, error) {
	var out Almanac

	var m *Map
	var ms []*Map
	for _, d := range data {
		if strings.HasPrefix(d, "seeds:") {
			seeds, err := ParseNumbers(d)
			if err != nil {
				return nil, err
			}
			if len(seeds)%2 != 0 {
				return nil, fmt.Errorf("invalid seeds '%v'", seeds)
			}
			for i := 0; i < len(seeds); i += 2 {
				out.SeedRanges = append(out.SeedRanges, &SeedRange{
					Num: seeds[i],
					Len: seeds[i+1],
				})
			}
			continue
		}

		if strings.Contains(d, "map:") {
			parts := strings.Split(strings.TrimSpace(strings.Replace(d, "map:", "", 1)), "-")
			if len(parts) != 3 || parts[1] != "to" {
				return nil, fmt.Errorf("invalid header '%s'", d)
			}
			m = &Map{
				Src: Category(parts[0]),
				Dst: Category(parts[2]),
			}
		} else if d == "" {
			if m != nil && m.Src != "" && m.Dst != "" {
				ms = append(ms, m)
			}
			m = &Map{}
		} else {
			n, err := ParseNumbers(d)
			if err != nil {
				return nil, err
			}
			if len(n) != 3 {
				return nil, fmt.Errorf("invalid map range '%v'", n)
			}
			m.Ranges = append(m.Ranges, &Range{
				DstNum: n[0],
				SrcNum: n[1],
				Len:    n[2],
			})
		}
	}
	if m != nil && m.Src != "" && m.Dst != "" {
		ms = append(ms, m)
	}

	getMap := func(src, dst Category) (*Map, error) {
		for _, m := range ms {
			if m.Src == src && m.Dst == dst {
				return m, nil
			}
		}
		return nil, fmt.Errorf("%s-to-%s map not found", src, dst)
	}

	curr, err := getMap(CategorySEED, CategorySOIL)
	if err != nil {
		return nil, err
	}
	out.Head = curr

	cat := CategorySOIL
	for _, next := range []Category{
		CategoryFERTILIZER,
		CategoryWATER,
		CategoryLIGHT,
		CategoryTEMPERATURE,
		CategoryHUMIDITY,
		CategoryLOCATION,
	} {
		m, err := getMap(cat, next)
		if err != nil {
			return nil, err
		}
		curr.Next = m
		m.Prev = curr
		curr = m
		cat = next
	}
	out.Tail = curr

	return &out, nil
}

func ParseNumbers(data string) ([]int, error) {
	var out []int
	for i := 0; i < len(data); i++ {
		var num []byte
		for i < len(data) && '0' <= data[i] && data[i] <= '9' {
			num = append(num, data[i])
			i++
		}
		if len(num) > 0 {
			n, err := strconv.Atoi(string(num))
			if err != nil {
				return nil, err
			}
			out = append(out, n)
		}
	}
	return out, nil
}

package d05

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Solve(t *testing.T) {
	data := ReadInput()
	almanac, err := ParseAlmanac(data)
	require.NoError(t, err)

	t.Run("part 1: seed locations", func(t *testing.T) {
		var minLocation int
		for i, seedRange := range almanac.SeedRanges {
			a := almanac.LookupDstID(CategorySEED, CategoryLOCATION, seedRange.Num)
			b := almanac.LookupDstID(CategorySEED, CategoryLOCATION, seedRange.Len)
			if i == 0 {
				minLocation = min(a, b)
			} else {
				minLocation = min(minLocation, a, b)
			}
		}
		assert.Equal(t, 324724204, minLocation)
	})

	t.Run("part 2: seed at lowest location", func(t *testing.T) {
		_, loc := almanac.SeedAtLowestLocation()
		assert.Equal(t, 104_070_862, loc)
	})
}

func Test_ParseAlmanac(t *testing.T) {
	t.Run("seed at lowest location", func(t *testing.T) {
		data := ReadExample1()
		almanac, err := ParseAlmanac(data)
		require.NoError(t, err)

		// override with part 1 seed format
		almanac.SeedRanges = []*SeedRange{
			{Num: 79, Len: 1},
			{Num: 14, Len: 1},
			{Num: 55, Len: 1},
			{Num: 13, Len: 1},
		}
		seed, loc := almanac.SeedAtLowestLocation()
		assert.Equal(t, 13, seed)
		assert.Equal(t, 35, loc)
	})

	data := ReadExample1()
	almanac, err := ParseAlmanac(data)
	require.NoError(t, err)

	t.Run("head is seed-to-soil map", func(t *testing.T) {
		assert.Equal(t, CategorySEED, almanac.Head.Src)
		assert.Equal(t, CategorySOIL, almanac.Head.Dst)
	})

	t.Run("tail is humidity-to-location map", func(t *testing.T) {
		assert.Equal(t, CategoryHUMIDITY, almanac.Tail.Src)
		assert.Equal(t, CategoryLOCATION, almanac.Tail.Dst)
	})

	t.Run("seed ranges", func(t *testing.T) {
		assert.Equal(t, []*SeedRange{
			{Num: 79, Len: 14},
			{Num: 55, Len: 13},
		}, almanac.SeedRanges)
	})

	t.Run("traverse nodes", func(t *testing.T) {
		t.Run("forward", func(t *testing.T) {
			curr := almanac.Head
			src := CategorySEED
			for _, dst := range []Category{
				CategorySOIL,
				CategoryFERTILIZER,
				CategoryWATER,
				CategoryLIGHT,
				CategoryTEMPERATURE,
				CategoryHUMIDITY,
				CategoryLOCATION,
			} {
				assert.Equal(t, src, curr.Src)
				assert.Equal(t, dst, curr.Dst)
				curr = curr.Next
				src = dst
			}
		})

		t.Run("reverse", func(t *testing.T) {
			curr := almanac.Head
			for {
				if curr.Next == nil {
					break
				}
				curr = curr.Next
			}
			dst := CategoryLOCATION
			for _, src := range []Category{
				CategoryHUMIDITY,
				CategoryTEMPERATURE,
				CategoryLIGHT,
				CategoryWATER,
				CategoryFERTILIZER,
				CategorySOIL,
				CategorySEED,
			} {
				assert.Equal(t, src, curr.Src)
				assert.Equal(t, dst, curr.Dst)
				curr = curr.Prev
				dst = src
			}
		})
	})

	t.Run("seed to location", func(t *testing.T) {
		assert.Equal(t, 82, almanac.LookupDstID(CategorySEED, CategoryLOCATION, 79))
		assert.Equal(t, 43, almanac.LookupDstID(CategorySEED, CategoryLOCATION, 14))
		assert.Equal(t, 86, almanac.LookupDstID(CategorySEED, CategoryLOCATION, 55))
		assert.Equal(t, 35, almanac.LookupDstID(CategorySEED, CategoryLOCATION, 13))
	})

	t.Run("location to seed", func(t *testing.T) {
		assert.Equal(t, 79, almanac.LookupSrcID(CategoryLOCATION, CategorySEED, 82))
		assert.Equal(t, 14, almanac.LookupSrcID(CategoryLOCATION, CategorySEED, 43))
		assert.Equal(t, 55, almanac.LookupSrcID(CategoryLOCATION, CategorySEED, 86))
		assert.Equal(t, 13, almanac.LookupSrcID(CategoryLOCATION, CategorySEED, 35))
	})

	t.Run("seed-to-soil", func(t *testing.T) {
		require.NoError(t, err)
		assert.Equal(t, []*Range{
			{SrcNum: 98, DstNum: 50, Len: 2},
			{SrcNum: 50, DstNum: 52, Len: 48},
		}, almanac.Head.Ranges)

		lookups := []struct {
			srcID int
			dstID int // expected
		}{
			{srcID: 79, dstID: 81},
			{srcID: 14, dstID: 14},
			{srcID: 55, dstID: 57},
			{srcID: 13, dstID: 13},
		}
		for _, tt := range lookups {
			t.Run(fmt.Sprintf("seed %d, soil %d", tt.srcID, tt.dstID), func(t *testing.T) {
				n := almanac.LookupDstID(CategorySEED, CategorySOIL, tt.srcID)
				assert.Equal(t, tt.dstID, n)
			})
		}
	})

	t.Run("contains all maps", func(t *testing.T) {
		expected := []string{
			"seed-to-soil",
			"soil-to-fertilizer",
			"fertilizer-to-water",
			"water-to-light",
			"light-to-temperature",
			"temperature-to-humidity",
			"humidity-to-location",
		}
		curr := almanac.Head
		var actual []string
		for {
			actual = append(actual, fmt.Sprintf("%s-to-%s", curr.Src, curr.Dst))
			curr = curr.Next
			if curr == nil {
				break
			}
		}
		assert.Len(t, actual, len(expected))
		assert.Equal(t, expected, actual)
	})
}

func Test_ParseNumbers(t *testing.T) {
	t.Run("seeds: 79 14 55 13", func(t *testing.T) {
		data := "seeds: 79 14 55 13"
		nums, err := ParseNumbers(data)
		require.NoError(t, err)
		assert.Equal(t, []int{79, 14, 55, 13}, nums)
	})

	t.Run("50 98 2", func(t *testing.T) {
		data := "50 98 2"
		nums, err := ParseNumbers(data)
		require.NoError(t, err)
		assert.Equal(t, []int{50, 98, 2}, nums)
	})

	t.Run("2650178765 568406716 612755541", func(t *testing.T) {
		data := "2650178765 568406716 612755541"
		nums, err := ParseNumbers(data)
		require.NoError(t, err)
		assert.Equal(t, []int{2650178765, 568406716, 612755541}, nums)
	})

	t.Run("50, 98; 2", func(t *testing.T) {
		data := "50, 98; 2"
		nums, err := ParseNumbers(data)
		require.NoError(t, err)
		assert.Equal(t, []int{50, 98, 2}, nums)
	})
}

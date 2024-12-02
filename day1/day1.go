package day1

import (
	"sort"

	"github.com/k0kubun/pp"
	"github.com/t0duf0du/AOC2024/cmd/utils"
)

func Day1() error {
	nums1, nums2, err := utils.ReadFromText("day1/data.txt")
	if err != nil {
		pp.Println(err.Error())
	}
	// pp.Println("Nums1")
	// pp.Println(nums1)

	// pp.Println("Nums2")
	// pp.Println(nums2)

	// for i := range nums1 {
	// 	pp.Println("a: ", nums1[i], "  b: ", nums2[i])
	// }

	sort.Ints(nums1)
	sort.Ints(nums2)

	distance, err := utils.Distance(nums1, nums2)

	pp.Println(distance)
	return err

}

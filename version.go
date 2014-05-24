package main

import (
	"fmt"
)

func versionsSubmajorGen(major, submajor, minor int) []string {
	var versions []string

	for i := 0; i <= minor; i++ {
		v := fmt.Sprintf("%d.%d.%d", major, submajor, i)
		versions = append(versions, v)
	}
	return versions
}

func versionsGen() []string {
	var versions []string

	versionsMinor0 := []int{45, 6, 61, 14, 38, 39, 69, 55, 7}
	versionsMinor1 := []int{15, 19, 9, 16, 7, 13, 0, 0}

	for i := 0; i < 9; i++ {
		versions = append(versions, versionsSubmajorGen(0, i+1, versionsMinor0[i])...)
	}

	for i := 0; i < 8; i++ {
		versions = append(versions, versionsSubmajorGen(1, i, versionsMinor1[i])...)
	}

	return versions
}

func showVersions() {
	versions := versionsGen()
	for _, v := range versions {
		fmt.Println(v)
	}
}
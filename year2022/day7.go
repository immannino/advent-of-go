package year2022

import (
	"advent-of-code/pkg/data"
	"fmt"
	"strconv"
	"strings"
)

const (
	PART_1_THRESHOLD = 100000
)

type File struct {
	name     string
	fileType string
	size     int
	children []*File
}

func ParentOf(name string, parent *File) *File {
	for _, c := range parent.children {
		if c.name == name {
			return parent
		}
	}

	for _, c := range parent.children {
		return ParentOf(name, c)
	}

	// For compiler
	return nil
}

func Find(name string, node *File) *File {
	// fmt.Printf("Looking for: %s, At: %s\n", name, node.name)

	if name == node.name {
		return node
	}

	for _, n := range node.children {
		// Only crawl dirs
		if n.fileType == "dir" {
			return Find(name, n)
		}
	}

	// For compiler
	return node
}

func Sum(current, threshold int, node *File) int {
	if node.fileType == "file" {
		if node.size < threshold {
			fmt.Println("PASSED: ", node)
			return node.size
		}

		return 0
	}

	if node.fileType == "dir" && len(node.children) == 0 {
		return 0
	}

	for _, c := range node.children {
		// fmt.Println("child loop ", c.name, c.fileType, c.size, len(node.children))
		current += Sum(current, threshold, c)
	}

	return current
}

func Sums(current, threshold int, node *File) int {
	if node.fileType == "file" {
		if node.size < threshold {
			fmt.Println("PASSED: ", node)
			return node.size
		}

		return 0
	}

	if node.fileType == "dir" && len(node.children) == 0 {
		return 0
	}

	for _, c := range node.children {
		// fmt.Println("child loop ", c.name, c.fileType, c.size, len(node.children))
		current += Sum(current, threshold, c)
	}

	return current
}

func Day7() []string {
	input := data.ReadAsString("data/2022/day7.txt")
	part1 := 0
	part2 := 0
	cmds := strings.Split(input, "\n")
	root := &File{
		name:     "root",
		fileType: "dir",
		children: []*File{},
	}

	// Build tree
	var curr *File

	for _, line := range cmds {
		if strings.Contains(line, "$ cd") {
			path := strings.Replace(line, "$ cd ", "", 1)

			if path == ".." {
				// traverse up
				// fmt.Println("Searching for Parent of: ", curr.name)
				curr = ParentOf(curr.name, root)
				// fmt.Println("Found Parent: ", curr.name)
			} else if path == "/" {
				// fmt.Println("Set root")
				curr = root
			} else {
				// Switch to previously declared node
				// fmt.Println("Finding node: ", path)
				curr = Find(path, root)
				// fmt.Println("Found node: ", curr.name)
			}

			continue
		}

		if strings.Contains(line, "$ ls") {
			// Not implemented
			continue
		}

		if strings.Contains(line, "dir ") {
			path := strings.Replace(line, "dir ", "", 1)
			// Append new dir to current children
			// fmt.Println(root, curr)
			curr.children = append(curr.children, &File{
				name:     path,
				fileType: "dir",
				children: []*File{},
			})
			continue
		}

		// Add child to tree
		parts := strings.Split(line, " ")
		size, _ := strconv.ParseInt(parts[0], 10, 32)
		curr.children = append(curr.children, &File{
			name:     parts[1],
			fileType: "file",
			size:     int(size),
		})
	}

	// Part1 Answer

	part1 = Sum(0, PART_1_THRESHOLD, root)

	return []string{"Day 7: No Space Left On Device", strconv.Itoa(part1), strconv.Itoa(part2)}
}

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
	parent   *File
	depth    int
}

func ParentOf(child *File, parent *File) *File {
	// fmt.Printf("Looking for parent of %s, have %s, type: %s\n", name, parent.name, parent.fileType)
	if parent.fileType == "file" {
		return nil
	}

	for _, c := range parent.children {
		if c.name == child.name && c.depth == child.depth {
			return parent
		}
	}

	for _, c := range parent.children {
		r := ParentOf(child, c)
		if r != nil {
			return r
		}
	}

	// For compiler
	return nil
}

func Find(name string, depth int, node *File) *File {
	// fmt.Printf("Looking for: %s, At: %s\n", name, node.name)

	if name == node.name && depth == node.depth {
		return node
	}

	for _, n := range node.children {
		// Only crawl dirs
		if n.fileType == "dir" {
			r := Find(name, depth, n)
			if r != nil {
				return r
			}
		}
	}

	// For compiler
	return nil
}

func Sum(sum uint64, node *File) uint64 {
	fmt.Println(sum, node.name)
	if node.fileType == "file" {
		return uint64(node.size)
	}

	if node.fileType == "dir" && len(node.children) == 0 {
		return 0
	}

	for _, c := range node.children {
		sum += Sum(sum, c)
	}

	return sum
}

func Walk(depth int, node *File) *File {
	// Pint stuf f
	return nil
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
		depth:    0,
		parent:   nil,
	}

	// Build tree
	var curr *File

	for _, line := range cmds {
		if curr != nil {
			fmt.Println("command: ", line, curr.name, len(curr.children))

			// time.Sleep(1 * time.Second)
			fmt.Println()
			if curr.name != "root" {
				PrintDir(ParentOf(curr, root))
			}
			fmt.Println()
		}
		if strings.Contains(line, "$ cd") {
			path := strings.Replace(line, "$ cd ", "", 1)

			if path == ".." {
				// traverse up
				if curr.name != "root" {
					// fmt.Println("Searching for Parent of: ", curr.name)
					// curr = ParentOf(curr, root)
					curr = curr.parent
					// fmt.Println("Found Parent: ", curr.name)
				}
				// if curr.name != "root" {
				// fmt.Println("Searching for Parent of: ", curr.name)
				// curr = ParentOf(curr.name, root)
				// fmt.Println("Found Parent: ", curr.name)
				// }
			} else if path == "/" {
				// fmt.Println("Set root")
				curr = root
			} else {
				// Switch to previously declared node
				curr = Find(path, curr.depth+1, root)
				if path != curr.name {
					panic(fmt.Sprintf("Found wrong node. Expected %s, got %s", path, curr.name))
				}
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
			curr.children = append(curr.children, &File{
				name:     path,
				fileType: "dir",
				children: []*File{},
				size:     0,
				parent:   curr,
				depth:    curr.depth + 1,
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
			depth:    curr.depth + 1,
			parent:   curr,
		})
	}

	// Part1 Answer
	queue := []*File{root}
	sums := []uint64{}
	names := []string{}

	c := 0
	for len(queue) > 0 {
		current := queue[0] // Fetch
		queue = queue[1:]   // Pop

		if current.fileType == "dir" {
			for _, c := range current.children {
				fmt.Printf("Child of %s - %s - %d\n", current.name, c.name, c.size)
			}
			s := Sum(0, current)
			fmt.Printf("Summing %v, val: %d \n", current.name, s)
			fmt.Println()
			sums = append(sums, s)
			names = append(names, current.name)
		}

		queue = append(queue, current.children...)
		// time.Sleep(1 * time.Second)
		c += 1
	}

	for i, v := range sums {
		if v <= PART_1_THRESHOLD {
			fmt.Println(names[i], v)
			part1 += int(v)
		}
	}

	return []string{"Day 7: No Space Left On Device", strconv.Itoa(part1), strconv.Itoa(part2)}
}

func PrintDir(f *File) *File {
	fmt.Printf("%s %s\n", strings.Repeat("\t", f.depth), f.name)

	if f.fileType == "file" {
		return nil
	}

	for _, c := range f.children {
		s := PrintDir(c)

		if s != nil {
			return s
		}
	}

	return nil
}

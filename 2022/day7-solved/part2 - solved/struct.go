package day

// directories = map[string]int

type file struct {
	name string
	size int
	dirs []string
}

type path struct {
	dirs []string
	full string
}

func (p path) cd(dir string) path {
	if dir == ".." {
		lastIndex := len(p.dirs) - 1
		p.dirs = p.dirs[0:lastIndex]
		p.full = p.dirs[lastIndex-1]
		return p
	}
	p.full = p.appendToFull(dir)
	p.dirs = append(p.dirs, p.full)
	return p
}

func (p path) appendToFull(dir string) string {
	if dir == "/" || p.full == "/" {
		return p.full + dir
	}
	return p.full + "/" + dir
}

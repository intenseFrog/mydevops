package common

import (
	"bufio"
	"os"
)

type Deployment struct {
	Nodes []*Node
	Myctl string

	Master *Node
}

func (d *Deployment) Create() error {
	for _, n := range d.Nodes {
		if err := n.Create(); err != nil {
			return err
		}
	}

	return nil
}

func (d *Deployment) Destroy() error {
	for _, n := range d.Nodes {
		if err := n.Destroy(); err != nil {
			return err
		}
	}

	return nil
}

// TODO: does YAML make more sense?
func (d *Deployment) Parse(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	d.Myctl = scanner.Text()
	d.Nodes = make([]*Node, 0)

	for scanner.Scan() {
		node := &Node{Deployment: d}
		if err := node.Parse(scanner.Text()); err != nil {
			return err
		}

		if node.Role == "master" {
			d.Master = node
		}

		d.Nodes = append(d.Nodes, node)
	}

	return nil
}
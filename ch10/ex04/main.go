package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"

	"encoding/json"

	"io"

	"fmt"

	"github.com/pkg/errors"
)

func main() {
	args := os.Args[1:]
	pkgs, err := execList(args...)
	if err != nil {
		log.Fatal(err)
	}
	deps := uniqPackages(pkgs...)
	pkgs, err = execList(deps...)
	if err != nil {
		log.Fatal(err)
	}
	deps = uniqPackages(pkgs...)
	for _, d := range deps {
		fmt.Println(d)
	}
}

type pkgInfo struct {
	Deps []string `json:"Deps"`
}

type AfterFunc func(p *pkgInfo) error

func uniqPackages(pkgs ...*pkgInfo) []string {
	var deps []string
	u := map[string]struct{}{}
	for _, p := range pkgs {
		deps = append(deps, p.Deps...)
	}
	for _, d := range deps {
		if _, ok := u[d]; !ok {
			u[d] = struct{}{}
		}
	}
	var keys []string
	for k := range u {
		keys = append(keys, k)
	}

	return keys
}

func execList(packages ...string) ([]*pkgInfo, error) {
	args := []string{"list", "-json"}
	args = append(args, packages...)
	cmd := exec.Command("go", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, errors.Wrap(err, "cmd output")
	}
	var pkgs []*pkgInfo
	decoder := json.NewDecoder(bytes.NewBuffer(out))

	for {
		pkg := &pkgInfo{}
		err := decoder.Decode(pkg)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			return pkgs, nil
		}
		pkgs = append(pkgs, pkg)
	}
	return nil, errors.New("unreachable")
}

package project

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/devexps/go-micro/cmd/micro/v2/internal/base"
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"regexp"
)

var repoAddIgnores = []string{
	".git", ".github",
}

func (p *Project) Add(ctx context.Context, dir string, layout string, branch string, mod string) error {
	to := filepath.Join(dir, p.Name)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the service.",
		}
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}
		if !override {
			return err
		}
		os.RemoveAll(to)
	}
	toProto := filepath.Join(dir, "api", "proto", p.Name)
	if _, err := os.Stat(toProto); !os.IsNotExist(err) {
		fmt.Printf("🚫 %s proto already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
			Help:    "Delete the existing folder and create new api proto.",
		}
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}
		if !override {
			return err
		}
		os.RemoveAll(toProto)
	}
	fmt.Printf("🚀 Add proto %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)

	repo := base.NewRepo(layout, branch)

	repoMod, err := base.ModulePath(filepath.Join(repo.Path(), "go.mod"))
	if err != nil {
		return err
	}
	from := filepath.Join(repo.Path(), serviceDefaultPath)
	if err := repo.CopyToV2(ctx, from, to, filepath.Join(mod, p.Path), repoAddIgnores, []string{
		filepath.Join(p.Path, "api"), "api",
		filepath.Join(repoMod, "api"), filepath.Join(mod, "api"),
		filepath.Join(repoMod, "pkg"), filepath.Join(mod, "pkg"),
		serviceDefaultPath, p.Name,
	}); err != nil {
		return err
	}

	protoPkgName := regexp.MustCompile(`[^a-zA-Z0-9._]+`).ReplaceAllString(p.Name, "")
	fromProto := filepath.Join(repo.Path(), "api", "proto", serviceDefaultPath)
	if err := repo.CopyToApiProto(ctx, fromProto, toProto, repoAddIgnores, []string{
		protoGoPackage + serviceDefaultPath, protoGoPackage + protoPkgName,
		protoJavaPackage + serviceDefaultPath, protoJavaPackage + protoPkgName,
		serviceDefaultPath, p.Name,
		filepath.Join(repoMod, "api"), filepath.Join(mod, "api"),
	}); err != nil {
		return err
	}

	base.Tree(to, dir)

	fmt.Printf("\n🍺 Service creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("💻 Use the following command to start a service 👇:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println(color.WhiteString("$ make all"))
	fmt.Println(color.WhiteString("$ make run\n"))
	fmt.Println("			🤝 Thanks for using Go-Micro")
	fmt.Println("	📚 Tutorial: https://devexps.com/go-micro/getting-started/start")
	return nil
}

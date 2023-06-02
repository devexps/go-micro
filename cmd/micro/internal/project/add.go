package project

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/devexps/go-micro/cmd/micro/v2/internal/base"
	"github.com/fatih/color"
	"os"
	"path/filepath"
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

	fmt.Printf("🚀 Add service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)

	repo := base.NewRepo(layout, branch)

	if err := repo.CopyToV2(ctx, to, serviceDefaultPath, filepath.Join(mod, p.Path), repoAddIgnores, []string{filepath.Join(p.Path, "api"), "api"}); err != nil {
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

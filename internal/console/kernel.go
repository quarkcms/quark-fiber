package console

import (
	"fmt"
	"io/fs"
	"reflect"
	"time"

	"github.com/fatih/color"
	"github.com/quarkcms/quark-fiber/internal/console/commands"
)

type Kernel struct{}

// 注册服务
var Commands = []interface{}{
	(&commands.Install{}).Init(),
	(&commands.Migrate{}).Init(),
	(&commands.DBSeed{}).Init(),
	(&commands.StorageLink{}).Init(),
}

// 执行命令
func (p *Kernel) Run(assets fs.FS) {

	// 启动时执行的命令
	p.command(assets)

	// 保持执行的命令
	p.keepRunningCommand(assets)
}

// 启动时执行的命令
func (p *Kernel) command(assets fs.FS) {

	for _, v := range Commands {

		// 命令标识
		signature := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Signature").String()

		// 命令描述
		description := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Description").String()

		if signature == "" {
			v.(interface{ Handle() }).Handle()
			if description != "" {
				color.Set(color.FgGreen)
				fmt.Println(description)
				color.Unset()
			}
		}
	}
}

// 保持执行的命令
func (p *Kernel) keepRunningCommand(assets fs.FS) {

	var (
		command         string
		commandExecuted bool
	)

	// 监听输入
	for {
		time.Sleep(time.Duration(500) * time.Millisecond)

		fmt.Scanln(&command)
		for _, v := range Commands {

			// 命令标识
			signature := reflect.
				ValueOf(v).
				Elem().
				FieldByName("Signature").String()

			// 命令描述
			description := reflect.
				ValueOf(v).
				Elem().
				FieldByName("Description").String()

			if signature == command && signature != "" {
				v.(interface{ Handle() }).Handle()
				commandExecuted = true
				if description != "" {
					color.Set(color.FgGreen)
					fmt.Println(description)
					color.Unset()
				}
			}
		}

		if commandExecuted != true && command != "" {
			color.Set(color.FgRed)
			fmt.Println("Error: The '" + command + "' command doesn't exist!")
		}

		// 重置颜色
		color.Unset()
		// 重置状态
		commandExecuted = false
		// 重置命令
		command = ""
	}
}

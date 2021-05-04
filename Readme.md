# BDIX Tester

BDIX Tester is a small CLI application for testing connectivity BDIX connected servers. It's primary purpose is to check whether a BDIX connected server is accessible from your current ISP or not.

<span style="color: red">**Warning:**</span> This tool doesn't ensure that you will be able to download content from the servers because some server owners use hyperlink to private IP so that only their user can actually download contents.

This is an experimental tool. I would appriciate constructive feedbacks.

## Download

You can find the binaries [here](https://github.com/sadmansakib/bdix-server-checker/releases/latest)

## Usage

You can run the tool as it is but you can also add optional parameters:

```sh
Usage:
  bdix-tester [flags]

Flags:
  -h, --help            help for bdix-tester
  -t, --timeout uint8   Define network timeout (in seconds) if network response is unresponsive (default 30)
  -v, --version         version for bdix-tester
  -w, --workers uint8   Define number of threads you want to assign to the program (default 2)
```

## Contribution

You can create an issue with server links you want bdix-tester to check. PR with Code optimizations are welcomed :grin:
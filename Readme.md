# BDIX Tester

BDIX Tester is a small CLI application for testing connectivity BDIX connected servers. It's primary purpose is to check whether a BDIX connected server is accessible from your current ISP or not.

<span style="color: red">**Warning:**</span> This tool doesn't ensure that you will be able to download content from the servers because some server owners use hyperlink to private IP so that only their user can actually download contents.

This is an experimental tool. I would appriciate constructive feedbacks.

## Download

You can find the binaries [here](https://github.com/sadmansakib/bdix-server-checker/releases/latest)

## Usage

### **Linux:**

Download `bdix-tester-linux-amd64.zip` and extract the zip. Open terminal inside the extracted folder and run: 

```sh
  chmod +x bdix-tester-linux-amd64 && ./bdix-tester-linux-amd64
```

### **MAC OSX:**
Download `bdix-tester-darwin-amd64.zip` or `bdix-tester-darwin-arm64.zip` **(for M1 users)** and extract the zip. Open terminal inside the extracted folder and run: 

```sh
  chmod +x bdix-tester-linux-amd64 && ./bdix-tester-linux-amd64
```
#### **For M1 users:**

```sh
  chmod +x bdix-tester-linux-arm64 && ./bdix-tester-linux-arm64
```

### **Windows:**
Download `bdix-tester-windows-amd64.zip` and extract the zip. Move the extracted `bdix-tester-windows-amd64.exe` to Desktop. Then open the cmd then run:

```sh
  cd Desktop && bdix-tester-windows-amd64.exe
```

## Additional Configuration (Optional):
You can pass additional configuration flags to the CLI. Currently supported flags are:

```sh
-t, --timeout uint8   Define network timeout (in seconds) (default 30)
-v, --version         version for bdix-tester
-w, --workers uint8   Define number of threads you want to assign to the program (default 2)
```

## Contribution

You can create an issue with server links you want bdix-tester to check. PR with Code optimizations are welcomed :grin:
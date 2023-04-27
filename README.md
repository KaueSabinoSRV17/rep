# rep

A CLI tool that replaces strings

## Usage

The tool is suposed to be used like this:

```bash
rep i-want-underscore-not-hyphens - _ # Prints out i_want_underscore_not_hyphens
```

## Instalation (Linux, Mac and WSL only)

As of now, you can run a git clone on project:

```bash
git clone https://github.com/KaueSabinoSRV17/rep
```
 
Build the Go Project (Must Have Go installed or using in Docker):

```bash
go build . # At the root of the project
```
  
And moving it to somewhere on your path:

```bash
sudo mv ./rep /usr/bin
```

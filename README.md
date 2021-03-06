consulctl
=========

Easily get and manage key/value pairs in a Consul cluster.

## Examples

```bash
$ consulctl put foo bar
$ consulctl put foo/bar baz
$ consulctl put foo/bar/baz qux

$ consulctl get foo
# => bar

$ consulctl get foo/bar
# => baz

$ consulctl delete foo
$ consulctl delete -r foo/bar

$ consulctl get foo/bar/baz
# =>
```

## Help

```
NAME:
   consulctl - A simple command line client for Consul k/v store.

USAGE:
   consulctl [global options] command [command options] [arguments...]

VERSION:
   0.3.0

AUTHOR:
  Jean Mertz - <jean@mertz.fm>

COMMANDS:
   get      retrieve the value of a key
   put      set the value of a key
   delete   delete a key
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --address, -A "127.0.0.1:8500"   the remote endpoint for the Consul cluster [$CONSUL_KV_ADDRESS]
   --wait, -w "0"                   request timeout value in seconds [$CONSUL_KV_WAIT]
   --help, -h                       show help
   --version, -v                    print the version
```

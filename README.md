# kyou

`kyou` is a CLI app to create your today's working directory.

```shell
# on 2021-11-26
❯ kyou foo # to create 'foo' in today's directory
/home/user/kyou/2021/11/26/foo

❯ kyou where  # to print today's directory 
/home/user/kyou/2021/11/26

❯ cd `kyou where` # to move to today's directory

❯ cd `kyou bar` # to create `bar` and move to it
```

## Installation

### Using [Nix](https://nixos.org)

#### Nix Flakes (Nix >= 2.4)

```
❯ nix profile install github:aumyf/kyou
```

You can try `kyou` to run:

```
❯ nix shell github:aumyf/kyou
```

#### Nix Channels

```
❯ nix-channel --add https://github.com/aumyf/kyou/archive/master.tar.gz kyou

❯ nix-channel --update

❯ nix-env -iA kyou
```

### Using Go

```
❯ go install github.com/aumyf/kyou@latest
```

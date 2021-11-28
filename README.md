# kyou

`kyou` is a CLI app to create your today's working directory.

```shell
# on 2021-11-26
❯ kyou recoil # to create 'recoil' in today's directory
creating /home/user/kyou/2021/11/26/recoil

❯ kyou where  # to print today's directory 
/home/user/kyou/2021/11/26

❯ cd `kyou where` # to move to today's directory
```

## Installation

### Using [Nix](https://nixos.org)

```
❯ nix profile install github:aumyf/kyou
```

You can try `kyou` to run:

```
❯ nix shell github:aumyf/kyou
```

### Using Go

```
❯ go install github.com/aumyf/kyou
```

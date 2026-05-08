# Terraform Provider SaltStack

[![Tests](https://github.com/imperva/terraform-provider-saltstack/actions/workflows/ci.yml/badge.svg)](https://github.com/imperva/terraform-provider-saltstack/actions/workflows/ci.yml)

The Terraform SaltStack provider is a plugin for Terraform that allows to create and accept SaltStack minion keys.   
This provider is maintained by Imperva Operations Engineering Team.

See: [Official documentation](https://registry.terraform.io/providers/imperva/saltstack/latest/docs) in the Terraform registry.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19

## Supported Salt Master versions
The provider is tested with versions `3006.0`, `3006.1`, `3006.10`, `3006.11`, `3006.12`, `3006.13`, `3006.14`, `3006.15`, `3006.16`, `3006.17`, `3006.18`, `3006.19`, `3006.2`, `3006.20`, `3006.21`, `3006.22`, `3006.23`, `3006.24`, `3006.3`, `3006.4`, `3006.5`, `3006.6`, `3006.7`, `3006.8`, `3006.9`, `3007.0`, `3007.1`, `3007.10`, `3007.11`, `3007.12`, `3007.13`, `3007.14`, `3007.2`, `3007.3`, `3007.4`, `3007.5`, `3007.6`, `3007.7`, `3007.8`, `3007.9`, `3005`, `3004.2`, `3004.1`, `3004`. Versions 3006.0 and newer are bootstrapped from Broadcom onedir artifacts. Older versions use the legacy `saltstack/salt` container image.

## Building The Provider

1. Clone the repository
1. Change to the repository directory
1. Build the provider using the `make install` command:

```sh
$ make install
```
## Using the provider

Here is a short example on how to use this provider:

```hcl
terraform {
  required_providers {
    saltstack = {
      version = "0.1.0"
      source  = "imperva/saltstack",
    }
  }
}

provider saltstack {
    host            = "10.20.30.40"
    port            =  8000
    username        =  "saltstack-api-user"
    password        =  "strongpassword"
}

resource saltstack_minion_key_pair single_minion_key {
    minion_id = "db-1.domain.com"
    key_size = 2048
}

resource saltstack_minion_key_pair few_minion_keys {
    count = 5
    minion_id = "web-${count.index+1}.domain.com"
    key_size = 2048
}
```
  
## Developing the Provider

If you wish to work on the provider, you need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `make install`. This will build the provider and put the provider binary in the `~/.terraform.d/plugins/registry.terraform.io/imperva/saltstack/<VERSION>/<OS_ARCH>` directory.

To generate or update documentation, run `go generate`.

In order to run the suite of unit tests, run `make test`.

In order to run the full suite of acceptance tests, run `make testacc`.

_Note:_ Acceptance tests create a docker compose stack on port 8000.

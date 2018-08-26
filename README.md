[![Build Status](https://travis-ci.org/hanks/terraform-variables-generator.svg?branch=master)](https://travis-ci.org/hanks/terraform-variables-generator)

# terraform-variables-generator

Simple Tool to Generate Variables file from Terraform Configuration. It will find all *.tf files in current directory, and generate variables.tf file. If you already have this file, it will ask to override it.

I want to create a tool to generate `variables.tf` from `*.tf`, and found this awesome project [terraform-variables-generator](https://github.com/alexandrst88/terraform-variables-generator).
So based on it, I add some enhancements that seems can not be merged back :joy:

## Prerequisite

* Docker, for development
* Linux/macOS, not test in Windows now

## Enhancements

1. Refactor dramatically for basic maintenance friendly
2. Use dep to do dependencies management
3. Add customized var settings, that can help to create more flexible variables.tf file
4. Full support variable block element, like 'type', 'description', 'default'
5. Use 'terraform fmt' api to do the final format
6. Add Makefile to do task management
7. Introduce docker environment for the development
8. Add more tests to do the better coverage
9. Update .travis.yml with docker setup

## Installation

```bash
make install
```

and uninstall by:

```bash
make uninstall
```

## Usage

```bash
tfvargen
```

It will find all *.tf files in current directory, and generate variables.tf file. If you already have this file, it will ask to override it.

and you can find help info by:

```bash
tfvargen -h
```

Also you can customized variable block with `vars.ylm` to be put into the `*.tf` directory, like

```text
vars:
  - public_subnets:
      type: list
      description: subnets for public
      default: |
        ["sub1", "sub2"]
  - tags:
      type: map
      default: |
        {
          Name = "Terraform"
        }
  - cidrs:
      type: list
      default: |
        ["10.0.0.0/16", "10.1.0.0/16"]
  - t1-var2:
      description: var for t1
      type: string
```

### Demo

![demo.gif](./docs/images/demo.gif)

### Example

```text
resource "aws_vpc" "vpc" {
  cidr_block           = "${var.cidr}"
  enable_dns_hostnames = "${var.enable_dns_hostnames}"
  enable_dns_support   = "${var.enable_dns_support}"

  tags {
    Name = "${var.name}"
  }
}

resource "aws_internet_gateway" "vpc" {
  vpc_id = "${aws_vpc.vpc.id}"

  tags {
    Name = "${var.name}-igw"
  }
}
```

 Will generate

 ```text
 variable "ami" {
   description  = ""
}

variable "instance_type" {
   description  = ""
}

variable "cidr" {
   description  = ""
}

variable "enable_dns_hostnames" {
   description  = ""
}

variable "enable_dns_support" {
   description  = ""
}

variable "name" {
   description  = ""
}
 ```

## Development

* `make dev`, build docker image used in dev, should be run before other commands
* `make test`, run unit test, coverage test, static analytics
* `make build`, cross compile binaries, and put into `dist/bin` directory
* `make debug`, use `dlv` to do the `gdb-style` debug

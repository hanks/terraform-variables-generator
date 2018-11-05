[![Build Status](https://travis-ci.org/hanks/terraform-variables-generator.svg?branch=master)](https://travis-ci.org/hanks/terraform-variables-generator) [![Coverage Status](https://coveralls.io/repos/github/hanks/terraform-variables-generator/badge.svg?branch=HEAD)](https://coveralls.io/github/hanks/terraform-variables-generator?branch=master)

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

Also you can customized variable block with `vars.yml` to be put into the `*.tf` directory, like

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
resource "aws_eip" "nat" {
  vpc   = true
  count = "${length(var.public_subnets)}"
}

resource "aws_nat_gateway" "nat" {
  allocation_id = "${element(aws_eip.nat.*.id, count.index)}"
  subnet_id     = "${element(aws_subnet.public.*.id, count.index)}"
  count         = "${length(var.public_subnets)}"
  tags          = "${var.tags}"
}

data "template_file" "template1" {
  template = "${file("${path.module}/template1.tpl")}"

  vars {
    t1_var1 = "${var.cidrs}"
    t1-var2 = "${var.t1-var2}"
    t1-var3 = "${var.t1-Var3}-${var.t1-inline}"
  }
}
```

Will generate:

```text
variable "cidrs" {}

variable "public_subnets" {}

variable "t1-Var3" {}

variable "t1-inline" {}

variable "t1-var2" {}

variable "tags" {}
```

And, if you add customized conf `vars.yml`:

```yaml
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

then run generator again, the result will be:

```text
variable "cidrs" {
  type = "list"

  default = ["10.0.0.0/16", "10.1.0.0/16"]
}

variable "public_subnets" {
  description = "subnets for public"
  type        = "list"

  default = ["sub1", "sub2"]
}

variable "t1-Var3" {}

variable "t1-inline" {}

variable "t1-var2" {
  description = "var for t1"
  type        = "string"
}

variable "tags" {
  type = "map"

  default = {
    Name = "Terraform"
  }
}
```

## Development

* `make test`, run unit test, coverage test, static analytics
* `make run`, run program to generate `variables.tf` in `tests` directory
* `make build`, cross compile binaries, and put into `dist/bin` directory
* `make debug`, use `dlv` to do the `gdb-style` debug
* `make dev`, build docker image used in dev

# Fake provider

This almost empty provider has been created to test a Terraform bug.


## Build and try

```bash
go build
terraform init
terrafom plan
```



## Empty string bug?

It seems that default value (or `Default` or `DefaultFunc`) is applied for string fields (at least)
when we provide an empty string but only when there's no existing state.

With the following `main.tf` file:

```hcl
resource "fake" "test" {
  value = ""
}
```

Starting from scratch (with no state):

```bash
$ terraform plan
Terraform will perform the following actions:

  # fake.test will be created
  + resource "fake" "test" {
      + id = (known after apply)
    }

Plan: 1 to add, 0 to change, 0 to destroy.

$ terraform apply -auto-approve
[...]

$ terraform state show fake.test
# fake.test:
resource "fake" "test" {
    id    = "toto"
    value = "plop"
}

```

Updating `main.tf` with:
```hcl
resource "fake" "test" {
  value = "tmp"
}
```

Then:

```bash
$ terraform plan
Terraform will perform the following actions:

  # fake.test will be updated in-place
  ~ resource "fake" "test" {
        id    = "toto"
      ~ value = "plop" -> "tmp"
    }

Plan: 0 to add, 1 to change, 0 to destroy.


$ terraform apply -auto-approve
[...]

$ terraform state show fake.test
# fake.test:
resource "fake" "test" {
    id    = "toto"
    value = "tmp"
}
```

Reverting `main.tf` to:

```hcl
resource "fake" "test" {
  value = ""
}
```

Then:

```bash
$ terraform plan
Terraform will perform the following actions:

  # fake.test will be updated in-place
  ~ resource "fake" "test" {
        id    = "toto"
      - value = "tmp" -> null
    }

Plan: 0 to add, 1 to change, 0 to destroy.


$ terraform apply -auto-approve
[...]

$ terraform state show fake.test
# fake.test:
resource "fake" "test" {
    id    = "toto"
}
```


### With null value it seems to work as expected

```hcl
resource "fake" "test" {
  value = null
}
```

```bash
$ terraform plan
Terraform will perform the following actions:

  # fake.test will be created
  + resource "fake" "test" {
      + id    = (known after apply)
      + value = "plop"
    }

Plan: 1 to add, 0 to change, 0 to destroy.

$ terraform apply -auto-approve
[...]

$ terraform state show fake.test
# fake.test:
resource "fake" "test" {
    id    = "toto"
    value = "plop"
}
```

```hcl
resource "fake" "test" {
  value = "tmp"
}
```

```bash
$ terraform plan
Terraform will perform the following actions:

  # fake.test will be updated in-place
  ~ resource "fake" "test" {
        id    = "toto"
      ~ value = "plop" -> "tmp"
    }

Plan: 0 to add, 1 to change, 0 to destroy.

$ terraform apply -auto-approve
[...]

$ terraform state show fake.test
# fake.test:
resource "fake" "test" {
    id    = "toto"
    value = "tmp"
}
```

```hcl
resource "fake" "test" {
  value = null
}
```

```bash
$ terraform plan
Terraform will perform the following actions:

  # fake.test will be updated in-place
  ~ resource "fake" "test" {
        id    = "toto"
      ~ value = "tmp" -> "plop"
    }

Plan: 0 to add, 1 to change, 0 to destroy.

$ terraform apply -auto-approve
[...]

$ terraform state show fake.test
# fake.test:
resource "fake" "test" {
    id    = "toto"
    value = "plop"
}
```


### In Terraform 0.11, default values are not applied for empty string

```
resource "fake" "test" {
  value = ""
}
```

```
$ terraform011 apply -auto-approve
[...]
$ terraform011 state show fake.test
id    = toto
```

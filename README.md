<p align="center">
    <img src="https://raw.githubusercontent.com/mishamyrt/checode/master/assets/logo@2x.png" width="538px">
<p><br>

[![Go Report Card](https://goreportcard.com/badge/github.com/mishamyrt/checode)][goreportcard]

Checode extracts, collects and reports TODOs, FIXMEs and other keywords in your code. If you don't know why you should write a comment, look at the [wiki note](https://github.com/mishamyrt/checode/wiki/Maybe-you-don't-really-want-TODO).

## Features

* Language agnostic
* Multithreaded
* CI ready

## Usage in command line

```sh
checode src/
```

## Usage in GitLab CI

```yaml
checode:
  stage: quality assurance
  image: mishamyrt/checode
  script:
    checode src/
```

## Configuring

When running, Checode checks if the `.checode.yaml` file is in current directory and applies it if it is. Default built-in config looks like this:

```yaml
keywords:
  TODO: warn
  FIXME: err
  STOPSHIP: err
  NOTE: info
```

Using the configuration file, you can add processing of any keywords.

```yaml
keywords:
  XXX: warn
  HACK: warn
```

To apply a configuration file with a different name, specify it as the parameter.

```sh
checode -c custom_config.yaml src/
```

## Commands

In addition to normal comments, you can make comments with logic. The commands are triggered as follows:

```ts
// <KEYWORD>:: <command> <argument>: <message> 
```

For example:
```ts
// FIXME:: due 2021-10-20: Remove this dirty hack after the demonstration.
```

### Available commands

* `due` — Limits the date until which this comment will not be considered an error. Raises the comment level to `err` if the specified date is greater than or equal to the current date. Accepts the date in `yyyy-mm-dd` format.

[goreportcard]: https://goreportcard.com/report/github.com/mishamyrt/checode

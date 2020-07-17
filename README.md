<p align="center">
    <img alt="Logo" src="https://raw.githubusercontent.com/mishamyrt/checode/master/assets/logo@2x.png" width="538px">
<p>
<br>
<p align="center">
    <a href="https://goreportcard.com/report/github.com/mishamyrt/checode">
        <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/mishamyrt/checode" />
    </a>
    <a href="https://github.com/mishamyrt/checode/blob/master/LICENSE">
        <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-blue.svg" />
    </a>
    <a href="https://github.com/mishamyrt/checode/actions?query=workflow%3Abuild">
        <img alt="Build status" src="https://github.com/mishamyrt/checode/workflows/build/badge.svg?branch=master" />
    </a>
</p>

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

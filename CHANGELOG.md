## Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2022-08-19
### Added
- `bh prem -add [ARG...]` to allow a user to add a new premises.
- `bh pres -list` to list all premises.
- Added the `filesys` to access the file system. 

## [0.0.0] - 2022-08-18
### Added
- Initial structure.
- Defined types `CommandSet`, `Commands` and `Command` to create modular
commands confined to files, within the `cmds` package.
- Defined all flags in `cmds/flags.go`.


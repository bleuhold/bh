## Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
## [0.11.0] - 2022-09-18
### Added 
- `bh transaction add` command to be able to add a new transaction.
- The add command automatically adds an item with matching description and
values.
- `bh transaction remove` command to be able to remove a transaction.
- The remove command, also remove all associated/nested items.

## [0.10.0] - 2022-09-12
### Added 
- `bh statement` command for the generating of a statement.
- `bh statement -uuid=<>` command to generate a statement for the given contract
UUID by default prints a statement preview to the console. Purposely for easy
editing.
- `bh statement -uuid=<> -pdf` only when the `-pdf` flag is passed with the
actual PDF be generated and save to the user's `Downloads` directory.

## [0.9.0] - 2022-09-12
### Added
- `bh contract` command to enable to user to manage their contracts.
- `bh contract -list` to view all the contracts.
- `bh contract -uuid` to get a detailed view of a specific contract.

## [0.8.0] - 2022-09-11
### Added
- `bh tenant` command enable to system to manage tenants.
- `bh tenant -list` to be able to list all tenants.

### Changed
- Added sorting to transactions just before saving.
- Added sorting to items just before saving.

## [0.7.4] - 2022-09-10
### Changed
- Re-introduced the `bh premises` command.
- As well as the `bh premises -list` to list all available premises.

## [0.7.3] - 2022-09-10
### Changed
- The account to also include an account holder name.
- The account to also include an alias for the account.

## [0.7.2] - 2022-09-10
### Changed
- Updated items to now have the same date as the transaction, to simplify the 
filtering of items.

## [0.7.1] - 2022-09-10
### Fixed
- The `bh info set` command would unset parameters if they were not specified
  or if their flags contained the zero value.

## [0.7.0] - 2022-09-09
### Added
- `bh item add` to create and add a new item to a transaction.
- `bh item remove` to permanently remove an item.

## [0.6.0] - 2022-08-29
### Added
- `bh item tag -add` to add tags to an item.
- `bh item tag -remove` to remove tags from an item.

## [0.5.0] - 2022-08-28
### Added 
- `bh item` command to manage the basic implementation of items.

## [0.4.1] - 2022-08-28
### Updated
- `bh upload` command now requires a valid account UUID with a `-uuid` flag
to be able to associate the transactions with an account.

## [0.4.0] - 2022-08-28
### Added
- `bh account` command to be able to manage all wallet/bank accounts.
- `bh account -list` to list all accounts.
- `bh account -help` to show account help description.
- `bh account add -number=... -provider-name=... -type=...` subcommand to be 
able to add a new account.
- `bh account remove -number=...` subcommand to be able to remove an account.

## [0.3.0] - 2022-08-27 
### Added
- `bh transactions` command to do transaction related tasks.
- `bh transactions -list` to view all uploaded transactions.

### Fixed
- Updated the duplicate transaction check since, currently a random bank account
UUID is added to new transactions.
- The switch with `bh info set` command did a wrong error check with is fixed.

## [0.2.0] - 2022-08-27
### Added
- `bh upload -f [PATH TO FILE]` to allow a user to upload a bank statement and 
marshal the transactions' data into the transactions of the command line tool.
- `bh upload -file [PATH TO FILE]` is the same as above.
- `bh upload --help` help flag to display the upload help information.
- Added `*Transactions.Save` method to save the transactions to the file system.
- Added `LoadTransactions` to load the transactions' data from the file system.

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


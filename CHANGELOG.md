# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic
Versioning](http://semver.org/spec/v2.0.0.html) except to the first release.

## [Unreleased]

### Added

- Constants from Tarantool 3.9.0:
  * `IPROTO_THREAD_ID`, `IPROTO_AUDIT_CONTEXT`, `IPROTO_BEGIN_KEY`,
    `IPROTO_END_KEY` keys.
  * `IPROTO_DELETE_RANGE` type.
  * `ER_ALIEN_ENGINE`, `ER_MVCC_UNAVAILABLE`,
    `ER_CANT_UPGRADE_INDEXED_FIELD`, `ER_SYNC_TIMEOUT`,
    `ER_FIELD_IRREPRESENTABLE_VALUE`,
    `ER_KEY_PART_IRREPRESENTABLE_VALUE`,
    `ER_INVALID_SORTDATA_FILE`, `ER_NO_SUCH_THREAD`,
    `ER_UNABLE_TO_PROCESS_IN_THREAD`, `ER_DICT_OVERFLOW`,
    `ER_NO_SUCH_THREAD_GROUP`, `ER_XLOG_NOT_FOUND`,
    `ER_RECOVERY_POINT_TXN_LAST_ROW`, `ER_NO_SUCH_READ_VIEW`,
    `ER_READ_VIEW_THROTTLED` error codes.

### Changed

- `ER_CROSS_ENGINE_TRANSACTION`, `ER_SQL_SYNTAX_WITH_POS` and
  `ER_SQL_PARSER_GENERIC_WITH_POS` are unused in Tarantool 3.9.0, the
  constants are kept for compatibility.

### Fixed

## [v1.1.0] - 2024-12-04

The release updates IPROTO constants for Tarantool 3.3.0.

### Added

- Constants from Tarantool 3.3.0.

## [v1.0.0] - 2023-12-27

The first stable release exports IPROTO constants from Tarantool 3.0.0.

### Added

- Constants from Tarantool 3.0.0 (#4).

### Changed

- `ER_UNUSED` with error code 248 renamed to `ER_DEFAULT_VALUE_TYPE` (#4).

## [v0.1.0] - 2023-06-01

The initial release adds IPROTO constants exported from Tarantool 2.11.0.

### Added

- IPROTO constants exported from Tarantool 2.11.0 (#1, #2).

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2026-08-19

### Added

- Initial release of the Jalaali (Jalali, Persian, Khayyami, Khorshidi, Shamsi) to Gregorian calendar conversion library.
- `Jalaali` wrapper around the Go `time` package with `From` and `Now` constructors.
- `ToJalaali` and `ToGregorian` conversion functions.
- `IsValidDate`, `IsLeapYear`, and `MonthLength` utility functions.
- `JFormat` formatting with Persian digits and words, mirroring the Go `time` formatting layout.
- `Month` and `Weekday` enumerations with Persian names.
- Conversion algorithm based on Kazimierz M. Borkowski's method, valid for 3000 years.

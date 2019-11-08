# Changelog

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.1] - 2019-11-08
### Changed
* If user name starts with a 'q' and the user name is not equal to the
  password, authentication fails.

  This is to test changing authentication, e.g. for software that does only
  allow a certain number of failed authentications within a certain time
  period.

## [1.0.0] - 2019-08-29
### Added
* Simple authentication, based solely on the user name.

# go-utils

<a name="v1.29.1"></a>
## [v1.29.1] - 2022-12-14
### New Features
- add GitHub action to run golangci-lint, unit test, and go build

<a name="v1.28.0"></a>
## [v1.28.0] - 2022-11-04
### New Features
- add DeleteByValue


<a name="v1.27.0"></a>
## [v1.27.0] - 2022-09-22
### New Features
- add GetCronNextAt


<a name="v1.26.0"></a>
## [v1.26.0] - 2022-09-21
### New Features
- create ValueOrDefault function


<a name="v1.25.1"></a>
## [v1.25.1] - 2022-08-25
### Other Improvements
- upgrade bluemonday


<a name="v1.25.0"></a>
## [v1.25.0] - 2022-08-11
### New Features
- escape quote


<a name="v1.24.0"></a>
## [v1.24.0] - 2022-07-15
### New Features
- nanoseconds check on rfc3339nano time formatting ([#38](https://github.com/kumparan/kumnats/issues/38))


<a name="v1.23.0"></a>
## [v1.23.0] - 2022-06-29
### New Features
- parse duration with default ([#36](https://github.com/kumparan/kumnats/issues/36))
- add int64 pointer to int64 conversion ([#33](https://github.com/kumparan/kumnats/issues/33))

### Other Improvements
- fix dependabot alert ([#37](https://github.com/kumparan/kumnats/issues/37))


<a name="v1.22.0"></a>
## [v1.22.0] - 2022-06-24
### New Features
- generate ULID from time


<a name="v1.21.0"></a>
## [v1.21.0] - 2022-06-24
### New Features
- add some functions using go generic ([#34](https://github.com/kumparan/kumnats/issues/34))


<a name="v1.20.1"></a>
## [v1.20.1] - 2022-03-15
### Fixes
- fix marshal issue on gorm.DeletedAt empty value ([#32](https://github.com/kumparan/kumnats/issues/32))


<a name="v1.20.0"></a>
## [v1.20.0] - 2022-03-11

<a name="v.1.20.0"></a>
## [v.1.20.0] - 2022-03-11
### New Features
- add constraint size gql directive ([#30](https://github.com/kumparan/kumnats/issues/30))


<a name="v1.19.3"></a>
## [v1.19.3] - 2022-02-25
### Fixes
- should handle error when marshal/unmarshal from gqlgen ([#29](https://github.com/kumparan/kumnats/issues/29))


<a name="v1.19.2"></a>
## [v1.19.2] - 2022-02-07
### Fixes
- handle index out of bound ([#28](https://github.com/kumparan/kumnats/issues/28))


<a name="v1.19.1"></a>
## [v1.19.1] - 2022-01-27
### Fixes
- add null string & handle unmarshal "null" values ([#27](https://github.com/kumparan/kumnats/issues/27))


<a name="v1.19.0"></a>
## [v1.19.0] - 2022-01-19
### New Features
- add gqlgen NullInt64, NullInt64ID & NullTime ([#26](https://github.com/kumparan/kumnats/issues/26))


<a name="v1.18.1"></a>
## [v1.18.1] - 2022-01-13
### Code Improvements
- generate the private key only once ([#25](https://github.com/kumparan/kumnats/issues/25))


<a name="v1.18.0"></a>
## [v1.18.0] - 2022-01-05

<a name="v1.17.0"></a>
## [v1.17.0] - 2022-01-03
### New Features
- add custom time for gqlgen ([#22](https://github.com/kumparan/kumnats/issues/22))


<a name="v1.16.0"></a>
## [v1.16.0] - 2022-01-03
### New Features
- add AESCryptor ([#23](https://github.com/kumparan/kumnats/issues/23))


<a name="v1.15.0"></a>
## [v1.15.0] - 2022-01-03
### New Features
- logrus sentry hook using sentry-go ([#18](https://github.com/kumparan/kumnats/issues/18))


<a name="v1.14.1"></a>
## [v1.14.1] - 2021-12-30
### Fixes
- gqlgen cannot resolve the package ([#21](https://github.com/kumparan/kumnats/issues/21))


<a name="v1.14.0"></a>
## [v1.14.0] - 2021-12-30
### New Features
- add custom scalara gqlgen Int64ID ([#20](https://github.com/kumparan/kumnats/issues/20))


<a name="v1.13.0"></a>
## [v1.13.0] - 2021-12-30
### New Features
- add GetDifferenceDaysForHumans ([#19](https://github.com/kumparan/kumnats/issues/19))


<a name="v1.12.0"></a>
## [v1.12.0] - 2021-04-12
### New Features
- strip HTML ([#16](https://github.com/kumparan/kumnats/issues/16))


<a name="v1.11.0"></a>
## [v1.11.0] - 2021-03-17
### New Features
- add int64 millis to time converter


<a name="v1.10.0"></a>
## [v1.10.0] - 2021-03-17
### New Features
- add string millis to time converter


<a name="v1.9.0"></a>
## [v1.9.0] - 2021-03-10
### New Features
- add money formatter for multiple currencies ([#13](https://github.com/kumparan/kumnats/issues/13))


<a name="v1.8.0"></a>
## [v1.8.0] - 2020-12-10

<a name="v1.7.1"></a>
## [v1.7.1] - 2020-12-10
### New Features
- add formatter for indonesian money and date


<a name="v1.7.0"></a>
## [v1.7.0] - 2020-11-19
### New Features
- add function to truncate string by length ([#11](https://github.com/kumparan/kumnats/issues/11))


<a name="v1.6.0"></a>
## [v1.6.0] - 2020-06-30
### Code Improvements
- rename func from GetStoryIDFromStorySlug to GetIDFromSlug

### New Features
- get story id from story slug


<a name="v1.5.0"></a>
## [v1.5.0] - 2020-06-17
### Code Improvements
- JoinURL to be more versatile and should be able to accept variadic variables ([#9](https://github.com/kumparan/kumnats/issues/9))


<a name="v1.4.0"></a>
## [v1.4.0] - 2020-06-11
### New Features
- create GenerateRandomAlphanumeric


<a name="v1.3.1"></a>
## [v1.3.1] - 2020-04-02
### Fixes
- missing RetryStopper constructor ([#6](https://github.com/kumparan/kumnats/issues/6))


<a name="v1.3.0"></a>
## [v1.3.0] - 2020-03-06
### New Features
- add lower map key


<a name="v1.2.0"></a>
## [v1.2.0] - 2020-02-04
### New Features
- new method to know the caller of the method in the runtime


<a name="v1.1.1"></a>
## [v1.1.1] - 2020-01-06

<a name="v1.1.0"></a>
## [v1.1.0] - 2020-01-06
### New Features
- generate media URL


<a name="v1.0.0"></a>
## v1.0.0 - 2019-12-23
### New Features
- init go-utils


[Unreleased]: https://github.com/kumparan/kumnats/compare/v1.28.0...HEAD
[v1.28.0]: https://github.com/kumparan/kumnats/compare/v1.27.0...v1.28.0
[v1.27.0]: https://github.com/kumparan/kumnats/compare/v1.26.0...v1.27.0
[v1.26.0]: https://github.com/kumparan/kumnats/compare/v1.25.1...v1.26.0
[v1.25.1]: https://github.com/kumparan/kumnats/compare/v1.25.0...v1.25.1
[v1.25.0]: https://github.com/kumparan/kumnats/compare/v1.24.0...v1.25.0
[v1.24.0]: https://github.com/kumparan/kumnats/compare/v1.23.0...v1.24.0
[v1.23.0]: https://github.com/kumparan/kumnats/compare/v1.22.0...v1.23.0
[v1.22.0]: https://github.com/kumparan/kumnats/compare/v1.21.0...v1.22.0
[v1.21.0]: https://github.com/kumparan/kumnats/compare/v1.20.1...v1.21.0
[v1.20.1]: https://github.com/kumparan/kumnats/compare/v1.20.0...v1.20.1
[v1.20.0]: https://github.com/kumparan/kumnats/compare/v.1.20.0...v1.20.0
[v.1.20.0]: https://github.com/kumparan/kumnats/compare/v1.19.3...v.1.20.0
[v1.19.3]: https://github.com/kumparan/kumnats/compare/v1.19.2...v1.19.3
[v1.19.2]: https://github.com/kumparan/kumnats/compare/v1.19.1...v1.19.2
[v1.19.1]: https://github.com/kumparan/kumnats/compare/v1.19.0...v1.19.1
[v1.19.0]: https://github.com/kumparan/kumnats/compare/v1.18.1...v1.19.0
[v1.18.1]: https://github.com/kumparan/kumnats/compare/v1.18.0...v1.18.1
[v1.18.0]: https://github.com/kumparan/kumnats/compare/v1.17.0...v1.18.0
[v1.17.0]: https://github.com/kumparan/kumnats/compare/v1.16.0...v1.17.0
[v1.16.0]: https://github.com/kumparan/kumnats/compare/v1.15.0...v1.16.0
[v1.15.0]: https://github.com/kumparan/kumnats/compare/v1.14.1...v1.15.0
[v1.14.1]: https://github.com/kumparan/kumnats/compare/v1.14.0...v1.14.1
[v1.14.0]: https://github.com/kumparan/kumnats/compare/v1.13.0...v1.14.0
[v1.13.0]: https://github.com/kumparan/kumnats/compare/v1.12.0...v1.13.0
[v1.12.0]: https://github.com/kumparan/kumnats/compare/v1.11.0...v1.12.0
[v1.11.0]: https://github.com/kumparan/kumnats/compare/v1.10.0...v1.11.0
[v1.10.0]: https://github.com/kumparan/kumnats/compare/v1.9.0...v1.10.0
[v1.9.0]: https://github.com/kumparan/kumnats/compare/v1.8.0...v1.9.0
[v1.8.0]: https://github.com/kumparan/kumnats/compare/v1.7.1...v1.8.0
[v1.7.1]: https://github.com/kumparan/kumnats/compare/v1.7.0...v1.7.1
[v1.7.0]: https://github.com/kumparan/kumnats/compare/v1.6.0...v1.7.0
[v1.6.0]: https://github.com/kumparan/kumnats/compare/v1.5.0...v1.6.0
[v1.5.0]: https://github.com/kumparan/kumnats/compare/v1.4.0...v1.5.0
[v1.4.0]: https://github.com/kumparan/kumnats/compare/v1.3.1...v1.4.0
[v1.3.1]: https://github.com/kumparan/kumnats/compare/v1.3.0...v1.3.1
[v1.3.0]: https://github.com/kumparan/kumnats/compare/v1.2.0...v1.3.0
[v1.2.0]: https://github.com/kumparan/kumnats/compare/v1.1.1...v1.2.0
[v1.1.1]: https://github.com/kumparan/kumnats/compare/v1.1.0...v1.1.1
[v1.1.0]: https://github.com/kumparan/kumnats/compare/v1.0.0...v1.1.0

# Changelog
All notable changes to this project will be documented in this file.

## [2.2.0] - 2020-06-23

### Added
- Project redirect settings. ([47b0626](https://github.com/paysuper/paysuper-management-api/commit/47b06265c4216ebaba93d4570d63ade6521deb6c))

### Changed
- Date filters refactoring ([86c84b7](https://github.com/paysuper/paysuper-management-api/commit/86c84b76605c97a2baf546517c7ae2b82d49c17d))
- Updated GO and Alpine Linux versions in the dockerfile. ([b61e367](https://github.com/paysuper/paysuper-management-api/commit/b61e36738f31de604dfff6f628e9ebc106699423))
- Add a filter parameter to exclude test orders from orders list. ([c5c4dbb](https://github.com/paysuper/paysuper-management-api/commit/c5c4dbb19bf8d264a1d7b69b3f7e315ee3401cc7))
- Add the merchantGetMerchantStatus permission for the merchant_view_only role. ([de17452](https://github.com/paysuper/paysuper-management-api/commit/de1745263798e9a132d09fcc6040995a737c6fae))
- Add more read permissions for the paylink for the merchant_view_only role. ([a6c6ce2](https://github.com/paysuper/paysuper-management-api/commit/a6c6ce230d01a42e0e144f50a818ba41f558b38d))
- Add the merchantGetPaylinksList permission for the merchant_view_only role. ([5ca26d7](https://github.com/paysuper/paysuper-management-api/commit/5ca26d7a6eb0eadac91a30b6fdd8dab0def99d12))
- Add a webhook testing method annotation. ([7071622](https://github.com/paysuper/paysuper-management-api/commit/7071622d707d0fd6928b54fad4976a53cd895f8f))
- Mark test and production orders with a special flag. ([915d991](https://github.com/paysuper/paysuper-management-api/commit/915d9910288fdd54bc04c100094b2ff9bb19afb4))
- Method to download a report file . ([a925fdc](https://github.com/paysuper/paysuper-management-api/commit/a925fdc953de1d32afff41c9bb88c56cc2750480)) ([1e92585](https://github.com/paysuper/paysuper-management-api/commit/1e92585167920fff659a01f454d58c13085a79b9))

### Fixed
- Buq with validation in the reports' download requests. ([7a50d1f](https://github.com/paysuper/paysuper-management-api/commit/7a50d1f08b1f6bdfb5839e3c8f98c9a33a781edd))
- Check merchant permission for the getPaylinkUrl method. ([#228](https://github.com/paysuper/paysuper-management-api/issues/228)) ([7dccf89](https://github.com/paysuper/paysuper-management-api/commit/7dccf8950e2ae08a6014789c21100036fe03cf1d))
- Fix a paylink URL. ([c43f85f](https://github.com/paysuper/paysuper-management-api/commit/c43f85fa147974e37b737c9cfd12d8026451c4ae)) ([39f79bb](https://github.com/paysuper/paysuper-management-api/commit/39f79bb6dd3e545e9e66baf57534b1527438f5c9)) ([c47756a](https://github.com/paysuper/paysuper-management-api/commit/c47756ac883c0a8be4050f44bc3aef1eab474b41)) ([8c35859](https://github.com/paysuper/paysuper-management-api/commit/8c35859be72ca4c7106450bf920b59d211cb3a89))
- Fix a transaction logs filter for a system administrator. ([1e1719b](https://github.com/paysuper/paysuper-management-api/commit/1e1719b946af72c64582e7ea5acc3a64ccbf5587))
- Fix a transaction log for an admin. ([22dde2b](https://github.com/paysuper/paysuper-management-api/commit/22dde2b91ddd0920880e63e028afca7173be73a7))

### Removed
- Remove unused constants. ([f58087c](https://github.com/paysuper/paysuper-management-api/commit/f58087c1e52ad05b1613fdbc73f89d57ed39f6e6))
- Remove unused endpoint /admin/api/v1/report_file. ([e772631](https://github.com/paysuper/paysuper-management-api/commit/e772631a1cf5c9f9a681b6e46ef8f229767e3eff))

***

## [2.1.0] - 2020-02-04

### Added
- The new parameter to filter the transactions log by the production or test mode.
- The new project's settings group to redirect user at the end of the payment process.
- Added API documentation using the OpenAPI specification in GO source files.

### Changed
- The merchant's license agreement (downloadable from PaySuper Dashboard) filename has been edited.
- Removed unused code.
- Updated GO and Alpine Linux versions in the dockerfile
- Update project's dependencies.

***

## [2.0.0] - 2019-12-23

### Changed
- Update project's dependencies.

### Removed
- Deleted the unused source code.
- Removed API endpoints for a payment processing and remained only API endpoints for the management functionality.

***

## [1.0.0] - 2019-12-19

### Added
- Send VAT value to the frontend.
- Set up a CORS configuration for restrictions.

### Changed
- Updated the Casbin policy for API methods.
- Updated dependencies.
- Updated README. 

### Removed
- Removed unused Onboarding API methods.
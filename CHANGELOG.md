# Changelog
All notable changes to this project will be documented in this file.


## [2.7.0] - 2020-12-03

### Added
- KYC procedure for merchant onboarding

### Fixed
- Some minor fixes added 
    
***


## [2.6.0] - 2020-11-18

### Added
- Routes for subscriptions functionality
- Routes for List of acts of completion

### Changed
- Update dependencies.

*** 

## [2.5.0] - 2020-09-10

### Added
- Remove the order's amount rounding. ([e4d6a95])(https://github.com/paysuper/paysuper-management-api/commit/e4d6a95804a20ce7e1bf423dc0971708591326a7)
- Add an endpoint for the act of completion report (#247). ([68b7e60](https://github.com/paysuper/paysuper-management-api/commit/68b7e60e588139666ad0f0cd71879534c879749d))

### Changed
- The URL path of a request to download an act of completion report (#251). ([1a50fca](https://github.com/paysuper/paysuper-management-api/commit/1a50fca9eb217bbc4b49fe006cd33219776185eb))
- Update dependencies. ([cfcfaab](https://github.com/paysuper/paysuper-management-api/commit/cfcfaab592493c944a279d20d85fcbfbe25e002d)) ([10c3955](https://github.com/paysuper/paysuper-management-api/commit/10c3955f65c356e707a7e194a240c94b60ed47a3)) ([2cf0942]()https://github.com/paysuper/paysuper-management-api/commit/2cf0942605da605bd73b1e835925ad184664953c)

### Fixed
- Use a string type value for dates in the ListOrdersRequest (#253). ([22d94f9](https://github.com/paysuper/paysuper-management-api/commit/22d94f93bef1dc7010f975e3746713bb93ba2531))
- If a company city contains an apostrophe then don't return an error (#249). ([9eca626](https://github.com/paysuper/paysuper-management-api/commit/9eca6265ccaff8b2a7f91215adc0299180364464))

### Removed

***

## [2.4.0] - 2020-08-26

### Added
- The PaySuper admin user obtains the list and detailed information about payouts. ([16d5e96](https://github.com/paysuper/paysuper-management-api/commit/16d5e969099f60d4c11a497531ebdba91acfe3d9)) ([eea79f9](https://github.com/paysuper/paysuper-management-api/commit/eea79f915ceea1710b8f0240ef1112c95b779c1f))
- Send notifications to the user's channel when the export has done. ([0e16a6d](https://github.com/paysuper/paysuper-management-api/commit/0e16a6d0dafeda357e94e3502001882e7b25a1a7))
- Flag to skip the post process. ([c06504d](https://github.com/paysuper/paysuper-management-api/commit/c06504d39c3d0c9bcd96054dfc6c91965c271337))
- The endpoint to return the order's private data (#244). ([18f7ed3](https://github.com/paysuper/paysuper-management-api/commit/18f7ed3643a44e11fa86c007f1c8fb63c2e1ba0c)) ([2dbab98](https://github.com/paysuper/paysuper-management-api/commit/2dbab98afc9019f269632df4ac6981272aac978c))
- S2S APIs. ([a73c4d7](https://github.com/paysuper/paysuper-management-api/commit/a73c4d7bf7852e4e58b1897e39928a873a73f145)) ([2daa00f](https://github.com/paysuper/paysuper-management-api/commit/2daa00ff7bb56382e6a179c5ab94660f5fe537d9))

### Changed
- Update dependencies. ([771c09e](https://github.com/paysuper/paysuper-management-api/commit/771c09edbcd7f804e8b365cd583c59dfddd5c205))

### Fixed
- GB zip regex. ([085f1ff](https://github.com/paysuper/paysuper-management-api/commit/085f1ff3995f4eefe5627b6795428caf06fde119))
- Fix a bug for a newly registered user. ([4a798ac](https://github.com/paysuper/paysuper-management-api/commit/4a798acfc06177bc3207487872189be441c153bb))
- Rounding amounts in the order view. ([2b8cd38](https://github.com/paysuper/paysuper-management-api/commit/2b8cd381f5f7f9ed842cd9137ea66f2834298f8f))

### Removed
- Remove the unused constant. ([687f88a](https://github.com/paysuper/paysuper-management-api/commit/687f88a23c33f0677fcfc007838552caf8c621f3))

***

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
- Buq with validation in the report download requests. ([7a50d1f](https://github.com/paysuper/paysuper-management-api/commit/7a50d1f08b1f6bdfb5839e3c8f98c9a33a781edd))
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

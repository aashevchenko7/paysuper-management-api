# Changelog
All notable changes to this project will be documented in this file.

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
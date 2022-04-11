# wallet-engine-api
## Wallet API
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#testing">Testing</a></li>
  </ol>
</details>


## About The Project
You are to build a wallet API that can be used to create, credit, debit and activate/deactivate a wallet.
## Getting Started

### Prerequisites

Before this project can be run locally you will need to install [Go](https://golang.org/doc/install)

### Installation

To utilize the project, the following needs to be done:
1. Clone this repository
2. Install the dependencies, using the following command:
```
go mod tidy
```

## Usage

1. To run the project locally, use the following command:
```
make run
```
2. To create a given wallet, use Postman to make a POST request to the following URL:
```
http://localhost:{port}/api/v1/createWallet
```
3. To credit money on a given wallet
   , use Postman to make a POST request to the following URL:
```
http://localhost:{port}/api/v1/creditWallet/:reference
```
4. To debit money from a given wallet reference, use Postman to make a POST request to the following URL:
```
http://localhost:{port}/api/v1/debitWallet/:reference
```
5. To activate a given wallet, use Postman to make a PUT request to the following URL:
```
http://localhost:{port}/api/v1/activateDeactivateWallet/:reference
```

## Testing
Tests can be run using the following command:
```
make test
```
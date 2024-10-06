# Avalanche-HyperSDK
# TokenVM Project

## Overview

The TokenVM Project is a demonstration of creating, minting, and managing a custom token using the TokenVM framework. This project showcases the functionality of token management, including asset creation, minting, order creation, and transaction observation on the Avalanche blockchain.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
  - [Create Your Asset](#create-your-asset)
  - [Mint Your Asset](#mint-your-asset)
  - [View Your Balance](#view-your-balance)
  - [Create an Order](#create-an-order)
  - [Fill Part of the Order](#fill-part-of-the-order)
  - [Close Order](#close-order)
  - [Watch Activity in Real-Time](#watch-activity-in-real-time)
- [Contributing](#contributing)
- [License](#license)

## Prerequisites

- [Avalanche CLI](https://docs.avax.network/learn/avalanche-cli) installed on your machine.
- A running Avalanche network (local or testnet).
- Basic understanding of command line operations.

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/tokenvm.git
   cd tokenvm
   ```

2. Build the project:
   ```bash
   ./build.sh
   ```

## Usage

### Create Your Asset

To create a new asset, run:
```bash
./build/token-cli action create-asset
```
- Follow the prompts to enter the asset metadata.

### Mint Your Asset

To mint tokens for your newly created asset, run:
```bash
./build/token-cli action mint-asset
```
- Provide the required asset ID, recipient address, and amount.

### View Your Balance

To check your token balance, run:
```bash
./build/token-cli key balance
```
- Specify the asset ID (use `TKN` for native token).

### Create an Order

To place a trade order, run:
```bash
./build/token-cli action create-order
```
- Follow the prompts to enter the necessary details.

### Fill Part of the Order

To fill an existing order, run:
```bash
./build/token-cli action fill-order
```
- Specify the asset IDs and amounts accordingly.

### Close Order

To cancel an order, run:
```bash
./build/token-cli action close-order
```
- Provide the required details to cancel the order.

### Watch Activity in Real-Time

To observe blockchain activity, run:
```bash
./build/token-cli chain watch
```
- Select the chain ID from the available options.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you find any bugs or have suggestions for improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

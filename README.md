# Aparok Sync Coordinator Node
Implementation of coordinator node from [`aparok-sync`](https://github.com/altercpre/aparok-sync) protocol.

## Building the source
To ensure compatibility, please use Go version `1.22`.

To build and run the Aparok Sync Coordinator Node on your own server, follow these technical steps:

1.  Clone the Aparok Sync Coordinator Node repository to your local machine.
2.  Navigate to the root directory of the repository, where you will find a `Makefile`.
3.  Run the following commands to install the required dependencies and build the Aparok Sync Coordinator Node.
    ```
    make deps
    make build
    ```
4.  If there are no errors, the Aparok Sync Coordinator Node will be built and can be found in the `./bin` directory.

## Running
You will need a MongoDB in replica set mode to run Aparok Sync Coordinator Node.

Aparok Sync Coordinator Node requires a configuration. You can generate configuration files for your nodes with [`aparok-sync-network`](https://github.com/altercpre/aparok-sync-tools) tool.

The following options are available for running the Aparok Sync Coordinator Node:

 - `-c` — path to config file (default `etc/aparok-sync-coordinator.yml`). 
 - `-v` — current version.
 - `-h` — help message.

## Contribution
Thank you for your desire to develop Aparok together!

❤️ This project and everyone involved in it is governed by the [Code of Conduct](https://github.com/altercpre/.github/blob/main/docs/CODE_OF_CONDUCT.md).

🧑‍💻 Check out our [contributing guide](https://github.com/altercpre/.github/blob/main/docs/CONTRIBUTING.md) to learn about asking questions, creating issues, or submitting pull requests.

🫢 For security findings, please email [security@altercpre.w8n.pw](mailto:security@altercpre.w8n.pw) and refer to our [security guide](https://github.com/altercpre/.github/blob/main/docs/SECURITY.md) for more information.

🤝 Follow us on [Github](https://github.com/altercpre) and join the [Contributors Community](https://github.com/orgs/altercpre/discussions).

---
Made by Altercpre — privacy-first, security-focused 🔒

Licensed under [MIT License](./LICENSE).

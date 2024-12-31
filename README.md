[![CoreDNS](https://coredns.io/images/CoreDNS_Colour_Horizontal.png)](https://coredns.io)

[![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/coredns/coredns)
[![Release](https://github.com/machsix/coredns-eh/actions/workflows/release.yml/badge.svg?event=release)](https://github.com/machsix/coredns-eh/actions/workflows/release.yml)
# CoreDNS with WGSD and CoreDNS-mDNS

This repository is a package of CoreDNS bundled with [WGSD](https://github.com/jwhited/wgsd) and [CoreDNS-mDNS](https://github.com/openshift/coredns-mdns).

## Introduction

CoreDNS is a DNS server that chains plugins. It is flexible and can be used in a variety of scenarios, from serving DNS records for a domain to acting as a proxy to other DNS servers.

### WGSD

[WGSD](https://github.com/jwhited/wgsd) is a plugin for CoreDNS that provides DNS-based service discovery for WireGuard peers. It allows WireGuard peers to discover each other using DNS queries.

### CoreDNS-mDNS

[CoreDNS-mDNS](https://github.com/openshift/coredns-mdns) is a plugin for CoreDNS that enables multicast DNS (mDNS) support. This allows CoreDNS to resolve local network names without needing a central DNS server.

## License

All three packages in this repository are licensed under the Apache License 2.0. You can find the full text of the license in the `LICENSE` file.

## Getting Started

To get started with this package, follow the instructions in the respective repositories for CoreDNS, WGSD, and CoreDNS-mDNS. You can find more detailed documentation and usage examples there.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request if you have any improvements or bug fixes.

## Contact

For any questions or support, please open an issue in this repository.
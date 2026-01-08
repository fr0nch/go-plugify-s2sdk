[![Русский](https://img.shields.io/badge/Русский-%F0%9F%87%B7%F0%9F%87%BA-green?style=for-the-badge)](README_ru.md)

# Go bindings for the Source 2 SDK Plugify plugin

This repository contains Go bindings for the [plugify-plugin-s2sdk](https://github.com/untrustedmodders/plugify-plugin-s2sdk) plugin. The bindings are automatically generated and kept in sync with the original plugin.

## API Documentation

Full API documentation can be found here on the [API Hub](https://api.plugify.net?file=https://raw.githubusercontent.com/untrustedmodders/plugify-plugin-s2sdk/refs/heads/main/plugify-plugin-s2sdk.pplugin.in) website.

## Installation

```bash
go get github.com/fr0nch/go-plugify-s2sdk/v2
```

## Updates

This repository uses GitHub Actions to automatically check for updates in [plugify-plugin-s2sdk](https://github.com/untrustedmodders/plugify-plugin-s2sdk) every day at 2:00 UTC. When a new version is detected, the bindings are updated, and a new release is created.

## License

This Go module for Plugify is licensed under the [MIT License](LICENSE).
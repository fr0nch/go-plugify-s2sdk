[![English](https://img.shields.io/badge/English-%F0%9F%87%AC%F0%9F%87%A7-blue?style=for-the-badge)](README.md)

# Go биндинги для Source 2 SDK Plugify плагина

Этот репозиторий содержит Go биндинги для плагина [plugify-plugin-s2sdk](https://github.com/untrustedmodders/plugify-plugin-s2sdk). Биндинги генерируются автоматически и синхронизируются с оригинальным плагином.

## Документация по API

Полную документацию по API можно найти здесь на сайте [API Hub](https://api.plugify.net?file=https://raw.githubusercontent.com/untrustedmodders/plugify-plugin-s2sdk/refs/heads/main/plugify-plugin-s2sdk.pplugin.in)

## Установка

```bash
go get github.com/fr0nch/go-plugify-s2sdk/v2
```

## Обновление

Этот репозиторий использует GitHub Actions для автоматической проверки обновлений в [plugify-plugin-s2sdk](https://github.com/untrustedmodders/plugify-plugin-s2sdk) в каждые 2:00 по UTC. При обнаружении новой версии, биндинги обновляются, и создается новый релиз.

## Лицензия

Этот модуль Go для Plugify лицензирован на условиях [MIT License](LICENSE).
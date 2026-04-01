# Tesla Home Assistant Add-ons

Custom Add-ons fuer Tesla-bezogene Dienste auf Home Assistant Green.

## Enthalten
- `tesla-fleet-telemetry`: Teslas offizieller `fleet-telemetry` Server als Home-Assistant-Add-on

## Installation
1. Dieses Repo in Home Assistant unter `Einstellungen -> Add-ons -> Add-on Store -> Repositories` hinzufuegen
2. Add-on `Tesla Fleet Telemetry` installieren
3. Optionen setzen
4. Zertifikate in `/ssl` hinterlegen
5. Router-Portforward auf den konfigurierten HTTPS-Port

## Erstes Ziel
Eingehende Tesla Fleet Telemetry auf `telemetry.mein-lila-tablett.at` annehmen, danach `fleet_telemetry_config` vom lokalen Tesla-Charging-App-Repo ans Fahrzeug pushen.

# Tesla Fleet Telemetry Add-on

Home-Assistant-Green-Add-on fuer Teslas offiziellen `fleet-telemetry` Server.

## Voraussetzungen
- `telemetry.mein-lila-tablett.at` zeigt per DNS auf deine Heim-IP
- Router forwardet den gewaehlten HTTPS-Port auf Home Assistant Green
- Gueltiges TLS-Zertifikat liegt in Home Assistant unter `/ssl/fullchain.pem` und `/ssl/privkey.pem`

## Installation
1. Repo in Home Assistant unter `Einstellungen -> Add-ons -> Add-on Store -> Repositories` hinzufuegen
2. Add-on `Tesla Fleet Telemetry` installieren
3. Optionen setzen:
   - `host`: `0.0.0.0`
   - `port`: `443`
   - `server_cert`: `/ssl/fullchain.pem`
   - `server_key`: `/ssl/privkey.pem`
4. Add-on starten

Wichtig:
- `host` ist die lokale Bind-Adresse des Servers, nicht dein öffentlicher DNS-Name
- dein öffentlicher Hostname bleibt `telemetry.mein-lila-tablett.at`
- beim späteren `fleet_telemetry_config` Push verwendest du weiter `telemetry.mein-lila-tablett.at`

## Troubleshooting
- Wenn vorher `cannot execute: required file not found` kam, Add-on aktualisieren oder neu bauen lassen. Das war der alte Binary-Base-Image-Mismatch.

## Danach
Wenn das Add-on laeuft, aus dem Tesla-Charging-App-Repo die Fahrzeug-Config pushen:

```bash
uv run python scripts/push_fleet_telemetry_config.py \
  --vin LRWYGCEK6NC350679 \
  --hostname telemetry.mein-lila-tablett.at \
  --ca-file deploy/fleet-telemetry/certs/fullchain.pem
```

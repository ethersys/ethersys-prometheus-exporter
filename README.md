# Prometheus exporter

Petit service qui sert à exposer la consommation de RAM des principeaux processe sur un compte d'hébergement.

## Usage

```bash
exporter '<IP>' '<PORT>' '<USER>' '<PASSWORD_BCRYPT_HASH>'
```

To get password hash

```bash
echo $(htpasswd -bnBC 10 "" "<PASSWORD>" | tr -d ':\n' | sed 's/$2y/$2a/' | sed 's/://')
```

## Build

```bash
git clone git@github.com:ethersys/ethersys-ansible-public.git
go install
go build
```

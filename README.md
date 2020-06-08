# aleph-exporter
![CI](https://github.com/ckluenter/aleph-exporter/workflows/CI/badge.svg)

Exposes some metrics from the aleph api (https://github.com/alephdata/aleph) as prometheus metrics.
![Screenshot](contrib/dashboard-screenshot.png)
## Test
`make test`

## RUN
```
./aleph-exporter --aleph-host some-host-running-aleph-api --aleph-token "some-token"`
```

## Docker

```
docker run --rm -p 8080:8080 -e ALEPH_HOST=$ALEPH_HOST -e ALEPH_TOKEN=$ALEPH_TOKEN ckluenter/alephexporter
```

## Grafana
There is an example dashboard in [contrib](contrib/)
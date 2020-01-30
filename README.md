# aleph-exporter

## Test
`make test`

## RUN
```
./aleph-exporter --aleph-host some-host-running-aleph-api --aleph-token "some-token"`
```

## Docker

```
docker run --rm -e ALEPH_HOST=$ALEPH_HOST -e ALEPH_TOKEN=$ALEPH_TOKEN ckluenter/alephexporter
```

# Patch API

A simple REST API for scheduling.

## API

Example of the API structure

`GET /patch/{ID}`

```json
{
  "ID": 15,
  "CreatedAt": "2023-06-05T13:16:29.418155715+01:00",
  "UpdatedAt": "2023-06-05T13:16:29.418155715+01:00",
  "DeletedAt": null,
  "Server": "server22.testing.lab",
  "PatchStart": "2023-06-12 20:45:00",
  "PreCheckScheduled": "1",
  "PreCheckStatus": "1",
  "PreCheckJobID": "",
  "PatchScheduled": "1",
  "PatchScheduleID": "323",
  "PatchJobID": "",
  "Status": "PatchScheduled"
}
```

`GET /patch/date/{date (%Y-%m-%d)}`

```json
[
  {
    "ID": 14,
    "CreatedAt": "2023-06-05T13:16:29.398938485+01:00",
    "UpdatedAt": "2023-06-05T13:16:29.398938485+01:00",
    "DeletedAt": null,
    "Server": "server19.testing.lab",
    "PatchStart": "2023-06-12 18:00:00",
    "PreCheckScheduled": "1",
    "PreCheckStatus": "1",
    "PreCheckJobID": "",
    "PatchScheduled": "1",
    "PatchScheduleID": "321",
    "PatchJobID": "",
    "Status": "PatchScheduled"
  },
  {
    "ID": 15,
    "CreatedAt": "2023-06-05T13:16:29.418155715+01:00",
    "UpdatedAt": "2023-06-05T13:16:29.418155715+01:00",
    "DeletedAt": null,
    "Server": "server22.testing.lab",
    "PatchStart": "2023-06-12 20:45:00",
    "PreCheckScheduled": "1",
    "PreCheckStatus": "1",
    "PreCheckJobID": "",
    "PatchScheduled": "1",
    "PatchScheduleID": "323",
    "PatchJobID": "",
    "Status": "PatchScheduled"
  }
]
```

`PUT /patch/{ID}`

Fields that can be updated:

```yaml
PreCheckScheduled:
PreCheckStatus:
PreCheckJobID:
PatchScheduled:
PatchScheduleID:
PatchJobID:
Status:
```

## To Build

Either run `go build -o patch-api cmd/main.go` or use `build_envs/` scripts to build for a specific environment.

## To deploy

Example systemd file `/etc/systemd/system/patch-api.service`:

```bash
[Unit]
Description=Patch API Service
Wants=network.target network-online.target
After=network.target network-online.target

[Service]
ExecStart=/usr/local/bin/patch-api --host 0.0.0.0 --port 8080
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

> SELinux context may need to be fixed in order for this to run.

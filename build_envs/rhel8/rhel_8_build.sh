#!/usr/bin/env bash

podman build -f Containerfile -t golang_patch_api_rhel8
podman create --name temp_rhel8 golang_patch_api_rhel8
podman cp temp_rhel8:/patch-api/patch-api_rhel8 .
podman rm -f temp_rhel8
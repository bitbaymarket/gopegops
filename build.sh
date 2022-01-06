#!/bin/sh
docker build . -f Dockerfile -t tmp/gopegops-check:latest
docker build . -f Dockerfile_a310 -t tmp/gopegops-check-a310:latest


#!/bin/bash

cp ~/.ssh/id_droplet .

docker build -t test:v0.0.1 .

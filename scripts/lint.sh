#!/bin/bash

set -e
set -x

cd src
./scripts/lint.sh
cd ..

cd doc
./scripts/lint.sh
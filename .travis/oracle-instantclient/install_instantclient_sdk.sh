#!/bin/sh -e

sudo dpkg --install `sudo alien --scripts --to-deb "$ORACLE_INSTANTCLIENT_SDK_FILE" | cut -d' ' -f1`

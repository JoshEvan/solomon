#!/bin/sh
# exec unset NODE_OPTIONS
# export NODE_OPTIONS=--openssl-legacy-provider
# export 
echo $(node --version)
# exec env NODE_OPTIONS=--openssl-legacy-provider npm run dev
exec npm run dev

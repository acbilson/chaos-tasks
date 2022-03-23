#!/bin/bash
. .env

ENVIRONMENT=$1

case $ENVIRONMENT in

dev)
  dotnet test -l "console;verbosity=detailed" Tests/*csproj
;;

*)
  echo "please provide one of the following as the first argument: dev."
  exit 1

esac

#!/usr/bin/env bash
go run .
cd $(cat /tmp/ProjManCd*)
rm /tmp/ProjManCd*
bash -i 

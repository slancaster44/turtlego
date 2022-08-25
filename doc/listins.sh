#!/usr/bin/bash

echo $1 >> tmp.s
as  tmp.s -o out.o
objdump -d --disassembler-options=intel-mnemonic out.o
rm -r tmp.s out.o
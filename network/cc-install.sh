#!/bin/bash

# Exit on first error
set -e

echo "=================== start ==================="
echo start at：$(date +%Y-%m-%d\ %H:%M:%S)
echo "============================================="
./cc.sh install acpc 1.0 go/acpc Synchro
./cc.sh install acc 1.0 go/acc Synchro
./cc.sh install odc 1.0 go/odc Synchro
./cc.sh install omc 1.0 go/omc Synchro
echo "=================== end ==================="
echo end at: $(date +%Y-%m-%d\ %H:%M:%S)
echo "==========================================="

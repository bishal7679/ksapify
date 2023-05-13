#!/bin/bash
bash <(curl -s https://raw.githubusercontent.com/bishal7679/ksapify/main/builder.sh)


echo -e "\033[32;40mINSTALLATION COMPLETED!\033[0m\n"


cd - || echo -e "\033[31;40mFailed to move to previous directory\033[0m\n"
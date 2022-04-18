#!/bin/bash

for FN in $@
do

    FN_NAME=$FN_handler
  echo "deploying $FN_NAME https://console.aws.amazon.com/lambda/home#/functions/$FN_NAME"


  # check if function exists, if not then skip
    aws --profile=onetwentyseven lambda get-function --function-name $FN_NAME > /dev/null 2> /dev/null
  if [ $? -ne 0 ]; then
    echo "$FN_NAME does not exist, skipping" && exit 1
  fi

  echo "deploying $FN_NAME https://console.aws.amazon.com/lambda/home#/functions/$FN_NAME"

  make deploy LAMBDA=$FN
  if [ $? -ne 0 ]; then
    echo "error deploying $FN" && exit 1;
  fi


done

echo -e '
 ___________________
< Deploy succeeded! >
 -------------------
   \
    \
        .--.
       |o_o |
       |:_/ |
      //   \ \
     (|     | )
    /"\_   _/"\
    \___)=(___/
'
#!/bin/bash

if [ $# -eq 0 ]; then
  echo "No parameters provided!"
  exit 1
fi

VAL1=$1

if [ $VAL1 = "git" ]; then
  USER_NAME="HyunsuChoi-git"
  USER_EMAIL="choihs1054@gmail.com"
elif [ $VAL1 = "bit" ]; then
  USER_NAME="hyunsuchoi"
  USER_EMAIL="hs.choi@okestro.com"
else
  echo "param error!"
  exit 1
fi


git config --global user.name ${USER_NAME}
git config --global user.email ${USER_EMAIL}

git config --global user.name
git config --global user.email
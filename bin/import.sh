#!/bin/bash

./oaem-RsltIndiv -bdd=reset
./oaem-RsltIndiv -c=import
./oaem-RsltIndiv -e=check
./oaem-RsltIndiv -r=import
./oaem-RsltIndiv -r=export
./oaem-RsltIndiv -r=team


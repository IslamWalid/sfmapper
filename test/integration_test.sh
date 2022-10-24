#!/bin/sh

go run main.go &
sleep 0.3

failure() {
    echo $1
    [[ $(mount | grep mnt) ]] && fusermount -zu ./mnt
    rmdir mnt
    exit 1
}

String=$(cat ./mnt/String)
if [[ $String != 'str' ]]; then
    failure 'TEST FAILED: file "String" does not match struct value'
fi

Int=$(cat ./mnt/Int)
if [[ $Int != '18' ]]; then
    failure 'TEST FAILED: file "Int" does not match struct value'
fi

Bool=$(cat ./mnt/Bool)
if [[ $Bool != 'true' ]]; then
    failure 'TEST FAILED: file "Bool" does not match struct value'
fi

find ./mnt/SubStructure >> /dev/null
if [[ $? != 0 ]]; then
    failure 'TEST FAILED: dir "SubStructure" does not exist'
fi

Float=$(cat ./mnt/SubStructure/Float)
if [[ $Float != 1.3 ]]; then
    failure 'TEST FAILED: file "Float" does not match struct value'
fi

sleep 1
String=$(cat ./mnt/String)
if [[ $String != 'new string' ]]; then
    failure 'TEST FAILED: file "String" was not modified'
fi

fusermount -zu ./mnt
rmdir mnt
echo 'TEST PASSED'

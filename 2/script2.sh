#!/bin/bash
input="input.txt"
horizontal=0
depth=0
aim=0
while IFS=" " read -r -a line
do
    word=${line[0]}
    number=${line[1]//[$'\t\r\n']}
    if [[ $word == "forward" ]]
    then
    let "horizontal+=number"
    ((multiplier=aim*number))
    let "depth+=multiplier"
    elif [[ $word == "up" ]]
    then
    let "aim-=number"    
    elif [[ $word == "down" ]] 
    then
    let "aim+=number"
    else
        echo $number
    fi
done < "$input"
echo $horizontal
echo $depth
echo "$((horizontal*depth))"
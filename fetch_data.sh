#!/usr/bin/env bash

[ -f data/ex2/ex2x.dat ] || {
    mkdir -p data/ex2
    f=`mktemp`
    curl -o $f http://openclassroom.stanford.edu/MainFolder/courses/MachineLearning/exercises/ex2materials/ex2Data.zip 
    unzip -d data/ex2 $f
    rm $f
}

[ -f data/ex2/ex2.m ] || {
    curl -o data/ex2/ex2.m http://openclassroom.stanford.edu/MainFolder/courses/MachineLearning/exercises/ex2materials/ex2.m
}

[ -f data/ex3/ex3x.dat ] || {
    mkdir -p data/ex3
    f=`mktemp`
    curl -o $f http://openclassroom.stanford.edu/MainFolder/courses/MachineLearning/exercises/ex3materials/ex3Data.zip
    unzip -d data/ex3 $f
    rm $f
}

[ -f data/ex3/ex3.m ] || {
    curl -o data/ex3/ex3.m http://openclassroom.stanford.edu/MainFolder/courses/MachineLearning/exercises/ex3materials/ex3.m
}

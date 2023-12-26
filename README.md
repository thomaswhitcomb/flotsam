# flotsam

## Overview

This is the main that uses the github.com/thomaswhitcomb/jetsom package to create a pipeline for reading CSV files and reducing the CSVs to two metrics.

The CSVs each have three columns:

- First Name
- Second Name
- Age

For the set of CSVs pulled through the jetsom pipeline, the flotsam reducer computes the average age and medium age.

The flotsam executable takes 1 parameter:

- **n** - the number of concurrent CSV loaders.

## Build

    > make

## Run

    > cat test_http.txt | ./bin/flotsam -n 3

# Apisearch - Juicer

This library is part of the Apisearch project.

[![Build Status](https://travis-ci.org/apisearch-io/juicer.svg?branch=master)](https://travis-ci.org/apisearch-io/juicer)
[![Join the chat](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/apisearch_io/general)

This is the official Apisearch Killer. We should be able to write the
most aggressive Apisearch bot, and even that, Apisearch should keep
working properly.

What is this script actually doing?

- Creates an incresciendo group of users hitting the website each 100 milliseconds
- Each second, more users enter the action. The more users ATST, the more goroutines
  are managed in parallel
- Each second, all requests captured the second before are analyzed and printed in
  a new line.
- Each line prints
    - The average response time
    - The minimum and maximum response time borders
    - The number of requests finished in this second
    - The average memory used by the server

[ 0ms] [}-------------------------------------------------] [  0 u] [mem:       0]  
[ 4ms] [={==---------------}------------------------------] [ 44 u] [mem:15384672]  
[ 3ms] [={=---------------}-------------------------------] [ 72 u] [mem:15384672]  
[ 4ms] [={==--------------------------}-------------------] [101 u] [mem:15384672]  
[ 3ms] [={=--------------}--------------------------------] [138 u] [mem:15384672]  
[ 3ms] [={=--------------------}--------------------------] [170 u] [mem:15384672]  
[ 2ms] [={-----------}------------------------------------] [200 u] [mem:15384672]  
[ 3ms] [={=---------------}-------------------------------] [234 u] [mem:15384672]  
[ 2ms] [={----------------}-------------------------------] [266 u] [mem:15404040]  
[ 3ms] [={=---------------------------}-------------------] [297 u] [mem:15384672]  
[ 3ms] [={=-------------------------}---------------------] [334 u] [mem:15384672]  
[ 3ms] [={=---------------------}-------------------------] [362 u] [mem:15384672]  
[ 4ms] [={==------------------------------}---------------] [386 u] [mem:15420032]  
[ 3ms] [={=--------------------}--------------------------] [428 u] [mem:15384672]  
[ 4ms] [={==-----------------------}----------------------] [449 u] [mem:15384672]  
[ 6ms] [={====-------------------------}------------------] [475 u] [mem:15420032]  
[ 5ms] [={===--------------------------}------------------] [512 u] [mem:15400664]  
[ 8ms] [={======--------------------------}---------------] [530 u] [mem:15420032]  
[ 8ms] [={======---------------------------}--------------] [554 u] [mem:15400664]  
[15ms] [========{======------------------------}----------] [555 u] [mem:15420032]  
[19ms] [==========={=======----------------------}--------] [562 u] [mem:15420032]  
[27ms] [================={=========----------------}------] [552 u] [mem:15420032]  
[22ms] [================={====-------------------}--------] [570 u] [mem:15420032]  
[24ms] [================={======--------------------}-----] [559 u] [mem:15420032]  
[31ms] [================={=============----------------->>] [533 u] [mem:15420032]  
[24ms] [================{=======---------------------}----] [561 u] [mem:15420032]  
[24ms] [=============={=========--------------------}-----] [565 u] [mem:15420032]  
[25ms] [=================={======-------------------}-----] [555 u] [mem:15420032]  
[ 4ms] [={==----------------------------------}-----------] [505 u] [mem:15400664]  
[ 2ms] [={----------------}-------------------------------] [331 u] [mem:15384672]  

Use your command line to execute the script properly. No Golang needed to execute it.

> The script attacks the port 8100 of your localhost, so make sure you have an stable
> Apisearch instance running in that port (regular distribution has).

Before executing this script, we need to create a valid index in our server. ATM, this
script will not test how heavy can be elasticsearch due to abusive and long queries, but
how the server responds in front of several requests per second, so an empty index will
be enough.

```bash
curl -XPOST "http://localhost:8100/v1/index?app_id=golang&index=test&token=0e4d75ba-c640-44c1-a745-06ee51db4e93
```

> This script works with the superadmin token distributed by the main Apisearch Server
> repository, with only testing purposes.


```bash
./juicer
```
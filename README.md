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

This would be an example of what this script would return

``` bash
[ 0ms] [}-------------------------------------------------] [  0 c] [mem:       0]  
[ 4ms] [={==---------------}------------------------------] [ 44 c] [mem:15384672]  
[ 3ms] [={=---------------}-------------------------------] [ 72 c] [mem:15384672]  
[ 4ms] [={==--------------------------}-------------------] [101 c] [mem:15384672]  
[ 3ms] [={=--------------}--------------------------------] [138 c] [mem:15384672]  
[ 3ms] [={=--------------------}--------------------------] [170 c] [mem:15384672]  
[ 2ms] [={-----------}------------------------------------] [200 c] [mem:15384672]  
[ 3ms] [={=---------------}-------------------------------] [234 c] [mem:15384672]  
[ 2ms] [={----------------}-------------------------------] [266 c] [mem:15404040]  
[ 3ms] [={=---------------------------}-------------------] [297 c] [mem:15384672]  
[ 3ms] [={=-------------------------}---------------------] [334 c] [mem:15384672]  
[ 3ms] [={=---------------------}-------------------------] [362 c] [mem:15384672]  
[ 4ms] [={==------------------------------}---------------] [386 c] [mem:15420032]  
[ 3ms] [={=--------------------}--------------------------] [428 c] [mem:15384672]  
[ 4ms] [={==-----------------------}----------------------] [449 c] [mem:15384672]  
[ 6ms] [={====-------------------------}------------------] [475 c] [mem:15420032]  
[ 5ms] [={===--------------------------}------------------] [512 c] [mem:15400664]  
[ 8ms] [={======--------------------------}---------------] [530 c] [mem:15420032]  
[ 8ms] [={======---------------------------}--------------] [554 c] [mem:15400664]  
[15ms] [========{======------------------------}----------] [555 c] [mem:15420032]  
[19ms] [==========={=======----------------------}--------] [562 c] [mem:15420032]  
[27ms] [================={=========----------------}------] [552 c] [mem:15420032]  
[22ms] [================={====-------------------}--------] [570 c] [mem:15420032]  
[24ms] [================={======--------------------}-----] [559 c] [mem:15420032]  
[31ms] [================={=============----------------->>] [533 c] [mem:15420032]  
[24ms] [================{=======---------------------}----] [561 c] [mem:15420032]  
[24ms] [=============={=========--------------------}-----] [565 c] [mem:15420032]  
[25ms] [=================={======-------------------}-----] [555 c] [mem:15420032]  
[ 4ms] [={==----------------------------------}-----------] [505 c] [mem:15400664]  
[ 2ms] [={----------------}-------------------------------] [331 c] [mem:15384672]  
```

Use your command line to execute the script properly. No Golang needed to execute it.  
In order to use it, create a valid urls.yml files with your testable urls. You can
take a template file from `urls.yml.dist`. The script will always take the first
url as the health check one.

``` yml
- '/health?token=xxx'
- '/v1?app_id=xxx&index=yyy&token=zzz&query={"q": "hello"}'
- '/v1?app_id=xxx&index=yyy&token=zzz&query={"q": "bye", "fields": ["metadata.*"]}'
```

Once defined your urls, just execute the script with these parameters (order matters)

- The Apisearch endpoint
- Number of concurrent users
- Number of seconds for the test

```bash
./juicer 0.0.0.0:8000 10 60
```

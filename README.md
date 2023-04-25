## httpHasher

### description
A demo for study</p>
This is a tool which makes http requests and prints the address of the request along with the MD5 hash of the response.



### how to use

get golang installed and clone the repository execute `go build` under the working directory

this tool accept a parallel parameter to set the upper limit of the concurrent tasks, which default value is 10

example: 

```
httpHasher -parallel 3 http://www.aaa.com www.google.com www.bing.com

httpHasher http://www.aaa.com www.google.com www.bing.com

```









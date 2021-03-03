# Build

All build-related scripts are placed in the "build" folder

1. ./build.sh - build app locally. It will be saved in dist/ folder
2. ./build_docker_image.sh - build docker image. To run image:
   docker run -p 8080:8080 -i -t detector 
   
# How to use it

There are 2 available http-calls:
- POST request to /task, which is expecting json-input: ["host1", "host2"]
- GET request to /task/:uuid which doesn't expect input. It returns the result
of execution for command above
  
I didn't implement immediate response for POST /task request. As it uses http 
connection to external resources so to save server resources I did only delayed response.

# How to scale it

So far I see 2 preferable ways to do that:
1. Set kubernetes deployment with ingress-nginx as reverse proxy which will send requests
to the services. To storage data I would use  MongoDB. Cons: 1) this system
   isn't fault tolerant. 
   
2. I would write a master node. Which do load balancing. It also would check status of tasks
and re-schedule in case of any worker node crashes
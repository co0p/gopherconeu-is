production readyness
--------------------

l1 hardware 
l2 comm network, service discoverty registry load balancing
l3 Application plattform - build deployment logging / monitoring
l4 applications configurations 

kuberentes
l1 controller 1... M Worker 1 ... N
l2 comm: etc, apicserver, scheduler, kubelet, kub-proxy, container runtime, ingress controlle
l3 helm, ci,cd monitoring, logging
l4 charts (config), application

ingress - routing 
helm - release manager for kubernetes

what means prod ready?
---------------------
stability, reliability, scalability, security, fault tolerance, monitoring, documentation, sla

12 factor apps
-------------
https://www.cdta.org/sites/default/files/awards/beyond_the_12-factor_app_pivotal.pdf

branching model
----------------
git flow
github flow
push to master, no tags with big teams? 
devops: http://blog.greglow.com/2018/04/27/devops-to-branch-or-not-to-branch/

go tooling action
-----------------

go get -> git clone to $GOPATH and compile 
gopath:
- .../bin/ executable binaries
- .../src/ other packages
- .../pkg/ compiles stuff
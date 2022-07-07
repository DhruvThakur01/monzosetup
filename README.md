# Monzo Egress Gateway Testing Setup

This repository contains all the files that were utilized in creating the testing setup for Monzo Egress Gateway. 

We worked with 2 AKS clusters:
- *server-aks-yg* for hosting the server application and its LoadBalancer services
- *monzo-aks-yg* for hosting both the gateway and non-gateway client applications and also for the monzo egress gateway 
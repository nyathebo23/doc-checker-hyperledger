Minikube Portion of the Readme
=====================================

## Kubernetes - Minikube (Local)
[Install Kubernetes and Minikube](https://kubernetes.io/docs/tasks/tools/)
[If OSX here is virtual box](https://www.virtualbox.org/wiki/Mac%20OS%20X%20build%20instructions)
[Kubernetes book](hhttps://www.amazon.com/Devops-2-3-Toolkit-Viktor-Farcic/dp/1789135508/ref=tmm_pap_swatch_0?_encoding=UTF8&sr=8-2)
[K8s Persistent Volumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/)

Okay, now that we've successfully ran the network locally, let's do this on a local kubernetes installation.
```bash
minikube start --memory=4g --cpus=4
sleep 5
kubectl apply -f minikube/storage/pvc.yaml
sleep 10
kubectl apply -f minikube/storage/tests
kubectl apply -f minikube/orgs/enspy/couchdb/peer0-couchdb-deployment.yaml
```

Now we have storage and we're going to test it. You can do a kubectl get pods to see what pods are up. Here's how I can connect to my containers. You should split your terminal and connect to both.
```bash
☁  k8s-hyperledger-fabric-2.2 [master] ⚡  kubectl get pods
NAME                        READY   STATUS    RESTARTS   AGE
example1-6858b4f776-5pgls   1/1     Running   0          17s
example1-6858b4f776-q92vv   1/1     Running   0          17s
example2-55fcbb9cbd-drzwn   1/1     Running   0          17s
example2-55fcbb9cbd-sv4c8   1/1     Running   0          17s
☁  k8s-hyperledger-fabric-2.2 [master] ⚡ 
```

We'll use one of these to setup the files for the network
```bash
kubectl exec -it $(kubectl get pods -o=name | grep example1 | sed "s/^.\{4\}//") -- mkdir -p /host/files/scripts
kubectl exec -it $(kubectl get pods -o=name | grep example1 | sed "s/^.\{4\}//") -- mkdir -p /host/files/chaincode
sleep 1
kubectl cp ./scripts $(kubectl get pods -o=name | grep example1 | sed "s/^.\{4\}//"):/host/files/
kubectl cp ./minikube/configtx.yaml $(kubectl get pods -o=name | grep example1 | sed "s/^.\{4\}//"):/host/files
kubectl cp ./minikube/config.yaml $(kubectl get pods -o=name | grep example1 | sed "s/^.\{4\}//"):/host/files
kubectl cp ./chaincode/organization $(kubectl get pods -o=name | grep example1 | sed "s/^.\{4\}//"):/host/files/chaincode
kubectl cp ./chaincode/beneficiary $(kubectl get pods -o=name | grep example1 | sed "s/^.\{4\}//"):/host/files/chaincode
kubectl cp ./chaincode/document_save $(kubectl get pods -o=name | grep example1 | sed "s/^.\{4\}//"):/host/files/chaincode
kubectl cp ../../fabric-samples/bin $(kubectl get pods -o=name | grep example1 | sed "s/^.\{4\}//"):/host/files
```

Let's bash into the container and make sure everything copied over properly
```bash
kubectl exec -it $(kubectl get pods -o=name | grep example1 | sed "s/^.\{4\}//") bash
```

Finally ready to start the ca containers
```bash
kubectl apply -f minikube/cas/orderers-ca-deployment.yaml
kubectl apply -f minikube/cas/enspy-ca-deployment.yaml
kubectl apply -f minikube/cas/enspd-ca-deployment.yaml

sleep 30

kubectl apply -f minikube/cas
```

Your containers should be up and running. You can check the logs like so and it should look liek this.
```bash
☁  k8s-hyperledger-fabric-2.2 [master] ⚡  kubectl logs -f orderers-ca-d69cbc664-dzk4f
2020/12/11 04:12:37 [INFO] Created default configuration file at /etc/hyperledger/fabric-ca-server/fabric-ca-server-config.yaml
2020/12/11 04:12:37 [INFO] Starting server in home directory: /etc/hyperledger/fabric-ca-server
...
2020/12/11 04:12:38 [INFO] generating key: &{A:ecdsa S:256}
2020/12/11 04:12:38 [INFO] encoded CSR
2020/12/11 04:12:38 [INFO] signed certificate with serial number 307836600921505839273746385963411812465330101584
2020/12/11 04:12:38 [INFO] Listening on https://0.0.0.0:7054
```

This should generate the crypto-config files necessary for the network. You can check on those files in any of the containers.
```bash
root@example1-6858b4f776-wmlth:/host# cd files
root@example1-6858b4f776-wmlth:/host/files# ls
bin  chaincode	config.yaml  configtx.yaml  crypto-config  scripts
root@example1-6858b4f776-wmlth:/host/files# cd crypto-config/
root@example1-6858b4f776-wmlth:/host/files/crypto-config# ls
ordererOrganizations  peerOrganizations
root@example1-6858b4f776-wmlth:/host/files/crypto-config# cd peerOrganizations/
root@example1-6858b4f776-wmlth:/host/files/crypto-config/peerOrganizations# ls
enspy  enspd
root@example1-6858b4f776-wmlth:/host/files/crypto-config/peerOrganizations# cd enspy/
root@example1-6858b4f776-wmlth:/host/files/crypto-config/peerOrganizations/enspy# ls
msp  peers  users
root@example1-6858b4f776-wmlth:/host/files/crypto-config/peerOrganizations/enspy# cd msp/
root@example1-6858b4f776-wmlth:/host/files/crypto-config/peerOrganizations/enspy/msp# ls
IssuerPublicKey  IssuerRevocationPublicKey  admincerts	cacerts  keystore  signcerts  tlscacerts  user
root@example1-6858b4f776-wmlth:/host/files/crypto-config/peerOrganizations/enspy/msp# cd tlscacerts/
```

Time to generate the artifacts inside one of the containers and in the files folder - NOTE: if you are on OSX you might have to load the proper libs `curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.2.1 1.4.7` (apt update then apt install curl) (you will also need to cp from (this path will depend on where you are in the container) cp bin/* /host/files/bin)
```bash
kubectl exec -it $(kubectl get pods -o=name | grep example1 | sed "s/^.\{4\}//") bash
...
cd /host/files

rm -rf orderer channels
mkdir -p orderer channels
bin/configtxgen -profile OrdererGenesis -channelID syschannel -outputBlock ./orderer/genesis.block
bin/configtxgen -profile MainChannel -outputCreateChannelTx ./channels/mainchannel.tx -channelID mainchannel
bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./channels/enspy-anchors.tx -channelID mainchannel -asOrg enspy
bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./channels/enspd-anchors.tx -channelID mainchannel -asOrg enspd
```

Let's try to start up the orderers
```bash
kubectl apply -f minikube/orderers/orderer0-deployment.yaml

kubectl apply -f minikube/orderers
```

Go ahead and check the logs and see that the orderers have selected a leader like so
```bash
 1 became follower at term 2 channel=syschannel node=1
2020-12-11 05:20:15.616 UTC [orderer.consensus.etcdraft] Step -> INFO 029 1 [logterm: 1, index: 3, vote: 0] cast MsgVote for 2 [logterm: 1, index: 3] at term 2 channel=syschannel node=1
2020-12-11 05:20:15.634 UTC [orderer.consensus.etcdraft] run -> INFO 02a raft.node: 1 elected leader 2 at term 2 channel=syschannel node=1
2020-12-11 05:20:15.639 UTC [orderer.consensus.etcdraft] run -> INFO 02b Raft leader changed: 0 -> 2 channel=syschannel node=1
```

We should be able to start the peers now
```bash
kubectl apply -f minikube/orgs/enspy/couchdb 
kubectl apply -f minikube/orgs/enspd/couchdb

kubectl apply -f minikube/orgs/enspy/
kubectl apply -f minikube/orgs/enspd/

kubectl apply -f minikube/orgs/enspy/cli
kubectl apply -f minikube/orgs/enspd/cli
kubectl delete -f minikube/storage/tests
kubectl delete -f minikube/cas/enspy-ca-client-deployment.yaml
kubectl delete -f minikube/cas/enspd-ca-client-deployment.yaml
kubectl delete -f minikube/cas/orderers-ca-client-deployment.yaml
```

NOTE: you can stop the cas if you don't need them anymore (don't do this if you want to continue making certs later)
- minikube only has so many resources so sometimes when testing you might need to decide what containers are more important
```bash
kubectl delete -f minikube/cas
```


Time to actually test the network
```bash
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer channel create -c mainchannel -f ./channels/mainchannel.tx -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'cp mainchannel.block ./channels/'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer channel join -b channels/mainchannel.block'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer channel join -b channels/mainchannel.block'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer channel join -b channels/mainchannel.block'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer channel join -b channels/mainchannel.block'

sleep 5

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer channel update -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem -c mainchannel -f channels/enspy-anchors.tx'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer channel update -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem -c mainchannel -f channels/enspd-anchors.tx'
```

Now we are going to install the chaincode - NOTE: Make sure you go mod vendor in each chaincode folder... might need to remove the go.sum depending
```bash
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package beneficiary.tar.gz --path /opt/gopath/src/beneficiary --lang golang --label beneficiary_1'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package beneficiary.tar.gz --path /opt/gopath/src/beneficiary --lang golang --label beneficiary_1'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package beneficiary.tar.gz --path /opt/gopath/src/beneficiary --lang golang --label beneficiary_1'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package beneficiary.tar.gz --path /opt/gopath/src/beneficiary --lang golang --label beneficiary_1'

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install beneficiary.tar.gz &> pkg.txt'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install beneficiary.tar.gz'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install beneficiary.tar.gz &> pkg.txt'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install beneficiary.tar.gz'

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode approveformyorg -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem --channelID mainchannel --collections-config /opt/gopath/src/beneficiary/collections-config.json --name beneficiary --version 1.0 --sequence 1 --package-id $(tail -n 1 pkg.txt | awk '\''NF>1{print $NF}'\'')'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode approveformyorg -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem --collections-config /opt/gopath/src/beneficiary/collections-config.json --channelID mainchannel --name beneficiary --version 1.0 --sequence 1 --package-id $(tail -n 1 pkg.txt | awk '\''NF>1{print $NF}'\'')'

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode commit -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem --channelID mainchannel --collections-config /opt/gopath/src/beneficiary/collections-config.json --name beneficiary --version 1.0 --sequence 1'


```
Lets go ahead and test this chaincode
```bash


kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode invoke -C mainchannel -n beneficiary -c '\''{"Args":["Create", "1","Nyatchou", "Franck", "franckhebo@gmail.com", "Jonathan23"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'
sleep 5
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode query -C mainchannel -n beneficiary -c '\''{"Args":["Index"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'



```
Chaincode for organization
```bash


kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package organization.tar.gz --path /opt/gopath/src/organization --lang golang --label organization_1'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package organization.tar.gz --path /opt/gopath/src/organization --lang golang --label organization_1'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package organization.tar.gz --path /opt/gopath/src/organization --lang golang --label organization_1'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package organization.tar.gz --path /opt/gopath/src/organization --lang golang --label organization_1'

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install organization.tar.gz &> pkg.txt'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install organization.tar.gz'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install organization.tar.gz &> pkg.txt'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install organization.tar.gz'

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode approveformyorg -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem --channelID mainchannel --collections-config /opt/gopath/src/organization/collections-config.json --name organization --version 1.0 --sequence 1 --package-id $(tail -n 1 pkg.txt | awk '\''NF>1{print $NF}'\'')'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode approveformyorg -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem --collections-config /opt/gopath/src/organization/collections-config.json --channelID mainchannel --name organization --version 1.0 --sequence 1 --package-id $(tail -n 1 pkg.txt | awk '\''NF>1{print $NF}'\'')'

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode commit -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem --channelID mainchannel --collections-config /opt/gopath/src/organization/collections-config.json --name organization --version 1.0 --sequence 1'


```
Lets go ahead and test this chaincode
```bash

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode invoke -C mainchannel -n organization -c '\''{"Args":["Create", "2","ENSPD", "Ecole Nationale Polytechnique", "true"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'
sleep 5
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode query -C mainchannel -n organization -c '\''{"Args":["Index"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'
```





Lets try the other chaincode
```bash
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package document_save.tar.gz --path /opt/gopath/src/document_save --lang golang --label document_save_1'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package document_save.tar.gz --path /opt/gopath/src/document_save --lang golang --label document_save_1'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package document_save.tar.gz --path /opt/gopath/src/document_save --lang golang --label document_save_1'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode package document_save.tar.gz --path /opt/gopath/src/document_save --lang golang --label document_save_1'


kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install document_save.tar.gz &> pkg.txt'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install document_save.tar.gz'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install document_save.tar.gz &> pkg.txt'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode install document_save.tar.gz'

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode approveformyorg -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem --channelID mainchannel --collections-config /opt/gopath/src/document_save/collections-config.json --name document_save --version 1.0 --sequence 1 --package-id $(tail -n 1 pkg.txt | awk '\''NF>1{print $NF}'\'')'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode approveformyorg -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem --channelID mainchannel --collections-config /opt/gopath/src/document_save/collections-config.json --name document_save --version 1.0 --sequence 1 --package-id $(tail -n 1 pkg.txt | awk '\''NF>1{print $NF}'\'')'

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer lifecycle chaincode commit -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem --channelID mainchannel --collections-config /opt/gopath/src/document_save/collections-config.json --name document_save --version 1.0 --sequence 1'

sleep 5

kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode invoke -C mainchannel -n document_save -c '\''{"Args":["Create","","CPUs","1"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode invoke -C mainchannel -n document_save -c '\''{"Args":["Create","","Database Servers","1"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode invoke -C mainchannel -n document_save -c '\''{"Args":["Create","","Mainframe Boards","1"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'
sleep 5
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode query -C mainchannel -n document_save -c '\''{"Args":["Index"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode query -C mainchannel -n document_save -c '\''{"Args":["Index"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode query -C mainchannel -n document_save -c '\''{"Args":["Index"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer1-enspd-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode query -C mainchannel -n document_save -c '\''{"Args":["Index"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'


kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode invoke -C mainchannel -n document_save -c '\''{"Args":["Create", "1","1", "1", "filepoh.pdf", "tyre44scsds545ds54fdsfjdsjhj"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'
sleep 5
kubectl exec -it $(kubectl get pods -o=name | grep cli-peer0-enspy-deployment | sed "s/^.\{4\}//") -- bash -c 'peer chaincode query -C mainchannel -n document_save -c '\''{"Args":["Index"]}'\'' -o orderer0-service:7050 --tls --cafile=/etc/hyperledger/orderers/msp/tlscacerts/orderers-ca-service-7054.pem'
```

Start the API
```bash
kubectl apply -f network/minikube/backend
```

Get the address for the nodeport
```bash
minikube service api-service-nodeport --url
```

## How to delete everything
```bash
kubectl delete -f network/minikube/backend
kubectl delete -f network/minikube/orgs/enspy/couchdb 
kubectl delete -f network/minikube/orgs/enspd/couchdb

kubectl delete -f network/minikube/orgs/enspy/
kubectl delete -f network/minikube/orgs/enspd/

kubectl delete -f network/minikube/orgs/enspy/cli
kubectl delete -f network/minikube/orgs/enspd/cli
kubectl delete -f network/minikube/cas
kubectl delete -f network/minikube/orderers
```

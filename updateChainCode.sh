version=$1;

echo "Installing siam chaincode to peer0.manufacturer.siam.com..."
docker exec -e "CORE_PEER_LOCALMSPID=manufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.siam.com/users/Admin@manufacturer.siam.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.siam.com:7051" cli peer chaincode install -n siamcc -v $version -p github.com/siam/go -l golang
echo "Installed siam chaincode to peer0.manufacturer.siam.com"

echo "Installing siam chaincode to peer0.rto.siam.com...."
docker exec -e "CORE_PEER_LOCALMSPID=rtoMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/rto.siam.com/users/Admin@rto.siam.com/msp" -e "CORE_PEER_ADDRESS=peer0.rto.siam.com:7051" cli peer chaincode install -n siamcc -v $version -p github.com/siam/go -l golang
echo "Installed siam chaincode to peer0.rto.siam.com"

echo "Installing siam chaincode to peer0.owner.siam.com..."
docker exec -e "CORE_PEER_LOCALMSPID=ownerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/owner.siam.com/users/Admin@owner.siam.com/msp" -e "CORE_PEER_ADDRESS=peer0.owner.siam.com:7051" cli peer chaincode install -n siamcc -v $version -p github.com/siam/go -l golang
echo "Installed siam chaincode to peer0.rto.siam.com"

echo "Installing siam chaincode to peer0.scrp.siam.com..."
docker exec -e "CORE_PEER_LOCALMSPID=scrpMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/scrp.siam.com/users/Admin@scrp.siam.com/msp" -e "CORE_PEER_ADDRESS=peer0.scrp.siam.com:7051" cli peer chaincode install -n siamcc -v $version -p github.com/siam/go -l golang
echo "Installed siam chaincode to peer0.scrp.siam.com"



sleep 5


echo "Updating siam chaincode.."

docker exec -e "CORE_PEER_LOCALMSPID=manufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.siam.com/users/Admin@manufacturer.siam.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.siam.com:7051" cli peer chaincode upgrade -o orderer.siam.com:7050 -C siamchannel -n siamcc -l golang -v $version -c '{"Args":[""]}' -P "OR ('manufacturerMSP.member','rtoMSP.member','ownerMSP.member')"

echo "Instantiated siam chaincode."

echo "Following is the docker network....."

docker ps

# Vehicle LifeCycle  on Hyperledger Fabric (*** Only CreateCar,RegisterCar,GetCarHistory,GetCarDetails and ScrapCar implemented ***)
Vehicle application on Hyperledger Fabric

*** Use sudo prefix to commands if you get permission denied error while executing any command, assumption is you already have  required software to run Hyperledger fabric network and node SDK *** 

## Start the Hyperledger Fabric Network 

1. cd vehicle-lc
2. ./start.sh (with this you will start docker-compose.yml up -d )

## Setup the Hyperledger Fabric Network

1. cd vehicle-lc
2. ./setup.sh (With this you will create the channel genesis block, add the peer0 to the channel created and instantiate tfbc chaincode.) 

*** In this usecase CA's are already generated. 

We **do not have to run** the following again:

1. "generate --config=crypto-config.yaml"
2. "siamOrgOrdererGenesis -outputBlock ./config/genesis.block" 
3. "siamOrgChannel -outputCreateChannelTx ./config/siamchannel.tx -channelID siamchannel". 

These three statements are part of the "generate.sh" file here. 


## Setup API users 

1. cd vehicle-lc/siam-api
2. npm install
3. rm hfc-key-store/*
4. node enrollBankUser.js
5. node enrollBuyerUser.js
6. node enrollSellerUser.js

## Run Node APIs  
1. node createCar.js [see chain code for argument]
2. node registerCar.js [see chain code for argument]
3. node getCarDetails [CHASIS NUMBER]
4. node getCarHistory [REGISTRATION NUMBER]
5. node scrapCar.js [REGISTRATION NUMBER] [Not tested]

1. cd vehicle-lc
2. ./stop.sh

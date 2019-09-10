rm -R crypto-config/*

./bin/cryptogen generate --config=crypto-config.yaml

rm config/*

./bin/configtxgen -profile SIAMOrgOrdererGenesis -outputBlock ./config/genesis.block

./bin/configtxgen -profile SIAMOrgChannel -outputCreateChannelTx ./config/siamchannel.tx -channelID siamchannel

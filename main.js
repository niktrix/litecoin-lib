
var express = require('express');
const app = express();
var explorers = require('bitcore-explorers');
const bitcore = require('bitcore-lib');
var insight = new explorers.Insight();
var insightTestnet = new explorers.Insight('testnet');
var BN = require('bn.js');


insightTestnet.getUnspentUtxos("myyyjh1D3P592vCa5JcJ5Kt19YTrrChM9y", function(err, utxos) {
      if (err) {
        next(err)
        console.log(err)
        res.send('UTXO detail not found');
      } else {
          const to = "mnjCg32kEHQZ8XSGJr9r7UJvwzAoZcxeze"
          const amount = 10000
          const fee = 5000


         const privateKey = new bitcore.PrivateKey("5JsjKubviP3TDfNfbE3qdxKuNqqSVCctEF3jzyw26qYzonGEgsE");// nicky
    // const privateKey = new bitcore.PrivateKey("cPmTTa8ctUckw7KYppLd1Vkx7jxjjRcpMqZ6Dm4n2FVfXRBRyirL");

        const fromAddress = privateKey.toAddress();
        console.log("From Address: "+fromAddress)
        console.log('To bitcoin address: '+ to)
        console.log('utxos: '+ utxos)
        // Create TX
        let tx = bitcore.Transaction();
            tx.from(utxos);
            tx.to(to, parseInt(amount));
            tx.change(fromAddress);
            tx.fee(fee);
            tx.sign(privateKey);

            try{
                tx.serialize();

            }catch(err){
                console.log(err)
            }

            console.log('tx: '+ tx)
            // Broadcast your transaction to the Bitcoin network

            console.log("tx.toString()",tx.toString())

        //     insightTestnet.broadcast(tx.toString(), (error, txid) => {
        //       if (error) {
        //           console.log("--",error)
        //       } else {
        //         // Bitcoin Transaction Id
        //           console.log(txid)
        //       }
        //   })
      }
    });

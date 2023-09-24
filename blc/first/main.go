package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const infraURL = "https://mainnet.infura.io/v3/771bdddaf5c74759ac261ddbf88a8b0a"
const txHash = "0x9fc7d10bb387aa735c4d00daa4aa8df71e697175a157f190ff1f35b8fb924cb8"

func main() {
	client, _ := connectToEthereum(infraURL)
	checkTrx(client, txHash)
	privateKeyByte := GeneratePrivateKey()

	_ = GetBalance(client, privateKeyByte)
	CheckContract(client, txHash)
}

func connectToEthereum(infraUrl string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(infraUrl)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return client, nil
}

func checkTrx(conn *ethclient.Client, txHash string) {
	cntx := context.Background()
	txn, pending, _ := conn.TransactionByHash(cntx, common.HexToHash(txHash))
	if pending {
		fmt.Println("txn is Pending ", txn)
	}
	fmt.Println("txn is not pending ", txn)
}

func GeneratePrivateKey() []byte {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal("Unable to generate Key")
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("private key bytes: " + hexutil.Encode(privateKeyBytes))
	return privateKeyBytes
}

func GetBalance(client *ethclient.Client, privateKeyByte []byte) error {
	account := common.HexToAddress(string(privateKeyByte))
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal("Unable to get balance")
		return err
	}
	fmt.Println("Balance: " + balance.String())
	return nil
}

func CheckContract(client *ethclient.Client, txContract string) error {
	regEx := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !regEx.MatchString(txContract) {
		log.Fatal("Wrong txContract!")
	}
	address := common.HexToAddress(txContract)
	byteCode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal("unable to check if its code at ", address)
		return err
	}
	isContract := len(byteCode) > 0
	fmt.Println(address, "is contract", isContract)
	return nil
}

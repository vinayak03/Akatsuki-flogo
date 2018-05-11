# Hyperledger-Fabric-ChainCode Flogo Activity
This activity allows you to execute and query the fabric chain code(smart contract)

## Pre-requisite
### Sample Hyperledger Network Setup
Note: Run the below steps in powershell on windows
#### Download Hyperledger network images
	Inside the directory [a relative link]../../fabric-setup/network/
	sh bootstrap.sh 1.1.0

#### Start a basic network with fabcar chaincode
	Inside directory [a relative link]../../fabric-setup/network/
	sh start.sh

#### Configure users in the network to users
	Inside [a relative link]../../fabric-setup/script
	//node dependency installtion
	npm install

	//Get the default admin user from ca
	node enrollAdmin.js

	//register the user to be used with the chaincode example.
	node registerUser.js

## Installation
###Flogo-webui
	Inside directory [a relative link]../../fabric-setup/network/

	sh flogo-activity-install.sh
	
	if the web-ui is already running it would need a restart.
### Flogo-cli
```
flogo add activity github.com/vinayak03/Akatsuki-flogo/activity/hyperledgerFabricChainCode
flogo ensure -add github.com/hyperledger/fabric-sdk-go@ba0e035b4a43d42233f0b40a3a35d0eb0e3a1c98
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "NetworkConfig",
      "type": "string"
    },
	{
      "name": "RequestType",
      "type": "string",
	  "required": true,
      "allowed" : ["Query", "Execute"]
    },
	{
      "name": "ChannelID",
      "type": "string"
    },
	{
      "name": "User",
      "type": "string"
    },
    {
      "name": "UserPasswd",
      "type": "string"
    },
    {
      "name": "UserOrg",
      "type": "string"
    },
	{
      "name": "ChainCodeID",
      "type": "string"
    },
	{
      "name": "FunctionName",
      "type": "string"
    },
	{
      "name": "Params",
      "type": "array"
    }
  ],
  "outputs": [
    {
      "name": "Status",
      "type": "string"
    },
    {
      "name": "ResponsePayload",
      "type": "string"
    },
    {
      "name": "ErrorMessage",
      "type": "string"
    }
    ,
    {
      "name": "TransactionID",
      "type": "string"
    }
    
  ]
}
```
## Settings
| Setting            | Required | Description |
|:-------------------|:---------|:------------|
| NetworkConfig      | True     | yaml configuration file describing the network and artifacts for network.
| RequestType        | False    | Query/Execute Operation
| ChannelID          | False    | On which channel the transaction/operation needs to be completed.
| User               | False    | User to be used for doing the transactio
| UserPasswd         | True     | Password for user
| UserOrg            | False    | Org for which the user belongs
| ChainCodeID        | False    | the chaincode which will be executed eg. fabcar
| FunctionName       | False    | name of the function in chaincode
| Params             | False    | the parameter array of strings , every function can have verying arguments.

## Examples
The below example chaincode queries 'queryAllCars' function:

```json
{
	"id": "hyperledgerFabricChainCode_2",
	"name": "hyperledgerFabricChainCode",
	"description": "Plugin to Invoke and Query Hyperledger Smart Contracts(ChainCodes)",
	"type": 1,
	"activityType": "hyperledgerFabricChainCode",
	"activityRef": "github.com/vinayak03/Akatsuki-flogo/activity/hyperledgerFabricChainCode",
	"attributes": [{
		"name": "NetworkConfig",
		"value": "/etc/hyperledger/client/config.yaml",
		"required": true,
		"type": "string"
	},
	{
		"name": "RequestType",
		"value": "Query",
		"required": true,
		"type": "string"
	},
	{
		"name": "ChannelID",
		"value": "mychannel",
		"required": true,
		"type": "string"
	},
	{
		"name": "User",
		"value": "user1",
		"required": true,
		"type": "string"
	},
	{
		"name": "UserPasswd",
		"value": "user1pw",
		"required": true,
		"type": "string"
	},
	{
		"name": "UserOrg",
		"value": "Org1",
		"required": true,
		"type": "string"
	},
	{
		"name": "ChainCodeID",
		"value": "fabcar",
		"required": true,
		"type": "string"
	},
	{
		"name": "FunctionName",
		"value": "queryAllCars",
		"required": true,
		"type": "string"
	},
	{
		"name": "Params",
		"value": "{\"params\" : []}",
		"required": true,
		"type": "any"
	}]
}
```

The below example chaincode queries 'queryCar' function:

```json
{
	"id": "hyperledgerFabricChainCode_2",
	"name": "hyperledgerFabricChainCode",
	"description": "Plugin to Invoke and Query Hyperledger Smart Contracts(ChainCodes)",
	"type": 1,
	"activityType": "hyperledgerFabricChainCode",
	"activityRef": "github.com/vinayak03/Akatsuki-flogo/activity/hyperledgerFabricChainCode",
	"attributes": [{
		"name": "NetworkConfig",
		"value": "/etc/hyperledger/client/config.yaml",
		"required": true,
		"type": "string"
	},
	{
		"name": "RequestType",
		"value": "Query",
		"required": true,
		"type": "string"
	},
	{
		"name": "ChannelID",
		"value": "mychannel",
		"required": true,
		"type": "string"
	},
	{
		"name": "User",
		"value": "user1",
		"required": true,
		"type": "string"
	},
	{
		"name": "UserPasswd",
		"value": "user1pw",
		"required": true,
		"type": "string"
	},
	{
		"name": "UserOrg",
		"value": "Org1",
		"required": true,
		"type": "string"
	},
	{
		"name": "ChainCodeID",
		"value": "fabcar",
		"required": true,
		"type": "string"
	},
	{
		"name": "FunctionName",
		"value": "queryCar",
		"required": true,
		"type": "string"
	},
	{
		"name": "Params",
		"value": "{\"params\" : [\"CAR1\"]}",
		"required": true,
		"type": "any"
	}]
}
```
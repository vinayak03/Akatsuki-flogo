# Flogo Activity Hyperledger-Fabric-List
This activity allows you to List the channels on the network within the org. It also supports listing out the chaincode which are installed and instantiated

## Pre-requisite
### Dependencies
1. Docker and Docker-Compose
2. Node.js 
3. Golang 1.9.4
4. fabric-sdk-go https://godoc.org/github.com/hyperledger/fabric-sdk-go
5. It is uses basic network setup and sample chaincode from https://github.com/hyperledger/fabric-samples

### Sample Hyperledger Network Setup
Note: Run the below steps in powershell on windows
#### Download Hyperledger network images
	Inside the directory ../../fabric-setup/network/
	sh bootstrap.sh 1.1.0

#### Start a basic network with fabcar chaincode
	Inside directory ../../fabric-setup/network/
	sh start.sh

#### Configure users in the network to users
	Inside ../../fabric-setup/script
	//node dependency installtion
	npm install

	//Get the default admin user from ca
	node enrollAdmin.js

	//register the user to be used with the chaincode example.
	node registerUser.js

## Installation
### Flogo-webui
	Inside directory ../../fabric-setup/network/

	sh flogo-activity-install.sh
	
	For Windows you can run the below bat file:
	
	flogo-activity-install.bat
	
	if the web-ui is already running it would need a restart.
### Flogo-cli
```
flogo add activity github.com/vinayak03/Akatsuki-flogo/activity/hyperledgerFabricList
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
      "name": "ResourceType",
      "type": "string",
	  "required": true,
      "allowed" : ["Channel","InstalledChainCode","InstantiatedChainCode"]
    },
	{
      "name": "AdminUser",
      "type": "string"
    },
	{
      "name": "AdminPasswd",
      "type": "string"
    },
    {
      "name": "Org",
      "type": "string"
    },
    {
      "name": "ChannelID",
      "type": "string"
    },
    {
      "name": "QueryParam",
      "type": "string"
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
  ]
}
```
## Input
| Input Parameter    | Required | Description                                                                 |
|:-------------------|:---------|:----------------------------------------------------------------------------|
| NetworkConfig      | True     | yaml configuration file describing the network and artifacts for network.
| ResourceType       | True     | Channel/InstalledChainCode/InstantiatedChainCode
| AdminUser          | True     | User which has privileges for resource management
| AdminPasswd        | True     | Password for user
| Org                | True     | Org for which we will query
| ChannelID          | True     | Incase of Instantiated Chaincode query we need to specify the channel name
| QueryParam         | True     | The URL of the any peer in the network to be queried eg. grpc://localhost:7051

## Output
| Output Parameter    | Required | Description                                                                 |
|:-------------------|:---------|:---------------------------------------------------------------------------- |
| Status             | True     | Successfull operation will return a status of 200 and error in processing will return 500
| ResponsePayload    | False    | Any response payload coming from the chaincode
| ErrorMessage       | False    | In case of error in processing it will return the detailed error message in this field.

## Examples
The below example chaincode queries 'queryAllCars' function of fabcar chaincode:

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

The below example chaincode queries 'queryCar' function of fabcar chaincode:

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
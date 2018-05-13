# Flogo Activity Hybris Product
This activity allows you to invoke the rest api for Hybris Product

## Installation
### Flogo-cli
```
flogo add activity github.com/vinayak03/Akatsuki-flogo/tree/master/activity/hybrisProduct
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "requestType",
      "type": "string",
	  "required": true,
      "allowed" : ["GetAllProducts", "GetProduct","CreateProduct","UpdateProduct","DeleteProduct","DeleteAllProducts"]
    },
    {
      "name": "APIKey",
      "type": "string",
      "required": true
    },
    {
      "name": "URL",
      "type": "string",
      "required": true
    },
    {
      "name": "tenant",
      "type": "string",
      "required": true
    },
    {
      "name": "productId",
      "type": "string",
      "required": false
    },
    {
      "name": "body",
      "type": "string",
      "required": false
    }
  ],
  "outputs": [
    {
      "name": "statusCode",
      "type": "string"
    },
    {
      "name": "responsePayload",
      "type": "string"
    },
    {
      "name": "errorMessage",
      "type": "string"
    }    
  ]
}
```
## Input
| Input Parameter    | Required | Description                                                                 |
|:-------------------|:---------|:----------------------------------------------------------------------------|
| requestType        | True     | Type of request for a product eg. Get , Delete , Update etc
| APIKey             | True     | APIKey or OAuth token to be used
| URL                | True     | BaseURL for Hybris 
| tenant             | True     | Tenant ID
| productId          | True     | ProductID
| body               | True     | JSON payload

## Output
| Output Parameter    | Required | Description                                                                 |
|:-------------------|:---------|:---------------------------------------------------------------------------- |
| statusCode         | True     | Successfull operation will return a status of 200 and error in processing will return 500
| responsePayload    | False    | Response Payload Coming from hybris
| errorMessage       | False    | In case of error in processing it will return the detailed error message in this field.

## Examples
The below example for Getting All Products.

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
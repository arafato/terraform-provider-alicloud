package gpdb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// AddBuDBInstanceRelation invokes the gpdb.AddBuDBInstanceRelation API synchronously
// api document: https://help.aliyun.com/api/gpdb/addbudbinstancerelation.html
func (client *Client) AddBuDBInstanceRelation(request *AddBuDBInstanceRelationRequest) (response *AddBuDBInstanceRelationResponse, err error) {
	response = CreateAddBuDBInstanceRelationResponse()
	err = client.DoAction(request, response)
	return
}

// AddBuDBInstanceRelationWithChan invokes the gpdb.AddBuDBInstanceRelation API asynchronously
// api document: https://help.aliyun.com/api/gpdb/addbudbinstancerelation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddBuDBInstanceRelationWithChan(request *AddBuDBInstanceRelationRequest) (<-chan *AddBuDBInstanceRelationResponse, <-chan error) {
	responseChan := make(chan *AddBuDBInstanceRelationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddBuDBInstanceRelation(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// AddBuDBInstanceRelationWithCallback invokes the gpdb.AddBuDBInstanceRelation API asynchronously
// api document: https://help.aliyun.com/api/gpdb/addbudbinstancerelation.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddBuDBInstanceRelationWithCallback(request *AddBuDBInstanceRelationRequest, callback func(response *AddBuDBInstanceRelationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddBuDBInstanceRelationResponse
		var err error
		defer close(result)
		response, err = client.AddBuDBInstanceRelation(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// AddBuDBInstanceRelationRequest is the request struct for api AddBuDBInstanceRelation
type AddBuDBInstanceRelationRequest struct {
	*requests.RpcRequest
	BusinessUnit string           `position:"Query" name:"BusinessUnit"`
	DBInstanceId string           `position:"Query" name:"DBInstanceId"`
	OwnerId      requests.Integer `position:"Query" name:"OwnerId"`
}

// AddBuDBInstanceRelationResponse is the response struct for api AddBuDBInstanceRelation
type AddBuDBInstanceRelationResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	BusinessUnit   string `json:"BusinessUnit" xml:"BusinessUnit"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
}

// CreateAddBuDBInstanceRelationRequest creates a request to invoke AddBuDBInstanceRelation API
func CreateAddBuDBInstanceRelationRequest() (request *AddBuDBInstanceRelationRequest) {
	request = &AddBuDBInstanceRelationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "AddBuDBInstanceRelation", "gpdb", "openAPI")
	return
}

// CreateAddBuDBInstanceRelationResponse creates a response to parse from AddBuDBInstanceRelation response
func CreateAddBuDBInstanceRelationResponse() (response *AddBuDBInstanceRelationResponse) {
	response = &AddBuDBInstanceRelationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
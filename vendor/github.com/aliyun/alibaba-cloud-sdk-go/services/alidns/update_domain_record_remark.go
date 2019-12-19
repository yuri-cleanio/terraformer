package alidns

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateDomainRecordRemark invokes the alidns.UpdateDomainRecordRemark API synchronously
// api document: https://help.aliyun.com/api/alidns/updatedomainrecordremark.html
func (client *Client) UpdateDomainRecordRemark(request *UpdateDomainRecordRemarkRequest) (response *UpdateDomainRecordRemarkResponse, err error) {
	response = CreateUpdateDomainRecordRemarkResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDomainRecordRemarkWithChan invokes the alidns.UpdateDomainRecordRemark API asynchronously
// api document: https://help.aliyun.com/api/alidns/updatedomainrecordremark.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDomainRecordRemarkWithChan(request *UpdateDomainRecordRemarkRequest) (<-chan *UpdateDomainRecordRemarkResponse, <-chan error) {
	responseChan := make(chan *UpdateDomainRecordRemarkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDomainRecordRemark(request)
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

// UpdateDomainRecordRemarkWithCallback invokes the alidns.UpdateDomainRecordRemark API asynchronously
// api document: https://help.aliyun.com/api/alidns/updatedomainrecordremark.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDomainRecordRemarkWithCallback(request *UpdateDomainRecordRemarkRequest, callback func(response *UpdateDomainRecordRemarkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDomainRecordRemarkResponse
		var err error
		defer close(result)
		response, err = client.UpdateDomainRecordRemark(request)
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

// UpdateDomainRecordRemarkRequest is the request struct for api UpdateDomainRecordRemark
type UpdateDomainRecordRemarkRequest struct {
	*requests.RpcRequest
	Remark       string `position:"Query" name:"Remark"`
	RecordId     string `position:"Query" name:"RecordId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// UpdateDomainRecordRemarkResponse is the response struct for api UpdateDomainRecordRemark
type UpdateDomainRecordRemarkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDomainRecordRemarkRequest creates a request to invoke UpdateDomainRecordRemark API
func CreateUpdateDomainRecordRemarkRequest() (request *UpdateDomainRecordRemarkRequest) {
	request = &UpdateDomainRecordRemarkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "UpdateDomainRecordRemark", "alidns", "openAPI")
	return
}

// CreateUpdateDomainRecordRemarkResponse creates a response to parse from UpdateDomainRecordRemark response
func CreateUpdateDomainRecordRemarkResponse() (response *UpdateDomainRecordRemarkResponse) {
	response = &UpdateDomainRecordRemarkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucket

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type CommitstatusesApiService service


/* CommitstatusesApiService 
 Returns the specified build status for a commit.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user. 
 @param node The commit&#39;s SHA1.
 @param key The build status&#39; unique key
 @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;. 
 @return Commitstatus*/
func (a *CommitstatusesApiService) RepositoriesUsernameRepoSlugCommitNodeStatusesBuildKeyGet(ctx context.Context, username string, node string, key string, repoSlug string) (Commitstatus,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Commitstatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{username}/{repo_slug}/commit/{node}/statuses/build/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"node"+"}", fmt.Sprintf("%v", node), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", fmt.Sprintf("%v", key), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* CommitstatusesApiService 
 Used to update the current status of a build status object on the specific commit.  This operation can also be used to change other properties of the build status:  * &#x60;state&#x60; * &#x60;name&#x60; * &#x60;description&#x60; * &#x60;url&#x60; * &#x60;refname&#x60;  The &#x60;key&#x60; cannot be changed.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user. 
 @param node The commit&#39;s SHA1.
 @param key The build status&#39; unique key
 @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;. 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (Commitstatus) The updated build status object
 @return Commitstatus*/
func (a *CommitstatusesApiService) RepositoriesUsernameRepoSlugCommitNodeStatusesBuildKeyPut(ctx context.Context, username string, node string, key string, repoSlug string, localVarOptionals map[string]interface{}) (Commitstatus,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Commitstatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{username}/{repo_slug}/commit/{node}/statuses/build/{key}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"node"+"}", fmt.Sprintf("%v", node), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key"+"}", fmt.Sprintf("%v", key), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(Commitstatus); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* CommitstatusesApiService 
 Creates a new build status against the specified commit.  If the specified key already exists, the existing status object will be overwritten.  When creating a new commit status, you can use a URI template for the URL. Templates are URLs that contain variable names that Bitbucket will evaluate at runtime whenever the URL is displayed anywhere similar to parameter substitution in [Bitbucket Connect](https://developer.atlassian.com/bitbucket/concepts/context-parameters.html). For example, one could use &#x60;https://foo.com/builds/{repository.full_name}&#x60; which Bitbucket will turn into &#x60;https://foo.com/builds/foo/bar&#x60; at render time. The context variables available are &#x60;repository&#x60; and &#x60;commit&#x60;.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user. 
 @param node The commit&#39;s SHA1.
 @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;. 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (Commitstatus) The new commit status object.
 @return Commitstatus*/
func (a *CommitstatusesApiService) RepositoriesUsernameRepoSlugCommitNodeStatusesBuildPost(ctx context.Context, username string, node string, repoSlug string, localVarOptionals map[string]interface{}) (Commitstatus,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  Commitstatus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{username}/{repo_slug}/commit/{node}/statuses/build"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"node"+"}", fmt.Sprintf("%v", node), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(Commitstatus); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* CommitstatusesApiService 
 Returns all statuses (e.g. build results) for a specific commit.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user. 
 @param node The commit&#39;s SHA1.
 @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;. 
 @return PaginatedCommitstatuses*/
func (a *CommitstatusesApiService) RepositoriesUsernameRepoSlugCommitNodeStatusesGet(ctx context.Context, username string, node string, repoSlug string) (PaginatedCommitstatuses,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PaginatedCommitstatuses
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{username}/{repo_slug}/commit/{node}/statuses"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"node"+"}", fmt.Sprintf("%v", node), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* CommitstatusesApiService 
 Returns all statuses (e.g. build results) for the given pull request.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param username This can either be the username or the UUID of the account, surrounded by curly-braces, for example: &#x60;{account UUID}&#x60;. An account is either a team or user. 
 @param pullRequestId The id of the pull request.
 @param repoSlug This can either be the repository slug or the UUID of the repository, surrounded by curly-braces, for example: &#x60;{repository UUID}&#x60;. 
 @return PaginatedCommitstatuses*/
func (a *CommitstatusesApiService) RepositoriesUsernameRepoSlugPullrequestsPullRequestIdStatusesGet(ctx context.Context, username string, pullRequestId int32, repoSlug string) (PaginatedCommitstatuses,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PaginatedCommitstatuses
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories/{username}/{repo_slug}/pullrequests/{pull_request_id}/statuses"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", fmt.Sprintf("%v", username), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pull_request_id"+"}", fmt.Sprintf("%v", pullRequestId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo_slug"+"}", fmt.Sprintf("%v", repoSlug), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}


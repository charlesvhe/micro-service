using Newtonsoft.Json;
using System;
using System.Collections.ObjectModel;
using System.Net;
using System.Net.Http;
using System.Reflection.Metadata;
using System.Text;
using System.Threading.Tasks;

namespace Common
{
    public class WebApiClient
    {
        public TimeSpan Timeout { get; set; }

        static object _lockObject = new object();
        static WebApiClient _instance;
        public static WebApiClient Instance
        {
            get
            {
                lock (_lockObject)
                {
                    if (_instance == null)
                    {
                        _instance = new WebApiClient();
                    }
                }
                return _instance;
            }
        }

        public WebApiClient()
        {
            Timeout = new TimeSpan(0, 1, 0);
        }

        #region Interface Method

        public string Get(string url, Collection<Cookie> cookies = null, int? timeout = null, params Parameter[] parms)
        {
            return HttpRequest(HttpMethod.Get, url, null, cookies, timeout, parms);
        }


        public string Delete(string url, Collection<Cookie> cookies = null, int? timeout = null, params Parameter[] parms)
        {
            return HttpRequest(HttpMethod.Delete, url, null, cookies, timeout, parms);
        }

        public string Post(string url, object data, Collection<Cookie> cookies = null, int? timeout = null)
        {
            return HttpRequest(HttpMethod.Post, url, data, cookies, timeout);
        }

        public string Put(string url, object data, Collection<Cookie> cookies = null, int? timeout = null)
        {
            return HttpRequest(HttpMethod.Put, url, data, cookies, timeout);
        }
        #endregion

        #region HttpMethod
        private string HttpRequest(HttpMethod httpMethod, string url, object data, Collection<Cookie> cookies = null, int? timeout = null, params Parameter[] parms)
        {
            using (var response = GetHttpResponse(httpMethod, ref url, data, cookies, timeout, parms))
            {
                var responseContent = response.Result.Content.ReadAsStringAsync().Result;
                return responseContent;
            }
        }

        private Task<HttpResponseMessage> GetHttpResponse(HttpMethod httpMethod, ref string url,
            object data, Collection<Cookie> cookies = null, int? timeout = null, params Parameter[] parms)
        {
            var cookieContainer = new CookieContainer();
            using (var handler = new HttpClientHandler { CookieContainer = cookieContainer })
            {
                using (var client = new HttpClient(handler))
                {
                    //Set TimeOut
                    client.Timeout = timeout.HasValue ? TimeSpan.FromMilliseconds(timeout.Value) : Timeout;
                    var requestBody = data == null ? "" : data.ToJsonString();
                    using (var requestContent = new StringContent(requestBody, Encoding.UTF8, "application/json"))
                    {
                        ServicePointManager.SecurityProtocol = SecurityProtocolType.Tls;
                        Task<HttpResponseMessage> response = null;
                        switch (httpMethod)
                        {
                            case HttpMethod.Get:
                                response = client.GetAsync(url);
                                break;
                            case HttpMethod.Delete:
                                response = client.DeleteAsync(url);
                                break;
                            case HttpMethod.Put:
                                response = client.PutAsync(url, requestContent);
                                break;
                            case HttpMethod.Post:
                                response = client.PostAsync(url, requestContent);
                                break;
                        }
                        if (response == null)
                            throw new Exception("HttpResponse Null Exception");

                        string responseContent = null;
                        Exception ex = null;
                        try
                        {
                            responseContent = response.Result.Content.ReadAsStringAsync().Result;
                        }
                        catch (Exception e)
                        {
                            ex = e;
                        }

                        if (response.Result.StatusCode != HttpStatusCode.OK)
                        {
                            int status = (int)response.Result.StatusCode;
                            string message = response.Result.StatusCode.ToString();
                            dynamic content = null;
                            try
                            {
                                content = responseContent?.DeserializeJsonObject<dynamic>();
                            }
                            catch
                            {
                            }

                            message = $"{message}:{content},url:{url},content:{responseContent}";
                            throw new Exception(message);
                        }
                        //SetResponseHeaders(response);
                        return response;
                    }
                }
            }
        }
        #endregion

        private enum HttpMethod
        {
            Get,
            Delete,
            Put,
            Post
        }
    }
}

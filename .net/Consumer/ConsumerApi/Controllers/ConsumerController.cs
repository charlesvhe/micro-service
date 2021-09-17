using Common;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace ConsumerService.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class ConsumerController : ControllerBase
    {
        private readonly ILogger<ConsumerController> _logger;

        public ConsumerController(ILogger<ConsumerController> logger)
        {
            _logger = logger;
        }

        [HttpGet]
        public string Get()
        {
            return "这是消费者的页面";
        }

        [HttpGet, Route("request")]
        public string RequestConsumerOnly(string content)
        {
            string str = string.Format("【{0}】请求了ConsumerApi服务！", content);
            _logger.LogDebug(str);
            return str;
        }

        [HttpGet, Route("requestProvider")]
        public string RequestConsumerToProvider(string content)
        {
            string str = string.Format("【{0}】先请求了ConsumerApi服务，然后", content);
            string url = string.Format("https://{0}:{1}/api/provider/consumerRequest?content={2}", "localhost", "51001", str);
            string result = WebApiClient.Instance.Get(url);
            _logger.LogDebug(result);
            return result;
        }

        [HttpGet, Route("providerRequest")]
        public string ProviderRequestConsumer(string content)
        {
            string str = string.Format("{0}请求了ConsumerApi服务！", content);
            _logger.LogDebug(str);
            return str;
        }
    }
}

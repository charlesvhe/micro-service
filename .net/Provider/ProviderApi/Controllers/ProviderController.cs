using Common;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System.Reflection.Metadata;

namespace ProviderService.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class ProviderController : ControllerBase
    {
        private readonly ILogger<ProviderController> _logger;

        public ProviderController(ILogger<ProviderController> logger)
        {
            _logger = logger;
        }

        [HttpGet, Route("Get")]
        public string Get()
        {
            return "这是生产者的页面";
        }

        [HttpGet, Route("request")]
        public string RequestProviderOnly(string content)
        {
            string str = string.Format("【{0}】请求了ProviderApi服务！", content);
            _logger.LogDebug(str);
            return str;
        }

        [HttpGet, Route("requestConsumer")]
        public string RequestProviderToConsumer(string content)
        {
            string str = string.Format("【{0}】先请求了Provider服务，然后", content);
            string url = string.Format("https://{0}:{1}/api/consumer/providerRequest?content={2}", "localhost", "52001", str);
            string result = WebApiClient.Instance.Get(url);
            _logger.LogDebug(result);
            return result;
        }

        [HttpGet, Route("consumerRequest")]
        public string ConsumerRequestProvider(string content)
        {
            string str = string.Format("{0}请求了ProviderApi服务！", content);
            _logger.LogDebug(str);
            return str;
        }
    }
}

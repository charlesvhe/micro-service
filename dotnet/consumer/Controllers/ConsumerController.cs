using System;
using System.Net.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

using Steeltoe.Discovery;
using Steeltoe.Common.Discovery;

namespace consumer.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ConsumerController : ControllerBase
    {
        private readonly ILogger<ConsumerController> _logger;
        private DiscoveryHttpClientHandler _handler;
        public ConsumerController(ILogger<ConsumerController> logger, IDiscoveryClient client)
        {
            _logger = logger;
             _handler = new DiscoveryHttpClientHandler(client);
        }

        [HttpGet]
        public String Get(String name)
        {
            _logger.LogInformation("Hi there");
            var task = new HttpClient(_handler, false).GetStringAsync("http://provider/provider?name=" + name);
            task.Wait();
            _logger.LogInformation("info "+task.ToString());
            return "Hello consumer " + task.Result;
        }
    }
}

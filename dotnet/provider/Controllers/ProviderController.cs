using System;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace provider.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ProviderController : ControllerBase
    {
        private readonly ILogger<ProviderController> _logger;

        public ProviderController(ILogger<ProviderController> logger)
        {
            _logger = logger;
        }

        [HttpGet]
        public String Get(String name)
        {
            _logger.LogInformation("Hi there");
            
            return "Hello " + name+"! "+DateTime.Now.ToString();
        }
    }
}

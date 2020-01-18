using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;
using CodeChatApp.Database;
using System.Threading.Tasks;
using CodeChatApp.Controllers.JsonRPC.Models;
using Microsoft.AspNetCore.Mvc;
using CodeChatApp.Services;
using Microsoft.Extensions.Logging;

namespace CodeChatApp.Controllers.JsonRPC
{
    /// <summary>
    /// Main controller
    /// </summary>
    [Route("CodeChat")]
    public class CodeChatController : Controller
    {
        private readonly IImplementor _implementor;
        private readonly ILogger _logger;

        public CodeChatController(IImplementor implementor, ILogger<CodeChatController> logger)
        {
            _implementor = implementor;
            _logger = logger;
        }

        /// <summary>
        /// HealthCheck method
        /// </summary>
        /// <returns> 'Stoit Vrode' message with 200 status code</returns>
        [HttpGet]
        public IActionResult Get()
        {
            _logger.LogWarning("Success Request to test get method");
            return new ObjectResult("Stoit vrode!!");
        }

        /// <summary>
        /// Main controller method for int ower JSON-RPC
        /// </summary>
        /// <param name="request">Request Body</param>
        /// <returns>Response body and status code 200</returns>
        [HttpPost]
        public IActionResult Post([FromBody] Request request)
        {
            Response response = new Response();

            if (request == null)
            {
                response.Status = 40000;
                _logger.LogWarning("Couldn't parse request body!!!!");
                response.Result = "Check your request body. Couldn't parse it!!!";
            }
            else
            {
                MethodInfo method = _implementor.GetType().GetMethod(request.Method);
                if (method != null)
                {
                    _logger.LogInformation($"Server method {request.Method} was called");
                    response = (Response)method.Invoke(_implementor, new object[] { request.Token, request.Params });
                }
                else
                {
                    response.Status = 40001;
                    _logger.LogWarning($"There is no such method as {request.Method}");
                    response.Result = "No such method in Backend!!!";
                }
            }

            return new ObjectResult(response);
        }
    }
}

using Newtonsoft.Json.Linq;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Controllers.JsonRPC.Models
{
    /// <summary>
    /// Represents request body class
    /// </summary>
    public class Request
    {
        /// <summary>
        /// Method to invoke in JSON-RPC
        /// </summary>
        public string Method { get; set; }

        /// <summary>
        /// Auth token
        /// </summary>
        public string Token { get; set; }

        /// <summary>
        /// others params. Could be everything
        /// </summary>
        public JObject Params { get; set; }
    }
}

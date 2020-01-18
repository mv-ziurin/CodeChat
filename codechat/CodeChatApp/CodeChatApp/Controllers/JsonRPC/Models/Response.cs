using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Controllers.JsonRPC.Models
{
    /// <summary>
    /// Represents response body
    /// </summary>
    public class Response
    {
        /// <summary>
        /// custom status code. Check readme.md to get more info 
        /// https://gitlab.com/codechat-bmstu/codechat/tree/master/codechat/CodeChatApp/CodeChatApp
        /// </summary>
        public int Status { get; set; }

        /// <summary>
        /// reqyest result body
        /// </summary>
        public object Result { get; set; }
    }
}

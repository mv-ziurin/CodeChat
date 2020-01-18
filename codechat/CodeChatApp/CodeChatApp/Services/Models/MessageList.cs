using CodeChatApp.Database.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Services.Models
{
    public class MessageList
    {
        public List<SingleMessage> Messages { get; set; }

        public MessageList()
        {
            this.Messages = new List<SingleMessage>();
        }
    }
}

using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Services.Models
{
    public class ChannelsList
    {
        public List<Channel> Channels { get; set; }

        public ChannelsList()
        {
            this.Channels = new List<Channel>();
        }
    }
}

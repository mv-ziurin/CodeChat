using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Services.Models
{
    public class Channel
    {
        public long ChatId { get; set; }

        public string Name { get; set; }

        public List<CodeChannel> CodeChats { get; set; }

        public Channel()
        {
            this.CodeChats = new List<CodeChannel>();
        }
    }
}

using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Services.Models
{
    public class SingleMessage
    {
        public long Id { get; set; }

        public string UserName { get; set; }

        public long ChatId { get; set; }

        public string Text { get; set; }

        public DateTime Time { get; set; }
    }
}

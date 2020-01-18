using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Database.Models
{
    public class Message
    {
        public long Id { get; set; }

        public string UserName { get; set; }

        public virtual Users User { get; set; }

        public long ChatId { get; set; }

        public virtual Chat Chat { get; set; }

        public string Text { get; set; }

        public DateTime Time { get; set; }
    }
}

using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Database.Models
{
    public class CodeChat
    {
        public long Id { get; set; }

        public string Name { get; set; }

        public long ChatId { get; set; }

        public virtual Chat Chat { get; set; }
    }
}

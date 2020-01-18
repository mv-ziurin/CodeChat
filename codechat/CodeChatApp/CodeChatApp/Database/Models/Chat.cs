using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Database.Models
{
    public class Chat
    {
        public long Id { get; set; }

        public string Name { get; set; }

        public virtual ICollection<UserChat> UserChats { get; set; }

        public virtual ICollection<Message> Messages { get; set; }

        public virtual ICollection<CodeChat> CodeChats { get; set; }

        public Chat()
        {
            this.UserChats = new HashSet<UserChat>();
            this.CodeChats = new HashSet<CodeChat>();
            this.Messages = new HashSet<Message>();
        }
    }
}

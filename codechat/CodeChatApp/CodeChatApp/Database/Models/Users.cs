using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Database.Models
{
    public class Users
    {
        public string Username { get; set; }

        public string Email { get; set; }

        public string Hash { get; set; }

        public string PhotoUrl { get; set; }

        public virtual ICollection<Message> Messages { get; set; }

        public virtual ICollection<UserChat> UserChats { get; set; }

        public ICollection<EventLog> EventLogEmailNavigation { get; set; }

        public ICollection<EventLog> EventLogUsernameNavigation { get; set; }

        public Users()
        {
            this.EventLogEmailNavigation = new HashSet<EventLog>();
            this.EventLogUsernameNavigation = new HashSet<EventLog>();
            this.UserChats = new HashSet<UserChat>();
            this.Messages = new HashSet<Message>();
        }
    }
}

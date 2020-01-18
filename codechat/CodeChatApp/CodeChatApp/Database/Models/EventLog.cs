using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Database.Models
{
    public partial class EventLog
    {
        public int Id { get; set; }

        public string Event { get; set; }

        public string Username { get; set; }

        public string Email { get; set; }

        public string Ip { get; set; }

        public string UserAgent { get; set; }

        public DateTime EventTime { get; set; }

        public Users EmailNavigation { get; set; }

        public Users UsernameNavigation { get; set; }
    }
}

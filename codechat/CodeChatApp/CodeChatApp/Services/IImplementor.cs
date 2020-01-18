using CodeChatApp.Database.Models;
using Newtonsoft.Json.Linq;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Services
{
    public interface IImplementor
    {
        object GetChats(string token, JObject jObject);

        object GetMessageHistory(string token, JObject jobject);

        object PostChat(string token, JObject jobject);

        object PostCodeChat(string token, JObject jobject);

        object AddUserToChat(string token, JObject jobject);

        object LeaveChannel(string token, JObject jobject);
    }
}

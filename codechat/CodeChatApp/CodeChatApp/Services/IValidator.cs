using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Services
{
    public interface IValidator
    {
        string GetUserName(string token);

        bool CheckUser(string username);

        bool Validate(string token);
    }
}

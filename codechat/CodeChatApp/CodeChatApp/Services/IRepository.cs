using CodeChatApp.Database.Models;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Services
{
    public interface IRepository
    {
        Chat GetChat(long id);

        CodeChat GetCodeChat(long id);

        Message GetMessage(long id);

        UserChat GetUserChat(long id);

        Users GetUser(string username);

        List<Chat> GetChats();

        List<CodeChat> GetCodeChats();

        List<Message> GetMessages();

        List<UserChat> GetUserChats();

        List<Users> GetUsers();

        List<Chat> GetChatsList(string userName);

        List<CodeChat> GetCodeChatsList(long chatId);

        List<Message> GetMessageHistory(long chatId);

        long PostChat(Chat chat);

        long PostCodeChat(CodeChat codechat);

        long PostMessage(Message message);

        long PostUserChat(UserChat userchat);

        void DeleteChat(Chat chat);

        void DeleteCodeChat(CodeChat codechat);

        void DeleteMessage(Message message);

        void DeleteUserChat(UserChat userchat);

        void DeleteUserChat(long chatId, string username);

        void DeleteFreeChats(long chatId, string username);

        bool CheckUserInChat(UserChat userchat);
    }
}

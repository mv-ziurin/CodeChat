using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using CodeChatApp.Database;
using CodeChatApp.Database.Models;
using Microsoft.EntityFrameworkCore;

namespace CodeChatApp.Services
{
    /// <summary>
    /// Service for intearacting with DB
    /// </summary>
    public class Repository : IRepository
    {
        private CodeChatContext _context;

        public Repository(CodeChatContext context)
        {
            _context = context;
        }

        #region GetMethods

        /// <summary>
        /// Get list of chats from DB
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public Chat GetChat(long id)
        {
            return _context.Chats.FirstOrDefault(c => c.Id == id);
        }

        /// <summary>
        /// Get codechats from DB
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public CodeChat GetCodeChat(long id)
        {
            return _context.CodeChats.FirstOrDefault(c => c.Id == id);
        }

        /// <summary>
        /// Get list of message from DB
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public Message GetMessage(long id)
        {
            return _context.Messages.FirstOrDefault(c => c.Id == id);
        }

        /// <summary>
        /// Get links users to chats from DB
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public UserChat GetUserChat(long id)
        {
            return _context.UserChats.FirstOrDefault(c => c.Id == id);
        }

        /// <summary>
        /// Get list of users from DB
        /// </summary>
        /// <param name="username"></param>
        /// <returns></returns>
        public Users GetUser(string username)
        {
            return _context.Users.FirstOrDefault(c => c.Username == username);
        }

        /// <summary>
        /// Get list of chats from DB
        /// </summary>
        /// <returns></returns>
        public List<Chat> GetChats()
        {
            return _context.Chats.ToList();
        }

        /// <summary>
        /// Get list of codechats from DB
        /// </summary>
        /// <returns></returns>
        public List<CodeChat> GetCodeChats()
        {
            return _context.CodeChats.ToList();
        }

        /// <summary>
        /// Get list of messages from DB
        /// </summary>
        /// <returns></returns>
        public List<Message> GetMessages()
        {
            return _context.Messages.ToList();
        }
        
        /// <summary>
        /// Get list of Userchats from DB
        /// </summary>
        /// <returns></returns>
        public List<UserChat> GetUserChats()
        {
            return _context.UserChats.ToList();
        }

        /// <summary>
        /// Get list of users from DB
        /// </summary>
        /// <returns></returns>
        public List<Users> GetUsers()
        {
            return _context.Users.ToList();
        }

        /// <summary>
        /// Get chat's list from db
        /// </summary>
        /// <param name="userName"></param>
        /// <returns></returns>
        public List<Chat> GetChatsList(string userName)
        {
            return _context.Chats
                .Where(d => d.UserChats.Any(c => c.UserName == userName))
                .OrderBy(d => d.GetType()
                .GetProperty("Name")
                .GetValue(d, null))
                .ToList();
        }

        /// <summary>
        /// Get list of codechat's from db
        /// </summary>
        /// <param name="chatId"></param>
        /// <returns></returns>
        public List<CodeChat> GetCodeChatsList(long chatId)
        {
            return _context.CodeChats
                .Where(d => d.ChatId == chatId)
                .OrderBy(d => d.GetType()
                .GetProperty("Name")
                .GetValue(d, null))
                .ToList();
        }

        public List<Message> GetMessageHistory(long chatId)
        {
            return _context.Messages
                    .Where(d => d.ChatId == chatId)
                    .OrderBy(d => d.GetType()
                    .GetProperty("Time")
                    .GetValue(d, null))
                    .ToList();
        }

        #endregion

        #region PostMethods

        /// <summary>
        /// Create new Chat
        /// </summary>
        /// <param name="chat"></param>
        /// <returns></returns>
        public long PostChat(Chat chat)
        {
            if (_context.Chats.ToList().Count == 0)
                chat.Id = 1;
            else
                chat.Id = _context.Chats.ToList().Max(t => t.Id) + 1;
            _context.Chats.Add(chat);
            _context.SaveChanges();

            return chat.Id;
        }

        /// <summary>
        /// Create new CodeChat
        /// </summary>
        /// <param name="codechat"></param>
        /// <returns></returns>
        public long PostCodeChat(CodeChat codechat)
        {
            if (_context.CodeChats.ToList().Count == 0)
                codechat.Id = 1;
            else
                codechat.Id = _context.CodeChats.ToList().Max(t => t.Id) + 1;
            _context.CodeChats.Add(codechat);
            _context.SaveChanges();

            return codechat.Id;
        }

        /// <summary>
        ///  Create new Message
        /// </summary>
        /// <param name="message"></param>
        /// <returns></returns>
        public long PostMessage(Message message)
        {
            if (_context.Messages.ToList().Count == 0)
                message.Id = 1;
            else
                message.Id = _context.Messages.ToList().Max(t => t.Id) + 1;
            _context.Messages.Add(message);
            _context.SaveChanges();

            return message.Id;
        }

        /// <summary>
        /// Create new UserChat link
        /// </summary>
        /// <param name="userchat"></param>
        /// <returns></returns>
        public long PostUserChat(UserChat userchat)
        {
            if (_context.UserChats.ToList().Count == 0)
                userchat.Id = 1;
            else
                userchat.Id = _context.UserChats.ToList().Max(t => t.Id) + 1;
            _context.UserChats.Add(userchat);
            _context.SaveChanges();

            return userchat.Id;
        }

        #endregion


        #region DeleteMethods
        
        /// <summary>
        /// Remove chat from DB
        /// </summary>
        /// <param name="chat"></param>
        public void DeleteChat(Chat chat)
        {
            _context.Chats.Remove(chat);
            _context.SaveChanges();
        }

        /// <summary>
        /// Remove CodeChat from DB
        /// </summary>
        /// <param name="codechat"></param>
        public void DeleteCodeChat(CodeChat codechat)
        {
            _context.CodeChats.Remove(codechat);
            _context.SaveChanges();
        }

        /// <summary>
        /// Remove message from DB
        /// </summary>
        /// <param name="message"></param>
        public void DeleteMessage(Message message)
        {
            _context.Messages.Remove(message);
            _context.SaveChanges();
        }

        /// <summary>
        /// Remove User to Chat link from DB
        /// </summary>
        /// <param name="userchat"></param>
        public void DeleteUserChat(UserChat userchat)
        {
            _context.UserChats.Remove(userchat);
            _context.SaveChanges();
        }

        /// <summary>
        /// Remove User to Chat link from DB
        /// </summary>
        /// <param name="chatId"></param>
        /// <param name="username"></param>
        public void DeleteUserChat(long chatId, string username)
        {
            UserChat userchat = _context.UserChats
                    .Where(d => d.ChatId == chatId && d.UserName == username)
                    .First();

            _context.UserChats.Remove(userchat);
            _context.SaveChanges();
        }

        /// <summary>
        /// Remove free chats without users
        /// </summary>
        /// <param name="chatId"></param>
        /// <param name="username"></param>
        public void DeleteFreeChats(long chatId, string username)
        {
            UserChat userchat = _context.UserChats
                    .FirstOrDefault(d => d.ChatId == chatId);

            if (userchat == null)
            {
                CodeChat codechat = _context.CodeChats
                    .FirstOrDefault(d => d.ChatId == chatId);
                while (codechat != null)
                {
                    _context.CodeChats.Remove(codechat);
                    _context.SaveChanges();
                    codechat = _context.CodeChats
                        .FirstOrDefault(d => d.ChatId == chatId);
                }
                Chat chat = _context.Chats
                    .FirstOrDefault(d => d.Id == chatId);
                _context.Chats.Remove(chat);

                _context.SaveChanges();
            }
        }

        #endregion

        /// <summary>
        /// Check if exist users in certain chat
        /// </summary>
        /// <param name="userchat"></param>
        /// <returns></returns>
        public bool CheckUserInChat(UserChat userchat)
        {
            if (_context.UserChats.FirstOrDefault(c => c.UserName == userchat.UserName && c.ChatId == userchat.ChatId) == null)
                return false;
            else
                return true;
        }

    }
}

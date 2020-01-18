using CodeChatApp.Controllers.JsonRPC.Models;
using CodeChatApp.Database.Models;
using Microsoft.IdentityModel.Tokens;
using Newtonsoft.Json.Linq;
using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Linq;
using System.Text;
using CodeChatApp.Database;
using System.Threading.Tasks;
using CodeChatApp.Services.Models;
using System.Net;
using System.IO;

namespace CodeChatApp.Services
{
    /// <summary>
    /// This service implemets main functional for JSON-RPC interaction
    /// </summary>
    public class Implementor : IImplementor
    {
        private IValidator _validator;
        private IRepository _repository;
        private Response _response;

        public Implementor(IRepository repository, IValidator validator)
        {
            _repository = repository;
            _validator = validator;
            _response = new Response();
        }

        /// <summary>
        /// Get list of user's chats
        /// </summary>
        /// <param name="token"></param>
        /// <param name="jObject"></param>
        /// <returns></returns>
        public object GetChats(string token, JObject jObject)
        {
            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                ChannelsList list = new ChannelsList();
                foreach (var chat in _repository.GetChatsList(userName))
                {
                    Channel channel = new Channel();
                    channel.ChatId = chat.Id;
                    channel.Name = chat.Name;
                    foreach (var codeChat in _repository.GetCodeChatsList(channel.ChatId))
                    {
                        CodeChannel codeChannel = new CodeChannel();
                        codeChannel.CodeChatId = codeChat.Id;
                        codeChannel.MainChatName = chat.Name;
                        codeChannel.Name = codeChat.Name;
                        channel.CodeChats.Add(codeChannel);
                    }
                    list.Channels.Add(channel);
                }

                _response.Status = 20000;
                _response.Result = list;
            }
            else
            {
                _response.Status = 40002;
                _response.Result = "Validation failed!!!";
            }
            return _response;
        }

        /// <summary>
        /// Get messageHistory of specific chat
        /// </summary>
        /// <param name="token"></param>
        /// <param name="jobject"></param>
        /// <returns></returns>
        public object GetMessageHistory(string token, JObject jobject)
        {
            long chatId;
            try
            {
                chatId = jobject.ToObject<ChatIdParam>().ChatId;
                if (chatId == 0)
                {
                    throw new Exception();
                }
            }
            catch (Exception e)
            {
                _response.Status = 40003;
                _response.Result = "Invalid params. long ChatId is needed";
                return _response;
            }

            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                MessageList messageList = new MessageList();
                foreach (var message in _repository.GetMessageHistory(chatId))
                {
                    SingleMessage singleMessage = new SingleMessage();
                    singleMessage.Id = message.Id;
                    singleMessage.ChatId = message.ChatId;
                    singleMessage.UserName = message.UserName;
                    singleMessage.Time = message.Time;
                    singleMessage.Text = message.Text;
                    messageList.Messages.Add(singleMessage);
                }

                _response.Status = 20000;
                _response.Result = messageList;
            }
            else
            {
                _response.Status = 40002;
                _response.Result = "Validation failed!!!";
            }
            return _response;
        }

        /// <summary>
        /// Create new chat
        /// </summary>
        /// <param name="token"></param>
        /// <param name="jobject"></param>
        /// <returns></returns>
        public object PostChat(string token, JObject jobject)
        {
            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                Chat chat;
                UserChat userChat = new UserChat();
                ChatIdParam chatId = new ChatIdParam();
                try
                {
                    chat = jobject.ToObject<Chat>();
                    if (chat.Name == null)
                    {
                        throw new Exception();
                    }
                }
                catch (Exception e)
                {
                    _response.Status = 40003;
                    _response.Result = "Wrong Chat entity was sent. Check Chat entity!";
                    return _response;
                }

                try
                {
                    chat.Id = _repository.PostChat(chat);
                    userChat.UserName = userName;
                    userChat.ChatId = chat.Id;
                    _repository.PostUserChat(userChat);
                    chatId.ChatId = chat.Id;

                    _response.Status = 20000;
                    _response.Result = chatId;
                }
                catch (Exception e)
                {
                    _response.Status = 40004;
                    _response.Result = "Error ocured in chat adding to DB";
                }
            }
            else
            {
                _response.Status = 40002;
                _response.Result = "Validation failed!!!";
            }
            return _response;
        }

        /// <summary>
        /// Create new codechat
        /// </summary>
        /// <param name="token"></param>
        /// <param name="jobject"></param>
        /// <returns></returns>
        public object PostCodeChat(string token, JObject jobject)
        {
            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                CodeChat codechat;
                CodeChatIdParam codechatId = new CodeChatIdParam();
                try
                {
                    codechat = jobject.ToObject<CodeChat>();
                    if (codechat.Name == null || codechat.ChatId == 0)
                    {
                        throw new Exception();
                    }
                }
                catch (Exception e)
                {
                    _response.Status = 40003;
                    _response.Result = "Wrong CodeChat entity was sent. Check Chat entity!";
                    return _response;
                }

                try
                {
                    codechat.Id = _repository.PostCodeChat(codechat);
                    codechatId.CodeChatId = codechat.Id;
                    _response.Status = 20000;
                    _response.Result = codechatId;
                }
                catch (Exception e)
                {
                    _response.Status = 40004;
                    _response.Result = "Error ocured in chat adding to DB";
                }
            }
            else
            {
                _response.Status = 40002;
                _response.Result = "Validation failed!!!";
            }
            return _response;
        }

        /// <summary>
        /// Add user to ur chat
        /// </summary>
        /// <param name="token"></param>
        /// <param name="jobject"></param>
        /// <returns></returns>
        public object AddUserToChat(string token, JObject jobject)
        {
            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                UserChat userchat;
                try
                {
                    userchat = jobject.ToObject<UserChat>();
                    if (userchat.UserName == null || userchat.ChatId == 0)
                    {
                        throw new Exception();
                    }
                }
                catch (Exception e)
                {
                    _response.Status = 40003;
                    _response.Result = "Wrong UserChat entity was sent. Check UserChat entity!";
                    return _response;
                }

                try
                {
                    if (_repository.CheckUserInChat(userchat))
                    {
                        throw new Exception();
                    }

                    _repository.PostUserChat(userchat);
                    _response.Status = 20000;
                    _response.Result = "User was successfully added to the chat";
                }
                catch (Exception e)
                {
                    _response.Status = 40004;
                    _response.Result = "Error ocured in userchat adding to DB";
                }
            }
            else
            {
                _response.Status = 40002;
                _response.Result = "Validation failed!!!";
            }
            return _response;
        }

        /// <summary>
        /// Leave channel event. If there no linked users to current chat, then chat would be removed from DB.
        /// </summary>
        /// <param name="token"></param>
        /// <param name="jobject"></param>
        /// <returns></returns>
        public object LeaveChannel(string token, JObject jobject)
        {
            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                long chatId;
                try
                {
                    chatId = jobject.ToObject<ChatIdParam>().ChatId;
                    if (chatId == 0)
                    {
                        throw new Exception();
                    }
                }
                catch (Exception e)
                {
                    _response.Status = 40003;
                    _response.Result = "Invalid params. long ChatId is needed";
                    return _response;
                }

                try
                {
                    _repository.DeleteUserChat(chatId, userName);
                    _repository.DeleteFreeChats(chatId, userName);
                    _response.Status = 20000;
                    _response.Result = "User has successfuly leaved the channel";
                }
                catch (Exception e)
                {
                    _response.Status = 40004;
                    _response.Result = "Selected user doesnt have such chat";
                }
            }
            else
            {
                _response.Status = 40002;
                _response.Result = "Validation failed!!!";
            }
            return _response;
        }

        public object GetUser(string token, JObject jobject)
        {
            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                _response.Status = 20000;
                _response.Result = userName;
            }
            else
            {
                _response.Status = 40002;
                _response.Result = "Validation failed!!!";
            }
            return _response;
        }
    }
}

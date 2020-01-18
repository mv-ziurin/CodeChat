using Microsoft.AspNetCore.SignalR;
using System.Threading.Tasks;
using CodeChatApp.Database;
using CodeChatApp.Database.Models;
using System.Collections.Generic;
using System.Linq;
using System;
using CodeChatApp.Services;
using CodeChatApp.Hubs.Models;
using System.Collections.Concurrent;

namespace CodeChatApp.Hubs
{
    /// <summary>
    /// Hub for real-time interaction
    /// </summary>
    public class CodeChatHub : Hub
    {
        private readonly CodeChatContext _context;
        private static readonly ConcurrentDictionary<string, CodeEditor> _codeEditors = new ConcurrentDictionary<string, CodeEditor>();
        private readonly IValidator _validator;

        public CodeChatHub(CodeChatContext context, IValidator validator)
        {
            _context = context;
            _validator = validator;
        }

        /// <summary>
        /// Send message in main chat
        /// </summary>
        /// <param name="token"></param>
        /// <param name="chatId"></param>
        /// <param name="text"></param>
        /// <returns></returns>
        public async Task SendToMainChat(string token, string chatId, string text)
        {
            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                await Clients.All.SendAsync("recieveMessageMainChat", userName, text, chatId);
                Message message = new Message();
                message.ChatId = Convert.ToInt64(chatId);
                message.Text = text;
                message.Time = DateTime.Now;
                message.UserName = userName;
                if (_context.Messages.ToList().Count == 0)
                    message.Id = 1;
                else
                    message.Id = _context.Messages.ToList().Max(t => t.Id) + 1;
                await _context.Messages.AddAsync(message);
                await _context.SaveChangesAsync();
            }
        }

        /// <summary>
        /// Update chat's list event
        /// </summary>
        /// <param name="userName"></param>
        /// <returns></returns>
        public async Task UpdateChats(string userName)
        {
            await Clients.All.SendAsync("recieveUpdateChats", userName);
        }

        /// <summary>
        /// Joint chat event
        /// </summary>
        /// <param name="userName"></param>
        /// <param name="chatId"></param>
        /// <returns></returns>
        public async Task JoinChat(string userName, string chatId)
        {
            await Clients.All.SendAsync("recieveUserJoinChat", userName, chatId);
        }

        /// <summary>
        /// Leave chat event
        /// </summary>
        /// <param name="token"></param>
        /// <param name="chatId"></param>
        /// <returns></returns>
        public async Task LeaveChat(string token, string chatId)
        {
            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                await Clients.All.SendAsync("recieveUserLeaveChat", userName, chatId);
            }
        }

        /// <summary>
        /// Join codeChat event
        /// </summary>
        /// <param name="token"></param>
        /// <param name="codeChatId"></param>
        /// <returns></returns>
        public async Task JoinCodeChat(string token, string codeChatId)
        {
            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                await Clients.All.SendAsync("recieveUserJoinCodeChat", userName, codeChatId);
            }
        }

        /// <summary>
        /// leave codechat event
        /// </summary>
        /// <param name="token"></param>
        /// <param name="codeChatId"></param>
        /// <returns></returns>
        public async Task LeaveCodeChat(string token, string codeChatId)
        {
            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                await Clients.All.SendAsync("recieveUserLeaveCodeChat", userName, codeChatId);
            }
        }

        /// <summary>
        /// This data required when u entered the codechat
        /// </summary>
        /// <param name="codeChatId"></param>
        /// <returns></returns>
        public async Task GetInitialData(string codeChatId)
        {
            CodeEditor codeEditor;
            if (!_codeEditors.TryGetValue(codeChatId, out codeEditor))
            {
                codeEditor = new CodeEditor();
                codeEditor.ModeValue = "";
                codeEditor.Text = "";
                AddOrUpdateCodeEditor(codeChatId, codeEditor);
            }
            await Clients.Caller.SendAsync("recieveInitialData", codeChatId, codeEditor.ModeValue, codeEditor.Text);
        }

        /// <summary>
        /// Send text to codeshare zone
        /// </summary>
        /// <param name="codeChatId"></param>
        /// <param name="text"></param>
        /// <returns></returns>
        public async Task SendToCodeShare(string codeChatId, string text)
        {
            await Clients.Others.SendAsync("recieveMessageCodeShare", text, codeChatId);
            CodeEditor codeEditor;
            _codeEditors.TryGetValue(codeChatId, out codeEditor);
            codeEditor.Text = text;
            AddOrUpdateCodeEditor(codeChatId, codeEditor);
        }

        /// <summary>
        /// Switch coding language in codeShare
        /// </summary>
        /// <param name="codeChatId"></param>
        /// <param name="mode"></param>
        /// <param name="value"></param>
        /// <returns></returns>
        public async Task SwitchModeCodeShare(string codeChatId, string mode, string value)
        {
            await Clients.All.SendAsync("recieveModeCodeShare", mode, value, codeChatId);
            CodeEditor codeEditor;
            _codeEditors.TryGetValue(codeChatId, out codeEditor);
            codeEditor.ModeValue = value;
            AddOrUpdateCodeEditor(codeChatId, codeEditor);
        }

        /// <summary>
        /// Send message in codechat
        /// </summary>
        /// <param name="token"></param>
        /// <param name="codeChatId"></param>
        /// <param name="text"></param>
        /// <returns></returns>
        public async Task SendToCodeChat(string token, string codeChatId, string text)
        {
            string userName = _validator.GetUserName(token);
            if (_validator.CheckUser(userName))
            {
                await Clients.All.SendAsync("recieveMessageCodeChat", userName, text, codeChatId);
            }
        }

        private void AddOrUpdateCodeEditor(string codeChatId, CodeEditor codeEditor)
        {
            _codeEditors.AddOrUpdate(codeChatId, codeEditor, (key, existingVal) =>
            {
                if (codeEditor != existingVal)
                {
                    existingVal.ModeValue = codeEditor.ModeValue;
                    existingVal.Text = codeEditor.Text;
                }
                return existingVal;
            });
        }
    }
}

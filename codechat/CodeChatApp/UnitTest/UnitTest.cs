using Microsoft.VisualStudio.TestTools.UnitTesting;
using Moq;
using CodeChatApp.Controllers;
using CodeChatApp.Database;
using CodeChatApp.Database.Models;
using CodeChatApp.Controllers.JsonRPC;
using CodeChatApp.Controllers.JsonRPC.Models;
using CodeChatApp.Services;
using CodeChatApp.Services.Models;
using Microsoft.Extensions.Logging;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json.Linq;
using System.Collections.Generic;
using System;

namespace UnitTest
{
    [TestClass]
    public class UnitTest
    {
        private readonly string _username = "Webber1580";
        private readonly string _rightToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IldlYmJlcjE1ODBAZ21haWwuY29tIiwiaWF0IjoxNTQyMzE0ODI1LCJwcm9qZWN0IjoiYXBpIiwidXNlcm5hbWUiOiJXZWJiZXIxNTgwIn0.yrVaLhXBejyu9V1zEPyj3BjHfjsPwTnAHPZ79QRay4A";
        private readonly string _wrongUserNameToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IldlYmJlcjE1ODBAZ21haWwuY29tIiwiaWF0IjoxNTQyMzE0ODI1LCJwcm9qZWN0IjoiYXBpIiwidXNlcm5hbWUiOiJOb3RXZWJiZXIifQ.hBnCX9bM8PhNXutDNssmQWm086JJFyLvFYPY50ThIaI";
        private readonly string _wrongKeyToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IldlYmJlcjE1ODBAZ21haWwuY29tIiwiaWF0IjoxNTQyMzE0ODI1LCJwcm9qZWN0IjoiYXBpIiwidXNlcm5hbWUiOiJXZWJiZXIxNTgwIn0.-zzExnZvg_v0HUNMUuleehYOWy3KK1NxCIEUJkBGOIU";

        // -------------------------------------- Controller ----------------------------------------- //
        [TestMethod]
        public void ControllerWrongRequestBody()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            var _implementor = new Mock<IImplementor>();
            var _logger = new Mock<ILogger<CodeChatController>>();
            CodeChatController controller = new CodeChatController(_implementor.Object, _logger.Object);
            ObjectResult result = (ObjectResult)controller.Post(null);
            var expectedValue = new Response();
            expectedValue.Status = 40000;
            expectedValue.Result = "Check your request body. Couldn't parse it!!!";
            var resultValue = (Response)result.Value;
            Assert.AreEqual(expectedValue.Status, resultValue.Status);
            Assert.AreEqual(expectedValue.Result, resultValue.Result);
        }

        [TestMethod]
        public void ControllerWrongMethod()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            var _implementor = new Mock<IImplementor>();
            var _logger = new Mock<ILogger<CodeChatController>>();
            CodeChatController controller = new CodeChatController(_implementor.Object, _logger.Object);
            Request req = new Request();
            req.Method = "LeaveMeAlone";
            ObjectResult result = (ObjectResult)controller.Post(req);
            var expectedValue = new Response();
            expectedValue.Status =  40001;
            expectedValue.Result = "No such method in Backend!!!";
            var resultValue = (Response)result.Value;
            Assert.AreEqual(expectedValue.Status, resultValue.Status);
            Assert.AreEqual(expectedValue.Result, resultValue.Result);
        }

        [TestMethod]
        public void ControllerRightMethod()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            var _implementor = new Mock<IImplementor>();
            _implementor.Setup(i => i.LeaveChannel(It.IsAny<string>(), It.IsAny<JObject>())).Returns(new Response() { Status = 20000, Result = "ok"});
            var _logger = new Mock<ILogger<CodeChatController>>();
            CodeChatController controller = new CodeChatController(_implementor.Object, _logger.Object);
            Request req = new Request();
            req.Method = "LeaveChannel";
            ObjectResult result = (ObjectResult)controller.Post(req);
            var expectedValue = new Response();
            expectedValue.Status = 20000;
            expectedValue.Result = "ok";
            var resultValue = (Response)result.Value;
            Assert.AreEqual(expectedValue.Status, resultValue.Status);
            Assert.AreEqual(expectedValue.Result, resultValue.Result);
        }

        // -------------------------------------- Validator ----------------------------------------- //
        [TestMethod]
        public void ValidatorGetUsernameSuccess()
        {
            var _repository = new Mock<IRepository>();
            var validator = new Validator(_repository.Object);
            var result = validator.GetUserName(_rightToken);
            Assert.AreEqual(result, _username);
        }

        [TestMethod]
        public void ValidatorGetUsernameFailed()
        {
            var _repository = new Mock<IRepository>();
            var validator = new Validator(_repository.Object);
            var result = validator.GetUserName(_wrongKeyToken);
            Assert.AreEqual(result, null);
        }

        [TestMethod]
        public void ValidatorCheckUserSuccess()
        {
            var _repository = new Mock<IRepository>();
            _repository.Setup(r => r.GetUsers()).Returns(() =>
            {
                var user = new Users();
                user.Username = _username;
                return new List<Users> { user };
            });
            var validator = new Validator(_repository.Object);
            var userName = _username;
            var result = validator.CheckUser(userName);
            Assert.AreEqual(result, true);
        }

        [TestMethod]
        public void ValidatorCheckUserFailed()
        {
            var _repository = new Mock<IRepository>();
            _repository.Setup(r => r.GetUsers()).Returns(() =>
            {
                var user = new Users();
                user.Username = _username;
                return new List<Users> { user };
            });
            var validator = new Validator(_repository.Object);
            var userName = "NotWebber";
            var result = validator.CheckUser(userName);
            Assert.AreEqual(result, false);
        }

        [TestMethod]
        public void ValidatorValidateSuccess()
        {
            var _repository = new Mock<IRepository>();
            _repository.Setup(r => r.GetUsers()).Returns(() =>
            {
                var user = new Users();
                user.Username = _username;
                return new List<Users> { user };
            });
            var validator = new Validator(_repository.Object);
            var result = validator.Validate(_rightToken);
            Assert.AreEqual(result, true);
        }

        [TestMethod]
        public void ValidatorValidateFalse()
        {
            var _repository = new Mock<IRepository>();
            _repository.Setup(r => r.GetUsers()).Returns(() =>
            {
                var user = new Users();
                user.Username = _username;
                return new List<Users> { user };
            });
            var validator = new Validator(_repository.Object);
            var result = validator.Validate(_wrongUserNameToken);
            Assert.AreEqual(result, false);
        }

        // -------------------------------------- Implementor ----------------------------------------- //

        //GetChats

        [TestMethod]
        public void GetChatsWrongValidation()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(false);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject();
            var result = (Response)implementor.GetChats(_wrongKeyToken, jObject);
            Assert.AreEqual(result.Result, "Validation failed!!!");
        }

        [TestMethod]
        public void GetChatsSuccess()
        {
            var _repository = new Mock<IRepository>();
            _repository.Setup(r => r.GetUsers()).Returns(() =>
            {
                var user = new Users();
                user.Username = _username;
                return new List<Users> { user };
            });
            _repository.Setup(r => r.GetChatsList(It.IsAny<string>())).Returns(() =>
            {
                return new List<Chat>();
            });
            _repository.Setup(r => r.GetCodeChatsList(It.IsAny<long>())).Returns(() =>
            {
                return new List<CodeChat>();
            });
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.GetUserName(It.IsAny<string>())).Returns(_username);
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject();
            Response result = (Response)implementor.GetChats(_rightToken, jObject);
            Assert.AreEqual(((ChannelsList)result.Result).Channels.Count, 0);
        }

        // GetMessageHistory

        [TestMethod]
        public void GetMessageHistoryWrongValidation()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(false);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("chatId", 4));
            var result = (Response)implementor.GetMessageHistory(_wrongKeyToken, jObject);
            Assert.AreEqual(result.Result, "Validation failed!!!");
        }

        [TestMethod]
        public void GetMessageHistoryWrongParams()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("name", "Hi"));
            var result = (Response)implementor.GetMessageHistory(_rightToken, jObject);
            Assert.AreEqual(result.Result, "Invalid params. long ChatId is needed");
        }

        [TestMethod]
        public void GetMessageHistorySuccess()
        {
            var _repository = new Mock<IRepository>();
            _repository.Setup(r => r.GetMessageHistory(It.IsAny<long>())).Returns(() =>
            {
                return new List<Message>();
            });
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("chatId", 4));
            var result = (Response)(implementor.GetMessageHistory(_rightToken, jObject));
            Assert.AreEqual(((MessageList)result.Result).Messages.Count, 0);
        }

        // PostChat

        [TestMethod]
        public void PostChatValidationFailed()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(false);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var result = (Response)implementor.PostChat(_wrongKeyToken, new JObject());
            Assert.AreEqual(result.Result, "Validation failed!!!");
        }

        [TestMethod]
        public void PostChatWrongParams()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("id", "hallo"),
                                      new JProperty("name", true));
            var result = (Response)implementor.PostChat(_rightToken, jObject);
            Assert.AreEqual(result.Result, "Wrong Chat entity was sent. Check Chat entity!");
        }

        [TestMethod]
        public void PostChatReposioryFailed()
        {
            var _repository = new Mock<IRepository>();
            _repository.Setup(r => r.PostChat(It.IsAny<Chat>())).Throws(new Exception("ooohhhh"));
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("name", "Hi"));
            var result = (Response)implementor.PostChat(_rightToken, jObject);
            Assert.AreEqual(result.Result, "Error ocured in chat adding to DB");
        }

        [TestMethod]
        public void PostChatSuccess()
        {
            var _repository = new Mock<IRepository>();
            _repository.Setup(r => r.PostChat(It.IsAny<Chat>())).Returns(() =>
            {
                return 1;
            });
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("name", "Hi"));
            var result = (Response)implementor.PostChat(_rightToken, jObject);
            Assert.AreEqual(((ChatIdParam)result.Result).ChatId, 1);
        }

        // PostCodeChat

        [TestMethod]
        public void PostCodeChatValidationFailed()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(false);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var result = (Response)implementor.PostCodeChat(_wrongKeyToken, new JObject());
            Assert.AreEqual(result.Result, "Validation failed!!!");
        }

        [TestMethod]
        public void PostCodeChatWrongParams()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("id", "hallo"),
                                      new JProperty("name", true));
            var result = (Response)implementor.PostCodeChat(_rightToken, jObject);
            Assert.AreEqual(result.Result, "Wrong CodeChat entity was sent. Check Chat entity!");
        }

        [TestMethod]
        public void PostCodeChatReposioryFailed()
        {
            var _repository = new Mock<IRepository>();
            _repository.Setup(r => r.PostCodeChat(It.IsAny<CodeChat>())).Throws(new Exception("ooohhhh"));
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("name", "Hi"),
                                      new JProperty("chatId", 4));
            var result = (Response)implementor.PostCodeChat(_rightToken, jObject);
            Assert.AreEqual(result.Result, "Error ocured in chat adding to DB");
        }

        [TestMethod]
        public void PostCodeChatSuccess()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("name", "Hi"),
                                      new JProperty("chatId", 4));
            var result = (Response)implementor.PostCodeChat(_rightToken, jObject);
            Assert.AreEqual(((CodeChatIdParam)result.Result).CodeChatId, 0);
        }

        // AddUserToChat

        [TestMethod]
        public void AddUserToChatValidationFailed()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(false);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var result = (Response)implementor.AddUserToChat(_wrongKeyToken, new JObject());
            Assert.AreEqual(result.Result, "Validation failed!!!");
        }

        [TestMethod]
        public void AddUserToChatWrongParams()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("id", "hallo"),
                                      new JProperty("name", true));
            var result = (Response)implementor.AddUserToChat(_rightToken, jObject);
            Assert.AreEqual(result.Result, "Wrong UserChat entity was sent. Check UserChat entity!");
        }

        [TestMethod]
        public void AddUserToChatReposioryFailed()
        {
            var _repository = new Mock<IRepository>();
            _repository.Setup(r => r.PostUserChat(It.IsAny<UserChat>())).Throws(new Exception("ooohhhh"));
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("userName", _username),
                                      new JProperty("chatId", 4));
            var result = (Response)implementor.AddUserToChat(_rightToken, jObject);
            Assert.AreEqual(result.Result, "Error ocured in userchat adding to DB");
        }

        [TestMethod]
        public void AddUserToChatSuccess()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("chatId", 4),
                                      new JProperty("username", _username));
            var result = (Response)implementor.AddUserToChat(_rightToken, jObject);
            Assert.AreEqual(result.Result, "User was successfully added to the chat");
        }

        // LeaveChannel

        [TestMethod]
        public void LeaveChannelValidationFailed()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(false);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var result = (Response)implementor.LeaveChannel(_wrongKeyToken, new JObject());
            Assert.AreEqual(result.Result, "Validation failed!!!");
        }

        [TestMethod]
        public void LeaveChannelWrongParams()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("id", "hallo"),
                                      new JProperty("name", true));
            var result = (Response)implementor.LeaveChannel(_rightToken, jObject);
            Assert.AreEqual(result.Result, "Invalid params. long ChatId is needed");
        }

        [TestMethod]
        public void LeaveChannelReposioryFailed()
        {
            var _repository = new Mock<IRepository>();
            _repository.Setup(r => r.DeleteUserChat(It.IsAny<long>(), It.IsAny<string>())).Throws(new Exception("ooohhhh"));
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("userName", _username),
                                      new JProperty("chatId", 4));
            var result = (Response)implementor.LeaveChannel(_rightToken, jObject);
            Assert.AreEqual(result.Result, "Selected user doesnt have such chat");
        }

        [TestMethod]
        public void LeaveChannelSuccess()
        {
            var _repository = new Mock<IRepository>();
            var _validator = new Mock<IValidator>();
            _validator.Setup(v => v.CheckUser(It.IsAny<string>())).Returns(true);
            var implementor = new Implementor(_repository.Object, _validator.Object);
            var jObject = new JObject(new JProperty("chatId", 4),
                                      new JProperty("username", _username));
            var result = (Response)implementor.LeaveChannel(_rightToken, jObject);
            Assert.AreEqual(result.Result, "User has successfuly leaved the channel");
        }
    }

}

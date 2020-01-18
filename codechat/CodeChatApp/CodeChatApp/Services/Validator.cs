using Microsoft.IdentityModel.Tokens;
using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CodeChatApp.Services
{
    /// <summary>
    /// This service implemets main functional for ayth token validation
    /// </summary>
    public class Validator : IValidator
    {
        private string _secret = "secretsecretsecret";
        private IRepository _repository;

        public Validator(IRepository repository)
        {
            _repository = repository;
        }

        /// <summary>
        /// Validate token
        /// </summary>
        /// <param name="token"></param>
        /// <returns></returns>
        public bool Validate(string token)
        {
            var username = GetUserName(token);
            return CheckUser(username);
        }

        /// <summary>
        /// Find out if user with such username exists in DB
        /// </summary>
        /// <param name="username"></param>
        /// <returns></returns>
        public bool CheckUser(string username)
        {
            if (username != null && _repository.GetUsers().Where(d => d.Username == username).ToList().Count != 0)
                return true;
            else
                return false;
        }

        /// <summary>
        /// Get user name from auth token
        /// </summary>
        /// <param name="token"></param>
        /// <returns></returns>
        public string GetUserName(string token)
        {
            try
            {
                var key = Encoding.ASCII.GetBytes(_secret);
                var handler = new JwtSecurityTokenHandler();
                var tokenSecure = handler.ReadToken(token) as SecurityToken;
                var validations = new TokenValidationParameters
                {
                    ValidateIssuerSigningKey = true,
                    IssuerSigningKey = new SymmetricSecurityKey(key),
                    ValidateLifetime = false,
                    ValidateIssuer = false,
                    ValidateAudience = false
                };
                var claims = handler.ValidateToken(token, validations, out tokenSecure);
                return claims.Claims.ToArray()[3].Value;
            }
            catch (Exception e)
            {
                return null;
            }
        }
    }
}

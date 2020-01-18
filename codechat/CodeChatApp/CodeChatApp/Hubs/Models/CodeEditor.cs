using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace CodeChatApp.Hubs.Models
{
    /// <summary>
    /// Represent codemirror entity on frontend
    /// </summary>
    public class CodeEditor
    {
        /// <summary>
        /// selected coding language
        /// </summary>
        public string ModeValue { get; set; }

        /// <summary>
        /// code listing
        /// </summary>
        public string Text { get; set; }
    }
}

namespace ChaosTasks.Data;

using ChaosTasks.Models;
using System.Collections.Immutable;
using System.IO;
using System.Text.RegularExpressions;
using System.Text;
using System.Linq;

public class TodoParserService {
  public ImmutableList<Todo> ToTodoList(string[] lines) {
    return lines.Select(line => this.ToTodo(line)).ToImmutableList();
  }

  public Todo ToTodo(string line) {
    var tokens = line.Split(' ');
    var priority = tokens.First(t => t.StartsWith('(') && t.EndsWith(')') && t.Length == 3);
    var projects = tokens.Where(t => t.StartsWith('+'));
    var contexts = tokens.Where(t => t.StartsWith('@'));
    var removedTokens = projects.Concat(contexts).Append(priority);
    var text = string.Join(' ', tokens.Where(t => !removedTokens.Contains(t)));

    return new Todo {
      Raw = line,
      Priority = priority?.Substring(1,1),
      Projects = projects.Select(p => p[1..]).ToArray(),
      Contexts = contexts.Select(c => c[1..]).ToArray(),
      Text = text
    };
  }
}

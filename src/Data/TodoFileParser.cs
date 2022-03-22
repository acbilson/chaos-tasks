namespace ChaosTasks.Data;

using System.IO;
using System.Collections.Immutable;
using ChaosTasks.Models;

public class TodoFileParser
{
  private string _rootPath { get; set; }

  public TodoFileParser(string rootPath) {
    _rootPath = rootPath;
  }

  public TodoFileResult ReadLines(string filePath) {
    var result = new TodoFileResult();
    try {
      result.Lines = File.ReadAllLines(Path.Join(_rootPath, filePath));
    } catch (DirectoryNotFoundException e) {
      result.Message = e.Message;
      return result;
    }
    result.Success = true;
    return result;
  }
}

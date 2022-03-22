namespace ChaosTasks.Models;

public class TodoFileResult {
  public bool Success { get; set; }
  public string Message { get; set; }
  public string[] Lines { get; set; }

  public TodoFileResult() {
    Success = false;
    Message = string.Empty;
    Lines = System.Array.Empty<string>();
  }
}


namespace ChaosTasks.Models;
public class Todo
{
  public string Raw { get; set; }
  public string Text { get; set; }
  public string Priority { get; set; }
  public string[] Projects { get; set; }
  public string[] Contexts { get; set; }
}

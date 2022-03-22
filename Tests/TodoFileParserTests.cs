namespace ChaosTasks.Tests;

using NUnit.Framework;
using ChaosTasks.Data;

public class TodoFileParserTests
{
    [SetUp]
    public void Setup()
    {
    }

    [Test]
    public void LoadsData()
    {
      var parser = new TodoFileParser("/Users/acbilson/source/chaos-tasks/src/Tests/data");
      var result = parser.ReadLines("basic.txt");
      Assert.True(result.Success, result.Message);
      Assert.Greater(result.Lines.Length, 0);
    }
}
